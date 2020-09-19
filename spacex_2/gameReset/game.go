package gameReset

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/riqueemn/jogos/spacex_2/buttons"
	"github.com/riqueemn/jogos/spacex_2/objects"
	"github.com/riqueemn/jogos/spacex_2/status"
	"github.com/riqueemn/jogos/spacex_2/window"
)

var (
	Nave                 objects.Nave
	Fire                 Fires
	Bomba                Bombas
	ExplosaoAr_Sprite    ExplosoesAr
	ExplosaoTerra_Sprite ExplosoesTerra
	fireImage            = objects.FireImage
	Tiros                Fires
	bombaImage           = objects.BombaImage
	BombasC              Bombas
	PlanetaLife          = 20
	Insta                = 0
	Barra                objects.BarraPlaneta
	BarraT               objects.BarraInsta
	Pontos               = 0
	Mode                 string
	QtdTiros             int
	Reset                bool
)

const (
	MinSprites = 0
	MaxSprites = 50000
)

type Sprites_Fire struct {
	Sprites []*objects.Fire
	Num     int
}

type Sprites_Bomba struct {
	Sprites []*objects.Bomba
	Num     int
}

type Sprites_ExplosaoAr struct {
	Sprites []*objects.ExplosaoAr
	Num     int
}
type Sprites_ExplosaoTerra struct {
	Sprites []*objects.ExplosaoTerra
	Num     int
}

type ExplosoesAr struct {
	Sprites Sprites_ExplosaoAr
	op      ebiten.DrawImageOptions
	inited  bool
}

type ExplosoesTerra struct {
	Sprites Sprites_ExplosaoTerra
	op      ebiten.DrawImageOptions
	inited  bool
}
type Fires struct {
	Sprites Sprites_Fire
	op      ebiten.DrawImageOptions
	inited  bool
}

type Bombas struct {
	Sprites Sprites_Bomba
	op      ebiten.DrawImageOptions
	inited  bool
}

func (s *Sprites_Fire) Update() {
	for i := 0; i < 50; i++ {
		s.Sprites[i].Update()
		colisaoTiroBomba(s.Sprites[i])
		colisaoTiroBomba(s.Sprites[i])
		if s.Sprites[i].Y < 0 {
			s.Sprites[i].Exists = false
			s.Sprites[i].X = 0
			s.Sprites[i].Y = 0
		}
	}
}

func init() {
	Reset = false
	Barra.X = 0
	Barra.Y = 20
	Barra.Width = 0.4
	Barra.Height = 0.4
	Barra.Life = PlanetaLife

	BarraT.X = 0
	BarraT.Y = 40
	BarraT.Width = 0.4
	BarraT.Height = 0.4
	BarraT.Life = Insta

	Nave.FireAllowd = true

}

func (f *Fires) init() {

	defer func() {
		f.inited = true
	}()

	f.Sprites.Sprites = make([]*objects.Fire, MaxSprites)
	f.Sprites.Num = 0
	for i := range f.Sprites.Sprites {
		w, h := 0.5, 0.5
		f.Sprites.Sprites[i] = &objects.Fire{
			Width:  w,
			Height: h,
			X:      0,
			Y:      0,
		}
	}
}

func (f *Fires) Update(screen *ebiten.Image) error {
	if !f.inited {
		f.init()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		f.Sprites.Num++
		if MaxSprites < f.Sprites.Num {
			f.Sprites.Num = MaxSprites
		}
	}

	f.Sprites.Update()
	return nil
}

func (f *Fires) Draw(screen *ebiten.Image) {

	//w, h := EbitenImage.Size()
	for i := 0; i < 50; i++ {
		if f.Sprites.Sprites[i].Exists {
			s := f.Sprites.Sprites[i]
			f.op.GeoM.Reset()
			f.op.GeoM.Scale(s.Width, s.Height)
			//f.op.GeoM.Translate((-float64(37)/2)+7, (-float64(70)/2)-2)
			//f.op.GeoM.Translate(window.ScreenWidth/2, window.ScreenHeight/2)
			f.op.GeoM.Translate(float64(s.X), float64(s.Y))
			screen.DrawImage(fireImage, &f.op)
			//fmt.Println(f.Sprites.Sprites[0].X, f.Sprites.Sprites[0].Y)
		}

	}
	/*
		msg := fmt.Sprintf(`TPS: %0.2f
			   FPS: %0.2f
			   Num of sprites: %d
			   Press <- or -> to change the number of sprites`, ebiten.CurrentTPS(), ebiten.CurrentFPS(), f.Sprites.Num)
		ebitenutil.DebugPrint(screen, msg)
	*/

}

