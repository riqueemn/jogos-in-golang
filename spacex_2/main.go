package main

import (
	"context"
	"log"
	"time"

	"github.com/riqueemn/jogos/spacex_2/telas"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/image/font"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/riqueemn/jogos/spacex_2/buttons"
	"github.com/riqueemn/jogos/spacex_2/game"
	"github.com/riqueemn/jogos/spacex_2/status"
	"github.com/riqueemn/jogos/spacex_2/window"
)

const ()

var (
	jogo            game.Jogo
	telaInicio      telas.TelaInicio
	telaVitoria     telas.TelaVitoria
	telaDerrotaH    telas.TelaDerrota
	telaDerrotaA    telas.TelaDerrotaArcade
	telaHistoria    telas.TelaHistoria
	mplusNormalFont font.Face
	P               *audio.Player
	som             bool

	//AudioContext    *audio.Context
)

type Game struct {
}

func init() {

	status.StatusGame = "Inicio"
	/*
		ctx, err := audio.NewContext(sampleRate)
		if err != nil {
			log.Fatal(err)
		}
		AudioContext = ctx
	*/
	som = true
}

func (g *Game) Update(screen *ebiten.Image) error {
	buttons.Update()

	if status.Play && status.StatusGame == "Historia" {
		telas.P.Pause()
		telas.P.Rewind()
		if som {

			P.Play()
		}
		game.Reset = false
		//jogo.Reset()
		jogo.Mode = "Historia"
		jogo.Update(screen)
		som = false
	} else if status.Play && status.StatusGame == "Arcade" {
		telas.P.Pause()
		telas.P.Rewind()
		if som {

			P.Play()
		}
		game.Reset = false
		//var jogoReset game.Jogo
		//game.Reset = false
		jogo.Mode = "Arcade"
		jogo.Update(screen)
		//telas.P.Rewind()
		//game.P.Play()
		som = false

	} else if status.StatusGame == "Inicio" {
		var telaInicioReset telas.TelaInicio
		telaInicio = telaInicioReset
		telas.P.Play()
		status.Pontos = 0
		//jogo.Reset()
		//game.Bomba.Init()
		telaInicio.Update(screen)
		//telas.P.Play()
		som = true
	} else if status.StatusGame == "Vitoria" {
		P.Pause()
		P.Rewind()
		jogo.Reset()
		telaVitoria.Update(screen)
	} else if status.StatusGame == "DerrotaH" {
		P.Pause()
		P.Rewind()
		jogo.Reset()
		telaDerrotaH.Update(screen)
	} else if status.StatusGame == "DerrotaA" {
		P.Pause()
		P.Rewind()
		//game.P.
		jogo.Reset()
		//jogo.Reset()
		telaDerrotaA.Update(screen)
	} else if status.StatusGame == "HistoriaI" {
		telaHistoria.Update(screen)
	}

	if buttons.Enter {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://adm:1234@cluster0.6nykz.mongodb.net/admin?retryWrites=true&w=majority"))
		if err != nil {
			log.Fatal(err)
		}
		status.Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(status.Ctx)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if status.Play {
		jogo.Draw(screen)
	} else if status.StatusGame == "Inicio" {

		telaInicio.Draw(screen)
		//jogo.Reset()
	} else if status.StatusGame == "Vitoria" {
		telaVitoria.Draw(screen)
	} else if status.StatusGame == "DerrotaH" {
		telaDerrotaH.Draw(screen)
	} else if status.StatusGame == "DerrotaA" {
		telaDerrotaA.Draw(screen)
		//jogo.Reset()
	} else if status.StatusGame == "HistoriaI" {
		telaHistoria.Draw(screen)
	}
	/*
		text := fmt.Sprintf(`                                  TPS: %0.2f
				                                       FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, text)
	*/
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return window.ScreenWidth, window.ScreenHeight
}

func main() {

	mp3F, err := ebitenutil.OpenFile("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/sons/tema_2.mp3")
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

	ebiten.SetWindowSize(window.ScreenWidth*2, window.ScreenHeight*2)
	ebiten.SetWindowTitle("SpaceX")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
