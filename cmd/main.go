package main

import (
	"github.com/fdsfss/go-bookstore-api"
	"github.com/fdsfss/go-bookstore-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(bookstore.Server)
	if err := srv.Start("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("server start error:%v", err)
	}
}
