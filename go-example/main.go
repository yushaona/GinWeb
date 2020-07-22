package main

import (
	"fmt"
	"net/http"

	"go-example/pkg/setting"
	"go-example/routers"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println(s.Addr)

	s.ListenAndServe()
}
