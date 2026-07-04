package main

import (
	"context"
	"log"
	"my-gin-app/internal/app"
	httpserver "my-gin-app/internal/httpServer"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Startup failed : %s", err)
	}
	router := httpserver.NewRouter(a)
	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Printf("Shutdown warning : %v", err)
		}
	}()
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: time.Second * 5,
	}

	log.Printf("API RUNNING ON %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("Server closed")
			return
		}
		log.Fatalf("Server error %v", err)
	}
}
