package main

import (
	controllers "qweqwe/Controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": "ok",
			"data":    "",
		})
	})
	server.GET("/words/:count", controllers.WordController)
	server.GET("/sentence/:count", controllers.SentenceController)
	server.GET("/paragraph/:count", controllers.ParagraphController)
	server.GET("/image/:imagesize", controllers.ImageController)
	server.GET("/random/:randomitem", controllers.RandomController)

	server.Run(":1234")
}
