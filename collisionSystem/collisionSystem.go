package collisionsystem

import (
	"github.com/quasilyte/gmath"
)

func getDividingVector(a, b gmath.Vec) gmath.Vec {
	return a.DirectionTo(b)

}

func getProjectionsOverlapping(point gmath.Vec, shape Shape, dividerAngle gmath.Rad) bool {

	point = point.Rotated(dividerAngle)
	maxY := shape.Points[0].Rotated(dividerAngle).Y
	minY := shape.Points[0].Rotated(dividerAngle).Y
	for i := 1; i < len(shape.Points); i++ {
		shape.Points[i] = shape.Points[i].Rotated(dividerAngle)
		if shape.Points[i].Y > maxY {
			maxY = shape.Points[i].Y
		}
		if shape.Points[i].Y < minY {
			minY = shape.Points[i].Y
		}
	}

	if !(point.Y < maxY && point.Y > minY) {

		return false

	} else {
		return true
	}
}

func CheckShapeCollisionWithPoint(point gmath.Vec, shape Shape) bool {
	var b gmath.Vec
	pointsNum := len(shape.Points) - 1
	for i := 0; i < len(shape.Points); i++ {
		shapeCopy := shape.Copy()
		a := shapeCopy.Points[i]

		if i == pointsNum {
			b = shapeCopy.Points[0]
		} else {
			b = shapeCopy.Points[i+1]
		}

		divider := getDividingVector(a, b)

		if !getProjectionsOverlapping(point, shapeCopy, -divider.Angle()) {

			return false
		}
	}
	return true

}

func CheckShapesCollision(shape1, shape2 Shape) bool {
	for j := 0; j < len(shape1.Points); j++ {
		point := shape1.Points[j]
		if CheckShapeCollisionWithPoint(point, shape2) {
			return true
		}
	}
	for j := 0; j < len(shape2.Points); j++ {
		point := shape2.Points[j]
		if CheckShapeCollisionWithPoint(point, shape1) {
			return true
		}
	}
	return false
}
