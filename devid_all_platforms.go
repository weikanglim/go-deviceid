package deviceid

import (
	"crypto/rand"
	"fmt"
	"os"
	"path"
)

// generateDeviceID generates values in the format of:
// `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`
// Where 'x' is any legal lowercased hex digit.
func generateDeviceID() (string, error) {
	randBytes := make([]byte, 4+2+2+2+6)

	if _, err := rand.Read(randBytes); err != nil {
		return "", err
	}

	return formatGUID(randBytes), nil
}

// formatGUID takes 16 bytes and formats it into a lowercased GUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
// NOTE: there's no error checking here, I just split this out from generateDeviceID so we can unit test it.
func formatGUID(randBytes []byte) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		randBytes[0:4],
		randBytes[4:6],
		randBytes[6:8],
		randBytes[8:10],
		randBytes[10:])
}

// readWriteDeviceIDFile reads a deviceid from a file in dir + "/deviceid" and returns it.
// If the file doesn't exist it creates the file with a newly generated device id and returns the new deviceid.
func readWriteDeviceIDFile(dir string) (string, error) {
	err := os.MkdirAll(dir, 0700)

	if err != nil {
		return "", err
	}

	filePath := path.Join(dir, "deviceid")

	contents, err := os.ReadFile(filePath)

	// TODO: scrub any PII from errors.
	if os.IsNotExist(err) {
		deviceID, err := generateDeviceID()

		if err != nil {
			return "", err
		}

		if err := os.WriteFile(filePath, []byte(deviceID), 0600); err != nil {
			return "", err
		}

		return deviceID, nil
	} else if err != nil {
		return "", err
	}

	return string(contents), nil
}
