package level

import (
	"bytes"
	"encoding/xml"

	"github.com/hajimehoshi/ebiten/v2"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

type locationInfo struct {
	scale  float64
	offset gmath.Vec
}
type Location struct {
	info   locationInfo
	shapes []collisionsystem.Shape

	scale float64
}

func (l *Location) LoadLocation(data resource.Raw, scale float64) {
	l.scale = scale
	xmlData := bytes.NewBufferString(string(data.Data))
	d := xml.NewDecoder(xmlData)
	j := -1
	for t, _ := d.Token(); t != nil; t, _ = d.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "Scale" {
				d.DecodeElement(&l.info.scale, &se)
			}
			if se.Name.Local == "Offset" {
				d.DecodeElement(&l.info.offset, &se)

			}
			if se.Name.Local == "Shape" {
				shape := collisionsystem.Shape{}
				l.shapes = append(l.shapes, shape)
				j++
			}
			if se.Name.Local == "Point" {
				vec := gmath.Vec{}
				d.DecodeElement(&vec, &se)
				vec = vec.Mulf(l.scale * l.info.scale)
				vec = vec.Add(l.info.offset)
				l.shapes[j].Points = append(l.shapes[j].Points, vec)

			}
		}
	}
}

func (l *Location) DrawLocation(screen *ebiten.Image) {
	for i := 0; i < len(l.shapes); i++ {
		l.shapes[i].DrawEntity(screen)
	}
}

func (l *Location) CheckCollisions(point gmath.Vec) bool {
	for i := 0; i < len(l.shapes); i++ {
		if collisionsystem.CheckCollisionWithShape(point, l.shapes[i]) {
			return true
		}
	}
	return false
}

// DebugTools
func (l *Location) DebugMoveCollisions() {
	speed := 0.5
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		l.info.offset.Y -= speed
		for i := range len(l.shapes) {

			for j := range len(l.shapes[i].Points) {
				l.shapes[i].Points[j].Y -= speed
			}

		}
		println(l.info.offset.String())
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		l.info.offset.Y += speed
		for i := range len(l.shapes) {

			for j := range len(l.shapes[i].Points) {
				l.shapes[i].Points[j].Y += speed
			}

		}
		println(l.info.offset.String())
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		l.info.offset.X += speed
		for i := range len(l.shapes) {

			for j := range len(l.shapes[i].Points) {
				l.shapes[i].Points[j].X += speed
			}

		}
		println(l.info.offset.String())
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		l.info.offset.X -= speed
		for i := range len(l.shapes) {

			for j := range len(l.shapes[i].Points) {
				l.shapes[i].Points[j].X -= speed
			}

		}
		println(l.info.offset.String())
	}

}
