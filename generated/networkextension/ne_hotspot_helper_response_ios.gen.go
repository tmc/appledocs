//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEHotspotHelperResponse


// Set the network that conveys the confidence level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResponse/setNetwork(_:)
func (n_ NEHotspotHelperResponse) SetNetwork(network INEHotspotNetwork) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNetwork:"), network)
}

// Set the list of handled networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResponse/setNetworkList(_:)
func (n_ NEHotspotHelperResponse) SetNetworkList(networkList []NEHotspotNetwork) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNetworkList:"), networkList)
}

// iOS-only properties





