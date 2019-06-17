package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
    r.Use(AppFilter)
	r.POST("/activities/invite-activity/invitee-click/", inviteeClick)
	s := &http.Server{
		Addr:           ":8112",
		Handler:        cors.AllowAll().Handler(r),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 21,
	}
	s.ListenAndServe()
}
