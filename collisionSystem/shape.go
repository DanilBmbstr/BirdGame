package collisionsystem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
)

type Shape struct {
	Points []gmath.Vec
}

func (s *Shape) DrawEntity(screen *ebiten.Image) {
	path := &vector.Path{}
	path.MoveTo(s.Points[0].AsVec32().X, s.Points[0].AsVec32().Y)
	for i := 1; i < len(s.Points); i++ {
		path.LineTo(float32(s.Points[i].X), float32(s.Points[i].Y))
	}

	path.Close()
	cs := ebiten.ColorScale{}
	cs.Scale(0, 1, 0, 1)
	vector.FillPath(screen, path, &vector.FillOptions{FillRule: vector.FillRuleEvenOdd}, &vector.DrawPathOptions{AntiAlias: true,
		ColorScale: cs})
}

func (s *Shape) Copy() Shape {
	var newPoints []gmath.Vec
	newPoints = append(newPoints, s.Points...)
	return Shape{newPoints}
}
