package internal

import (
    "log"
    "github.com/kardianos/osext"
)

// GetInstallLocation() to get the install location fo binary
func GetInstallLocation() string {
    dirpath, err := osext.ExecutableFolder()
    if err != nil {
        log.Fatal(err)
    }
    return dirpath
}
