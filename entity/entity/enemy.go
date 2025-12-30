package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"
	"github.com/quasilyte/gmath"
)

type Enemy interface {
	Init(x, y float64, img *ebiten.Image)
	DrawEntity(screen *ebiten.Image) error
	Update(player *player.Player)
	IsDestroyed() bool
	GetPosition() gmath.Vec
	GetCollider() *collisionsystem.Shape
}

type Chaser interface {
	Enemy
	GetSpeed() float64
	SetSpeed(float64)
	GetVelocity() float64
	SetVelocity(float64)
	PassBy(Chaser)
}
