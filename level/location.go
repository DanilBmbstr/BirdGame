package level

import (
	"bytes"
	"encoding/xml"

	"github.com/hajimehoshi/ebiten/v2"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

type Location struct {
	shapes []collisionsystem.Shape
}

func (l *Location) LoadLocation(data resource.Raw) {
	xmlData := bytes.NewBufferString(string(data.Data))
	d := xml.NewDecoder(xmlData)
	j := -1
	for t, _ := d.Token(); t != nil; t, _ = d.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "Shape" {
				shape := collisionsystem.Shape{}
				l.shapes = append(l.shapes, shape)
				j++
			}
			if se.Name.Local == "Point" {
				vec := gmath.Vec{}
				d.DecodeElement(&vec, &se)
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
