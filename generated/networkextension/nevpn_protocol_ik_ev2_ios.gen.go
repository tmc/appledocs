//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for NEVPNProtocolIKEv2


// iOS-only properties

// A property to enable the use of cellular data when Wi-Fi connectivity is poor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enableFallback
func (n_ NEVPNProtocolIKEv2) EnableFallback() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enableFallback"))
	return rv
}
func (n_ NEVPNProtocolIKEv2) SetEnableFallback(value bool) {
	n_.ID.Send(objc.RegisterName("setEnableFallback:"), value)
}





