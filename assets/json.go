package assets

import (
	_ "os"

	resource "github.com/quasilyte/ebitengine-resource"
)

const (
	JsonNone resource.RawID = iota
	PlayerSettings
)

func RegisterJsonResources(loader *resource.Loader) {
	jsonResources := map[resource.RawID]resource.RawInfo{
		PlayerSettings: {Path: "settings/playerSettings.json"},
	}

	for id, res := range jsonResources {
		loader.RawRegistry.Set(id, res)
	}
}