func (s *Sprites_Bomba) Update() {
	for i := 0; i < 50; i++ {
		s.Sprites[i].Update()

		if s.Sprites[i].Y > 320 {
			PlanetaLife -= 10
			s.Sprites[i].Exists = false
			bombaOffsetLeft := s.Sprites[i].X

			ExplosaoTerra_Sprite.Sprites.Sprites[i].X = bombaOffsetLeft - 20
			ExplosaoTerra_Sprite.Sprites.Sprites[i].Y = 320 - 0.5*94
			ExplosaoTerra_Sprite.Sprites.Sprites[i].Exists = true
			ExplosaoTerra_Sprite.Sprites.Num++

			s.Sprites[i].X = 0
			s.Sprites[i].Y = 0
		}

	}
}

func (f *Bombas) init() {
	go f.cont()
	defer func() {
		f.inited = true
	}()

	f.Sprites.Sprites = make([]*objects.Bomba, MaxSprites)
	f.Sprites.Num = 0
	for i := range f.Sprites.Sprites {
		w, h := 0.5, 0.5
		f.Sprites.Sprites[i] = &objects.Bomba{
			Width:  w,
			Height: h,
			X:      0,
			Y:      0,
			Exists: false,
		}
	}
}

func (f *Bombas) cont() {
	if !Reset {
		ticker := time.NewTicker(time.Second)
		//defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if !Reset {
					f.Sprites.Num++
					updateBomba()
				}
			}
		}
	}
}

func (f *Bombas) Update(screen *ebiten.Image) error {
	if !f.inited {
		f.init()
	}

	if f.Sprites.Num < MinSprites {
		f.Sprites.Num = MinSprites
	}
	f.Sprites.Update()
	return nil
}

func (f *Bombas) Draw(screen *ebiten.Image) {

	for i := 0; i < 50; i++ {
		if f.Sprites.Sprites[i].Exists {
			s := f.Sprites.Sprites[i]
			f.op.GeoM.Reset()
			f.op.GeoM.Scale(s.Width, s.Height)
			f.op.GeoM.Translate(float64(s.X), float64(s.Y))
			screen.DrawImage(bombaImage, &f.op)
		}
	}
	/*
		msg := fmt.Sprintf(`TPS: %0.2f
						FPS: %0.2f
						Num of sprites: %d
						Press <- or -> to change the number of sprites`, ebiten.CurrentTPS(), ebiten.CurrentFPS(), f.Sprites.Num)
		ebitenutil.DebugPrint(screen, msg)
	*/
}

//****EXPLOSÕES****

func (e *Sprites_ExplosaoAr) Update(screen *ebiten.Image) {
	for i := 0; i < 1000; i++ {
		if e.Sprites[i].Exists {
			e.Sprites[i].Update(screen)
		}
	}
	if e.Num >= 40 {
		e.Num = 0
	}
}

func (e *ExplosoesAr) init() {
	//go e.cont()
	defer func() {
		e.inited = true
	}()

	e.Sprites.Sprites = make([]*objects.ExplosaoAr, MaxSprites)
	e.Sprites.Num = 0
	for i := range e.Sprites.Sprites {
		w, h := 0.5, 0.5
		e.Sprites.Sprites[i] = &objects.ExplosaoAr{
			Width:      w,
			Height:     h,
			Exists:     false,
			CountExpAr: 0,
		}
	}
}

func (e *ExplosoesAr) Update(screen *ebiten.Image) error {
	if !e.inited {
		e.init()
	}

	if e.Sprites.Num < MinSprites {
		e.Sprites.Num = MinSprites
	}
	e.Sprites.Update(screen)
	return nil
}

func (e *ExplosoesAr) Draw(screen *ebiten.Image) {

	for i := 0; i < 1000; i++ {
		if e.Sprites.Sprites[i].Exists {
			//s := e.Sprites.Sprites[i]
			//e.op.GeoM.Reset()
			//e.op.GeoM.Scale(s.Width, s.Height)
			//e.op.GeoM.Translate(float64(s.X), float64(s.Y))
			//screen.DrawImage(explosaoImage, &e.op)
			e.Sprites.Sprites[i].Draw(screen)
		}
	}
	/*
		msg := fmt.Sprintf(`TPS: %0.2f
					FPS: %0.2f
					Num of sprites: %d
					Press <- or -> to change the number of sprites`, ebiten.CurrentTPS(), ebiten.CurrentFPS(), e.Sprites.Num)
		ebitenutil.DebugPrint(screen, msg)
	*/
}

func (e *Sprites_ExplosaoTerra) Update(screen *ebiten.Image) {
	for i := 0; i < 1000; i++ {
		if e.Sprites[i].Exists {
			e.Sprites[i].Update(screen)
		}
	}
	if e.Num >= 40 {
		e.Num = 0
	}
}

