package devices

import (
	"encoding/json"
	"fmt"
	"os"
)

type DeviceModel string

// GetDevicesMap reads the devEUI -> device model registry from a JSON file.
func GetDevicesMap(path string) (map[string]DeviceModel, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read device registry file: %w", err)
	}

	var devicesEUI map[string]DeviceModel
	if err := json.Unmarshal(contents, &devicesEUI); err != nil {
		return nil, fmt.Errorf("parse device registry file: %w", err)
	}

	return devicesEUI, nil
}
