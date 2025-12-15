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
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/camera"
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

func (p *Player) DrawEntity(screen *ebiten.Image) {
	xOffset := camera.X_offset - camera.WindowSizeX/2
	yOffset := camera.Y_offset - camera.WindowSizeY/2

	p.PlayerTail.DrawTail(screen)
	left, right := p.getFlanks(tailStartWidth)
	path := &vector.Path{}
	path.MoveTo(p.Center.AsVec32().X-float32(xOffset), p.Center.AsVec32().Y-float32(yOffset))

	path.LineTo(right.AsVec32().X-float32(xOffset), right.AsVec32().Y-float32(yOffset))
	path.LineTo(float32(p.PlayerTail.segments[tailStartSegmentsLenght].getSegmentEnd(p.TailSegmentLenght+tailStartLenght).X-xOffset), float32(p.PlayerTail.segments[3].getSegmentEnd(p.TailSegmentLenght).Y-yOffset))
	path.LineTo(left.AsVec32().X-float32(xOffset), left.AsVec32().Y-float32(yOffset))
	path.Close()

	vector.FillPath(screen, path, &vector.FillOptions{FillRule: vector.FillRuleEvenOdd}, &vector.DrawPathOptions{AntiAlias: true,
		ColorScale: p.playerColor})

	p.PlayerSprite.Position = p.Center
	p.PlayerSprite.DrawSprite(screen)

}

func (p *Player) Rotate(dest gmath.Vec, speed float64) {

	p.Direction = p.Direction.LinearInterpolate(gmath.Vec{camera.WindowSizeX / 2, camera.WindowSizeY / 2}.DirectionTo(dest), speed)

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
