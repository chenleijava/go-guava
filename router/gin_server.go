package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	DebugMode   = "debug" //  debug
	ReleaseMode = "release"
	TestMode    = "test"
)

//start http server
//base on gin framework
//this method must be hook!!!
//routeBindImpl : bind-reset or bind-view ---> route config!?
func Start(serverPort int, mode string, routeBindImpl func(route *gin.Engine)) {

	//router
	router := router(mode)


	var address = fmt.Sprintf(":%d", serverPort)
	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}


	//router logic bind-rest or bind-view ?!
	routeBindImpl(router)


	log.Printf("listen port:%s", address)
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()



	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	//
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
