package pkg

import (
	"fmt"
	"os"
	"path"

	"github.com/wakeful-cloud/vdf"
)

var SHORTCUTS = path.Join(os.Getenv("HOME"), fmt.Sprintf(".local/share/Steam/userdata/%s/config/shortcuts.vdf", GetUserID()))

type VDF struct {
	vdfMap vdf.Map
}

func (v *VDF) AddEntry(name, filePath, baseDir string) {

	newEntry := vdf.Map{
		"appid":               GetNextAppID(),
		"AllowDesktopConfig":  uint32(1),
		"AllowOverlay":        uint32(1),
		"AppName":             name,
		"Devkit":              uint32(0),
		"DevkitOverrideAppID": uint32(0),
		"Exe":                 fmt.Sprintf("\"%s\"", filePath),
		"IsHidden":            uint32(0),
		"LastPlayTime":        uint32(0),
		"OpenVR":              uint32(0),
		"StartDir":            fmt.Sprintf("%s/", baseDir),
	}

	shortcuts := v.vdfMap["shortcuts"].(vdf.Map)
	shortcuts[fmt.Sprintf("%d", len(shortcuts))] = newEntry
	v.vdfMap["shortcuts"] = shortcuts

}

func (v *VDF) Save() {
	rawVdf, err := vdf.WriteVdf(v.vdfMap)

	if err != nil {
		panic(err)
	}

	//Write the file
	err = os.WriteFile(SHORTCUTS, rawVdf, 0666)

	if err != nil {
		panic(err)
	}
}

func GetShortcuts() *VDF {
	//Read the file
	bytes, err := os.ReadFile(SHORTCUTS)

	if err != nil {
		panic(err)
	}

	//Read VDF
	vdfMap, err := vdf.ReadVdf(bytes)

	if err != nil {
		panic(err)
	}

	return &VDF{vdfMap: vdfMap}
}
