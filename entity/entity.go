package entity

import "github.com/hajimehoshi/ebiten/v2"

/*
⢠⣶⠶⢶⡆⠀⠀⠐⢶⣆⢠⣶⠆⢠⣶⠿⢷⣦⡀⢰⡆⠀⢰⣶⠀⢶⣆⢀⣶⠆
⠘⢿⣦⣼⡇⠀⠀⠀⠈⣿⣿⡃⠀⣿⡇⠀⠀⣿⡇⢸⣷⣤⣼⣿⠀⠈⣿⣾⡟⠀
⢠⣿⠃⢸⡇⠀⠀⢀⣼⡟⠹⣷⡀⠻⣷⣤⣴⡿⠃⠀⠉⠉⢹⣿⠀⣤⣼⡟⠀⠀
⢀⣴⡶⠶⢶⠀⢰⡶⠶⣶⣄⠀⠀⢠⣶⣶⡀⠀⠰⠶⢶⣶⠶⠦⢰⣶⠀⠀⠀⠀
⣾⡏⠀⠀⠀⠀⢸⡇⠀⣸⣿⠀⠀⣾⡇⢻⣧⠀⠀⠀⢸⣿⠀⠀⢸⣿⣤⣤⣄⡀
⢿⣇⠀⠀⠀⠀⢸⡟⠛⠛⠁⠀⢰⣿⣤⣼⣿⡆⠀⠀⢸⣿⠀⠀⢸⣿⠁⠈⢹⣷
⠈⠻⠷⠶⠾⠀⠸⠇⠀⠀⠀⠠⠿⠃⠀⠀⠸⠿⠀⠀⠸⠿⠀⠀⠸⠿⠶⠶⠿
*/
type Entity interface {
	SetPosition(x, y float64)
	DrawEntity(screen *ebiten.Image)
}
