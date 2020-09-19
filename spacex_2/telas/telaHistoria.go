package telas

import (
	"context"
	"fmt"
	"log"
	"time"

	"image/color"
	_ "image/png"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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
	botaoComecar      objects.BotaoComecar
	mplusNormalFontHI font.Face
	mplusNormalFontHT font.Face
)

type TelaHistoria struct {
	text    string
	counter int
}

func init() {

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 150
	mplusNormalFontHI = truetype.NewFace(tt, &truetype.Options{
		Size:    8,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	mplusNormalFontHT = truetype.NewFace(tt, &truetype.Options{
		Size:    5,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	botaoComecar.Width = 0.5
	botaoComecar.Height = 0.5
	botaoComecar.X = window.ScreenWidth/2 - 70
	botaoComecar.Y = window.ScreenHeight / 2
	botaoComecar.Selecionado = true
}

func (t *TelaHistoria) Update(screen *ebiten.Image) error {
	botaoComecar.Update(screen)
	t.text += string(ebiten.InputChars())

	if buttons.Enter {

	}

	return nil
}

func (t *TelaHistoria) Draw(screen *ebiten.Image) {
	const dpi = 150

	msg := fmt.Sprintf("O ano é 2020...")

	text.Draw(screen, msg, mplusNormalFontHI, 5, 40, color.White)

	msg = "Várias coisas anormais aconteceram..."
	text.Draw(screen, msg, mplusNormalFontHT, 5, 55, color.White)
	msg = "No mês de novembro um ataque alienígena veio para destruir\na Terra!!"

	text.Draw(screen, msg, mplusNormalFontHT, 5, 65, color.White)
	msg = "Salve o planeta contra os bombardeiros das naves alienígenas!!!"

	text.Draw(screen, msg, mplusNormalFontHT, 5, 85, color.White)

	botaoComecar.Draw(screen)

	if buttons.Enter {
		status.Play = true
		status.StatusGame = "Historia"

		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://adm:1234@cluster0.6nykz.mongodb.net/admin?retryWrites=true&w=majority"))
		if err != nil {
			log.Fatal(err)
		}
		status.Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(status.Ctx)
		if err != nil {
			log.Fatal(err)
		}

		status.Db = client.Database("spaceX")

	}

	//ebitenutil.DebugPrint(screen, tx)
	/*
		text := fmt.Sprintf(`TPS: %0.2f
		FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, text)
	*/
}
