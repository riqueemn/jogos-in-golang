package telas

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/riqueemn/jogos/spacex_2/buttons"
	"github.com/riqueemn/jogos/spacex_2/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/riqueemn/jogos/spacex_2/window"

	"image/color"
	_ "image/png"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/riqueemn/jogos/spacex_2/objects"
	"golang.org/x/image/font"
)

var (
	mplusNormalFont        font.Face
	mplusNormalFontRanking font.Face
	botaoHistoria          objects.BotaoHistoria
	botaoArcade            objects.BotaoArcade
	collection_2           *mongo.Collection
	rankings               []Player
	P                      *audio.Player
)

type TelaInicio struct {
}

type Player struct {
	Nome   string `json:"nome,omitempty"`
	Pontos int    `json:"pontos,omitempty"`
}

func init() {

	botaoHistoria.Width = 0.5
	botaoHistoria.Height = 0.5
	botaoHistoria.X = window.ScreenWidth/2 - 70
	botaoHistoria.Y = window.ScreenHeight/2 - 50
	botaoHistoria.Selecionado = true

	botaoArcade.Width = 0.5
	botaoArcade.Height = 0.5
	botaoArcade.X = window.ScreenWidth/2 - 70
	botaoArcade.Y = window.ScreenHeight / 2
	botaoArcade.Selecionado = false

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 150
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    15,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	mplusNormalFontRanking = truetype.NewFace(tt, &truetype.Options{
		Size:    5,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://adm:1234@cluster0.6nykz.mongodb.net/admin?retryWrites=true&w=majority"))
	if err != nil {
		//log.Fatal(err)
	}
	status.Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(status.Ctx)
	if err != nil {
		//log.Fatal(err)
	}
	//defer client.Disconnect(status.Ctx)

	status.Db = client.Database("spaceX")

	collection := status.Db.Collection("nomes")

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		//log.Fatal(err)
	}

	if err != nil {
		//log.Fatal(err)
	} else {

		for cur.Next(status.Ctx) {
			var jogador Player
			err := cur.Decode(&jogador)

			if err == nil {
				rankings = append(rankings, jogador)
			}
		}
		//defer cur.Close(context.TODO())

		//log.Fatal(err)
	}
	ctx, err := audio.NewContext(status.SampleRate)
	if err != nil {
		log.Fatal(err)
	}
	status.AudioContext = ctx

	mp3F, err := ebitenutil.OpenFile("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/sons/tema_1.mp3")
	if err != nil {
		log.Fatal(err)
	}

	s, err := mp3.Decode(status.AudioContext, mp3F)
	if err != nil {
		log.Fatal(err)
	}

	P, err = audio.NewPlayer(status.AudioContext, s)
	if err != nil {
		log.Fatal(err)
	}
	P.SetVolume(0.4)
	//P.Play()

}

func (t *TelaInicio) Update(screen *ebiten.Image) error {
	for i := 0; i < len(rankings); i++ {
		for j := i; j < len(rankings); j++ {
			if rankings[j].Pontos > rankings[i].Pontos {
				auxN := rankings[i].Nome
				aux := rankings[i].Pontos
				rankings[i].Pontos = rankings[j].Pontos
				rankings[i].Nome = rankings[j].Nome
				rankings[j].Pontos = aux
				rankings[j].Nome = auxN
			}
		}
	}

	if buttons.Down {
		botaoArcade.Selecionado = true
		botaoHistoria.Selecionado = false
	}
	if buttons.Up {
		botaoArcade.Selecionado = false
		botaoHistoria.Selecionado = true

	}

	if botaoHistoria.Selecionado && buttons.Enter {
	} else {

		botaoArcade.Update(screen)
		botaoHistoria.Update(screen)
	}
	return nil
}

func (t *TelaInicio) Draw(screen *ebiten.Image) {
	botaoHistoria.Draw(screen)
	botaoArcade.Draw(screen)
	if botaoHistoria.Selecionado && buttons.Enter {
		status.StatusGame = "HistoriaI"
	} else if botaoArcade.Selecionado && buttons.Enter {
		status.Play = true
		status.StatusGame = "Arcade"
	}

	const x = 20
	msg := "SpaceX"

	text.Draw(screen, msg, mplusNormalFont, x, 40, color.White)

	msg = "Ranking"

	text.Draw(screen, msg, mplusNormalFontRanking, 10, 200, color.White)

	msg = fmt.Sprintf("1. %v %v\n2. %v %v\n3. %v %v\n4. %v %v\n5. %v %v",
		rankings[0].Nome, rankings[0].Pontos, rankings[1].Nome, rankings[1].Pontos,
		rankings[2].Nome, rankings[2].Pontos, rankings[3].Nome, rankings[3].Pontos,
		rankings[4].Nome, rankings[4].Pontos)

	text.Draw(screen, msg, mplusNormalFontRanking, 10, 220, color.White)
	/*
		text := fmt.Sprintf(`TPS: %0.2f
		FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, text)
	*/
}
