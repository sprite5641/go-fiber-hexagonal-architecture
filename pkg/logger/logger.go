package logger

import (
	"fmt"
	"go-hexagonal/pkg/utils"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ILogger interface {
	Print() ILogger
	Save()
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResponse(res any)
}

type Logger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func InitLogger(c *fiber.Ctx, res any, code int) ILogger {
	log := &Logger{
		Time:       time.Now().Local().Format("02/01/2006 15:04:05"),
		Ip:         c.IP(),
		Method:     c.Method(),
		Path:       c.Path(),
		StatusCode: code,
	}
	log.SetQuery(c)
	log.SetBody(c)
	log.SetResponse(res)
	return log
}

func (l *Logger) Print() ILogger {
	utils.Debug(l)
	return l
}

func (l *Logger) Save() {
	data := utils.Output(l)

	filename := fmt.Sprintf("./assets/logs/logger_%v.txt", strings.ReplaceAll(time.Now().Format("02-01-2006"), "-", ""))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	file.WriteString(string(data) + "\n")
}

func (l *Logger) SetQuery(c *fiber.Ctx) {
	var body any

	if len(c.Request().URI().QueryString()) == 0 {
		return
	}

	if err := c.QueryParser(&body); err != nil {
		log.Printf("query parser error: %v", err)
	}
	l.Query = body
}

func (l *Logger) SetBody(c *fiber.Ctx) {
	var body any

	if len(c.Request().Body()) == 0 {
		return
	}

	if err := c.BodyParser(&body); err != nil {
		log.Printf("body parser error: %v", err)
	}

	switch l.Path {
	case "v1/users/signup":
		l.Body = "never gonna give you up"
	default:
		l.Body = body
	}
}

func (l *Logger) SetResponse(res any) {
	l.Response = res
}
