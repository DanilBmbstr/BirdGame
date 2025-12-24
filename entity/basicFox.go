package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

type BasicFox struct {
	Parent Enemy
	Speed  float64
}

func (bs *BasicFox) Init(x, y float64, img *ebiten.Image) {
	bs.Speed = 5
	sprite := Sprite{Position: gmath.Vec{X: x, Y: y}}
	sprite.SetOriginAtCenter()
	sprite.SetImage(img)
	sprite.SetScale(0.3)
	bs.Parent.Init(x, y, sprite)
}

func (bs *BasicFox) Move(pos gmath.Vec) {
	bs.Parent.MoveTowards(pos, bs.Speed)

}
