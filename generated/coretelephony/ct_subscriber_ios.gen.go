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


// iOS-only properties

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





