package main

import (
	"os"
	"path/filepath"

	"github.com/alcmoraes/arch-steam-utils/internal/vdf"
)

func main() {
	filePath := os.Args[1]

	fileName := filepath.Base(filePath)
	fileDir := filepath.Dir(filePath)

	vdf := vdf.GetShortcuts()
	vdf.AddEntry(fileName, filePath, fileDir)
	vdf.Save()
}
