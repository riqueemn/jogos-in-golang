package objects

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var (
	botaoArcadeImage       *ebiten.Image
	botaoArcadeSelecionado *ebiten.Image
)

type BotaoArcade struct {
	Width       float64
	Height      float64
	X           float64
	Y           float64
	Selecionado bool
}

func init() {

	reader, err := os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/botao_ModoArcade.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := origEbitenImage.Size()
	botaoArcadeImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	botaoArcadeImage.DrawImage(origEbitenImage, op)

	reader, err = os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/botao_ModoArcadeT.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h = origEbitenImage.Size()
	botaoArcadeSelecionado, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op = &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	botaoArcadeSelecionado.DrawImage(origEbitenImage, op)
}

func (b *BotaoArcade) Update(screen *ebiten.Image) error {
	return nil
}

func (b *BotaoArcade) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.Width, b.Height)
	op.GeoM.Translate(b.X, b.Y)
	if b.Selecionado {
		screen.DrawImage(botaoArcadeSelecionado, op)
	} else {
		screen.DrawImage(botaoArcadeImage, op)
	}

}
