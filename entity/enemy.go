package entity

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
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
)

type Enemy struct {
	Center     gmath.Vec
	Direction  gmath.Vec
	Lenght     float64
	PlayerTail *Tail
}

func (e *Enemy) Init(x, y float64) {
	e.SetPosition(x, y)
}

func (e *Enemy) SetPosition(x, y float64) {
	e.Center.X = x
	e.Center.Y = y
}

func (e *Enemy) MoveTowards(dest gmath.Vec, speed float64) {
	e.Center = e.Center.Add(e.Center.DirectionTo(dest).Normalized().Mulf(speed))
}

func (e *Enemy) DrawEntity(screen *ebiten.Image) {

	vector.FillCircle(screen, float32(e.Center.X), float32(e.Center.Y), 10, color.RGBA{255, 0, 0, 255}, true)

}
