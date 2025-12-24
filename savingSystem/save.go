package savingsystem

import "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity"

var (
	slot gameSave
)

type gameSave struct {
	player entity.Player
}

func performSaving(gs gameSave) {

	slot = gameSave{
		gs.player,
	}
}
