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
	collisionsystem "github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/collisionSystem"
	"github.com/hajimehoshi/ebiten/v2/cmd/BirdGame/entity"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

var debugCollisions bool

type BirdGame struct {
	windowWidth  int
	windowHeight int
	loader       *resource.Loader
	player       *entity.Player

	inputSystem input.System
	handler     *input.Handler
	testShape   collisionsystem.Shape
}

func (g *BirdGame) init() {
	g.loader = assets.CreateLoader()
	//imgA := g.loader.LoadImage(assets.ImageA).Data
	g.player = &entity.Player{TailLenght: 50, Speed: 4}
	img := g.loader.LoadImage(assets.Bird).Data
	g.player.PlayerSprite.SetImage(img)
	g.player.PlayerSprite.SetScale(0.1)
	g.player.PlayerSprite.Origin = g.player.PlayerSprite.GetImageSize().Mulf(0.5).Add(gmath.Vec{X: 0, Y: -8})
	g.player.TailSegmentLenght = 5

	cs := ebiten.ColorScale{}
	cs.Scale(0.5, 0.5, 0.5, 0.5)
	g.player.TailColor = cs
	//При инициализации игрока создаётся хвост, необходимо выставить параметры влияющие на его создание до вызова InitPlayer
	g.player.InitPlayer(250, 250)

	g.player.SetDirection(1, 1)
	g.player.Lenght = 100
	g.player.PlayerTail.MovingSpeed = 4
	g.player.PlayerTail.RotationSpeed = 0.25

	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	g.handler = g.inputSystem.NewHandler(0, input.Keymap{})

	g.testShape.Points = append(g.testShape.Points, gmath.Vec{X: 500, Y: 500})
	g.testShape.Points = append(g.testShape.Points, gmath.Vec{X: 300, Y: 750})
	g.testShape.Points = append(g.testShape.Points, gmath.Vec{X: 450, Y: 850})
	g.testShape.Points = append(g.testShape.Points, gmath.Vec{X: 700, Y: 700})

}

func (g *BirdGame) Update() error {
	//Обновляем камеру
	camera.X_offset = g.player.Center.X
	camera.Y_offset = g.player.Center.Y
	camera.WindowSizeX = float64(g.windowWidth)
	camera.WindowSizeY = float64(g.windowHeight)
	//

	x, y := ebiten.CursorPosition()

	MousePosition := gmath.Vec{X: float64(x), Y: float64(y)}

	g.player.Rotate(MousePosition, 0.1)

	g.player.MoveForward()
	debugCollisions = collisionsystem.CheckCollisionWithShape(g.player.Center, g.testShape)
	return nil
}

func (g *BirdGame) Draw(screen *ebiten.Image) {

	g.testShape.DrawEntity(screen)

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
