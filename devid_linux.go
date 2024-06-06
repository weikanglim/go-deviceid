//go:build linux

package deviceid

import (
	"fmt"
	"os"
	"path"
)

// 1.2 Linux
// * The folder path will be <RootPath>/Microsoft/DeveloperTools where <RootPath> is $XDG_CACHE_HOME if it is set and not empty, else use $HOME/.cache.
// * The file will be called 'deviceid'.
// * The value should be stored in plain text, UTF-8, and in the format specified in Section 1.

func deviceID() (string, error) {
	xdgCacheHome := os.Getenv("XDG_CACHE_HOME")
	home := os.Getenv("HOME")

	const devToolsSubPath = `Microsoft/DeveloperTools`

	switch {
	case xdgCacheHome != "":
		dir := path.Join(xdgCacheHome, devToolsSubPath)
		return readWriteDeviceIDFile(dir)
	case home != "":
		dir := path.Join(home, ".cache", devToolsSubPath)
		return readWriteDeviceIDFile(dir)
	default:
		return "", fmt.Errorf("neither XDG_CACHE_HOME or HOME are set")
	}
}
