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
type CollisionMask struct {
	info   locationInfo
	shapes []collisionsystem.Shape

	scale float64
}

func (cm *CollisionMask) SetInfo(scale float64, offset gmath.Vec) {
	cm.info.scale = scale
	cm.info.offset = offset
}

func (cm *CollisionMask) AddShape(shape collisionsystem.Shape) {
	cm.shapes = append(cm.shapes, shape)
}

func (l *CollisionMask) LoadLocation(data resource.Raw, scale float64) {
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

func (l *CollisionMask) DrawLCollisionMask(screen *ebiten.Image) {
	for i := 0; i < len(l.shapes); i++ {
		l.shapes[i].DrawEntity(screen)
	}
}

func (l *CollisionMask) CheckCollisions(point gmath.Vec) bool {
	for i := 0; i < len(l.shapes); i++ {
		if collisionsystem.CheckShapeCollisionWithPoint(point, l.shapes[i]) {
			return true
		}
	}
	return false
}
func (l *CollisionMask) CheckCollisionsWithShape(shape collisionsystem.Shape) bool {
	for i := 0; i < len(l.shapes); i++ {
		if collisionsystem.CheckShapesCollision(shape, l.shapes[i]) {
			return true
		}
	}
	return false
}

// DebugTools
func (l *CollisionMask) DebugMoveCollisions() {
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
