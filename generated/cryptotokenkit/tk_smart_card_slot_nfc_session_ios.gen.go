//go:build darwin && ios

// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for TKSmartCardSlotNFCSession


// Ends the NFC slot session and dismisses the system-presented NFC UI (if present).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotNFCSession/end()
func (t_ TKSmartCardSlotNFCSession) EndSession() {
	objc.Send[objc.ID](t_.ID, objc.Sel("endSession"))
}

// iOS-only properties

// Smart card slot name of the NFC slot that was created together with this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotNFCSession/slotName
func (t_ TKSmartCardSlotNFCSession) SlotName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("slotName"))
	return rv
}





