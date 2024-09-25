package controllers

import (
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RandomController(ctx *gin.Context) {
	item := strings.Split(ctx.Param("randomitem"), ",")
	if len(item) < 1 {
		ctx.JSON(400, gin.H{
			"status":  400,
			"message": "parameter error",
		})
		return
	}
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	rand.Shuffle(len(item), func(i, j int) { item[i], item[j] = item[j], item[i] })
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "random array",
		"data":    item,
	})
}