func (e *ExplosoesTerra) init() {
	defer func() {
		e.inited = true
	}()

	e.Sprites.Sprites = make([]*objects.ExplosaoTerra, MaxSprites)
	e.Sprites.Num = 0
	for i := range e.Sprites.Sprites {
		w, h := 0.5, 0.5
		e.Sprites.Sprites[i] = &objects.ExplosaoTerra{
			Width:         w,
			Height:        h,
			Exists:        false,
			CountExpTerra: 0,
		}
	}
}

func (e *ExplosoesTerra) Update(screen *ebiten.Image) error {
	if !e.inited {
		e.init()
	}

	if e.Sprites.Num < MinSprites {
		e.Sprites.Num = MinSprites
	}
	e.Sprites.Update(screen)
	return nil
}

func (e *ExplosoesTerra) Draw(screen *ebiten.Image) {

	for i := 0; i < 1000; i++ {
		if e.Sprites.Sprites[i].Exists {
			e.Sprites.Sprites[i].Draw(screen)
		}
	}

	/*
		msg := fmt.Sprintf(`TPS: %0.2f
					FPS: %0.2f
					Num of sprites: %d
					Press <- or -> to change the number of sprites`, ebiten.CurrentTPS(), ebiten.CurrentFPS(), e.Sprites.Num)
		ebitenutil.DebugPrint(screen, msg)
	*/

}

//****FIM EXPLOSÕES****

func naveCustom() {
	Nave.Width = 0.2
	Nave.Height = 0.2
}

func teclaDw() {

	if buttons.Up { //Cima
		Nave.DiryJ = -1
	} else if buttons.Down { //Baixo
		Nave.DiryJ = 1
	}
	if buttons.Left { //Esquerda
		Nave.DirxJ = -1
	} else if buttons.Right { //Direita
		Nave.DirxJ = 1
	}
	if buttons.Space { //Tiro
		Nave.IsFire = true
	}

	if buttons.Up && buttons.Down {
		Nave.DiryJ = 0
		Nave.DirxJ = 0
	}
	if buttons.Left && buttons.Right {
		Nave.DirxJ = 0
		Nave.DiryJ = 0
	}

}

func teclaUp() {

	if !buttons.Up || !buttons.Down {
		Nave.DiryJ = 0
	}
	if !buttons.Left || !buttons.Right {
		Nave.DirxJ = 0
	}

	if !inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		Nave.IsFire = false
	}

}

func controleNave() {
	teclaUp()
	teclaDw()
	if Nave.Pjx >= 160 && Nave.DirxJ == 1 {
		Nave.DirxJ = 0
	} else if Nave.Pjx <= -145 && Nave.DirxJ == -1 {
		Nave.DirxJ = 0
	} else if Nave.Pjy >= 165 && Nave.DiryJ == 1 {
		Nave.DiryJ = 0
	} else if Nave.Pjy <= -130 && Nave.DiryJ == -1 {
		Nave.DiryJ = 0
	}

	Nave.Vel = 3
	//go contVelocidade()

	Nave.Pjx += Nave.DirxJ * float64(Nave.Vel)
	Nave.Pjy += Nave.DiryJ * float64(Nave.Vel)

}

/*
func contVelocidade() {

	if Nave.Lentidao {
		ticker := time.NewTicker(time.Second)
		Nave.Vel = 2
		for {
			select {
			case <-ticker.C:
				ticker.Stop()
				Nave.Lentidao = false
				Nave.Vel = 2
			}
		}
	} else {
		Nave.Vel = 2
	}
}
*/
var a bool

func instaBarra(screen *ebiten.Image) {
	a = Nave.IsFire

	if BarraT.Life > 100 {
		Nave.IsFire = false
		Nave.FireAllowd = false
		a = true
	}
	if !Nave.FireAllowd {
		Nave.IsFire = false
	}

	if Nave.IsFire && Nave.FireAllowd {
		QtdTiros++
	}
	if a {
		go contInsta()
	}

	//fmt.Println("Nave.IsFire:", Nave.IsFire)
	//fmt.Println("Nave.FireAllowd:", Nave.FireAllowd)

	Insta = QtdTiros * 10
	BarraT.Life = Insta
	BarraT.Update(screen)
}

func contInsta() {
	if a {
		ticker2 := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker2.C:
				go tempoBomba()
				ticker2.Stop()
			}
		}
	}
}

func tempoBomba() {
	if !Nave.IsFire {
		ticker3 := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker3.C:
				if QtdTiros > 0 {
					QtdTiros--
				} else {
					Nave.FireAllowd = true
				}
				fmt.Println(QtdTiros)
				ticker3.Stop()
			}
		}
	}
}

