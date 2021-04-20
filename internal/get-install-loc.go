package internal

import (
	"os"
	"path/filepath"
	"runtime"
)

// GetInstallLocation() to get the install location fo binary
func GetInstallLocation() string {
	var tokenDir string
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		tokenDir = os.Getenv("APPDATA")
	} else {
		tokenDir = filepath.Join(home, ".config")
	}
	dname := filepath.Join(tokenDir, "gtodo")
	return dname
}
