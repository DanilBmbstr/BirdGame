package enemy

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/camera"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/level"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
)

var (
	outside_circle color.Color
	inside_circle  color.Color
)

type Frog_Projecttile struct {
	EnemyBasic
	Speed float64

	direction      gmath.Vec
	target         gmath.Vec
	colliderRadius float64
	corpseMode     bool
	scale          float32
}

func (fp *Frog_Projecttile) Init(x, y float64) {
	fp.colliderRadius = 10
	fp.SetPosition(x, y)
	fp.Speed = 5
	outside_circle = color.RGBA{20, 20, 20, 0}
	inside_circle = color.RGBA{80, 80, 80, 50}
	fp.scale = 1
}

func (f *Frog_Projecttile) die() {
	f.scale = (f.scale * 0.9)
	if f.scale < 0.0001 {

		f.isDestroyed = true
	}
}

func (f *Frog_Projecttile) Update(player *player.Player, location *level.CollisionMask) {
	if f.corpseMode {
		f.die()
	}
	if !f.corpseMode {
		f.updateCollider()
		if collisionsystem.CheckShapesCollision(f.collider, *player.GetCollider()) {
			player.Reset()
			f.isDestroyed = true
		}
		if player.BattleSystem.HitTime > 1 {
			if f.CheckCollisionsWithTail(player.GetTail()) {
				f.corpseMode = true
			}
		}

		newpos := f.GetPosition().Add(f.GetPosition().DirectionTo(player.GetPosition()).Normalized().Mulf(f.Speed))

		f.SetPosition(newpos.X, newpos.Y)
	}

}

func (f *Frog_Projecttile) Draw(screen *ebiten.Image) {
	xOffset := camera.X_offset - camera.WindowSizeX/2
	yOffset := camera.Y_offset - camera.WindowSizeY/2
	vector.FillCircle(screen, f.GetPosition().AsVec32().X-float32(xOffset), f.GetPosition().AsVec32().Y-float32(yOffset), float32(f.colliderRadius)*f.scale, outside_circle, true)
	vector.FillCircle(screen, f.GetPosition().AsVec32().X-float32(xOffset), f.GetPosition().AsVec32().Y-float32(yOffset), (float32(f.colliderRadius)/2)*f.scale, inside_circle, false)

}

func (f *Frog_Projecttile) updateCollider() {
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

func (f *Frog_Projecttile) SetSpeed(speed float64) {
	f.Speed = speed
}

func (f *Frog_Projecttile) GetSpeed() float64 {
	return f.Speed
}
