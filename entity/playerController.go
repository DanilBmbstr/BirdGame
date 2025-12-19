package entity

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/camera"
	"github.com/quasilyte/gmath"
)

type PlayerController struct {
	Speed         float64
	RotationSpeed float64
}

func (pc *PlayerController) getRotationAngle(p *Player, dest gmath.Vec) gmath.Rad {
	return gmath.Vec{X: dest.X - camera.WindowSizeX/2, Y: dest.Y - camera.WindowSizeY/2}.Angle() - p.direction.Angle()
}
func (pc *PlayerController) newDir(p *Player, speed float64) gmath.Vec {
	return p.direction.Rotated(gmath.Rad(speed))
}

func (pc *PlayerController) Rotate(p *Player, dest gmath.Vec) {

	windowCenter := gmath.Vec{X: camera.WindowSizeX / 2, Y: camera.WindowSizeY / 2}
	rotationVector := windowCenter.DirectionTo(dest)
	mouseAngle := gmath.Vec{X: dest.X - camera.WindowSizeX/2, Y: dest.Y - camera.WindowSizeY/2}.Angle()
	maxAngle := float64(-10)

	if windowCenter.DistanceTo(dest) < 10 {
		return
	}
	if (math.Abs(float64(rotationVector.Angle()-p.direction.Angle()))) < float64(gmath.DegToRad(90)) || (math.Abs(float64(rotationVector.Angle()-p.direction.Angle()))) > float64(gmath.DegToRad(300)) {
		dir := p.direction.LinearInterpolate(rotationVector, pc.RotationSpeed)

		p.direction = dir
	} else {

		var dir gmath.Vec

		if pc.getRotationAngle(p, dest) > gmath.Rad(pc.RotationSpeed) {

			dir = pc.newDir(p, pc.RotationSpeed)

			if mouseAngle > (gmath.DegToRad(90)) && p.direction.Angle() < gmath.DegToRad(maxAngle) {
				dir = pc.newDir(p, -pc.RotationSpeed)
			}

		} else if pc.getRotationAngle(p, dest) < gmath.Rad(-pc.RotationSpeed) {
			dir = pc.newDir(p, -pc.RotationSpeed)

			if mouseAngle < (gmath.DegToRad(-90)) && p.direction.Angle() > gmath.DegToRad(maxAngle) {
				dir = pc.newDir(p, pc.RotationSpeed)
			}
		} else {
			dir = p.direction
		}
		p.direction = dir
	}

}

func (pc *PlayerController) MoveForward(p *Player) {
	p.center = p.center.Add(p.direction.Normalized().Mulf(pc.Speed))
	if p.playerTail.GetTailDistanceTo(p.center) > p.playerTail.SegmentLenght {
		p.playerTail.MoveTail(p.center)
	}

	p.playerSprite.Rotation = p.direction.Angle() + gmath.DegToRad(90)
}

/*func (p *Player) RotateTowards(dest gmath.Vec, speed float64) {

	p.direction = p.direction.LinearInterpolate(p.center.DirectionTo(dest), speed)

}*/
