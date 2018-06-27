package main

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

/* gin app */
var app *gin.Engine

func httpServer() {

	log.Printf("[INFO] Starting up API")
	gin.DefaultWriter = ioutil.Discard
	gin.SetMode(gin.ReleaseMode)
	app = gin.Default()
	//app.GET("/", serverInit)

	var err error
	err = app.Run("59105")

	if err != nil {
		log.Fatal(err)
	}

}
