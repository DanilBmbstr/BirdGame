package enemymanager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/enemy"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/entity"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"
	resource "github.com/quasilyte/ebitengine-resource"
)

var (
	enemies []entity.Enemy = []entity.Enemy{&enemy.BasicFox{}, &enemy.BasicFox{}}

	rLoader *resource.Loader
)

func InitEnemies(loader *resource.Loader) {
	rLoader = loader
	enemies[0].Init(1500, 1500, loader.LoadImage(assets.BasicFox).Data)
	enemies[1].Init(3000, 1500, loader.LoadImage(assets.BasicFox).Data)

}

func Update(player *player.Player) {
	for i := range enemies {
		if enemies[i].IsDestroyed() {
			enemies[i] = &enemy.BasicFox{}
			enemies[i].Init(1500, 1500, rLoader.LoadImage(assets.BasicFox).Data)
		}
		enemies[i].Update(player)
	}
}

func Draw(screen *ebiten.Image) {
	enemies[0].DrawEntity(screen)
	enemies[1].DrawEntity(screen)
}
