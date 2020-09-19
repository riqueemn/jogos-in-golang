package objects

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/riqueemn/jogos/spacex_2/window"
)

var (
	FireImage *ebiten.Image
	fire      = window.Fire_1
)

type Fire struct {
	Width  float64
	Height float64
	X      float64
	Y      float64
	Vel    float64
	Exists bool
	Player string
}

func init() {

	reader, err := os.Open(fire)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := origEbitenImage.Size()
	FireImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	FireImage.DrawImage(origEbitenImage, op)
}

func (f *Fire) Update() {
	if f.Exists {
		f.Y -= f.Vel
	}
}
