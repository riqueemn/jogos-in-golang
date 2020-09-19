package objects

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/riqueemn/jogos/spacex_2/window"
)

var (
	BombaImage *ebiten.Image
	bomba      = window.Bomba_1
)

type Bomba struct {
	Width  float64
	Height float64
	X      float64
	Y      float64
	Exists bool
	Vel    float64
}

func init() {

	reader, err := os.Open(bomba)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := origEbitenImage.Size()
	BombaImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	BombaImage.DrawImage(origEbitenImage, op)
}

func (b *Bomba) Update() {
	if b.Exists {
		//b.Vel = 2
		b.Y += b.Vel
	}
}
