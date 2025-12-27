package enemy

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"
	"github.com/quasilyte/gmath"
)

type basicFoxMode int

const (
	chase basicFoxMode = iota
	pursue
	corpse
)

type BasicFox struct {
	EnemyBasic
	Speed          float64
	RotationSpeed  float64
	mode           basicFoxMode
	colliderRadius float64
	isDestroyed    bool
}

func (bs *BasicFox) Init(x, y float64, img *ebiten.Image) {
	bs.isDestroyed = false
	bs.RotationSpeed = 0.01
	bs.colliderRadius = 31

	bs.mode = basicFoxMode(chase)
	bs.Speed = 5
	sprite := assets.Sprite{Position: gmath.Vec{X: x, Y: y}}
	sprite.SetOriginAtCenter()
	sprite.SetImage(img)
	sprite.SetScale(0.3)
	bs.initBase(x, y, sprite)
}

func (bs *BasicFox) IsDestroyed() bool {
	return bs.isDestroyed
}

func (bs *BasicFox) Destroy() {
	bs.isDestroyed = true
}

func (bs *BasicFox) Update(player *player.Player) {

	bs.sprite.Rotation += gmath.Rad(bs.RotationSpeed)

	playerRight := player.GetPosition().Add(player.GetDirection().Rotated(80).Mulf(200))
	playerLeft := player.GetPosition().Add(player.GetDirection().Rotated(-80).Mulf(200))

	if bs.Center.DistanceTo(playerLeft) < bs.Speed || bs.Center.DistanceTo(playerRight) < bs.Speed {
		bs.mode = basicFoxMode(pursue)
	}

	if bs.mode == basicFoxMode(chase) {
		if bs.Center.DistanceTo(playerLeft) < bs.Center.DistanceTo(playerRight) {
			bs.MoveTowards(playerLeft, bs.Speed)
		} else {
			bs.MoveTowards(playerRight, bs.Speed)
		}
	} else if bs.mode == basicFoxMode(pursue) {
		if collisionsystem.CheckShapeCollisionWithPoint(player.GetPosition(), bs.collider) {
			player.Reset()
			bs.Destroy()
		}
		bs.MoveTowards(player.GetPosition(), bs.Speed*2)
	} else if bs.mode == corpse {
		bs.sprite.SetScale(bs.sprite.GetScale() * 0.9)
		if bs.sprite.GetScale() < 0.01 {
			bs.sprite.SetScale(0.3)
			bs.Destroy()
		}
	}
	if bs.CheckCollisionsWithTail(player.GetTail()) && player.BattleSystem.HitTime > 1 {
		bs.mode = corpse
	}
	coef := 1.5
	bs.collider.Points = []gmath.Vec{

		gmath.Vec{X: bs.Center.X + bs.colliderRadius, Y: bs.Center.Y + bs.colliderRadius},
		gmath.Vec{X: bs.Center.X + 0, Y: bs.Center.Y + bs.colliderRadius*coef},
		gmath.Vec{X: bs.Center.X - bs.colliderRadius, Y: bs.Center.Y + bs.colliderRadius},
		gmath.Vec{X: bs.Center.X - bs.colliderRadius*coef, Y: bs.Center.Y + 0},
		gmath.Vec{X: bs.Center.X - bs.colliderRadius, Y: bs.Center.Y - bs.colliderRadius},
		gmath.Vec{X: bs.Center.X - 0, Y: bs.Center.Y - bs.colliderRadius*coef},
		gmath.Vec{X: bs.Center.X + bs.colliderRadius, Y: bs.Center.Y - bs.colliderRadius},
		gmath.Vec{X: bs.Center.X + bs.colliderRadius*coef, Y: bs.Center.Y - 0}}

}
