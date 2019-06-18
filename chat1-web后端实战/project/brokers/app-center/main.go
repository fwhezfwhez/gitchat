package main

import (
	"fmt"
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/activity/activityPb"
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/activity/activityRouter"
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/activity/activityService"
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/props/propPb"
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/props/propRouter"
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/props/propService"
	"github.com/fwhezfwhez/tcpx"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"time"
)

func main() {
	go myHttp()
	go myGrpc()
	go myTcp()
	select {}
}

func myHttp() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// prop
	propRouter.HTTPRouter(r)
	// activity
	activityRouter.HTTPRouter(r)

	s := &http.Server{
		Addr:           "8001",
		Handler:        cors.AllowAll().Handler(r),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 21,
	}
	fmt.Println("http listens on 8001")
	s.ListenAndServe()
}

func myGrpc() {
	lis, err := net.Listen("tcp", ":6001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := grpc.NewServer()
	// prop
	propPb.RegisterPropServiceServer(s, &propService.PropService{})
	// activity
	activityPb.RegisterActivityServiceServer(s, &activityService.ActivityService{})

	fmt.Println("grpc listens on 6001")
	s.Serve(lis)
}

func myTcp() {
	srv := tcpx.NewTcpX(tcpx.ProtobufMarshaller{})
	srv.HeartBeatMode(true, 10*time.Second)
	srv.AddHandler(1, func(c *tcpx.Context) {
		// HeartBeat
		c.RecvHeartBeat()
	})
	go func() {
		fmt.Println("kcp listens on 7002")
		_ = srv.ListenAndServe("kcp","7002")
	}()

	fmt.Println("tcp listens on 7001")
	_ = srv.ListenAndServe("tcp", ":7001")


}
