package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"<%- repourl%>/config"
	"<%- repourl%>/<%- packagename%>"
	"github.com/ricardo-ch/go-logger"
	tracing "github.com/ricardo-ch/go-tracing"
)

const appName = "<%- projectname%>"

func init() {
	// initialization (optional)
	logger.InitLogger(false)
}

func main() {

	//Zipkin Connection
	tracing.SetGlobalTracer(appName, config.SvcTracingZipkin)
	defer tracing.FlushCollector()

	// Errors channel
	errc := make(chan error)

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// <%- packagename%> endpoint
	<%- packagename%>Service := <%- packagename%>.NewService(<%- packagename%>.NewRepository(nil))
	<%- packagename%>Service = <%- packagename%>.NewTracing(<%- packagename%>Service)
	<%- packagename%>Handler := <%- packagename%>.NewHandler(<%- packagename%>Service)

	go func() {

		httpAddr := ":" + config.AppPort
		router := mux.NewRouter()

		// index endpoint
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Welcome to the demo-yogo API!")
		})

		// healthz endpoint
		router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
		})

		router.Handle("/<%- packagename%>/", tracing.HTTPMiddleware("<%- packagename%>-handler", http.HandlerFunc(<%- packagename%>Handler.Get)))

		httpServer := &http.Server{
			Addr:    httpAddr,
			Handler: router,
		}

		logger.Info(fmt.Sprintf("The microservice demo-yogo is started on port %s", config.AppPort), zap.String("port", config.AppPort))
		errc <- httpServer.ListenAndServe()

	}()

	logger.Error("exit", zap.Error(<-errc))
}
