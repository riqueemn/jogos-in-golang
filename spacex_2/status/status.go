package status

import (
	"context"

	"github.com/hajimehoshi/ebiten/audio"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Play       bool = false
	StatusGame string
	Pontos     int
	Player     string
	Db         *mongo.Database
	Ctx        context.Context

	AudioContext *audio.Context
)

const (
	SampleRate = 22050
)
