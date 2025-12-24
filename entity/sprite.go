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
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/camera"
	"github.com/quasilyte/gmath"
)

type Sprite struct {
	spriteImage *ebiten.Image
	scale       float64
	origin      gmath.Vec
	Position    gmath.Vec
	Rotation    gmath.Rad
	imageSize   gmath.Vec
}

func (s *Sprite) DrawSprite(screen *ebiten.Image) {
	xOffset := camera.X_offset - camera.WindowSizeX/2
	yOffset := camera.Y_offset - camera.WindowSizeY/2
	rotationOffset := s.origin.Rotated(s.Rotation)

	imageOptions := ebiten.DrawImageOptions{}

	imageOptions.GeoM.Rotate(float64(s.Rotation))

	imageOptions.GeoM.Scale(s.scale, s.scale)
	imageOptions.GeoM.Translate(-rotationOffset.X, -rotationOffset.Y)
	imageOptions.GeoM.Translate(s.Position.X-xOffset, s.Position.Y-yOffset)
	screen.DrawImage(s.spriteImage, &imageOptions)

}

func (s *Sprite) SetImage(img *ebiten.Image) {
	s.imageSize = gmath.Vec{X: float64(img.Bounds().Max.X), Y: float64(img.Bounds().Max.Y)}
	s.scale = 1
	s.origin = gmath.Vec{X: 0, Y: 0}
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

func (s *Sprite) GetOrigin() gmath.Vec {
	return s.origin
}

func (s *Sprite) SetOrigin(pos gmath.Vec) {
	s.origin = pos
}

func (s *Sprite) SetOriginAtCenter() error {
	if s.spriteImage == nil {
		return fmt.Errorf("Sprite image is not set")
	}
	s.origin = s.GetImageSize().Mulf(0.5)
	return nil

}

func (s *Sprite) MoveOrigin(vec gmath.Vec) {
	s.origin = s.origin.Add(vec)
}
