package main

import (
	_ "BPTree_Web/docs"
	"BPTree_Web/routers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	routers.Init()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/create", routers.CreateTree)
	r.POST("/show", routers.ShowTree)
	r.POST("/get", routers.Get)
	r.POST("/update", routers.Update)
	r.POST("/remove", routers.Remove)
	r.POST("/marshal", routers.Marshal)
	r.POST("/travel", routers.Travel)

	r.Run(":8080")
}
