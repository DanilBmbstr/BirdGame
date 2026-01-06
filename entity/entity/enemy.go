package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/level"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

type Enemy interface {
	Init(x, y float64, loader *resource.Loader)
	DrawEntity(screen *ebiten.Image) error
	Update(player *player.Player, location *level.CollisionMask)
	IsDestroyed() bool
	GetPosition() gmath.Vec
	GetCollider() *collisionsystem.Shape

	GetSpeed() float64
	SetSpeed(float64)
	GetVelocity() float64
	SetVelocity(float64)
	PassBy(Enemy)
}
