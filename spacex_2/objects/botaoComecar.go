package objects

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var (
	botaoComecarImage       *ebiten.Image
	botaoComecarSelecionado *ebiten.Image
)

type BotaoComecar struct {
	Width       float64
	Height      float64
	X           float64
	Y           float64
	Selecionado bool
}

func init() {

	reader, err := os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/botao_Comecar.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := origEbitenImage.Size()
	botaoComecarImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	botaoComecarImage.DrawImage(origEbitenImage, op)

	reader, err = os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/botao_ComecarT.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h = origEbitenImage.Size()
	botaoComecarSelecionado, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op = &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	botaoComecarSelecionado.DrawImage(origEbitenImage, op)
}

func (b *BotaoComecar) Update(screen *ebiten.Image) error {
	return nil
}

func (b *BotaoComecar) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.Width, b.Height)
	op.GeoM.Translate(b.X, b.Y)
	if b.Selecionado {
		screen.DrawImage(botaoComecarImage, op)
	} else {
		screen.DrawImage(botaoComecarSelecionado, op)
	}

}
