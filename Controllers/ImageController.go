package controllers

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/softwarebackend"
)

func ImageController(ctx *gin.Context) {
	Image := strings.Split(ctx.Param("imagesize"), ".")
	Size := strings.Split(Image[0], "x")
	if len(Image) != 2 || len(Size) != 2 {
		ctx.JSON(400, gin.H{
			"status":  400,
			"message": "parameter error",
		})
		return
	}

	width, werr := strconv.Atoi(Size[0])
	height, herr := strconv.Atoi(Size[1])
	if werr != nil || herr != nil {
		ctx.JSON(400, gin.H{
			"status":  400,
			"message": "parameter error",
		})
		return
	}

	back := softwarebackend.New(width, height)
	canvas := canvas.New(back)
	canvas.SetFillStyle("#198964")
	canvas.FillRect(0, 0, float64(width), float64(height))

	fontsize := width / 10
	canvas.SetFillStyle("#ffffff")
	canvas.SetFont("static/HackRegular.ttf", float64(fontsize))
	text := fmt.Sprintf("%dx%d", width, height)
	textWidth := canvas.MeasureText(text).Width
	canvas.FillText(text, float64(width)/2-textWidth/2, float64(height)/2+float64(fontsize)/2)

	ctx.Header("Content-Type", fmt.Sprintf("image/%s", Image[1]))
	if Image[1] == "png" {
		err := png.Encode(ctx.Writer, back.Image)
		if err != nil {
			ctx.JSON(400, gin.H{
				"status":  400,
				"message": "generate failed",
			})
			return
		}
	} else {
		err := jpeg.Encode(ctx.Writer, back.Image, &jpeg.Options{Quality: 100})
		if err != nil {
			ctx.JSON(400, gin.H{
				"status":  400,
				"message": "generate failed",
			})
			return
		}
	}
}
