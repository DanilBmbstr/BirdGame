package main

/*
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
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	debugtools "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/DebugTools"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/assets"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/camera"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/enemy"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity/player"

	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/inputEvents"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/level"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/optionsLoaders"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

const (
	windowHeingtMin = 1000
	windowWidthMin  = 1000
)

var (
	debugMode       = false
	debugCollisions bool
	windowWidth     int
	windowHeight    int
	playerDTO       optionsLoaders.PlayerDTO
	levelScale      float64 = 0.7
)

type BirdGame struct {
	loader          *resource.Loader
	player          *player.Player
	test_enemy      enemy.BasicFox
	loadedLocations [2]level.Location

	handler *input.Handler
}

func switchFullscreenMode(g *BirdGame) {

	if inputEvents.OnKeyPressed(ebiten.KeyAlt) {
		if !ebiten.IsFullscreen() {
			ebiten.SetFullscreen(true)
			mx, my := ebiten.Monitor().Size()
			windowWidth = mx
			windowHeight = my
		} else {
			ebiten.SetFullscreen(false)
			windowWidth = windowHeingtMin
			windowHeight = windowHeingtMin
		}
	}

}

func (g *BirdGame) init() {
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetVsyncEnabled(true)

	if len(os.Args) > 1 && os.Args[1] == "debug" {
		debugMode = true
	}

	inputEvents.InitInputEvents()
	g.loader = assets.CreateLoader()

	//Инициализация игрока
	g.player = &player.Player{}
	img := g.loader.LoadImage(assets.Bird).Data
	g.player.GetSprite().SetImage(img)
	playerDTO = optionsLoaders.LoadPlayer(g.loader)
	g.player.InitPlayer(1000, 1500, playerDTO)

	//////////////////////////

	g.loadedLocations[0].LoadLocation(g.loader.LoadRaw(assets.Test_level), levelScale)
	g.loadedLocations[1].LoadLocation(g.loader.LoadRaw(assets.Test_level2), levelScale)

	g.test_enemy.Init(1000, 1000, g.loader.LoadImage(assets.BasicFox).Data)

}

func (g *BirdGame) Update() error {

	switchFullscreenMode(g)

	//Обновляем камеру
	camera.X_offset = g.player.GetPosition().X
	camera.Y_offset = g.player.GetPosition().Y
	camera.WindowSizeX = float64(windowWidth)
	camera.WindowSizeY = float64(windowHeight)
	//

	x, y := ebiten.CursorPosition()

	MousePosition := gmath.Vec{X: float64(x), Y: float64(y)}

	g.player.Controller.Rotate(g.player, MousePosition)
	if debugMode {
		debugtools.Debug(g.player, g.loadedLocations[0])
	} else {
		g.player.Controller.MoveForward(g.player)
	}

	for i := 0; i < len(g.loadedLocations); i++ {
		debugCollisions = g.loadedLocations[i].CheckCollisions(g.player.GetPosition())
		if debugCollisions {
			g.player.InitPlayer(1000, 1500, playerDTO)
		}
	}

	g.test_enemy.Move(g.player.GetPosition())

	inputEvents.UpdateInputEvents()
	return nil
}

func (g *BirdGame) Draw(screen *ebiten.Image) {

	imageOptions := ebiten.DrawImageOptions{}
	imageOptions.GeoM.Scale(levelScale, levelScale)
	imageOptions.GeoM.Translate(-(camera.X_offset - camera.WindowSizeX/2), -(camera.Y_offset - camera.WindowSizeY/2))

	screen.DrawImage(g.loader.LoadImage(assets.Back).Data, &imageOptions)
	if debugMode {
		for i := 0; i < len(g.loadedLocations); i++ {
			g.loadedLocations[i].DrawLocation(screen)
		}
	}

	g.player.DrawEntity(screen)
	g.test_enemy.Parent.DrawEntity(screen)
	ebitenutil.DebugPrint(screen, strconv.FormatFloat(ebiten.ActualFPS(), 'f', 6, 64)+"\n"+strconv.FormatBool(debugCollisions))

}

func (g *BirdGame) Layout(w, h int) (int, int) {

	return windowWidth, windowHeight
}

func main() {

	g := &BirdGame{}
	windowHeight = windowHeingtMin
	windowWidth = windowWidthMin
	g.init()

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Ebitengine")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
