package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/samuelsih/fakeapi/cmd/builder"
)

type App struct {
	Server               *http.Server
	Data                 builder.Builder
	infoLog              *log.Logger
	idleConnectionClosed chan struct{}
	port                 string
}

func RunServer(port string, data builder.Builder) {
	if port == "" {
		port = "7000"
	}

	app := createServer(port, data)

	go app.shutdown()

	app.serve()

	<-app.idleConnectionClosed

	app.infoLog.Println("App stopped successfully.")
}

func createServer(port string, data builder.Builder) *App {
	if port == "" {
		port = "7000"
	}

	return &App{
		port:                 port,
		Data:                 data,
		infoLog:              log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		idleConnectionClosed: make(chan struct{}),
	}
}

func (app *App) serve() {
	app.Server = &http.Server{
		Addr:     ":" + app.port,
		Handler:  app.Routes(),
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	figure.NewColorFigure("- FAKE API -", "", "blue", true).Print()
	println()
	app.infoLog.Printf("Listening on http://localhost:%s\n", app.port)

	if err := app.Server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			app.infoLog.Fatal("Server failed to start:", err)
		}
	}

}

func (app *App) shutdown() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	<-sigint

	app.infoLog.Println("Shutdown order received!")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		app.infoLog.Printf("Server shutdown error: %v", err)
	}

	app.infoLog.Println("Shutdown complete")
	close(app.idleConnectionClosed)
	close(sigint)
}
