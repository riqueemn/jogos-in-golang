package buttons

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var (
	Up    bool
	Down  bool
	Right bool
	Left  bool
	Space bool
	Enter bool
)

/*
type Game struct {
	pressed []ebiten.Key
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.pressed = nil
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			g.pressed = append(g.pressed, k)
		}
	}
	return nil
}
*/

func teclaDw() {

	if ebiten.IsKeyPressed(ebiten.KeyW) { //Cima
		Up = true
	} else {
		Up = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) { //Baixo
		Down = true
	} else {
		Down = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) { //Esquerda
		Left = true
	} else {
		Left = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) { //Direita
		Right = true
	} else {
		Right = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) { //Tiro
		Space = true
	} else {
		Space = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) { //Enter
		Enter = true
	} else {
		Enter = false
	}

}

func Update() error {
	teclaDw()

	return nil
}
