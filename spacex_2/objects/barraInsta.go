package objects

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var (
	barraInstaImage  *ebiten.Image
	barraInstaImageT *ebiten.Image
)

type BarraInsta struct {
	Width   float64
	Height  float64
	X       float64
	Y       float64
	Life    int
	lifeInt int
}

func init() {

	reader, err := os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/barraInsta_png.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := origEbitenImage.Size()
	barraInstaImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	barraInstaImage.DrawImage(origEbitenImage, op)

	reader, err = os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/barraPlaneta_t.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h = origEbitenImage.Size()
	barraInstaImageT, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op = &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	barraInstaImageT.DrawImage(origEbitenImage, op)
}

func (b *BarraInsta) Update(screen *ebiten.Image) error {
	b.lifeInt = (296 * b.Life) / 100
	return nil
}

func (b *BarraInsta) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.Width, b.Height)
	op.GeoM.Translate(b.X, b.Y)
	screen.DrawImage(barraInstaImage.SubImage(image.Rect(0, 0, b.lifeInt, 28)).(*ebiten.Image), op)
	screen.DrawImage(barraInstaImageT, op)

}
