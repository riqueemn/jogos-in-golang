package objects

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/riqueemn/jogos/spacex_2/window"
)

const (
	frameOX     = 13
	frameOY     = 0
	frameWidth  = 37
	frameHeight = 70
	frameNum    = 3
)

var (
	naveImage *ebiten.Image
	count     int
)

type Nave struct {
	Pjx        float64
	Pjy        float64
	DiryJ      float64
	DirxJ      float64
	Width      float64
	Height     float64
	IsFire     bool
	Player     string
	Vel        int
	Lentidao   bool
	FireAllowd bool
}

func init() {

	reader, err := os.Open("C:/Users/Henrique/go/src/github.com/riqueemn/jogos/spacex_2/imagens/nave_png.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := origEbitenImage.Size()
	naveImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.8)
	naveImage.DrawImage(origEbitenImage, op)
}

func (n *Nave) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Reset()
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(window.ScreenWidth/2, window.ScreenHeight/2)
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(n.Pjx, n.Pjy)
	i := (count / 6) % frameNum
	sx, sy := frameOX+i*frameWidth+i*27, frameOY
	//op.GeoM.Scale(n.Width, n.Height)

	screen.DrawImage(naveImage.SubImage(image.Rect(sx, frameOY, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

	//10, 0, 50, 70
	//75, 0, 115, 70
	//140, 0, 185, 70
}

func (n *Nave) Update(screen *ebiten.Image) error {
	count++

	return nil
}

/*
	func (n *Nave) Executar(screen *ebiten.Image) {
		n.MoveNave()
		n.DrawNave(screen)
		n.Count++
	}
*/
