# Device ID for Go

`devid` provides a device ID for a given system, based on the `DevDeviceId` specification.

## Installation

`go get github.com/richardpark-msft/go-deviceid`

## Usage

```golang
import devid "github.com/richardpark-msft/go-deviceid"

deviceId, err := devid.DeviceID()
if err != nil {
  // handle error
}

fmt.Println("Device ID is: ", deviceId)
```
