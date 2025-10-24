//go:build darwin && ios

// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CloudServiceSetupViewController


// iOS-only properties

// The cloud service view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceSetupViewController/delegate
func (c_ CloudServiceSetupViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}
func (c_ CloudServiceSetupViewController) SetDelegate(value unsafe.Pointer) {
	c_.ID.Send(objc.RegisterName("setDelegate:"), value)
}





