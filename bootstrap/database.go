package bootstrap

import (
	"context"
	"fmt"
	"go-hexagonal/internal/infrastructure/repository/mongo"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=============================\n", sql)
}

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	println("dbPort", dbPort)
	mongodbURI := fmt.Sprintf("mongodb+srv://%s:%s@%s.nfn7pei.mongodb.net/?retryWrites=true&w=majority", dbUser, dbPass, dbHost)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb+srv://%s:%s@%s.nfn7pei.mongodb.net/?retryWrites=true&w=majority", dbUser, dbPass, dbHost)
	}

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}

func NewPostgresDatabase(env *Env) *gorm.DB {
	var dsn string

	if env.RunEnv == "local" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok", env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName)
	} else {
		instanceConnName := env.InstanceConnectionName
		dbUser := env.DBUser
		dbName := env.DBName
		dbPass := env.DBPass
		var dbURI string

		if env.RunEnv == "cloud-dev" {
			dbURI = fmt.Sprintf("postgres://%s:%s@/%s?host=/cloudsql/%s", dbUser, dbPass, dbName, instanceConnName)
		} else if env.RunEnv == "cloud-prod" {
			dbURI = fmt.Sprintf("host=/cloudsql/%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Bangkok", instanceConnName, dbUser, dbName, env.DBPass)
		} else {
			log.Fatalf("Unknown RUN_ENV: %s", env.RunEnv)
		}

		dsn = dbURI
	}
	fmt.Println("dsn", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: SqlLogger{logger.Default.LogMode(logger.Info)},
		DryRun: false,
	})
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	// postgresDB.MigratePostgresDB(db)

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
