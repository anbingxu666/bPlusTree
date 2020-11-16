package main

import (
	"BPTree_Web/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.Init()

	r.POST("/create", routers.CreateTree)
	r.POST("/show", routers.ShowTree)
	r.POST("/get", routers.Get)
	r.POST("/update", routers.Update)
	r.POST("/remove", routers.Remove)
	r.POST("/marshal", routers.Marshal)
	r.POST("/travel", routers.Travel)

	r.Run(":8080")
}
