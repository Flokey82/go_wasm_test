package main

import (
	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetFullscreen(true)
	//ebiten.SetWindowSize(dungeon.ScreenWidth*10, dungeon.ScreenHeight*10)
	ebiten.SetWindowTitle("Dungeon Crawler")
}
