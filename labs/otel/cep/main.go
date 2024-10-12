package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yodalis/labs/otel/cep/webapp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initTracer() (*sdktrace.TracerProvider, error) {
	serviceName := "cep-service"

	exporter, err := zipkin.New("http://zipkin-all-in-one:9411/api/v2/spans")
	if err != nil {
		log.Fatal(err)
	}

	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
		),
	)

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	return tp, nil
}

func main() {
	tp, err := initTracer()
	if err != nil {
		panic(err)
	}
	defer tp.Shutdown(context.Background())

	cepRouter := chi.NewRouter()
	cepRouter.Use(middleware.Logger)
	cepRouter.Post("/", webapp.CEPHandler)

	fmt.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", cepRouter); err != nil {
		log.Fatalf("Erro ao iniciar o servidor na porta 8080: %v", err)
	}
}
