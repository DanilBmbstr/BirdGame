package entity

/*██████████████████████████
█░░░░░░░░█░░░░░░░█░░░░░░░█
█░█████░░█░█████░█░█████░█
█░█░░░█░░█░█░░░█░█░█░░░█░█
█░█░█░█░░█░█░█░█░█░█░█░█░█
█░█░░░█░░█░█░░░█░█░█░░░█░█
█░█████░░█░█████░█░█████░█
█░█░░█░░░█░█░░░█░█░█░░░░░█
█░█░░█░░░█░█░░░█░█░█░█████
█░█░░█░░░█░█░░░█░█░█░█████
█░█░░███░█░█░░░█░█░█░█████
█░░░░░░░░█░░░░░░░█░░░█████
██████████████████████████
█░░░░░░▄▄░░░░░░░░░░░░░░░░█
█░░░░▄▀░░█░░░░░░░░░░░░░░░█
█░░░░█░░░█░░░░░░░░▄▄▀▄░░░█
█░░░░█░░░█░░░░░░░░█░░░█░░█
█░░░░█░░░█▄▄▄▄░▄▄▄█░░░█░░█
█░░░░█░░░█▀░░▀█░░▀█░░░█░░█
█░░░░█░▄▄█▄░░░█░░░█░░░█░░█
█░░░▄█▀░░░░█░░█░░░█░░░█░░█
█░░▄▀░░░░░▄▀░░█░░░█░░░█░░█
█░░█░░░░█▀█░░░█░░░█░░░█░░█
█░░█░░░░▀░▀▄▄▀▀▄▄▄▀░░░█░░█
█░░█░░░░░░░░░░░░░░░░░░█░░█
█░░▀▄░░░░░░░░░░░░░░░░▄▀░░█
█░░░▀▄▄░░░░░░░░░░░░▄▄▀░░░█
█░░░░░▀█▄▄▄▄▄▄▄▄▄▄█▀░░░░░█
█░░░░░░░░▀▀▀▀▀▀▀▀░░░░░░░░█
██████████████████████████*/
import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/camera"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
)

const (
	tailStartLenght         = 0
	tailStartSegmentsLenght = 1

	tailStartWidth  = 15
	tailWidth       = 9
	tailWidthNarrow = 0.95
)

type Player struct {
	PlayerSprite      Sprite
	Center            gmath.Vec
	Direction         gmath.Vec
	Lenght            float64
	TailColor         ebiten.ColorScale
	PlayerTail        *Tail
	TailLenght        int
	TailSegmentLenght float64
	Speed             float64
	playerColor       ebiten.ColorScale
}

func (p *Player) InitPlayer(xPos, yPos float64) {
	p.SetPosition(xPos, yPos)
	t := Tail{tailWidth: tailWidth, widthNarrow: tailWidthNarrow}
	p.PlayerTail = &t
	p.PlayerTail.SegmentLenght = p.TailSegmentLenght
	p.PlayerTail.CreateTail(p.Center, p.TailLenght)

}

func (p *Player) SetPosition(x, y float64) {
	p.Center.X = x
	p.Center.Y = y
}

func (p *Player) SetDirection(x, y float64) {
	p.Direction.X = x
	p.Direction.Y = y
	p.Direction = p.Direction.Normalized()
}

func (p *Player) getFlanks(width float64) (gmath.Vec, gmath.Vec) {

	left := p.Center.Add(p.Direction.Normalized().Rotated(gmath.DegToRad(-90)).Mulf(float64(width)))
	right := p.Center.Add(p.Direction.Normalized().Rotated(gmath.DegToRad(90)).Mulf(float64(width)))

	return left, right
}

