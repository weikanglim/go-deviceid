package devid

import (
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDevID(t *testing.T) {
	devID, err := GetDeviceID()
	require.NoError(t, err)
	require.NotEmpty(t, devID)

	secondDevID, err := GetDeviceID()
	require.NoError(t, err)
	require.Equal(t, devID, secondDevID)
}

func TestGenerateDevID(t *testing.T) {
	t.Run("BasicRandomnessCheck", func(t *testing.T) {
		allIDs := map[string]bool{}

		id, err := generateDeviceID()
		require.NoError(t, err)
		require.NotEmpty(t, id)

		allIDs[id] = true

		for i := 0; i < 100; i++ {
			id2, err := generateDeviceID()
			require.NoError(t, err)
			require.False(t, allIDs[id2])
			allIDs[id2] = true
		}
	})

	t.Run("ProperlyFormattedGUID", func(t *testing.T) {
		id, err := generateDeviceID()
		require.NoError(t, err)

		t.Logf("Generated ID = %s", id)
		requireValidGUID(t, id)
	})
}

func TestGenerateDeviceIDFile(t *testing.T) {
	now := time.Now().UnixNano()

	tempRoot := path.Join(os.TempDir(), strconv.FormatInt(now, 10))
	defer os.RemoveAll(tempRoot)

	defer func() {
		err := os.RemoveAll(tempRoot)
		require.NoError(t, err)
	}()

	t.Run("CreateDirAndFile", func(t *testing.T) {
		devID, err := readWriteDeviceIDFile(tempRoot)
		require.NoError(t, err)
		requireValidGUID(t, devID)
	})

	t.Run("FileAlreadyExists", func(t *testing.T) {
		devID, err := readWriteDeviceIDFile(tempRoot)
		require.NoError(t, err)
		requireValidGUID(t, devID)

		cachedDevID, err := readWriteDeviceIDFile(tempRoot)
		require.NoError(t, err)
		require.Equal(t, devID, cachedDevID)
	})

	t.Run("DirAlreadyExistsButNoFile", func(t *testing.T) {
		origDevID, err := readWriteDeviceIDFile(tempRoot)
		require.NoError(t, err)
		requireValidGUID(t, origDevID)

		err = os.Remove(path.Join(tempRoot, "deviceid"))
		require.NoError(t, err)

		newDevID, err := readWriteDeviceIDFile(tempRoot)
		require.NoError(t, err)
		requireValidGUID(t, newDevID)

		// and this is a new ID
		require.NotEqual(t, origDevID, newDevID)
	})

}

func requireValidGUID(t *testing.T, id string) {
	parts := strings.Split(id, "-")
	require.Equal(t, 5, len(parts))

	// 8-4-4-4-12
	require.Equal(t, 8, len(parts[0]))
	require.Equal(t, 4, len(parts[1]))
	require.Equal(t, 4, len(parts[2]))
	require.Equal(t, 4, len(parts[3]))
	require.Equal(t, 12, len(parts[4]))

	// all lowercased hex
	for _, part := range parts {
		for _, c := range part {
			require.True(t, (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f'))
		}
	}
}
