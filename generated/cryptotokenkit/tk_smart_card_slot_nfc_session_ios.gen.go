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

// Updates the message of the system-presented NFC UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotNFCSession/update(message:)
func (t_ TKSmartCardSlotNFCSession) UpdateWithMessageError(message objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("updateWithMessage:error:"), message, error_)
	return rv
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





