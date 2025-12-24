package savingsystem

import "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"

var (
	slot gameSave
)

type gameSave struct {
	player player.Player
}

func performSaving(gs gameSave) {

	slot = gameSave{
		gs.player,
	}
}
