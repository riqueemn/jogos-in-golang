package objects

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var (
	botaoHistoriaImage       *ebiten.Image
	botaoHistoriaSelecionado *ebiten.Image
)

type BotaoHistoria struct {
	Width       float64
	Height      float64
	X           float64
	Y           float64
	Selecionado bool
}

func init() {

	reader, err := os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/botao_ModoHistoria.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := origEbitenImage.Size()
	botaoHistoriaImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	botaoHistoriaImage.DrawImage(origEbitenImage, op)

	reader, err = os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/botao_ModoHistoriaT.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h = origEbitenImage.Size()
	botaoHistoriaSelecionado, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op = &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	botaoHistoriaSelecionado.DrawImage(origEbitenImage, op)
}

func (b *BotaoHistoria) Update(screen *ebiten.Image) error {
	return nil
}

func (b *BotaoHistoria) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.Width, b.Height)
	op.GeoM.Translate(b.X, b.Y)
	if b.Selecionado {
		screen.DrawImage(botaoHistoriaSelecionado, op)
	} else {
		screen.DrawImage(botaoHistoriaImage, op)
	}

}
