package objects

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var (
	botaoMenuImage       *ebiten.Image
	botaoMenuSelecionado *ebiten.Image
)

type BotaoMenu struct {
	Width       float64
	Height      float64
	X           float64
	Y           float64
	Selecionado bool
}

func init() {

	reader, err := os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/botao_Menu.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := origEbitenImage.Size()
	botaoMenuImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	botaoMenuImage.DrawImage(origEbitenImage, op)

	reader, err = os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/botao_MenuT.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h = origEbitenImage.Size()
	botaoMenuSelecionado, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op = &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	botaoMenuSelecionado.DrawImage(origEbitenImage, op)
}

func (b *BotaoMenu) Update(screen *ebiten.Image) error {
	return nil
}

func (b *BotaoMenu) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.Width, b.Height)
	op.GeoM.Translate(b.X, b.Y)
	if b.Selecionado {
		screen.DrawImage(botaoMenuImage, op)
	} else {
		screen.DrawImage(botaoMenuSelecionado, op)
	}

}
