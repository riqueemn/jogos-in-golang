package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240

	sampleRate = 22050
)

type Game struct {
}

func (g *Game) Update(screen *ebiten.Image) error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Audio (Ebiten Demo)")

	mp3F, err := ebitenutil.OpenFile("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/sons/explosion.mp3")
	if err != nil {
		log.Fatal(err)
	}

	audioContext, err := audio.NewContext(sampleRate)

	som, err := mp3.Decode(audioContext, mp3F)
	if err != nil {
		log.Fatal(err)
	}
	p, err := audio.NewPlayer(audioContext, som)
	p.Play()

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
