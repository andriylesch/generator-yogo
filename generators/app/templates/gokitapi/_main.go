package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	logger "github.com/ricardo-ch/go-logger"
	tracing "github.com/ricardo-ch/go-tracing"
	"<%- repourl%>/config"
	"<%- repourl%>/<%- packagename%>"
)

var appName = "<%- projectname%>"

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

	// create <%- packagename%> service
	<%- packagename%>Service, err := <%- packagename%>.NewService()
	if err != nil {
		logger.Error(err.Error())
	}

	<%- packagename%>Service = <%- packagename%>.NewTracing(<%- packagename%>Service)	
	// create <%- packagename%> endpoints
	<%- packagename%>Endpoint := <%- packagename%>.Endpoints{
		GetEndpoint: <%- packagename%>.MakeEndpoint(<%- packagename%>Service),
	}

	// HTTP Transport
	go func() {

		httpAddr := ":" + config.AppPort
		mux := http.NewServeMux()

		// index endpoint
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Welcome to the <%- projectname%> API!")
		})

		// healthz endpoint
		mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
		})

		mux.Handle("/<%- packagename%>/", tracing.HTTPMiddleware("<%- packagename%>_handler", <%- packagename%>.MakeHTTPHandler(<%- packagename%>Endpoint)))

		httpServer := &http.Server{
			Addr:    httpAddr,
			Handler: mux,
		}

		logger.Info(fmt.Sprintf("The microservice <%- projectname%> is started on port %s", config.AppPort), zap.String("port", config.AppPort))
		errc <- httpServer.ListenAndServe()

	}()

	logger.Error("exit", zap.Error(<-errc))
}