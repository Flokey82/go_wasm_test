package main

import (
	"log"

	"github.com/Flokey82/go_gens/gencellular"
	"github.com/hajimehoshi/ebiten"
)

const (
	height = 128
	width  = 128
)

func main() {
	g := newGame()
	ebiten.SetWindowSize(width*2, height*2)
	ebiten.SetWindowTitle("Dungeon Crawler")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	*gencellular.Culture
	pixels []byte
}

func newGame() *Game {
	return &Game{
		Culture: gencellular.New(width, height),
	}
}

// Update game state by one tick.
func (w *Game) Update() error {
	w.Culture.Tick()
	return nil
}

func (w *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

// Draw paints current game state.
func (w *Game) Draw(screen *ebiten.Image) {
	if w.pixels == nil {
		w.pixels = make([]byte, width*height*4)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if w.Cells[(w.Generation+1)%2][x][y] {
				w.pixels[(y*width+x)*4] = 0xff
				w.pixels[(y*width+x)*4+1] = 0xff
				w.pixels[(y*width+x)*4+2] = 0xff
				w.pixels[(y*width+x)*4+3] = 0xff
			} else {
				w.pixels[(y*width+x)*4] = 0x00
				w.pixels[(y*width+x)*4+1] = 0x00
				w.pixels[(y*width+x)*4+2] = 0x00
				w.pixels[(y*width+x)*4+3] = 0x00
			}
		}
	}
}
