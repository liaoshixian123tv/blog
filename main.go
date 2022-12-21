package main

import (
	"blog/internal/routers"
	"net/http"
	"time"
)

func main() {

	router := routers.BlogRouter()
	s := http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
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
