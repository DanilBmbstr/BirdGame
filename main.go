package main

/*░░░░░░░░░░░░░░░░░░▄▄░░░░░░░░░░░░░░░░░░
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
	"github.com/hajimehoshi/ebiten/v2/cmd/mygame/assets"
	"github.com/hajimehoshi/ebiten/v2/cmd/mygame/entity"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

type myGame struct {
	windowWidth  int
	windowHeight int
	loader       *resource.Loader
	player       *entity.Player

	inputSystem input.System
	handler     *input.Handler
}

func (g *myGame) init() {
	g.loader = assets.CreateLoader()
	//imgA := g.loader.LoadImage(assets.ImageA).Data
	g.player = &entity.Player{TailLenght: 50, Speed: 2}
	img := g.loader.LoadImage(assets.Bird).Data
	g.player.PlayerSprite.SetImage(img)
	g.player.PlayerSprite.SetScale(0.1)
	g.player.PlayerSprite.Origin = g.player.PlayerSprite.GetImageSize().Mulf(0.5).Add(gmath.Vec{X: 0, Y: -8})
	g.player.TailSegmentLenght = 5

	//При инициализации игрока создаётся хвост, необходимо выставить параметры влияющие на его создание до вызова InitPlayer
	g.player.InitPlayer(250, 250)

	g.player.SetDirection(1, 1)
	g.player.Lenght = 100
	g.player.PlayerTail.MovingSpeed = 2
	g.player.PlayerTail.RotationSpeed = 0.25

	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	g.handler = g.inputSystem.NewHandler(0, input.Keymap{})

}

func (g *myGame) Update() error {

	//g.player.Center.X += 16 * (1.0 / 60.0)
	x, y := ebiten.CursorPosition()

	MousePosition := gmath.Vec{X: float64(x), Y: float64(y)}

	g.player.RotateTowards(MousePosition, 0.1)

	g.player.MoveForward()

	//g.player.SetPosition(float64(x), float64(y))
	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	g.player.DrawEntity(screen)
	ebitenutil.DebugPrint(screen, strconv.FormatFloat(ebiten.ActualFPS(), 'f', 6, 64))

	/*


		vector.FillCircle(screen, 60, 70, 50, color.RGBA{255, 255, 255, 255}, true)
		path := &vector.Path{}

		// Define a triangle
		path.MoveTo(50, 50)
		path.LineTo(150, 50)
		path.LineTo(100, 150)
		path.Close() // Close the path to form a complete shape
		cs := ebiten.ColorScale{}
		cs.Scale(1, 0.5, 1, 1)

		img := g.player.Img

		var options ebiten.DrawImageOptions
		options.GeoM.Translate(g.player.Pos.X, g.player.Pos.Y)
		screen.DrawImage(img, &options)

		vector.FillPath(screen, path, &vector.FillOptions{FillRule: vector.FillRuleEvenOdd}, &vector.DrawPathOptions{AntiAlias: true,
			ColorScale: cs})
	*/
}

func (g *myGame) Layout(w, h int) (int, int) {

	return g.windowWidth, g.windowHeight
}

func main() {
	g := &myGame{
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
