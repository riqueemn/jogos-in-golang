package telas

import (
	"log"

	"image/color"
	_ "image/png"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/riqueemn/jogos/spacex_2/buttons"
	"github.com/riqueemn/jogos/spacex_2/objects"
	"github.com/riqueemn/jogos/spacex_2/status"
	"github.com/riqueemn/jogos/spacex_2/window"
	"golang.org/x/image/font"
)

type TelaVitoria struct {
}

var (
	mplusNormalFontV font.Face
	botaoMenu2       objects.BotaoMenu
)

func init() {

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 100
	mplusNormalFontV = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	botaoMenu2.Width = 0.5
	botaoMenu2.Height = 0.5
	botaoMenu2.X = window.ScreenWidth/2 - 70
	botaoMenu2.Y = window.ScreenHeight / 2
	botaoMenu2.Selecionado = true
}

func (t *TelaVitoria) Update(screen *ebiten.Image) error {
	botaoMenu2.Draw(screen)
	return nil
}

func (t *TelaVitoria) Draw(screen *ebiten.Image) {

	const x = 20
	msg := "Parabéns!!! Você Salvou o Planeta!!!"

	text.Draw(screen, msg, mplusNormalFontV, x, 140, color.White)

	botaoMenu2.Draw(screen)

	if buttons.Enter {
		status.StatusGame = "Inicio"

	}

	/*
		text := fmt.Sprintf(`TPS: %0.2f
		FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, text)
	*/
}
