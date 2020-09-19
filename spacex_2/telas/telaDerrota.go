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

var (
	botaoMenu        objects.BotaoMenu
	mplusNormalFontD font.Face
)

type TelaDerrota struct {
}

func init() {

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 100
	mplusNormalFontD = truetype.NewFace(tt, &truetype.Options{
		Size:    8.5,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	botaoMenu.Width = 0.5
	botaoMenu.Height = 0.5
	botaoMenu.X = window.ScreenWidth/2 - 70
	botaoMenu.Y = window.ScreenHeight / 2
	botaoMenu.Selecionado = true
}

func (t *TelaDerrota) Update(screen *ebiten.Image) error {
	botaoMenu.Update(screen)

	return nil
}

func (t *TelaDerrota) Draw(screen *ebiten.Image) {

	const x = 20
	msg := "Infelizmente Você não conseguiu salvar o planeta.\nQuem sabe em outro muiltiverso!"

	text.Draw(screen, msg, mplusNormalFontD, x, 120, color.White)

	botaoMenu.Draw(screen)

	if buttons.Enter {
		status.StatusGame = "Inicio"

	}

	/*
		text := fmt.Sprintf(`TPS: %0.2f
		FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, text)
	*/
}