func (p *Player) DrawEntity(screen *ebiten.Image) {
	xOffset := camera.X_offset - camera.WindowSizeX/2
	yOffset := camera.Y_offset - camera.WindowSizeY/2

	p.PlayerTail.DrawTail(screen, p.TailColor)
	left, right := p.getFlanks(tailStartWidth)
	path := &vector.Path{}
	path.MoveTo(p.Center.AsVec32().X-float32(xOffset), p.Center.AsVec32().Y-float32(yOffset))

	path.LineTo(right.AsVec32().X-float32(xOffset), right.AsVec32().Y-float32(yOffset))
	path.LineTo(float32(p.PlayerTail.segments[tailStartSegmentsLenght+1].start.X-xOffset), float32(p.PlayerTail.segments[tailStartSegmentsLenght+1].start.Y-yOffset))
	path.LineTo(left.AsVec32().X-float32(xOffset), left.AsVec32().Y-float32(yOffset))
	path.Close()

	vector.FillPath(screen, path, &vector.FillOptions{FillRule: vector.FillRuleEvenOdd}, &vector.DrawPathOptions{AntiAlias: true,
		ColorScale: p.playerColor})

	path.MoveTo(right.AsVec32().X-float32(xOffset), right.AsVec32().Y-float32(yOffset))
	sLeft, sRight := p.PlayerTail.segments[0].getSegmentFlanks()

	path.LineTo(float32(sRight.X)-float32(xOffset), float32(sRight.Y-yOffset))
	path.LineTo(float32(sLeft.X)-float32(xOffset), float32(sLeft.Y-yOffset))
	path.LineTo(left.AsVec32().X-float32(xOffset), left.AsVec32().Y-float32(yOffset))
	vector.FillPath(screen, path, &vector.FillOptions{FillRule: vector.FillRuleEvenOdd}, &vector.DrawPathOptions{AntiAlias: true,
		ColorScale: p.TailColor})
	p.PlayerSprite.Position = p.Center
	p.PlayerSprite.DrawSprite(screen)

}

func (p *Player) getRotationAngle(dest gmath.Vec) gmath.Rad {
	return gmath.Vec{X: dest.X - camera.WindowSizeX/2, Y: dest.Y - camera.WindowSizeY/2}.Angle() - p.Direction.Angle()
}
func (p *Player) newDir(speed float64) gmath.Vec { return p.Direction.Rotated(gmath.Rad(speed)) }

func (p *Player) Rotate(dest gmath.Vec, speed float64) {

	windowCenter := gmath.Vec{X: camera.WindowSizeX / 2, Y: camera.WindowSizeY / 2}
	rotationVector := windowCenter.DirectionTo(dest)
	mouseAngle := gmath.Vec{X: dest.X - camera.WindowSizeX/2, Y: dest.Y - camera.WindowSizeY/2}.Angle()
	maxAngle := float64(-10)

	if windowCenter.DistanceTo(dest) < 10 {
		return
	}
	if (math.Abs(float64(rotationVector.Angle()-p.Direction.Angle()))) < float64(gmath.DegToRad(90)) || (math.Abs(float64(rotationVector.Angle()-p.Direction.Angle()))) > float64(gmath.DegToRad(300)) {
		dir := p.Direction.LinearInterpolate(rotationVector, speed)

		p.Direction = dir
	} else {

		var dir gmath.Vec

		if p.getRotationAngle(dest) > gmath.Rad(speed) {

			dir = p.newDir(speed)

			if mouseAngle > (gmath.DegToRad(90)) && p.Direction.Angle() < gmath.DegToRad(maxAngle) {
				dir = p.newDir(-speed)
			}

		} else if p.getRotationAngle(dest) < gmath.Rad(-speed) {
			dir = p.newDir(-speed)

			if mouseAngle < (gmath.DegToRad(-90)) && p.Direction.Angle() > gmath.DegToRad(maxAngle) {
				dir = p.newDir(speed)
			}
		} else {
			dir = p.Direction
		}
		p.Direction = dir
	}

}

func (p *Player) RotateTowards(dest gmath.Vec, speed float64) {

	p.Direction = p.Direction.LinearInterpolate(p.Center.DirectionTo(dest), speed)

}

func (p *Player) MoveForward() {
	p.Center = p.Center.Add(p.Direction.Normalized().Mulf(p.Speed))
	if p.PlayerTail.GetTailDistanceTo(p.Center) > p.PlayerTail.SegmentLenght {
		p.PlayerTail.MoveTail(p.Center)
	}

	p.PlayerSprite.Rotation = p.Direction.Angle() + gmath.DegToRad(90)
}
