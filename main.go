package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type Game struct{}

var img *ebiten.Image

func init() {
	img = ebiten.NewImage(10, 10)
	img.Fill(color.RGBA{0xff, 0, 0, 0xff})
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 780
}

func main() {
	ebiten.SetWindowSize(1280, 780)
	ebiten.SetWindowTitle("By Corentin")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
