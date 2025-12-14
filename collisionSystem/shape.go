package collisionsystem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
)

type Shape struct {
	Points []gmath.Vec
}

func (s *Shape) DrawEntity(screen *ebiten.Image, xOffset, yOffset int) {
	path := &vector.Path{}
	path.MoveTo(s.Points[0].AsVec32().X-float32(xOffset), s.Points[0].AsVec32().Y-float32(yOffset))
	for i := 1; i < len(s.Points); i++ {
		path.LineTo(float32(s.Points[i].X-float64(xOffset)), float32(s.Points[i].Y-float64(yOffset)))
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
