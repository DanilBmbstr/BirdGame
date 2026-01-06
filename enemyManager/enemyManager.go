package enemymanager

import (
	"github.com/hajimehoshi/ebiten/v2"
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/enemy"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/entity"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/level"
	resource "github.com/quasilyte/ebitengine-resource"
)

var (
	enemyes []entity.Enemy = []entity.Enemy{&enemy.BasicFox{}, &enemy.BasicFox{}, &enemy.Ameba{}, &enemy.Ameba{}, &enemy.Ameba{}, &enemy.Ameba{}, &enemy.Ameba{}}

	rLoader *resource.Loader
)

func InitEnemies(loader *resource.Loader) {

	rLoader = loader

	for _, i := range enemyes {
		i.Init(1500, 1500, loader)
	}

}

func Update(player *player.Player, location *level.CollisionMask) {

	for i := range enemyes {
		if enemyes[i].IsDestroyed() {
			enemyes = append(enemyes[:i], enemyes[i+1:]...)
			break
		}
	}

	for i := range enemyes {
		for j := range enemyes {

			if collisionsystem.CheckShapesCollision(*enemyes[i].GetCollider(), *enemyes[j].GetCollider()) && i != j {

				enemyes[i].PassBy(enemyes[j])

			}
		}

		enemyes[i].Update(player, location)
	}

}

func Draw(screen *ebiten.Image) {
	for _, i := range enemyes {
		i.DrawEntity(screen)
		i.GetCollider().DrawEntity(screen)
	}

}
