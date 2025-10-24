//go:build darwin && ios

// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Subscriber


// Attempts to refresh the carrier token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriber/refreshCarrierToken()
func (s_ Subscriber) RefreshCarrierToken() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("refreshCarrierToken"))
	return rv
}

// iOS-only properties

// A data object containing authorization information about the subscriber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriber/carrierToken
func (s_ Subscriber) CarrierToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](s_.ID, objc.Sel("carrierToken"))
	return rv
}

// A delegate that receives updates on the subscriber information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriber/delegate
func (s_ Subscriber) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}
func (s_ Subscriber) SetDelegate(value unsafe.Pointer) {
	s_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// An implementation-defined identifier used to correlate this subscriber with information vended by other APIs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriber/identifier
func (s_ Subscriber) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("identifier"))
	return rv
}

// A Boolean property that indicates whether a SIM is present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriber/isSIMInserted
func (s_ Subscriber) SIMInserted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("SIMInserted"))
	return rv
}





