package main

import (
	"gitchat/chat1-web后端实战/project/extern/externRouter"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// extern
    externRouter.ExternRouter(r)
	// prop

	s := &http.Server{
		Addr:           "8112",
		Handler:        cors.AllowAll().Handler(r),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 21,
	}
	s.ListenAndServe()
}
