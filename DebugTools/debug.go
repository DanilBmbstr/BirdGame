package debugtools

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/level"
)

func Debug(player *entity.Player, location level.Location) {
	location.DebugMoveCollisions()
	if ebiten.IsKeyPressed(ebiten.KeyControl) {
		player.Controller.MoveForward(player)
	}
}
