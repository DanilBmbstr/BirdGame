package collisionsystem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/camera"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
)

type Shape struct {
	Points []gmath.Vec
}

func (s *Shape) DrawEntity(screen *ebiten.Image) {
	xOffset := camera.X_offset - camera.WindowSizeX/2
	yOffset := camera.Y_offset - camera.WindowSizeY/2
	path := &vector.Path{}
	path.MoveTo(s.Points[0].AsVec32().X-float32(xOffset), s.Points[0].AsVec32().Y-float32(yOffset))
	for i := 1; i < len(s.Points); i++ {
		path.LineTo(float32(s.Points[i].X-xOffset), float32(s.Points[i].Y-yOffset))
	}

	path.Close()
	cs := ebiten.ColorScale{}
	cs.Scale(0.25, 0.25, 0.5, 0.5)
	vector.FillPath(screen, path, &vector.FillOptions{FillRule: vector.FillRuleEvenOdd}, &vector.DrawPathOptions{AntiAlias: true,
		ColorScale: cs})
}

func (s *Shape) Copy() Shape {
	var newPoints []gmath.Vec
	newPoints = append(newPoints, s.Points...)
	return Shape{newPoints}
}
