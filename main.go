package main

import (
	// "fmt"
	"image"
	"image/color"
	"log"
	// "math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	img  *image.RGBA
	grid *Grid
}

func (g *Game) Update() error {
	for _, coords := range g.grid.coordinates {
		g.img.Set(coords.x, coords.y, color.White)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.img.Pix)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	grid := makeGrid(screenWidth, screenHeight)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game of Life")
	g := &Game{
		img:  image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight)),
		grid: grid,
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
