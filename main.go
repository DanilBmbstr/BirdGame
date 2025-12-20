package main

/*
░░░░░░░░░░░░░░░░░░▄▄░░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░░▄██▄░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░▄████▄░░░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░██████░░░░░░░░░░░░░░░░
███▄▄░░░░░░░░░░████████░░░░░░░░░░▄▄▄██
▀██████▄▄░░░░░░████████░░░░░░▄▄██████▀
░█████████▄░░░░████████░░░░▄█████████░
░░▀▀▀▀▀▀▀▀▀▀░░░▀▀▀▀▀▀▀▀░░░▀▀▀▀▀▀▀▀▀▀░░
░░░▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄░░░
░░░░▀████████████████████████████▀░░░░
░░░░░░▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄░░░░░░
░░░░░░░▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀░░░░░░░
░░░░░░░░░░░▄▄▄▄▄▄░▄▄░▄▄▄▄▄▄░░░░░░░░░░░
░░░░░░░░░░░░░░▀▀▀░▀▀░▀▀▀░░░░░░░░░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░░░░░░░░░██░▀▀░░░░░░██░░░░░░░░░░░░░
░▄█▀▀██░▄█▀▀██░██░▄█▀▀███░▄█▀▀██░█▀▀█▄
░█░░░██░█░░░██░██░█░░░███░█░░░██░▀███▄
░▀█▄▄██░▀█▄▄██░██░▀█▄▄███░▀█▄▄██░█▄▄█▀*/

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/camera"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/level"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/optionsLoaders"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

var debugCollisions bool

type BirdGame struct {
	windowWidth     int
	windowHeight    int
	loader          *resource.Loader
	player          *entity.Player
	loadedLocations [2]level.Location
	inputSystem     input.System
	handler         *input.Handler
}

func (g *BirdGame) init() {

	g.loader = assets.CreateLoader()

	//Инициализация игрока
	g.player = &entity.Player{}
	img := g.loader.LoadImage(assets.Bird).Data
	g.player.GetSprite().SetImage(img)
	playerDTO := optionsLoaders.LoadPlayer(g.loader)
	g.player.InitPlayer(250, 250, playerDTO)
	g.player.GetSprite().Origin = g.player.GetSprite().GetImageSize().Mulf(0.5).Add(gmath.Vec{X: 0, Y: -8})
	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	//////////////////////////

	g.loadedLocations[0].LoadLocation(g.loader.LoadRaw(assets.Test_level))
	g.loadedLocations[1].LoadLocation(g.loader.LoadRaw(assets.Test_level2))
	g.handler = g.inputSystem.NewHandler(0, input.Keymap{})

}

func (g *BirdGame) Update() error {
	//Обновляем камеру
	camera.X_offset = g.player.GetPosition().X
	camera.Y_offset = g.player.GetPosition().Y
	camera.WindowSizeX = float64(g.windowWidth)
	camera.WindowSizeY = float64(g.windowHeight)
	//

	x, y := ebiten.CursorPosition()

	MousePosition := gmath.Vec{X: float64(x), Y: float64(y)}

	g.player.Controller.Rotate(g.player, MousePosition)

	g.player.Controller.MoveForward(g.player)

	for i := 0; i < len(g.loadedLocations); i++ {
		debugCollisions = g.loadedLocations[i].CheckCollisions(g.player.GetPosition())
		if debugCollisions {
			break
		}
	}

	return nil
}

func (g *BirdGame) Draw(screen *ebiten.Image) {

	for i := 0; i < len(g.loadedLocations); i++ {
		g.loadedLocations[i].DrawLocation(screen)
	}

	g.player.DrawEntity(screen)

	ebitenutil.DebugPrint(screen, strconv.FormatFloat(ebiten.ActualFPS(), 'f', 6, 64)+"\n"+strconv.FormatBool(debugCollisions))

}

func (g *BirdGame) Layout(w, h int) (int, int) {

	return g.windowWidth, g.windowHeight
}

func main() {
	g := &BirdGame{
		windowWidth:  1000,
		windowHeight: 1000,
	}

	g.init()

	ebiten.SetWindowSize(g.windowWidth, g.windowHeight)
	ebiten.SetWindowTitle("Ebitengine")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
