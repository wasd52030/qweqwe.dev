package controllers

import (
	"math/rand"
	"strconv"
	"time"

	lorem "github.com/drhodes/golorem"
	"github.com/gin-gonic/gin"
)

func SentenceController(ctx *gin.Context) {
	rand.Seed(time.Now().Unix())
	count, _ := strconv.Atoi(ctx.Param("count"))
	sentence := []string{}
	for i := 0; i < count; i++ {
		sentence = append(sentence, lorem.Sentence(1, rand.Intn(20)+1))
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "some sentence",
		"data":    sentence,
	})
}
