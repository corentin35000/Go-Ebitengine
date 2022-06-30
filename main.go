package main

import (
	_ "github.com/codecat/go-enet"
	"github.com/corentin35000/Go-Ebitengine/src/sceneManager"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/text"
	_ "golang.org/x/image/font"
	_ "golang.org/x/image/font/opentype"
	"log"
)

type Game struct{}

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
)

var fpsCurrent float64 = ebiten.CurrentFPS()

func init() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("By Corentin")

	//sceneManager.Init(SCREENWIDTH, SCREENHEIGTH)
}

func (g *Game) Update() error {
	sceneManager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	sceneManager.Draw()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 780
}

func main() {
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
