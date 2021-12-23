package main

import (
	"image"
	"image/color"
	"log"

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
	for i, entity := range g.grid.entities {
		if !shouldLive(&entity, i, g.grid) {
			g.grid.entities[i].alive = 0
		} else {
			g.grid.entities[i].alive = 1
		}
	}

	for _, entity := range g.grid.entities {
		if entity.alive == 1 {
			g.img.Set(entity.x, entity.y, color.White)
		} else {
			g.img.Set(entity.x, entity.y, color.Black)
		}
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
	// ebiten.SetMaxTPS(5)
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
