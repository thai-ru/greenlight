package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// version number of our build
const version = "1.0.0"

// config holds configuration settings for our application, e.g version, network port and environment vriables
type config struct {
	port int
	env  string
}

// application Struct holds the dependencies for HTTP handlers, helpers and middleware
type application struct {
	config config
	logger *slog.Logger
}

// main is the main application
func main() {
	// Declare an instance of the config struct
	var cfg config

	//read the value of the port and ENV COMMAND-LINE flags into the config struct
	// default to port 4000 if no flags are provided
	flag.IntVar(&cfg.port, "port", 4000, "API server port")

	// default to environment development if no flags are provided
	flag.StringVar(&cfg.env, "env", "development", "environment(..development|staging|production)")
	flag.Parse()

	// Initialize a new structured logger, writes logs to the standard output stream
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Declare an instance of th application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare a new servermux which dispatches requests to the healthcheckHandler
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Declare a http server which listens on the port provided in the config struct
	// write logs to the logger at Error level

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// Start the server
	logger.Info("Starting server on port ", "port", srv.Addr, "env", cfg.env)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
