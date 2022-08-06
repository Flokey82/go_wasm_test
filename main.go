package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/Flokey82/go_gens/gencellular"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	height = 512
	width  = 512
)

const sampleText = `The quick brown fox jumps over the lazy dog.`

var (
	mplusNormalFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    14,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

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
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		rand.Seed(time.Now().UnixNano())
		w.Culture.Reset()
	}
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

	var col byte
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if w.Cells[(w.Generation+1)%2][x][y] {
				col = 0xFF
			} else {
				col = 0x00
			}
			w.pixels[(y*width+x)*4] = col
			w.pixels[(y*width+x)*4+1] = col
			w.pixels[(y*width+x)*4+2] = col
			w.pixels[(y*width+x)*4+3] = col
		}
	}
	screen.ReplacePixels(w.pixels)
	// Draw info
	const x = 2
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	text.Draw(screen, msg, mplusNormalFont, x, 17, color.White)
	text.Draw(screen, "Press R to reset", mplusNormalFont, x, 17+17, color.White)
}
