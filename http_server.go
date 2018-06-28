package main

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

/* gin app */
var app *gin.Engine

func serverInit(c *gin.Context) {
	//c.Header("Access-Control-Allow-Origin", "*")
	//get, _ = url.ParseQuery(string(c.Request.URL.RawQuery))

	token := c.DefaultQuery("token", "")
	obj := lru.Get(token)
	var hosts []string
	if obj != nil {
		hosts = obj.([]string)
	}
	c.JSON(200, gin.H{"ns": hosts, "ip": "xxxx"})
	return
}

func httpServer(addr string) {

	log.Printf("Starting up HTTP Server, LISTEN : %s", addr)
	gin.DefaultWriter = ioutil.Discard
	gin.SetMode(gin.ReleaseMode)
	app = gin.Default()
	app.GET("/", serverInit)

	var err error
	err = app.Run(addr)

	if err != nil {
		log.Fatal(err)
	}

}
