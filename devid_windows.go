//go:build windows

package devid

import (
	"errors"

	"golang.org/x/sys/windows/registry"
)

func GetDeviceID() (string, error) {
	return getDeviceIDImpl(devToolsSubPath)
}

const devToolsSubPath = `SOFTWARE\Microsoft\DeveloperTools`
const deviceIDValueName = "deviceid"

func getDeviceIDImpl(subKeyPath string) (string, error) {
	// 1.1 Windows
	// * The value is cached in the 64-bit Windows Registry under HKeyCurrentUser\SOFTWARE\Microsoft\DeveloperTools.
	// * The key should be named 'deviceid' and should be of type REG_SZ (String value).
	// * The value should be stored in plain text and in the format specified in Section 1.
	key, _, err := registry.CreateKey(registry.CURRENT_USER, subKeyPath, registry.READ|registry.WRITE)

	if err != nil {
		return "", err
	}

	defer key.Close()

	value, _, err := key.GetStringValue(deviceIDValueName)

	if err == nil {
		return value, nil
	}

	if !errors.Is(err, registry.ErrNotExist) {
		return "", err
	}

	devID, err := generateDeviceID()

	if err != nil {
		return "", err
	}

	if err := key.SetStringValue(deviceIDValueName, devID); err != nil {
		return "", err
	}

	return devID, nil
}
