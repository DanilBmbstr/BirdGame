package enemy

/*──────────██
─────────█▒▒█
─────────█▒▒█
─────────█▒▒█
──────██▒█▒▒█
────██▒▒██▒▒█
──██▒▒█▒█████
─█▒▒█▒▒█▒▒▒▒▒█
█▒█▒▒█▒▒███▒▒▒█
─█▒█▒▒█▒▒█▒█▒▒█
─█▒█▒▒███▒▒▒▒█
──█▒██▒▒▒▒▒▒█
───█▒▒▒▒▒▒▒█
────█▒▒▒▒▒▒█
───██████████


*/
import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"

	"github.com/quasilyte/gmath"
)

type EnemyBasic struct {
	sprite     assets.Sprite
	Center     gmath.Vec
	Direction  gmath.Vec
	Lenght     float64
	PlayerTail *player.Tail
	collider   collisionsystem.Shape
}

func (e *EnemyBasic) GetPosition() gmath.Vec {
	return e.Center
}

func (e *EnemyBasic) GetCollider() *collisionsystem.Shape {
	return &e.collider
}

func (e *EnemyBasic) initBase(x, y float64, sprite assets.Sprite) {
	e.SetPosition(x, y)
	e.sprite = sprite
	e.sprite.SetOriginAtCenter()
}

func (e *EnemyBasic) SetPosition(x, y float64) {
	e.Center.X = x
	e.Center.Y = y
}

func (e *EnemyBasic) MoveTowards(dest gmath.Vec, speed float64) {
	e.Center = e.Center.Add(e.Center.DirectionTo(dest).Mulf(speed))
}

func (e *EnemyBasic) DrawEntity(screen *ebiten.Image) error {
	if (e.Center == gmath.Vec{}) {

		return fmt.Errorf("enemy position is not initialized")
	}

	e.sprite.Position = e.Center
	e.sprite.DrawSprite(screen)
	return nil
}

func (e *EnemyBasic) CheckCollisionsWithTail(tail *player.Tail) bool {
	return (tail.GetCollider().CheckCollisionsWithShape(e.collider))

}
