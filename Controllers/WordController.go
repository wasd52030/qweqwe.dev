package controllers

import (
	"math/rand"
	"strconv"
	"time"

	lorem "github.com/drhodes/golorem"
	"github.com/gin-gonic/gin"
)

func WordController(ctx *gin.Context) {
	rand.Seed(time.Now().Unix())
	count, _ := strconv.Atoi(ctx.Param("count"))
	words := []string{}
	for i := 0; i < count; i++ {
		words = append(words, lorem.Word(1, rand.Intn(20)+1))
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "some words",
		"data":    words,
	})
}
