package enemy

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/entity"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"

	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/level"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

type Ameba struct {
	EnemyBasic
	Speed          float64
	RotationSpeed  float64
	direction      gmath.Vec
	target         gmath.Vec
	colliderRadius float64

	steps int

	corpseMode bool

	postmp  gmath.Vec
	targtmp gmath.Vec
}

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

func (a *Ameba) Init(x, y float64, loader *resource.Loader) {
	a.corpseMode = false
	a.steps = 0
	a.Speed = 0.15
	a.direction = gmath.Vec{X: 1, Y: 1}
	a.isDestroyed = false
	a.RotationSpeed = 0.01
	a.colliderRadius = 31

	sprite := assets.Sprite{Position: gmath.Vec{X: x, Y: y}}
	sprite.SetOriginAtCenter()
	sprite.SetImage(loader.LoadImage(assets.Ameba).Data)
	sprite.SetScale(0.3)
	a.initBase(x, y, sprite)
	a.target = a.GetPosition()
}

func (a *Ameba) die() {
	a.sprite.SetScale(a.sprite.GetScale() * 0.9)
	if a.sprite.GetScale() < 0.01 {

		a.isDestroyed = true
	}
}

func (a *Ameba) Update(player *player.Player, location *level.CollisionMask) {
	if a.corpseMode {
		a.die()
		return
	}
	dir := gmath.Vec{}
	if a.target.DistanceTo(a.GetPosition()) < a.Speed*3 {
		dir = gmath.Vec{X: float64(r.Intn(200)) - 100, Y: float64(r.Intn(200)) - 100}.Normalized()
		a.targtmp = a.target
		a.target = a.target.Add(dir.Mulf(float64(r.Intn(250) + 50)))

	}
	a.postmp = a.GetPosition()
	a.SetPosition((a.GetPosition().Add(a.GetPosition().DirectionTo(a.target))).X, a.GetPosition().Add(a.GetPosition().DirectionTo(a.target)).Y)

	a.updateCollider()
	if location.CheckCollisions(a.collider.Points[1]) ||

		location.CheckCollisions(a.collider.Points[5]) {

		a.SetPosition(a.postmp.X, a.postmp.Y)
		a.target = a.targtmp
	}
	if collisionsystem.CheckShapesCollision(a.collider, *player.GetCollider()) {
		player.Reset()
	}
	if player.BattleSystem.HitTime > 1 {
		if a.CheckCollisionsWithTail(player.GetTail()) {
			a.corpseMode = true
		}
	}
}

func (a *Ameba) updateCollider() {
	coef := 1.5
	a.collider.Points = []gmath.Vec{

		gmath.Vec{X: a.Center.X + a.colliderRadius, Y: a.Center.Y + a.colliderRadius + 10},
		gmath.Vec{X: a.Center.X - 10, Y: a.Center.Y + a.colliderRadius*(coef+0.4)},
		gmath.Vec{X: a.Center.X - a.colliderRadius, Y: a.Center.Y + a.colliderRadius + 10},
		gmath.Vec{X: a.Center.X - a.colliderRadius*coef, Y: a.Center.Y + 0},
		gmath.Vec{X: a.Center.X - a.colliderRadius, Y: a.Center.Y - a.colliderRadius - 10},
		gmath.Vec{X: a.Center.X + 10, Y: a.Center.Y - a.colliderRadius*(coef+0.4)},
		gmath.Vec{X: a.Center.X + a.colliderRadius, Y: a.Center.Y - a.colliderRadius - 10},
		gmath.Vec{X: a.Center.X + a.colliderRadius*coef, Y: a.Center.Y - 0}}
}

func (a *Ameba) SetSpeed(speed float64) {
	a.Speed = speed
}

func (a *Ameba) GetSpeed() float64 {
	return a.Speed
}

func (a *Ameba) SetVelocity(speed float64) {
	//a.Velocity = speed
}

func (a *Ameba) PassBy(bs2 entity.Enemy) {

}

func (a *Ameba) GetVelocity() float64 {
	//return bs.Velocity
	return 0
}
