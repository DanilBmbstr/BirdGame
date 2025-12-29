package player

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
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/camera"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"

	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/optionsLoaders"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
)

type Player struct {
	PlayerDto    optionsLoaders.PlayerDTO
	BattleSystem PlayerBattleSystem
	playerSprite assets.Sprite
	center       gmath.Vec
	direction    gmath.Vec

	tailColor         ebiten.ColorScale
	playerTail        *Tail
	tailLenght        int
	tailSegmentLenght float64
	playerColor       ebiten.ColorScale

	//Public
	//Lenght     float64
	Controller PlayerController

	//Константы

	tailStartSegmentsLenght int
	tailStartWidth          float32
	tailWidth               float32
	tailWidthNarrow         float32
}

func (p *Player) Reset() {
	p.InitPlayer(1000, 1500, p.PlayerDto)
}

func (p *Player) GetCollider() *collisionsystem.Shape {
	size := 50.0
	collider := *&collisionsystem.Shape{}

	collider.Points = append(collider.Points, p.GetPosition().Add(gmath.Vec{X: 0, Y: size}.Rotated(p.direction.Angle())))
	collider.Points = append(collider.Points, p.GetPosition().Add(gmath.Vec{X: size, Y: 0}.Rotated(p.direction.Angle())))
	collider.Points = append(collider.Points, p.GetPosition().Add(gmath.Vec{X: 0, Y: -size}.Rotated(p.direction.Angle())))

	collider.Points = append(collider.Points, p.GetPosition().Add(gmath.Vec{X: -size / 2, Y: 0}.Rotated(p.direction.Angle())))
	return &collider
}

func (p *Player) GetDirection() gmath.Vec {
	return p.direction
}

func (p *Player) GetTail() *Tail {
	return p.playerTail
}

func (p *Player) GetSprite() *assets.Sprite {
	return &p.playerSprite
}

func (p *Player) SetSprite(s assets.Sprite) {
	p.playerSprite = s
}

func (p *Player) InitPlayer(xPos, yPos float64, dto optionsLoaders.PlayerDTO) {
	p.PlayerDto = dto
	p.direction = gmath.Vec{X: 1, Y: 1}
	p.SetPosition(xPos, yPos)

	//Инициализируем цвета

	tailColor := ebiten.ColorScale{}
	tailColor.Scale(dto.TailColor[0], dto.TailColor[1], dto.TailColor[2], dto.TailColor[3])
	p.tailColor = tailColor
	playerColor := ebiten.ColorScale{}
	playerColor.Scale(dto.PlayerColor[0], dto.PlayerColor[1], dto.PlayerColor[2], dto.PlayerColor[3])
	p.playerColor = playerColor
	p.playerSprite.SetScale(dto.SpriteScale)

	p.tailLenght = dto.TailLenght
	p.tailSegmentLenght = dto.TailSegmentLenght
	p.tailStartSegmentsLenght = dto.TailStartSegmentsLenght
	p.tailStartWidth = dto.TailStartWidth
	p.tailWidth = dto.TailWidth
	p.tailWidthNarrow = dto.TailWidthNarrow

	t := Tail{tailWidth: p.tailWidth, widthNarrow: p.tailWidthNarrow}
	p.playerTail = &t
	p.playerTail.segmentLenght = p.tailSegmentLenght
	p.playerTail.CreateTail(p.center, p.tailLenght)
	p.playerTail.movingSpeed = dto.TailMovingSpeed
	p.playerTail.rotationSpeed = dto.TailRotationSpeed
	p.Controller = PlayerController{Speed: dto.Speed, RotationSpeed: dto.RotationSpeed}

	p.playerSprite.SetOriginAtCenter()
	p.playerSprite.MoveOrigin(gmath.Vec{X: 0, Y: -8})

	p.BattleSystem.Init()

}

func (p *Player) SetPosition(x, y float64) {
	p.center.X = x
	p.center.Y = y
}
func (p *Player) GetPosition() gmath.Vec {
	return p.center
}

func (p *Player) SetDirection(x, y float64) {
	p.direction.X = x
	p.direction.Y = y
	p.direction = p.direction.Normalized()
}

func (p *Player) getFlanks(width float64) (gmath.Vec, gmath.Vec) {

	left := p.center.Add(p.direction.Normalized().Rotated(gmath.DegToRad(-90)).Mulf(float64(width)))
	right := p.center.Add(p.direction.Normalized().Rotated(gmath.DegToRad(90)).Mulf(float64(width)))

	return left, right
}

func (p *Player) DrawEntity(screen *ebiten.Image) {
	xOffset := camera.X_offset - camera.WindowSizeX/2
	yOffset := camera.Y_offset - camera.WindowSizeY/2
	cs := p.tailColor
	cs.SetA((0.49 - (float32(p.BattleSystem.cooldown) / 100)))

	p.playerTail.DrawTail(screen, cs)

	left, right := p.getFlanks(float64(p.tailStartWidth))
	path := &vector.Path{}
	path.MoveTo(p.center.AsVec32().X-float32(xOffset), p.center.AsVec32().Y-float32(yOffset))

	path.LineTo(right.AsVec32().X-float32(xOffset), right.AsVec32().Y-float32(yOffset))
	path.LineTo(float32(p.playerTail.segments[p.tailStartSegmentsLenght+1].start.X-xOffset), float32(p.playerTail.segments[p.tailStartSegmentsLenght+1].start.Y-yOffset))
	path.LineTo(left.AsVec32().X-float32(xOffset), left.AsVec32().Y-float32(yOffset))
	path.Close()

	vector.FillPath(screen, path, &vector.FillOptions{FillRule: vector.FillRuleEvenOdd}, &vector.DrawPathOptions{AntiAlias: true,
		ColorScale: p.playerColor})

	path.MoveTo(right.AsVec32().X-float32(xOffset), right.AsVec32().Y-float32(yOffset))
	sLeft, sRight := p.playerTail.segments[0].getSegmentFlanks()

	path.LineTo(float32(sRight.X)-float32(xOffset), float32(sRight.Y-yOffset))
	path.LineTo(float32(sLeft.X)-float32(xOffset), float32(sLeft.Y-yOffset))
	path.LineTo(left.AsVec32().X-float32(xOffset), left.AsVec32().Y-float32(yOffset))
	vector.FillPath(screen, path, &vector.FillOptions{FillRule: vector.FillRuleEvenOdd}, &vector.DrawPathOptions{AntiAlias: true,
		ColorScale: cs})
	p.playerSprite.Position = p.center
	p.playerSprite.DrawSprite(screen)

}
