package assets

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
	"image"
	"math"

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
	Rect        image.Rectangle

	Anim AnimationOptions
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
	if s.Rect != image.Rect(0, 0, s.spriteImage.Bounds().Max.X, s.spriteImage.Bounds().Max.Y) {
		rect := s.Rect
		rect.Min.X += s.Anim.animationOffset
		rect.Max.X += s.Anim.animationOffset
		screen.DrawImage(s.spriteImage.SubImage(rect).(*ebiten.Image), &imageOptions)

		return
	}
	screen.DrawImage(s.spriteImage, &imageOptions)

}

func (s *Sprite) SetImage(img *ebiten.Image) {
	s.Anim.animationTime = 0

	s.Anim.animationOffset = 0
	s.imageSize = gmath.Vec{X: float64(img.Bounds().Max.X), Y: float64(img.Bounds().Max.Y)}
	if (s.Rect == image.Rectangle{}) {
		s.Rect = image.Rect(0, 0, int(s.imageSize.X), int(s.imageSize.Y))
	}
	s.scale = 1
	s.origin = gmath.Vec{X: 0, Y: 0}
	s.Rotation = 0

	s.spriteImage = img
	if s.Anim.AnimationRate == 0 {
		s.Anim.AnimationRate = 1

	}
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

func (s *Sprite) Animate() {

	if s.Anim.animationTime == s.Anim.AnimationRate {

		if s.Anim.animationOffset < s.imageSize.ToStd().X-s.Rect.Dx()*2 {
			s.Anim.animationOffset += s.Rect.Dx()
		} else {
			if s.Anim.Loop {
				s.Anim.animationOffset = 0
			}
		}
		s.Anim.animationTime = 0
	}
	s.Anim.animationTime++

}

func (s *Sprite) SetFrame(frame int) {
	if frame >= 0 {
		s.Anim.animationOffset = s.Rect.Dx() * frame

	} else {
		s.Anim.animationOffset = s.imageSize.ToStd().X - s.Rect.Dx()*int(math.Abs(float64(frame)))

	}

}
func (s *Sprite) GetAnimOffset() int {
	return s.Anim.animationOffset
}

func (s *Sprite) AnimateBackwards() {

	if s.Anim.animationTime == s.Anim.AnimationRate {

		if s.Anim.animationOffset > s.Rect.Dx() {

			s.Anim.animationOffset -= s.Rect.Dx()
		} else {
			if s.Anim.Loop {
				s.Anim.animationOffset = s.imageSize.ToStd().X - s.Rect.Dx()*2
			}
		}
		s.Anim.animationTime = 0
	}
	s.Anim.animationTime++

}

func (s *Sprite) SetOriginAtCenter() error {
	if s.spriteImage == nil {
		return fmt.Errorf("Sprite image is not set")
	}

	if s.Rect != image.Rect(0, 0, s.imageSize.ToStd().X, s.imageSize.ToStd().Y) {
		s.origin = gmath.Vec{X: (float64(s.Rect.Max.X-s.Rect.Min.X) / 2) * s.scale, Y: (float64(s.Rect.Max.Y-s.Rect.Min.Y) / 2) * s.scale}
	} else {

		s.origin = s.GetImageSize().Mulf(0.5)

	}
	return nil

}

func (s *Sprite) MoveOrigin(vec gmath.Vec) {
	s.origin = s.origin.Add(vec)
}
