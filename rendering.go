package rendering

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawCircle(sceen *ebiten.Image, centerX, centerY, radius float64, clr color.Color) {
	for angle := 0.0; angle < 2*math.Pi; angle += 0.01 {
		var x = centerX + radius*math.Cos(angle)
		var y = centerY + radius*math.Cos(angle)
		sceen.Set(int(x), int(y), clr)
	}
}
