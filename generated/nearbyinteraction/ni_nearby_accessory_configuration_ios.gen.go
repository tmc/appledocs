//go:build darwin && ios

// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for NINearbyAccessoryConfiguration


// iOS-only properties

// An identifier for the accessory in a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyAccessoryConfiguration/accessoryDiscoveryToken
func (n_ NINearbyAccessoryConfiguration) AccessoryDiscoveryToken() INIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](n_.ID, objc.Sel("accessoryDiscoveryToken"))
	return rv
}

// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyAccessoryConfiguration/isCameraAssistanceEnabled
func (n_ NINearbyAccessoryConfiguration) CameraAssistanceEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("cameraAssistanceEnabled"))
	return rv
}
func (n_ NINearbyAccessoryConfiguration) SetCameraAssistanceEnabled(value bool) {
	n_.ID.Send(objc.RegisterName("setCameraAssistanceEnabled:"), value)
}




