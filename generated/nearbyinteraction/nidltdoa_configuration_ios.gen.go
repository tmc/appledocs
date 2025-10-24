//go:build darwin && ios

// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for NIDLTDOAConfiguration


// iOS-only properties

// A unique identifier for a Downlink Time-Difference-of-Arrival network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAConfiguration/networkIdentifier
func (n_ NIDLTDOAConfiguration) NetworkIdentifier() int {
	rv := objc.Send[int](n_.ID, objc.Sel("networkIdentifier"))
	return rv
}
func (n_ NIDLTDOAConfiguration) SetNetworkIdentifier(value int) {
	n_.ID.Send(objc.RegisterName("setNetworkIdentifier:"), value)
}




