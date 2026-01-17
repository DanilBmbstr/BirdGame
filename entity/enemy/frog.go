package enemy

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/entity"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/level"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

const (
	shooting_chance = 5
)

type Frog struct {
	EnemyBasic
	Speed float64

	direction      gmath.Vec
	target         gmath.Vec
	colliderRadius float64

	corpseMode bool

	postmp  gmath.Vec
	targtmp gmath.Vec

	time int

	areaMin gmath.Vec
	areaMax gmath.Vec

	projecttiles []Frog_Projecttile
}

func (f *Frog) SetArea(x1, y1, x2, y2 float64) {
	f.areaMin = gmath.Vec{X: x1, Y: y1}
	f.areaMax = gmath.Vec{X: x2, Y: y2}

}

func (f *Frog) GetArea() (x1, y1, x2, y2 float64) {
	return f.areaMin.X, f.areaMin.Y, f.areaMax.X, f.areaMax.Y

}

func (f *Frog) Init(x, y float64, loader *resource.Loader) {
	f.time = 0
	f.corpseMode = false

	f.Speed = 10

	f.isDestroyed = false

	f.colliderRadius = 31

	sprite := assets.Sprite{Position: gmath.Vec{X: x, Y: y}}
	sprite.Rect = image.Rectangle{image.Point{0, 0}, image.Point{300, 342}}
	sprite.SetOriginAtCenter()
	sprite.SetImage(loader.LoadImage(assets.Frog).Data)
	sprite.Anim.AnimationRate = 2
	sprite.Anim.Loop = false
	sprite.SetScale(0.3)

	f.initBase(x, y, sprite)

}

func (f *Frog) die() {
	f.sprite.SetScale(f.sprite.GetScale() * 0.9)
	if f.sprite.GetScale() < 0.01 {

		f.isDestroyed = true
	}
}

func (e *Frog) DrawEntity(screen *ebiten.Image) error {
	if (e.Center == gmath.Vec{}) {

		return fmt.Errorf("enemy position is not initialized")
	}

	e.sprite.Position = e.Center
	e.sprite.DrawSprite(screen)
	for _, i := range e.projecttiles {
		i.Draw(screen)

	}
	return nil

}

const (
	shooting_time = 1
	max_jump_time = 1
	landing_time  = 2

	max_jump_distange = 450
	min_jump_distance = 150
)

func (f *Frog) Update(player *player.Player, location *level.CollisionMask) {
	f.updateCollider()
	if collisionsystem.CheckShapesCollision(f.collider, *player.GetCollider()) {
		player.Reset()
	}
	if player.BattleSystem.HitTime > 1 {
		if f.CheckCollisionsWithTail(player.GetTail()) {
			f.corpseMode = true
		}
	}

	for i := range f.projecttiles {
		if f.projecttiles[i].isDestroyed {
			f.projecttiles = append(f.projecttiles[:i], f.projecttiles[i+1:]...)
			break
		}
		f.projecttiles[i].Update(player, location)

	}

	f.time++

	if f.corpseMode {
		f.die()
		return
	}
	if f.time > int(Seconds(shooting_time+max_jump_time+landing_time)) {

		f.time = 0

	} else if f.time > int(Seconds(shooting_time+max_jump_time)) || (f.GetPosition().DistanceTo(f.target) <= f.Speed && f.target != gmath.Vec{}) {
		if f.time < int((Seconds(shooting_time + max_jump_time))) {
			f.time = int(Seconds(shooting_time+max_jump_time)) + 1
		}
		f.direction = gmath.Vec{}
		f.target = gmath.Vec{}
		f.sprite.AnimateBackwards()

	} else if f.time > int(Seconds(shooting_time)) {
		if r.Intn(100) == shooting_chance {
			f.projecttiles = append(f.projecttiles, Frog_Projecttile{})
			f.projecttiles[len(f.projecttiles)-1].Init(f.GetPosition().X, f.GetPosition().Y)
		}
		if (f.direction == gmath.Vec{}) {
			f.direction = gmath.Vec{X: float64(r.Intn(200)) - 100, Y: float64(r.Intn(200)) - 100}.Normalized()

			f.targtmp = f.target
			f.target = f.GetPosition().Add(f.direction.Mulf(float64(r.Intn(max_jump_distange-min_jump_distance) + min_jump_distance)))
			x1, y1, x2, y2 := f.GetArea()
			if location.CheckCollisions(f.target) ||

				f.target.X < x1 || f.target.X > x2 || f.target.Y < y1 || f.target.Y > y2 {

				f.target = gmath.Vec{}
				f.direction = gmath.Vec{}
				f.time = 0
				return

			}

		}

		f.SetPosition((f.GetPosition().Add(f.GetPosition().DirectionTo(f.target).Mulf(f.Speed))).X, f.GetPosition().Add(f.GetPosition().DirectionTo(f.target).Mulf(f.Speed)).Y)
		f.sprite.Animate()
	}

}

func (f *Frog) updateCollider() {
	coef := 1.5
	f.collider.Points = []gmath.Vec{
		gmath.Vec{X: f.Center.X + f.colliderRadius, Y: f.Center.Y + f.colliderRadius},
		gmath.Vec{X: f.Center.X + 0, Y: f.Center.Y + f.colliderRadius*coef},
		gmath.Vec{X: f.Center.X - f.colliderRadius, Y: f.Center.Y + f.colliderRadius},
		gmath.Vec{X: f.Center.X - f.colliderRadius*coef, Y: f.Center.Y + 0},
		gmath.Vec{X: f.Center.X - f.colliderRadius, Y: f.Center.Y - f.colliderRadius},
		gmath.Vec{X: f.Center.X - 0, Y: f.Center.Y - f.colliderRadius*coef},
		gmath.Vec{X: f.Center.X + f.colliderRadius, Y: f.Center.Y - f.colliderRadius},
		gmath.Vec{X: f.Center.X + f.colliderRadius*coef, Y: f.Center.Y - 0}}
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
