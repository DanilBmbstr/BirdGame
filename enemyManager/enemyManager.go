package enemymanager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/enemy"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/entity"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/level"
	resource "github.com/quasilyte/ebitengine-resource"
)

var (
	enemyes []entity.Enemy

	rLoader *resource.Loader
)

func InitEnemies(loader *resource.Loader) {
	enemyes = []entity.Enemy{&enemy.Ameba{}, &enemy.BasicFox{}}
	frog := &enemy.Frog{}
	frog.SetArea(1000, 1000, 2500, 2500)
	enemyes = append(enemyes, frog)
	frog2 := &enemy.Frog{}
	frog2.SetArea(1000, 1000, 2500, 2500)
	enemyes = append(enemyes, frog2)

	rLoader = loader

	for _, i := range enemyes {
		i.Init(1500, 1500, loader)
	}

}

func ReInit() {
	InitEnemies(rLoader)
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

			if i != j {

				enemyes[i].PassBy(enemyes[j])

			}
		}

		enemyes[i].Update(player, location)
	}

}

func Draw(screen *ebiten.Image) {
	for _, i := range enemyes {
		i.DrawEntity(screen)
		//i.GetCollider().DrawEntity(screen)
	}

}
