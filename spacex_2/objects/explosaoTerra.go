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
	frameOXExpTerra     = 0
	frameOYExpTerra     = 0
	frameWidthExpTerra  = 110
	frameHeightExpTerra = 94
	frameNumExpTerra    = 5
)

var (
	explosaoTerraImage *ebiten.Image
	player_2           Player
)

type ExplosaoTerra struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	Exists        bool
	CountExpTerra int
	op            ebiten.DrawImageOptions
	Player        Player
}

func init() {

	reader, err := os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/explosaoTerra_png.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := origEbitenImage.Size()
	explosaoTerraImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.8)
	explosaoTerraImage.DrawImage(origEbitenImage, op)
}

func (e *ExplosaoTerra) Draw(screen *ebiten.Image) {

	//op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	//op.GeoM.Translate(window.ScreenWidth/2, window.ScreenHeight/2)
	e.op.GeoM.Reset()
	e.op.GeoM.Scale(0.5, 0.5)
	e.op.GeoM.Translate(e.X, e.Y)
	i := (e.CountExpTerra / 5) % frameNumExpTerra
	sx, sy := frameOXExpTerra+i*frameWidthExpTerra, frameOYExpTerra

	screen.DrawImage(explosaoTerraImage.SubImage(image.Rect(sx, frameOYExpTerra, sx+frameWidthExpTerra, sy+frameHeightExpTerra)).(*ebiten.Image), &e.op)

	if i == 4 {
		e.Exists = false
		e.CountExpTerra = 0
		player_2.toque = false
	}
	//10, 0, 50, 70
	//75, 0, 115, 70
	//140, 0, 185, 70
}

func (e *ExplosaoTerra) NewContext() {
	e.Player.toque = true

}

func (e *ExplosaoTerra) NewPlayer(audioContext *audio.Context) {
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
	e.Player.audioPlayer = p

	e.Player.audioPlayer.Play()

}

func (e *ExplosaoTerra) SomExplosao() {

	e.NewPlayer(status.AudioContext)
	e.Player.toque = true

}

func (e *ExplosaoTerra) Update(screen *ebiten.Image) error {
	if e.Exists {
		e.CountExpTerra++
		if !e.Player.toque {
			e.SomExplosao()
		}
	}

	return nil
}
