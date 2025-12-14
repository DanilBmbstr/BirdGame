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
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
)

const (
	tailStartLenght         = 10
	tailStartSegmentsLenght = 3

	tailStartWidth  = 15
	tailWidth       = 9
	tailWidthNarrow = 0.95
)

type Player struct {
	PlayerSprite Sprite
	Center       gmath.Vec
	Direction    gmath.Vec
	Lenght       float64

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

	left := p.Center.Add(p.Direction.Rotated(gmath.DegToRad(-90)).Mulf(float64(width)))
	right := p.Center.Add(p.Direction.Rotated(gmath.DegToRad(90)).Mulf(float64(width)))

	return left, right
}

func (p *Player) DrawEntity(screen *ebiten.Image, xOffset, yOffset int) {
	p.PlayerTail.DrawTail(screen, xOffset, yOffset)
	left, right := p.getFlanks(tailStartWidth)
	path := &vector.Path{}
	path.MoveTo(p.Center.AsVec32().X-float32(xOffset), p.Center.AsVec32().Y-float32(yOffset))

	path.LineTo(right.AsVec32().X-float32(xOffset), right.AsVec32().Y-float32(yOffset))
	path.LineTo(float32(p.PlayerTail.segments[tailStartSegmentsLenght].getSegmentEnd(p.TailSegmentLenght+tailStartLenght).X-float64(xOffset)), float32(p.PlayerTail.segments[3].getSegmentEnd(p.TailSegmentLenght).Y-float64(yOffset)))
	path.LineTo(left.AsVec32().X-float32(xOffset), left.AsVec32().Y-float32(yOffset))
	path.Close()

	vector.FillPath(screen, path, &vector.FillOptions{FillRule: vector.FillRuleEvenOdd}, &vector.DrawPathOptions{AntiAlias: true,
		ColorScale: p.playerColor})

	p.PlayerSprite.Position = p.Center
	p.PlayerSprite.DrawSprite(screen, xOffset, yOffset)

	/*
		//Создание фигуры птицы
		path := &vector.Path{}

		path.MoveTo(float32(p.Center.X)+float32(p.Direction.X*p.Lenght), float32(p.Center.Y)+float32(p.Direction.Y*p.Lenght))
		path.LineTo(float32(p.Center.X)+float32(p.Direction.X*p.Lenght-0.022871), float32(p.Center.Y)+float32(p.Direction.Y*p.Lenght+0.930223))
		path.LineTo(float32(p.Center.X)+float32(p.Direction.X*p.Lenght-0.305671), float32(p.Center.Y)+float32(p.Direction.Y*p.Lenght+0.706639))
		path.LineTo(float32(p.Center.X)+float32(p.Direction.X*p.Lenght-0.230515), float32(p.Center.Y)+float32(p.Direction.Y*p.Lenght+0.679934))
		leftFlank, rightFlank := p.PlayerTail.segments[0].getSegmentFlanks()
		path.LineTo(leftFlank.AsVec32().X, leftFlank.AsVec32().Y)
		path.LineTo(rightFlank.AsVec32().X, rightFlank.AsVec32().Y)
		path.Close()
		cs := ebiten.ColorScale{}
		cs.Scale(1, 1, 1, 1)
		vector.FillPath(screen, path, &vector.FillOptions{FillRule: vector.FillRuleEvenOdd}, &vector.DrawPathOptions{AntiAlias: true,
			ColorScale: cs})
	*/

}

func (p *Player) Rotate(dest gmath.Vec, speed float64) {

	p.Direction = p.Direction.LinearInterpolate(dest, speed)

}

func (p *Player) RotateTowards(dest gmath.Vec, speed float64) {

	p.Direction = p.Direction.LinearInterpolate(p.Center.DirectionTo(dest), speed)

}

func (p *Player) MoveForward() {
	p.Center = p.Center.Add(p.Direction.Mulf(p.Speed))
	if p.PlayerTail.GetTailDistanceTo(p.Center) > p.PlayerTail.SegmentLenght {
		p.PlayerTail.MoveTail(p.Center)
	}

	p.PlayerSprite.Rotation = p.Direction.Angle() + gmath.DegToRad(90)
}
