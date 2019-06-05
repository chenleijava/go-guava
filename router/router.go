package router

import (
	"github.com/chenleijava/go-guava"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"time"
)

//gin rourter
//gin run mode!
func router(mode string) *gin.Engine {
	//init go std log
	guava.LogFormatInit()
	gin.SetMode(mode) //

	//router
	var router *gin.Engine
	if mode == gin.DebugMode {
		router = gin.Default()
	} else {
		router = gin.New() //release
		router.Use(gin.Recovery())
	}

	//gin gzip
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// cors config
	{
		//corsConfig := cors.DefaultConfig()
		//corsConfig.AddAllowHeaders("Authorization", "Access-Control-Allow-Origin", "Origin") //Allow header
		//corsConfig.AllowOrigins = []string{"*"}
		//router.Use(cors.New(corsConfig))
		router.Use(cors.New(cors.Config{
			AllowOriginFunc:  func(origin string) bool { return true },
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
			AllowHeaders:     []string{"Authorization", "Access-Control-Allow-Origin", "Origin", "Content-Length", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	//bind router jwt ,session or ws
	//here init session rout bind

	return router
}

//router bind  view and  reset_ful interface
type RoutBind interface {
	//bind view
	bindView(routeGroup *gin.RouterGroup)
	//bind reset
	bindRest(routeGroup *gin.RouterGroup)
}
