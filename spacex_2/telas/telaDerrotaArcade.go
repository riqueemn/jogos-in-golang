package telas

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"image/color"
	_ "image/png"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/riqueemn/jogos/spacex_2/buttons"
	"github.com/riqueemn/jogos/spacex_2/objects"
	"github.com/riqueemn/jogos/spacex_2/status"
	"github.com/riqueemn/jogos/spacex_2/window"
	"golang.org/x/image/font"
)

var (
	botaoEnter objects.BotaoEnter
)

type TelaDerrotaArcade struct {
	text    string
	counter int
}

type Dados struct {
	Nome   string
	Pontos int
}

func init() {

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    20,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	botaoEnter.Width = 0.5
	botaoEnter.Height = 0.5
	botaoEnter.X = window.ScreenWidth/2 - 70
	botaoEnter.Y = window.ScreenHeight / 2
	botaoEnter.Selecionado = true
}

func (t *TelaDerrotaArcade) Update(screen *ebiten.Image) error {
	botaoEnter.Update(screen)
	t.text += string(ebiten.InputChars())

	// Adjust the string to be at most 10 lines.
	ss := strings.Split(t.text, "\n")
	if len(ss) > 10 {
		t.text = strings.Join(ss[len(ss)-10:], "\n")
	}

	// If the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
		//t.text += "\n"
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(t.text) >= 1 {
			t.text = t.text[:len(t.text)-1]
		}
	}

	if buttons.Enter {
		//status.Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

		//defer client.Disconnect(ctx)

		collection := status.Db.Collection("nomes")

		_, err := collection.InsertOne(context.TODO(), &Dados{Nome: t.text, Pontos: status.Pontos})
		if err != nil {
			log.Fatal(err)
		}
		status.Pontos = 0
		t.text = ""
		fmt.Println(t.text)
	}
	t.counter++

	return nil
}

func (t *TelaDerrotaArcade) Draw(screen *ebiten.Image) {

	msg := fmt.Sprintf("Pontos: %d", status.Pontos)

	text.Draw(screen, msg, mplusNormalFont, 20, 40, color.White)
	msg = "Digite Seu nome:"

	text.Draw(screen, msg, mplusNormalFont, 20, 100, color.White)
	tx := t.text
	if t.counter%60 < 30 {
		tx += "_"
	}
	text.Draw(screen, tx, mplusNormalFont, 20, 150, color.White)
	botaoEnter.Draw(screen)

	if buttons.Enter {

		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://adm:1234@cluster0.6nykz.mongodb.net/admin?retryWrites=true&w=majority"))
		if err != nil {
			//log.Fatal(err)
		}
		status.Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(status.Ctx)
		if err != nil {
			//log.Fatal(err)
		}

		status.Db = client.Database("spaceX")

		var a []Player
		rankings = a
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

			//log.Fatal(err)
		}
		status.StatusGame = "Inicio"

	}

	//ebitenutil.DebugPrint(screen, tx)
	/*
		text := fmt.Sprintf(`TPS: %0.2f
		FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, text)
	*/
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
