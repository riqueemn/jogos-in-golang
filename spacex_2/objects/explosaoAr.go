package objects

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/riqueemn/jogos/spacex_2/status"
)

const (
	frameOXExpAr     = 0
	frameOYExpAr     = 0
	frameWidthExpAr  = 50
	frameHeightExpAr = 50
	frameNumExpAr    = 7
	sampleRate       = 22050
)

var (
	ExplosaoArImage *ebiten.Image
)

type ExplosaoAr struct {
	X          float64
	Y          float64
	Width      float64
	Height     float64
	Exists     bool
	CountExpAr int
	Player     Player
}

type Player struct {
	audioPlayer *audio.Player
	mp3F        ebitenutil.ReadSeekCloser
	toque       bool
}

func init() {
	/*
		ctx, err := audio.NewContext(sampleRate)
		if err != nil {
			log.Fatal(err)
		}
		status.AudioContext = ctx
		//NewContext()
	*/
	reader, err := os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/explosaoAr_png.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := origEbitenImage.Size()
	ExplosaoArImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.8)
	ExplosaoArImage.DrawImage(origEbitenImage, op)
}

func (e *ExplosaoAr) Draw(screen *ebiten.Image) {
	//exp.audioPlayer.Play()

	op := &ebiten.DrawImageOptions{}

	//op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	//op.GeoM.Translate(window.ScreenWidth/2, window.ScreenHeight/2)
	op.GeoM.Translate(e.X, e.Y)
	i := (e.CountExpAr / 6) % frameNumExpAr

	sx, sy := frameOXExpAr+i*frameWidthExpAr, frameOYExpAr
	//op.GeoM.Scale(e.Width, e.Height)

	screen.DrawImage(ExplosaoArImage.SubImage(image.Rect(sx, frameOYExpAr, sx+frameWidthExpAr, sy+frameHeightExpAr)).(*ebiten.Image), op)

	if i == 6 {
		e.Exists = false
		e.CountExpAr = 0
		e.Player.toque = false

	}
	//0, 0, 50, 50
	//50, 0, 100, 50
	//100, 0, 150, 50
	//150, 0, 200, 50
	//200, 0, 250, 50
	//250, 0, 300, 50
	//300, 0, 350, 50
}

func (e *ExplosaoAr) NewContext() {
	e.Player.toque = true

}

func (e *ExplosaoAr) NewPlayer(audioContext *audio.Context) {
	mp3F, err := ebitenutil.OpenFile("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/sons/explosion_2.mp3")
	if err != nil {
		log.Fatal(err)
	}

	e.Player.mp3F = mp3F

	s, err := mp3.Decode(audioContext, e.Player.mp3F)
	if err != nil {
		log.Fatal(err)
	}

	p, err := audio.NewPlayer(audioContext, s)
	if err != nil {
		log.Fatal(err)
	}
	p.SetVolume(0.4)
	e.Player.audioPlayer = p

	e.Player.audioPlayer.Play()

}

func (e *ExplosaoAr) SomExplosao() {

	e.NewPlayer(status.AudioContext)
	e.Player.toque = true

}

func (e *ExplosaoAr) Update(screen *ebiten.Image) error {
	if e.Exists {
		e.CountExpAr++
		if !e.Player.toque {
			e.SomExplosao()
		}
	}

	return nil
}
