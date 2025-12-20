package inputEvents

import "github.com/hajimehoshi/ebiten/v2"

var dict map[ebiten.Key]bool

func InitInputEvents() {
	dict = make(map[ebiten.Key]bool)

	dict[ebiten.KeyF] = false
	dict[ebiten.KeyAlt] = false

}

func OnKeyPressed(key ebiten.Key) bool {
	if ebiten.IsKeyPressed(key) {
		return !dict[key]
	}
	return false
}

func UpdateInputEvents() {
	for k, _ := range dict {
		if !ebiten.IsKeyPressed(k) {
			dict[k] = false

		} else {
			dict[k] = true
		}
	}
}
