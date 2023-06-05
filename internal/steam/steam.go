package steam

import (
	"os"
	"path"
	"strconv"
)

var COMPATDATA = path.Join(os.Getenv("HOME"), ".local/share/Steam/steamapps/compatdata")

// Get UserID
func GetUserID() string {
	const steam_path = ".local/share/Steam/userdata"
	h := os.Getenv("HOME")
	entries, _ := os.ReadDir(path.Join(h, steam_path))
	if len(entries) == 0 {
		panic("No Steam User ID found")
	}
	return entries[0].Name()
}

// Get Next AppID
func GetNextAppID() uint32 {
	d, _ := os.ReadDir(COMPATDATA)
	gtN := 10
	for _, v := range d {
		if asN, err := strconv.Atoi(v.Name()); err == nil {
			if asN > gtN {
				gtN = asN
			}
		}
	}
	return uint32(gtN * 10)
}
