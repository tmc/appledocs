//go:build darwin && ios

// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NINearbyObject


// iOS-only properties

// A vector that points from the user’s device in the direction of the peer device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/direction-5xcld
func (n_ NINearbyObject) Direction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("direction"))
	return rv
}

// A unique identifier for a peer device in the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/discoveryToken
func (n_ NINearbyObject) DiscoveryToken() INIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](n_.ID, objc.Sel("discoveryToken"))
	return rv
}

// The distance from the user’s device to the peer device in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/distance-9atp7
func (n_ NINearbyObject) Distance() float32 {
	rv := objc.Send[float32](n_.ID, objc.Sel("distance"))
	return rv
}

// An angle in radians that indicates the azimuthal direction to the nearby object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/horizontalAngle-9ibky
func (n_ NINearbyObject) HorizontalAngle() float32 {
	rv := objc.Send[float32](n_.ID, objc.Sel("horizontalAngle"))
	return rv
}

// The estimation of a nearby object’s vertical position as it relates to the user’s device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/verticalDirectionEstimate-swift.property
func (n_ NINearbyObject) VerticalDirectionEstimate() NINearbyObjectVerticalDirectionEstimate {
	rv := objc.Send[NINearbyObjectVerticalDirectionEstimate](n_.ID, objc.Sel("verticalDirectionEstimate"))
	return rv
}





