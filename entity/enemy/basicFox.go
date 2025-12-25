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
)

type BasicFox struct {
	Parent         Enemy
	Speed          float64
	RotationSpeed  float64
	mode           basicFoxMode
	colliderRadius float64
}

func (bs *BasicFox) Init(x, y float64, img *ebiten.Image) {

	bs.colliderRadius = 30

	bs.Parent.collider = collisionsystem.Shape{}

	bs.mode = basicFoxMode(chase)
	bs.Speed = 5
	sprite := assets.Sprite{Position: gmath.Vec{X: x, Y: y}}
	sprite.SetOriginAtCenter()
	sprite.SetImage(img)
	sprite.SetScale(0.3)
	bs.Parent.Init(x, y, sprite)
}

func (bs *BasicFox) Destroy() {
	bs.Parent.Center = gmath.Vec{X: 500, Y: 500}
	bs.mode = basicFoxMode(chase)
}

func (bs *BasicFox) Move(player *player.Player) {

	bs.Parent.sprite.Rotation += gmath.Rad(bs.RotationSpeed)

	playerRight := player.GetPosition().Add(player.GetDirection().Rotated(80).Mulf(200))
	playerLeft := player.GetPosition().Add(player.GetDirection().Rotated(-80).Mulf(200))

	if bs.Parent.Center.DistanceTo(playerLeft) < bs.Speed || bs.Parent.Center.DistanceTo(playerRight) < bs.Speed {
		bs.mode = basicFoxMode(pursue)
	}

	if bs.mode == basicFoxMode(chase) {
		if bs.Parent.Center.DistanceTo(playerLeft) < bs.Parent.Center.DistanceTo(playerRight) {
			bs.Parent.MoveTowards(playerLeft, bs.Speed)
		} else {
			bs.Parent.MoveTowards(playerRight, bs.Speed)
		}
	} else if bs.mode == basicFoxMode(pursue) {
		if bs.Parent.Center.DistanceTo(player.GetPosition()) < bs.Speed {
			player.Reset()
			bs.Destroy()
		}
		bs.Parent.MoveTowards(player.GetPosition(), bs.Speed*2)
	}
	if bs.Parent.CheckCollisionsWithTail(player.GetTail()) {
		bs.Destroy()
	}

	bs.Parent.collider.Points = []gmath.Vec{gmath.Vec{X: bs.Parent.Center.X + bs.colliderRadius, Y: bs.Parent.Center.Y + bs.colliderRadius},
		gmath.Vec{X: bs.Parent.Center.X - bs.colliderRadius, Y: bs.Parent.Center.Y + bs.colliderRadius},
		gmath.Vec{X: bs.Parent.Center.X - bs.colliderRadius, Y: bs.Parent.Center.Y - bs.colliderRadius},
		gmath.Vec{X: bs.Parent.Center.X + bs.colliderRadius, Y: bs.Parent.Center.Y - bs.colliderRadius}}
}
