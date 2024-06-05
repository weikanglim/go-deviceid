//go:build darwin

package devid

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetDeviceID_darwin(t *testing.T) {
	t.Run("HOME", func(t *testing.T) {
		homeDir := filepath.Join(os.TempDir(), strconv.FormatInt(time.Now().UnixNano(), 10), "home")

		t.Setenv("HOME", homeDir)

		deviceID, err := GetDeviceID()
		require.NoError(t, err)
		requireValidGUID(t, deviceID)

		// validate it went to the right spot
		bytes, err := os.ReadFile(filepath.Join(homeDir, "Library/Application Support/Microsoft/DeveloperTools", "deviceid"))
		require.NoError(t, err)
		require.Equal(t, deviceID, string(bytes))
	})

	t.Run("HOME is not set", func(t *testing.T) {
		t.Setenv("HOME", "")
		deviceID, err := GetDeviceID()
		require.Empty(t, deviceID)
		require.EqualError(t, err, "environment variable HOME is not set")
	})
}
