//go:build darwin && ios

// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for TKSmartCardSlotManager


// Creates an NFC smart card slot using the device’s hardware and presents a system UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/createNFCSlot(message:completion:)
func (t_ TKSmartCardSlotManager) CreateNFCSlotWithMessageCompletion(message objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("createNFCSlotWithMessage:completion:"), message, completion)
}

// Determines whether NFC (Near Field Communication) is supported on this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/isNFCSupported()
func (t_ TKSmartCardSlotManager) IsNFCSupported() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isNFCSupported"))
	return rv
}

// iOS-only properties





