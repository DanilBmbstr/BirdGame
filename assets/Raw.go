package assets

import (
	_ "os"

	resource "github.com/quasilyte/ebitengine-resource"
)

const (
	RawNone resource.RawID = iota
	PlayerSettings
	Test_level
	Test_level2
)

func RegisterRawResources(loader *resource.Loader) {
	rawResources := map[resource.RawID]resource.RawInfo{
		PlayerSettings: {Path: "settings/playerSettings.json"},
		Test_level:     {Path: "levels/test_level.xml"},
		Test_level2:    {Path: "levels/test_level2.xml"},
	}

	for id, res := range rawResources {
		loader.RawRegistry.Set(id, res)
	}
}
