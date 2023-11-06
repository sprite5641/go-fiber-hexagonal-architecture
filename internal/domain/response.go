package domain

import (
	"go-hexagonal/pkg/logger"
	"os"

	"github.com/gofiber/fiber/v2"
)

const (
	ProductionEnv       = "production"
	InternalServerError = "Internal Server Error"
)

var appEnv string

func init() {
	appEnv = os.Getenv("RUN_ENV")
}

type HandleError struct {
	Code  int    `json:"code,omitempty"`
	Error string `json:"error"`
}

type ErrorResponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
	Success    bool   `json:"success"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
}

type IResponse interface {
	SendSuccess(code int, data interface{}) IResponse
	SendSuccessWitOutData(code int) IResponse
	SendError(code int, msg string) IResponse
	Res() error
}

type Response struct {
	StatusCode int
	Data       interface{}
	Success    bool
	Context    *fiber.Ctx
	ErrorRes   *ErrorResponse
	SuccessRes *SuccessResponse
}

func NewResponse(c *fiber.Ctx) IResponse {
	return &Response{
		Context: c,
	}
}

func (r *Response) SendSuccess(code int, data interface{}) IResponse {
	logger.InitLogger(r.Context, &r.SuccessRes, code).Print()
	r.SuccessRes = &SuccessResponse{
		Message: "success",
		Data:    data,
		Success: true,
	}
	r.StatusCode = code
	r.Success = true
	return r
}

func (r *Response) SendSuccessWitOutData(code int) IResponse {
	logger.InitLogger(r.Context, &r.SuccessRes, code).Print()
	r.SuccessRes = &SuccessResponse{
		Message: "success",
		Success: true,
	}
	r.StatusCode = code
	r.Success = true
	return r
}

func (r *Response) SendError(code int, msg string) IResponse {
	log := logger.InitLogger(r.Context, &r.ErrorRes, code)
	log.Print()

	if appEnv == ProductionEnv && code == 500 {
		msg = InternalServerError
	}

	r.ErrorRes = &ErrorResponse{
		Message:    msg,
		StatusCode: code,
		Success:    false,
	}
	r.StatusCode = code
	r.Success = false

	return r
}

func (r *Response) Res() error {
	return r.Context.Status(r.StatusCode).JSON(func() any {
		if !r.Success {
			return &r.ErrorRes
		}
		return &r.SuccessRes
	}())
}
