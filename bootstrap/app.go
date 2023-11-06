package bootstrap

import (
	"go-hexagonal/internal/infrastructure/repository/mongo"
	"go-hexagonal/internal/infrastructure/repository/redis"

	"gorm.io/gorm"
)

type Application struct {
	Env      *Env
	Mongo    *mongo.Client
	Postgres *gorm.DB
	Redis    *redis.RedisClient
}

type NewApplication interface {
	LoadConfig() Application
	GetEnv() Env
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	// app.Mongo = NewMongoDatabase(app.Env)
	app.Postgres = NewPostgresDatabase(app.Env)
	app.Redis = redis.NewRedisClient(app.Env.RedisHost, app.Env.RedisPassword, app.Env.RedisDB)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(*app.Mongo)
}

func (app *Application) ClosePostgresDBConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.Close()

	return nil
}

func (app *Application) CloseRedisConnection(redis *redis.RedisClient) {
	if err := redis.Close(); err != nil {
		panic(err)
	}
}
