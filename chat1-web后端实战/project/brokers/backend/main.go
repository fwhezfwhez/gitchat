package main

import (
	"fmt"
	"gitchat/chat1-web后端实战/project/brokers/backend/module/activity/activityRouter"
	"gitchat/chat1-web后端实战/project/brokers/backend/module/props/propRouter"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	// activity
	activityRouter.Router(r)
	// prop
	propRouter.HTTPRouter(r)
	s := &http.Server{
		Addr:           "8001",
		Handler:        cors.AllowAll().Handler(r),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 21,
	}
	fmt.Println("backend http srv  listens on 8002")
	s.ListenAndServe()
}
