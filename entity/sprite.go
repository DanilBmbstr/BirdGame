package entity

/*⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⠛⠛⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⡿⠋⢀⣠⣤⣶⣶⣶⣤⡀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⡟⠄⢰⣿⣿⡿⠛⠉⠉⣉⣉⠄⠄⠘⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⠇⠄⢸⣿⣿⠁⠘⣿⣿⣿⣿⣿⣿⣦⡀⢹⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⠿⠟⠛⠛⠄⠄⣼⣿⣿⡀⠄⠈⠛⠛⠛⠛⠛⠋⠁⢸⣿⣿⣿⣿⣿⣿⣿⣿
⣿⡇⢀⣶⣿⠂⠄⢸⣿⣿⣿⣧⣤⣄⣀⣀⣀⣀⣤⠄⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⡇⠄⣿⣿⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣄⣈⠙⠛⠿⣿⣿⣿⣿⣿
⣿⡇⠄⣿⣿⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣤⣈⠻⣿⣿⣿
⣿⡇⠄⠈⠉⠄⣠⣿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⠈⢻⣿
⣿⣿⠆⠄⠄⠄⢸⣿⡆⠙⢿⣿⣿⣿⣿⣿⣿⣿⡿⠿⠛⣿⣿⣿⡿⠟⠋⢀⣾⣿
⣿⣿⠄⠄⠄⠄⣼⣿⣿⣆⡀⠈⠉⠉⠉⠉⠉⠄⢀⣀⣀⠄⠁⠄⢀⣠⣾⣿⣿⣿
⣿⣿⡆⠄⠄⠄⢿⣿⣿⣿⣿⣶⣤⣤⣤⣤⣶⣶⣾⣿⣿⠁⠄⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⡆⠄⠄⠈⠻⣿⣿⣿⣿⠟⠉⠉⠉⠁⣼⣿⣿⠏⠄⢠⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⡄⠄⠄⠄⠈⠛⠛⠁⢠⣄⠄⠄⠄⠄⠄⠄⠄⣰⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣷⡄⠄⠄⠄⠄⠄⠄⠛⠛⠂⠄⠄⠄⠄⠄⠴⢿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿*/
import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

type Sprite struct {
	spriteImage *ebiten.Image
	scale       float64
	Origin      gmath.Vec
	Position    gmath.Vec
	Rotation    gmath.Rad
	imageSize   gmath.Vec
}

func (s *Sprite) DrawSprite(screen *ebiten.Image, xOffset, yOffset int) {

	rotationOffset := s.Origin.Rotated(s.Rotation)

	imageOptions := ebiten.DrawImageOptions{}

	imageOptions.GeoM.Rotate(float64(s.Rotation))

	imageOptions.GeoM.Scale(s.scale, s.scale)
	imageOptions.GeoM.Translate(-rotationOffset.X, -rotationOffset.Y)
	imageOptions.GeoM.Translate(s.Position.X-float64(xOffset), s.Position.Y-float64(yOffset))
	screen.DrawImage(s.spriteImage, &imageOptions)

}

func (s *Sprite) SetImage(img *ebiten.Image) {
	s.imageSize = gmath.Vec{X: float64(img.Bounds().Max.X), Y: float64(img.Bounds().Max.Y)}
	s.scale = 1
	s.Origin = gmath.Vec{X: 0, Y: 0}
	s.Rotation = 0
	s.spriteImage = img
}

func (s *Sprite) GetImage() *ebiten.Image {
	return s.spriteImage
}

func (s *Sprite) SetScale(scale float64) {
	s.scale = scale
}

func (s *Sprite) GetScale() float64 {
	return s.scale
}

func (s *Sprite) GetImageSize() gmath.Vec {
	return s.imageSize.Mulf(s.scale)
}
