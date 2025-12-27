package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/inputEvents"
)

type PlayerBattleSystem struct {
	HitTime      int
	hitMaxTime   int
	cooldownTime int
	cooldown     int
}

func (pbs *PlayerBattleSystem) Init() {
	pbs.HitTime = 1
	pbs.hitMaxTime = 30
	pbs.cooldownTime = 40
	pbs.cooldown = 0
}

func (pbs *PlayerBattleSystem) Update() {
	if inputEvents.OnKeyPressed(ebiten.KeySpace) && pbs.HitTime == 1 && pbs.cooldown == 0 {
		pbs.cooldown = pbs.cooldownTime
		pbs.HitTime = pbs.hitMaxTime
	}
	if pbs.HitTime > 1 {
		pbs.HitTime--
	}
	if pbs.cooldown > 0 {
		pbs.cooldown--
	}

}
