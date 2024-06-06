# Device ID for Go

`deviceid` provides a device ID for a given system, based on the `DevDeviceId` specification.

## Installation

`go get github.com/richardpark-msft/go-deviceid`

## Usage

```golang
import deviceid "github.com/richardpark-msft/go-deviceid"

deviceId, err := deviceid.Get()
if err != nil {
  // handle error
}

fmt.Println("Device ID is: ", deviceId)
```
