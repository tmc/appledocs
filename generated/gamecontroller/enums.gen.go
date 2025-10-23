// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

// Enum types and constants
// GCDeviceBatteryState - A state that indicates whether a device’s battery has power and is charging.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/State
type GCDeviceBatteryState uint

const (
	// GCDeviceBatteryStateCharging - The device’s battery has power and is charging, but isn’t fully charged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/State/charging
	GCDeviceBatteryStateCharging GCDeviceBatteryState = 0
	// GCDeviceBatteryStateDischarging - The device’s battery is discharging.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/State/discharging
	GCDeviceBatteryStateDischarging GCDeviceBatteryState = 0
	// GCDeviceBatteryStateFull - The device’s battery has power and is fully charged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/State/full
	GCDeviceBatteryStateFull GCDeviceBatteryState = 0
	// GCDeviceBatteryStateUnknown - The state of the device’s battery is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/State/unknown
	GCDeviceBatteryStateUnknown GCDeviceBatteryState = 0
)


