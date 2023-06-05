package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alcmoraes/arch-steam-utils/pkg"
)

func main() {
	filePath := os.Args[1]
	addToSteam(filePath)
}

func addToSteam(filePath string) {

	fileName := filepath.Base(filePath)
	fileDir := filepath.Dir(filePath)

	fmt.Println(fileName)
	fmt.Println(fileDir)

	vdf := pkg.GetShortcuts()
	vdf.AddEntry(fileName, filePath, fileDir)
	vdf.Save()

}
