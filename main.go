package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	img *image.RGBA
}

func (g *Game) Update() error {
	const pixels = screenWidth * screenHeight
	for i := 0; i < pixels; i++ {
		random := rand.Uint64()
		g.img.Pix[4*i] = uint8(random)
		g.img.Pix[4*i+1] = uint8(random)
		g.img.Pix[4*i+2] = uint8(random)
		g.img.Pix[4*i+3] = 0xff
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
	fmt.Println(grid)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game of Life")
	g := &Game{
		img: image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight)),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
