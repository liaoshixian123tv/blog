package main

import (
	global "blog/global/setting"

	"blog/internal/routers"
	"net/http"
)

func main() {

	router := routers.BlogRouter()
	s := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

// func ginServer() {
// 	router := gin.Default()

// 	router.GET("/ping", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, "PING SUCCESS")
// 	})
// 	router.Run()
// }
