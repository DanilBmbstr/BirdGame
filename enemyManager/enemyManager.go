package enemymanager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/enemy"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/entity"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"
	resource "github.com/quasilyte/ebitengine-resource"
)

var (
	chaserEnemyes []entity.Chaser = []entity.Chaser{&enemy.BasicFox{}, &enemy.BasicFox{}}
	enemyes       []entity.Enemy

	rLoader *resource.Loader
)

func InitEnemies(loader *resource.Loader) {

	rLoader = loader

	for _, i := range chaserEnemyes {
		i.Init(1500, 1500, loader.LoadImage(assets.BasicFox).Data)
	}

}

func Update(player *player.Player) {
	for i := range chaserEnemyes {
		for j := range chaserEnemyes {

			if collisionsystem.CheckShapesCollision(*chaserEnemyes[i].GetCollider(), *chaserEnemyes[j].GetCollider()) && i != j {

				chaserEnemyes[i].PassBy(chaserEnemyes[j])

			}
		}
		if chaserEnemyes[i].IsDestroyed() {
			chaserEnemyes[i] = &enemy.BasicFox{}
			chaserEnemyes[i].Init(1500, 1500, rLoader.LoadImage(assets.BasicFox).Data)
		}
		chaserEnemyes[i].Update(player)
	}

}

func Draw(screen *ebiten.Image) {
	for _, i := range chaserEnemyes {
		i.DrawEntity(screen)
	}
}
