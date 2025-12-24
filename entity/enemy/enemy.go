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
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"

	"github.com/quasilyte/gmath"
)

type Enemy struct {
	sprite     assets.Sprite
	Center     gmath.Vec
	Direction  gmath.Vec
	Lenght     float64
	PlayerTail *player.Tail
}

func (e *Enemy) Init(x, y float64, sprite assets.Sprite) {
	e.SetPosition(x, y)
	e.sprite = sprite
}

func (e *Enemy) SetPosition(x, y float64) {
	e.Center.X = x
	e.Center.Y = y
}

func (e *Enemy) MoveTowards(dest gmath.Vec, speed float64) {
	e.Center = e.Center.Add(e.Center.DirectionTo(dest).Normalized().Mulf(speed))
}

func (e *Enemy) DrawEntity(screen *ebiten.Image) error {
	if (e.Center == gmath.Vec{}) {

		return fmt.Errorf("enemy position is not initialized")
	}

	e.sprite.Position = e.Center
	e.sprite.DrawSprite(screen)
	return nil
}

func (e *Enemy) CheckCollisionsWithTail(tail *player.Tail) {

}
