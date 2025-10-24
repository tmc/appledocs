//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEVPNProtocol


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/sliceUUID
func (n_ NEVPNProtocol) SliceUUID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("sliceUUID"))
	return rv
}
func (n_ NEVPNProtocol) SetSliceUUID(value objc.IObject /* cross-framework: NSString */) {
	n_.ID.Send(objc.RegisterName("setSliceUUID:"), value)
}





