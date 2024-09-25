package controllers

import (
	"math/rand"
	"strconv"
	"time"

	lorem "github.com/drhodes/golorem"
	"github.com/gin-gonic/gin"
)

func ParagraphController(ctx *gin.Context) {
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	count, _ := strconv.Atoi(ctx.Param("count"))
	paragraph := []string{}
	for i := 0; i < count; i++ {
		paragraph = append(paragraph, lorem.Paragraph(1, rand.Intn(20)+1))
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "some paragraph",
		"data":    paragraph,
	})
}
