package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"
)

type Enemy interface {
	Init(x, y float64, img *ebiten.Image)
	DrawEntity(screen *ebiten.Image) error
	Update(player *player.Player)
	IsDestroyed() bool
}
