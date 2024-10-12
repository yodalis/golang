package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/yodalis/labs/otel/weather/webapp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func initTracer() (*sdktrace.TracerProvider, error) {
	serviceName := "weather-service"

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

	temperatureRouter := chi.NewRouter()
	temperatureRouter.Use(middleware.Logger)
	temperatureRouter.Get("/weather", webapp.TemperatureHandler)

	fmt.Println("Servidor iniciado na porta 9090")
	if err := http.ListenAndServe(":9090", temperatureRouter); err != nil {
		log.Fatalf("Erro ao iniciar o servidor na porta 9090: %v", err)
	}
}
