package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
	"testing"
)

func TestPprof(t *testing.T) {
	r := gin.New()
	pprof.Register(r, &pprof.Options{RoutePrefix: "debug/pprof"})
	log.Printf("http://localhost:8080/debug/pprof/heap")
	r.Run(":8080")
}
