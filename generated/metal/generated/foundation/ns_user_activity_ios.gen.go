//go:build darwin && ios

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for UserActivity


// iOS-only properties

// An object containing the payload information that launches an App Clip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/appClipActivationPayload
func (u_ UserActivity) AppClipActivationPayload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("appClipActivationPayload"))
	return rv
}

// A unique identifier from the app’s media content catalog for the currently displayed media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/externalMediaContentIdentifier
func (u_ UserActivity) ExternalMediaContentIdentifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("externalMediaContentIdentifier"))
	return rv
}
func (u_ UserActivity) SetExternalMediaContentIdentifier(value IString) {
	u_.ID.Send(objc.RegisterName("setExternalMediaContentIdentifier:"), value)
}

// A Boolean value that determines whether Siri can suggest the user activity as a shortcut to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForPrediction
func (u_ UserActivity) EligibleForPrediction() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("eligibleForPrediction"))
	return rv
}
func (u_ UserActivity) SetEligibleForPrediction(value bool) {
	u_.ID.Send(objc.RegisterName("setEligibleForPrediction:"), value)
}

// The NDEF message read by the system in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/ndefMessagePayload
func (u_ UserActivity) NdefMessagePayload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("ndefMessagePayload"))
	return rv
}

// A set of defined contexts in which an intent or activity might be relevant to a user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/shortcutAvailability
func (u_ UserActivity) ShortcutAvailability() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("shortcutAvailability"))
	return rv
}
func (u_ UserActivity) SetShortcutAvailability(value unsafe.Pointer) {
	u_.ID.Send(objc.RegisterName("setShortcutAvailability:"), value)
}





