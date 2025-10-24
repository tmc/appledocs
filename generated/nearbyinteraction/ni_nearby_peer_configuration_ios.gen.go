//go:build darwin && ios

// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for NINearbyPeerConfiguration


// iOS-only properties

// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyPeerConfiguration/isCameraAssistanceEnabled
func (n_ NINearbyPeerConfiguration) CameraAssistanceEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("cameraAssistanceEnabled"))
	return rv
}
func (n_ NINearbyPeerConfiguration) SetCameraAssistanceEnabled(value bool) {
	n_.ID.Send(objc.RegisterName("setCameraAssistanceEnabled:"), value)
}




