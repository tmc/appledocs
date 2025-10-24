//go:build darwin && ios

// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NISession


// iOS-only properties

// A temporary, random identifier for a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/discoveryToken
func (n_ NISession) DiscoveryToken() objc.IObject /* cross-framework: NIDiscoveryToken */ {
	rv := objc.Send[NIDiscoveryToken](n_.ID, objc.Sel("discoveryToken"))
	return rv
}






