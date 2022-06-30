package sceneManager

import (
	"github.com/corentin35000/Go-Ebitengine/src/sceneGameplay"
	"github.com/corentin35000/Go-Ebitengine/src/sceneMenu"
)

var sceneCurrent string = "Menu"

func Init(ScreenWidth int, ScreenHeight int) {
	sceneMenu.Init()
	sceneGameplay.Init()
}

func Update() {
	switch sceneCurrent {
	case "Menu":
		sceneMenu.Update()
	case "Gameplay":
		sceneGameplay.Update()
	}
}

func Draw() {
	switch sceneCurrent {
	case "Menu":
		sceneMenu.Draw()
	case "Gameplay":
		sceneGameplay.Draw()
	}
}

func UpdateSocketsUDP() {
	switch sceneCurrent {
	case "Menu":
		sceneMenu.UpdateSocketsUDP()
	case "Gameplay":
		sceneGameplay.UpdateSocketsUDP()
	}
}
