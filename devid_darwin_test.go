//go:build darwin

package deviceid

import (
	"os"
	"path"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDeviceID_darwin(t *testing.T) {
	t.Run("HOME", func(t *testing.T) {
		homeDir := path.Join(os.TempDir(), strconv.FormatInt(time.Now().UnixNano(), 10), "home")

		t.Setenv("HOME", homeDir)

		id, err := deviceID()
		require.NoError(t, err)
		requireValidGUID(t, id)

		// validate it went to the right spot
		bytes, err := os.ReadFile(path.Join(homeDir, "Library/Application Support/Microsoft/DeveloperTools", "deviceid"))
		require.NoError(t, err)
		require.Equal(t, id, string(bytes))
	})

	t.Run("HOME is not set", func(t *testing.T) {
		t.Setenv("HOME", "")
		id, err := deviceID()
		require.Empty(t, id)
		require.EqualError(t, err, "environment variable HOME is not set")
	})
}
