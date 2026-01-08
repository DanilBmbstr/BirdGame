package enemy

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/entity"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/level"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

type Frog struct {
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

func (f *Frog) Init(x, y float64, loader *resource.Loader) {

	f.corpseMode = false
	f.steps = 0
	f.Speed = 0.15
	f.direction = gmath.Vec{X: 1, Y: 1}
	f.isDestroyed = false
	f.RotationSpeed = 0.01
	f.colliderRadius = 31

	sprite := assets.Sprite{Position: gmath.Vec{X: x, Y: y}}
	sprite.Rect = image.Rectangle{image.Point{0, 0}, image.Point{300, 342}}
	sprite.SetOriginAtCenter()
	sprite.SetImage(loader.LoadImage(assets.Frog).Data)
	sprite.Anim.AnimationRate = 4
	sprite.Anim.Loop = false
	sprite.SetScale(0.3)

	f.initBase(x, y, sprite)
	f.sprite.SetFrame(-1)

	f.target = f.GetPosition()

}

func (f *Frog) die() {
	f.sprite.SetScale(f.sprite.GetScale() * 0.9)
	if f.sprite.GetScale() < 0.01 {

		f.isDestroyed = true
	}
}

func (f *Frog) Update(player *player.Player, location *level.CollisionMask) {
	f.sprite.AnimateBackwards()

}

func (f *Frog) updateCollider() {

}

func (f *Frog) SetSpeed(speed float64) {
	f.Speed = speed
}

func (f *Frog) GetSpeed() float64 {
	return f.Speed
}

func (f *Frog) SetVelocity(speed float64) {
	//f *FrogVelocity = speed
}

func (f *Frog) PassBy(bs2 entity.Enemy) {

}

func (f *Frog) GetVelocity() float64 {
	//return bs.Velocity
	return 0
}