func updateFire(screen *ebiten.Image) {

	if Fire.Sprites.Num >= 49 {
		Fire.Sprites.Num = 0
	}

	i := Fire.Sprites.Num
	if Nave.IsFire {

		Fire.Sprites.Sprites[i].X = Nave.Pjx + 149
		Fire.Sprites.Sprites[i].Y = Nave.Pjy + 120
		Fire.Sprites.Sprites[i].Vel = 4
		Fire.Sprites.Sprites[i].Exists = true
	}

}

func updateBomba() {
	rand.Seed(time.Now().UnixNano())

	if Bomba.Sprites.Num >= 49 {
		Bomba.Sprites.Num = 0
	}

	i := Bomba.Sprites.Num
	if i > 0 {
	}
	Bomba.Sprites.Sprites[i].Exists = true

	Bomba.Sprites.Sprites[i].X = rand.Float64() * float64(window.ScreenWidth*0.95)
	Bomba.Sprites.Sprites[i].Y = 0
	Bomba.Sprites.Sprites[i].Vel = 3

}

func colisaoTiroBomba(tiro *objects.Fire) {
	//var k = len(Fire.Sprites.Sprites)
	for i := 0; i < 50; i++ {
		if Bomba.Sprites.Sprites[i].Exists && tiro.Exists {
			tiroOffsetTop := tiro.Y
			tiroOffsetLeft := tiro.X
			bombaOffsetTop := Bomba.Sprites.Sprites[i].Y
			bombaOffsetLeft := Bomba.Sprites.Sprites[i].X

			//if (((tiroOffsetTop) <= (bombaOffsetTop + 41*0.5)) && ((tiroOffsetTop + 10*0.5) >= (bombaOffsetTop))) && (((tiroOffsetLeft) <= (bombaOffsetLeft + 50*0.5)) && ((tiroOffsetLeft + 10*0.5) >= (bombaOffsetLeft))) {
			if (((tiroOffsetTop) <= (bombaOffsetTop + 41*0.5)) && ((tiroOffsetTop + 10*0.5) >= (bombaOffsetTop))) && (((tiroOffsetLeft) <= (bombaOffsetLeft + 50*0.5)) && ((tiroOffsetLeft + 10*0.5) >= (bombaOffsetLeft))) {

				Bomba.Sprites.Sprites[i].Exists = false
				tiro.Exists = false

				ExplosaoAr_Sprite.Sprites.Sprites[i].X = bombaOffsetLeft - 20
				ExplosaoAr_Sprite.Sprites.Sprites[i].Y = tiroOffsetTop - 20
				ExplosaoAr_Sprite.Sprites.Sprites[i].Exists = true
				ExplosaoAr_Sprite.Sprites.Num++
				status.Pontos++

			}

		}
	}
}

type Jogo struct {
	Mode string
}

func (j *Jogo) Update(screen *ebiten.Image) error {
	Bomba.Update(screen)
	ExplosaoAr_Sprite.Update(screen)
	ExplosaoTerra_Sprite.Update(screen)
	Fire.Update(screen)
	Nave.Update(screen)
	updateFire(screen)
	controleNave()
	instaBarra(screen)
	Barra.Update(screen)
	Barra.Life = PlanetaLife
	if j.Mode == "Historia" {
	} else if j.Mode == "Arcade" {

	}

	if PlanetaLife <= 0 && j.Mode == "Arcade" {
		status.Play = false
		status.StatusGame = "DerrotaA"
		PlanetaLife = 20
		//j.Reset()
		//status.Pontos = Pontos
	} else if status.Pontos >= 30 && j.Mode == "Historia" {
		status.Play = false
		status.StatusGame = "Vitoria"
		//j.Reset()
	} else if PlanetaLife <= 0 && j.Mode == "Historia" {
		status.Play = false
		status.StatusGame = "DerrotaH"
		//j.Reset()
		//status.Pontos = Pontos
	}

	return nil
}

func (j *Jogo) Draw(screen *ebiten.Image) {
	naveCustom()
	Nave.Draw(screen)
	Fire.Draw(screen)
	Bomba.Draw(screen)
	ExplosaoAr_Sprite.Draw(screen)
	ExplosaoTerra_Sprite.Draw(screen)
	Barra.Draw(screen)
	BarraT.Draw(screen)
	if j.Mode == "Historia" {
		msg := fmt.Sprintf("Planeta")
		ebitenutil.DebugPrint(screen, msg)
	} else if j.Mode == "Arcade" {
		msg := fmt.Sprintf(`Pontos: %d`, status.Pontos)
		ebitenutil.DebugPrint(screen, msg)

	}

}

func (j *Jogo) Reset() {
	Bomba.init()
	Reset = true
	Pontos = 0
	PlanetaLife = 20
	Insta = 0

	var a objects.Nave
	var b Fires
	//var c Bombas

	Nave = a
	Fire = b
	//Bomba.init()
	//QtdTiros = 0
}
