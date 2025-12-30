package enemy

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/entity"
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
	Velocity       float64
	RotationSpeed  float64
	mode           basicFoxMode
	colliderRadius float64
	isDestroyed    bool
}

func (bs *BasicFox) Init(x, y float64, img *ebiten.Image) {
	bs.Speed = 5
	bs.Velocity = bs.Speed
	bs.isDestroyed = false
	bs.RotationSpeed = 0.01
	bs.colliderRadius = 31

	bs.mode = basicFoxMode(chase)

	sprite := assets.Sprite{Position: gmath.Vec{X: x, Y: y}}
	sprite.SetOriginAtCenter()
	sprite.SetImage(img)
	sprite.SetScale(0.3)
	bs.initBase(x, y, sprite)
}

func (bs *BasicFox) SetSpeed(speed float64) {
	bs.Speed = speed
}

func (bs *BasicFox) GetSpeed() float64 {
	return bs.Speed
}

func (bs *BasicFox) SetVelocity(speed float64) {
	bs.Velocity = speed
}

func (bs *BasicFox) PassBy(bs2 entity.Chaser) {
	if math.Abs(bs.GetVelocity()-bs2.GetVelocity()) < 0.01 {
		bs.Velocity = bs.Velocity / 2
	}
}

func (bs *BasicFox) GetVelocity() float64 {
	return bs.Velocity
}

func (bs *BasicFox) IsDestroyed() bool {
	return bs.isDestroyed
}

func (bs *BasicFox) Destroy() {
	bs.isDestroyed = true
}

func (bs *BasicFox) Update(player *player.Player) {
	if bs.Velocity < bs.Speed {
		bs.Velocity *= 1.001
	}
	bs.sprite.Rotation += gmath.Rad(bs.RotationSpeed)

	playerRight := player.GetPosition().Add(player.GetDirection().Rotated(80).Mulf(200))
	playerLeft := player.GetPosition().Add(player.GetDirection().Rotated(-80).Mulf(200))

	if bs.Center.DistanceTo(playerLeft) < bs.Speed || bs.Center.DistanceTo(playerRight) < bs.Speed {
		bs.mode = basicFoxMode(pursue)
	}

	if bs.mode == basicFoxMode(chase) {
		if bs.Center.DistanceTo(playerLeft) < bs.Center.DistanceTo(playerRight) {
			bs.MoveTowards(playerLeft, bs.Velocity)
		} else {
			bs.MoveTowards(playerRight, bs.Velocity)
		}
	} else if bs.mode == basicFoxMode(pursue) {
		if collisionsystem.CheckShapesCollision(bs.collider, *player.GetCollider()) {
			player.Reset()
			bs.Destroy()
		}
		bs.MoveTowards(player.GetPosition(), bs.Velocity*2)

	} else if bs.mode == basicFoxMode(corpse) {
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
