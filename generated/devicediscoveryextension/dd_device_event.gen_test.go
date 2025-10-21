// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

package devicediscoveryextension_test

import (
	"github.com/tmc/appledocs/generated/devicediscoveryextension"
)

// Suppress unused import errors
var _ = devicediscoveryextension.NewDDDeviceEvent

// ExampleNewDDDeviceEventWithEventTypeDevice demonstrates how to create a DDDeviceEvent instance using NewDDDeviceEventWithEventTypeDevice.
// Creates an event object that conveys status for a discovered device of interest.
func ExampleNewDDDeviceEventWithEventTypeDevice() {
	_ = devicediscoveryextension.NewDDDeviceEventWithEventTypeDevice(
		devicediscoveryextension.DDEventType{}, // type DDEventType
		devicediscoveryextension.DDDevice{}, // device DDDevice
	)
	// Output:
}
