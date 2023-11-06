package bootstrap

import (
	"context"
	"fmt"
	"go-hexagonal/internal/infrastructure/repository/mongo"
	"log"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"

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
	var gormDialector gorm.Dialector

	if env.RunEnv == "local" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok", env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName)

		gormDialector = postgres.Open(dsn)
	} else {
		if env.RunEnv == "cloud-dev" || env.RunEnv == "cloud-prod" {
			dsn = fmt.Sprintf("host=/cloudsql/%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Bangkok", env.InstanceConnectionName, env.DBUser, env.DBName, env.DBPass)
		} else {
			log.Fatalf("Unknown RUN_ENV: %s", env.RunEnv)
		}

		gormDialector = postgres.New(postgres.Config{
			DriverName:           "cloudsqlpostgres",
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		})
	}

	db, err := gorm.Open(gormDialector, &gorm.Config{
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
