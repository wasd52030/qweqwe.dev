package main

import (
	"log"
	"net/http"
	"os"

	controllers "qweqwe.dev/Controllers"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

func main() {

	readmeMD, err := os.ReadFile("readme.md")
	if err != nil {
		log.Fatal(err)
	}
	readmeHTM := blackfriday.MarkdownCommon(readmeMD)

	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", readmeHTM)
	})
	server.GET("/word/:count", controllers.WordController)
	server.GET("/sentence/:count", controllers.SentenceController)
	server.GET("/paragraph/:count", controllers.ParagraphController)
	server.GET("/image/:imagesize", controllers.ImageController)
	server.GET("/random/:randomitem", controllers.RandomController)

	server.Run(":1234")
}
