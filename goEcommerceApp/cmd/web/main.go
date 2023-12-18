package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// using a css version we can append to .css and .js files to force browsers
// to get the new version, rather than clearing their cache
const cssVersion = "1"

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secretKey string
		publicKey string
	}
}

type application struct {
	config        config
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	templateCache map[string]*template.Template
	version       string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLogger.Println("Starting HTTP server in %s mode on port %d", app.config.env, app.config.port)

	return srv.ListenAndServe()
}

func main() {
	var config config

	flag.IntVar(&config.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&config.env, "env", "development", "Application enviroment {development|production}")
	flag.StringVar(&config.api, "api", "http://localhost:4001", "URL to api")

	flag.Parse()

	config.stripe.publicKey = os.Getenv("STRIPE_KEY")
	config.stripe.secretKey = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	tc := make(map[string]*template.Template)

	app := &application{
		config:        config,
		infoLogger:    infoLog,
		errorLogger:   errorLog,
		templateCache: tc,
		version:       version,
	}

	err := app.serve()
	if err != nil {
		app.errorLogger.Println(err)
		log.Fatal()
	}
}
