//go:build darwin && ios

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for URLSessionConfiguration


// iOS-only properties

// A service type that specifies the Multipath TCP connection policy for transmitting data over Wi-Fi and cellular interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/multipathServiceType-swift.property
func (u_ URLSessionConfiguration) MultipathServiceType() URLSessionMultipathServiceType {
	rv := objc.Send[URLSessionMultipathServiceType](u_.ID, objc.Sel("multipathServiceType"))
	return rv
}
func (u_ URLSessionConfiguration) SetMultipathServiceType(value URLSessionMultipathServiceType) {
	u_.ID.Send(objc.RegisterName("setMultipathServiceType:"), value)
}




