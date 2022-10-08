package main

import (
	"fmt"
	"os"

	"github.com/LordRahl90/auth-manager-ot/domain/users/repository/database"
	"github.com/LordRahl90/auth-manager-ot/servers"

	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" || env == "development" {
		if err := godotenv.Load(".envs/.env"); err != nil {
			panic(err)
		}
	}
	db, err := setupDB()
	if err != nil {
		panic(err)
	}
	if err := migrate(db); err != nil {
		panic(err)
	}

	tp, err := newExporter(os.Getenv("TELEMETRY_ENDPOINT"))
	if err != nil {
		fmt.Printf("telemetry not well configured")
	}
	if tp == nil {
		panic("TP is wrongly setup")
	}
	if tp != nil {
		otel.SetTracerProvider(tp)
		tp.Tracer("component-http")
	}

	srv := servers.New(db)
	if err := srv.Start(":4200"); err != nil {
		panic(err)
	}
}

func newExporter(endpoint string) (*sdktrace.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("auth-manager"),
				attribute.String("environment", "development"), attribute.Int64("ID", 1))))

	return tp, nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(&database.User{})
}

func setupDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbase := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, dbase)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
