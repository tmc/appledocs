//go:build darwin && ios

// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NIDiscoveryToken


// iOS-only properties

// A protocol object that describes the nearby interaction capabilities of a person’s device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDiscoveryToken/deviceCapabilities
func (n_ NIDiscoveryToken) DeviceCapabilities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("deviceCapabilities"))
	return rv
}





