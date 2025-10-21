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
)

// GCDevicePhysicalInputElementChange - Possible values that describe whether the input value of an element changes.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputElementChange
type GCDevicePhysicalInputElementChange uint

const (
	// GCDevicePhysicalInputElementNoChange - There’s no change to the input value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputElementChange/noChange
	GCDevicePhysicalInputElementNoChange GCDevicePhysicalInputElementChange = 0
)


