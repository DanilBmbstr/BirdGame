package assets

/*.
⠄⠄⠄⠄⠄⠄⢀⣀⣤⣤⣤⢤⣤⣤⣀⡀⠄⠄⠄⠄⠄⠄⠄⠄
⠄⠄⠄⣠⣴⣿⣿⣿⣿⣿⣿⡇⠄⢸⣿⣿⣷⡀⠄⠄⠄⠄⠄⠄
⠄⠄⠈⣿⣿⣿⣿⣿⣿⣿⣿⣧⠄⢸⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄
⠄⠄⠄⢸⣿⣿⣿⣿⣿⣿⣿⣿⡄⢸⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄
⠄⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⡇⢸⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄
⠄⠄⠄⠄⢸⣿⣿⣿⣿⡿⠿⠿⠿⠾⠛⠛⣉⣇⣀⠄⠄⢀⡀⠄
⠄⠄⠄⠄⢈⣇⠄⣀⣀⣠⣤⣤⣴⡶⢚⣫⣭⣷⣶⣿⣿⣿⣿⠇
⠄⣠⣶⣿⣿⣿⣿⣿⡿⢟⣛⣭⣶⢿⣿⠟⠻⠭⠛⡟⠛⠉⠄⠄
⠸⠿⢿⣿⣿⣭⣿⠶⠾⡟⠛⣉⠘⡀⠄⠄⠄⠄⠄⡸⠄⠄⠄⠄
⠄⠄⠄⠄⠄⢸⠄⠄⠄⣰⣿⣿⣷⠈⠒⢒⢊⠤⠊⠁⠄⠄⠄⠄
⠄⠄⠄⠄⠄⢿⣾⣷⢶⡹⠛⠛⠁⠄⠄⢣⠼⡄⠄⠄⠄⠄⠄⠄
⠄⠄⠄⠄⠄⠄⢹⠋⢸⡿⠶⠄⢹⡶⢖⠉⠄⠄⠄⠄⠄⠄⠄⠄
⠄⠄⠄⠄⠄⠄⠘⠤⢤⣤⠄⠄⢨⢧⠸⡀⠄⠄⠄⠄⠄⠄⠄⠄
⠄⠄⠄⠄⠄⠄⠄⠄⠄⢡⢻⣙⣵⡿⡆⠱⡀⠄⠄⠄⠄⠄⠄⠄
⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢣⡉⠁⠄⠁⡠⠁⠄⠄⠄⠄⠄⠄⠄
⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⡇⠄⡸⠋⠄⠄⠄⠄⠄⠄⠄⠄⠄
⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠉⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄*/
import (
	"embed"
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio"

	resource "github.com/quasilyte/ebitengine-resource"
)

//go:embed all:_data
var gameAssets embed.FS

func OpenAsset(path string) io.ReadCloser {
	// Функция OpenAsset могла бы работать как с данными внутри бинарника,
	// так и с внешними. Для этого ей нужно распознавать ресурс по его пути.
	// Самым простым вариантом является использование префиксов в пути,
	// типа "$music/filename.ogg" вместо "filename.ogg", когда мы ищем
	// файл во внешнем каталоге (а не в бинарнике).
	//
	// Но на данном этапе у нас только один источник ассетов - бинарник.
	f, err := gameAssets.Open("_data/" + path)
	if err != nil {
		panic(err)
	}
	return f
}

func CreateLoader() *resource.Loader {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = OpenAsset
	RegisterResources(loader)
	return loader
}

func RegisterResources(loader *resource.Loader) {
	registerImageResources(loader)
}
