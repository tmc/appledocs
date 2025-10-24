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
func (n_ NINearbyAccessoryConfiguration) AccessoryDiscoveryToken() objc.IObject /* cross-framework: NIDiscoveryToken */ {
	rv := objc.Send[NIDiscoveryToken](n_.ID, objc.Sel("accessoryDiscoveryToken"))
	return rv
}




