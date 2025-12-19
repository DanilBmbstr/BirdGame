package optionsLoaders

import (
	"encoding/json"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	resource "github.com/quasilyte/ebitengine-resource"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type PlayerDTO struct {
	Lenght            float64
	TailColor         [4]float32
	TailLenght        int
	TailSegmentLenght float64
	PlayerColor       [4]float32
	Speed             float64
	RotationSpeed     float64

	TailMovingSpeed   float64
	TailRotationSpeed float64

	//Константы

	TailStartSegmentsLenght int
	TailStartWidth          float64
	TailWidth               float64
	TailWidthNarrow         float64

	SpriteScale  float64
	SpriteOrigin [2]float64
}

func LoadPlayer(l *resource.Loader) PlayerDTO {
	dto := PlayerDTO{}
	dat := l.LoadRaw(assets.PlayerSettings)
	err := json.Unmarshal(dat.Data, &dto)
	check(err)

	tailColor := ebiten.ColorScale{}
	tailColor.Scale(dto.TailColor[0], dto.TailColor[1], dto.TailColor[2], dto.TailColor[3])

	playerColor := ebiten.ColorScale{}
	playerColor.Scale(dto.PlayerColor[0], dto.PlayerColor[1], dto.PlayerColor[2], dto.PlayerColor[3])

	return dto
}
