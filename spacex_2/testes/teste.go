package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/vorbis"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	raudio "github.com/hajimehoshi/ebiten/examples/resources/audio"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 240

	sampleRate = 22050
)

var (
	playerBarColor     = color.RGBA{0x80, 0x80, 0x80, 0xff}
	playerCurrentColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
)

type musicType int

const (
	typeOgg musicType = iota
	typeMP3
)

func (t musicType) String() string {
	switch t {
	case typeOgg:
		return "Ogg"
	case typeMP3:
		return "MP3"
	default:
		panic("not reached")
	}
}

// Player represents the current audio state.
type Player struct {
	audioContext *audio.Context
	audioPlayer  *audio.Player
	current      time.Duration
	total        time.Duration
	seBytes      []byte
	seCh         chan []byte
	volume128    int
	musicType    musicType
}

func playerBarRect() (x, y, w, h int) {
	w, h = 300, 4
	x = (screenWidth - w) / 2
	y = screenHeight - h - 16
	return
}

func NewPlayer(audioContext *audio.Context, musicType musicType) (*Player, error) {
	type audioStream interface {
		audio.ReadSeekCloser
		Length() int64
	}

	const bytesPerSample = 4 // TODO: This should be defined in audio package

	var s audioStream

	switch musicType {
	case typeOgg:
		var err error
		s, err = vorbis.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Ragtime_ogg))
		if err != nil {
			return nil, err
		}
	case typeMP3:
		var err error
		s, err = mp3.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Classic_mp3))
		if err != nil {
			return nil, err
		}
	default:
		panic("not reached")
	}
	p, err := audio.NewPlayer(audioContext, s)
	if err != nil {
		return nil, err
	}
	player := &Player{
		audioContext: audioContext,
		audioPlayer:  p,
		total:        time.Second * time.Duration(s.Length()) / bytesPerSample / sampleRate,
		volume128:    128,
		seCh:         make(chan []byte),
		musicType:    musicType,
	}
	if player.total == 0 {
		player.total = 1
	}
	player.audioPlayer.Play()
	go func() {
		s, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Jab_wav))
		if err != nil {
			log.Fatal(err)
			return
		}
		b, err := ioutil.ReadAll(s)
		if err != nil {
			log.Fatal(err)
			return
		}
		player.seCh <- b
	}()
	return player, nil
}

func (p *Player) Close() error {
	return p.audioPlayer.Close()
}

func (p *Player) update() error {
	select {
	case p.seBytes = <-p.seCh:
		close(p.seCh)
		p.seCh = nil
	default:
	}

	if p.audioPlayer.IsPlaying() {
		p.current = p.audioPlayer.Current()
	}
	p.seekBarIfNeeded()
	p.switchPlayStateIfNeeded()
	p.playSEIfNeeded()
	p.updateVolumeIfNeeded()

	if inpututil.IsKeyJustPressed(ebiten.KeyU) {
		b := ebiten.IsRunnableOnUnfocused()
		ebiten.SetRunnableOnUnfocused(!b)
	}
	return nil
}

func (p *Player) playSEIfNeeded() {
	if p.seBytes == nil {
		// Bytes for the SE is not loaded yet.
		return
	}

	if !inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return
	}
	sePlayer, _ := audio.NewPlayerFromBytes(p.audioContext, p.seBytes)
	sePlayer.Play()
}

func (p *Player) updateVolumeIfNeeded() {
	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		p.volume128--
	}
	if ebiten.IsKeyPressed(ebiten.KeyX) {
		p.volume128++
	}
	if p.volume128 < 0 {
		p.volume128 = 0
	}
	if 128 < p.volume128 {
		p.volume128 = 128
	}
	p.audioPlayer.SetVolume(float64(p.volume128) / 128)
}

func (p *Player) switchPlayStateIfNeeded() {
	if !inpututil.IsKeyJustPressed(ebiten.KeyS) {
		return
	}
	if p.audioPlayer.IsPlaying() {
		p.audioPlayer.Pause()
		return
	}
	p.audioPlayer.Play()
}

func (p *Player) seekBarIfNeeded() {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return
	}

	// Calculate the next seeking position from the current cursor position.
	x, y := ebiten.CursorPosition()
	bx, by, bw, bh := playerBarRect()
	const padding = 4
	if y < by-padding || by+bh+padding <= y {
		return
	}
	if x < bx || bx+bw <= x {
		return
	}
	pos := time.Duration(x-bx) * p.total / time.Duration(bw)
	p.current = pos
	p.audioPlayer.Seek(pos)
}

func (p *Player) draw(screen *ebiten.Image) {
	// Draw the bar.
	x, y, w, h := playerBarRect()
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(w), float64(h), playerBarColor)

	// Draw the cursor on the bar.
	c := p.current
	cw, ch := 4, 10
	cx := int(time.Duration(w)*c/p.total) + x - cw/2
	cy := y - (ch-h)/2
	ebitenutil.DrawRect(screen, float64(cx), float64(cy), float64(cw), float64(ch), playerCurrentColor)

	// Compose the curren time text.
	m := (c / time.Minute) % 100
	s := (c / time.Second) % 60
	currentTimeStr := fmt.Sprintf("%02d:%02d", m, s)

	// Draw the debug message.
	msg := fmt.Sprintf(`TPS: %0.2f
Press S to toggle Play/Pause
Press P to play SE
Press Z or X to change volume of the music
Press U to switch the runnable-on-unfocused state
Press A to switch Ogg and MP3
Current Time: %s
Current Volume: %d/128
Type: %s`, ebiten.CurrentTPS(), currentTimeStr, int(p.audioPlayer.Volume()*128), p.musicType)
	ebitenutil.DebugPrint(screen, msg)
}

type Game struct {
	musicPlayer   *Player
	musicPlayerCh chan *Player
	errCh         chan error
}

func NewGame() (*Game, error) {
	audioContext, err := audio.NewContext(sampleRate)
	if err != nil {
		return nil, err
	}

	m, err := NewPlayer(audioContext, typeOgg)
	if err != nil {
		return nil, err
	}

	return &Game{
		musicPlayer:   m,
		musicPlayerCh: make(chan *Player),
		errCh:         make(chan error),
	}, nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	select {
	case p := <-g.musicPlayerCh:
		g.musicPlayer = p
	case err := <-g.errCh:
		return err
	default:
	}

	if g.musicPlayer != nil && inpututil.IsKeyJustPressed(ebiten.KeyA) {
		//var s musicType
		var t musicType
		switch g.musicPlayer.musicType {
		case typeOgg:
			t = typeMP3
		case typeMP3:
			t = typeOgg
		default:
			panic("not reached")
		}
		fmt.Println(t)

		g.musicPlayer.Close()
		g.musicPlayer = nil

		go func() {
			p, err := NewPlayer(audio.CurrentContext(), typeMP3)
			if err != nil {
				g.errCh <- err
				return
			}
			g.musicPlayerCh <- p
		}()
	}

	if g.musicPlayer != nil {
		if err := g.musicPlayer.update(); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.musicPlayer != nil {
		g.musicPlayer.draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Audio (Ebiten Demo)")
	g, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
