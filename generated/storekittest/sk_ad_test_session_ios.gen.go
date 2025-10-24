//go:build darwin && ios

// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/storekit"
)

// iOS-only methods for AdTestSession


// Sends the test postbacks and handles the responses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/flushPostbacks(responses:)
func (a_ AdTestSession) FlushPostbacksWithResponses(responses ANTestPostbackResponseHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("flushPostbacksWithResponses:"), responses)
}

// Add test postbacks to the test session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/setPostbacks(_:)
func (a_ AdTestSession) SetPostbacksError(postbacks []AdTestPostback, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPostbacks:error:"), postbacks, error_)
	return rv
}

// Validates an impression for a view-through ad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/validate(_:publicKey:)
func (a_ AdTestSession) ValidateImpressionPublicKeyError(impression storekit.AdImpression, publicKey objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("validateImpression:publicKey:error:"), impression, publicKey, error_)
	return rv
}

// Validates an impression for a StoreKit-rendered ad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/validateImpression(parameters:publicKey:)
func (a_ AdTestSession) ValidateImpressionWithParametersPublicKeyError(parameters foundation.IDictionary, publicKey objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("validateImpressionWithParameters:publicKey:error:"), parameters, publicKey, error_)
	return rv
}

// Validates an impression for a web ad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/validateWebAdImpressionPayload(_:publicKey:)
func (a_ AdTestSession) ValidateWebAdImpressionPayloadPublicKeyError(impressionData objc.IObject /* cross-framework: NSData */, publicKey objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("validateWebAdImpressionPayload:publicKey:error:"), impressionData, publicKey, error_)
	return rv
}

// iOS-only properties

// The URL that SKAdNetwork computes to send copies of winning postbacks to the advertised app’s developer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/developerPostbackURL
func (a_ AdTestSession) DeveloperPostbackURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("developerPostbackURL"))
	return rv
}

// An array of test postbacks you set in the testing environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/postbacks
func (a_ AdTestSession) Postbacks() []AdTestPostback {
	rv := objc.Send[[]AdTestPostback](a_.ID, objc.Sel("postbacks"))
	return rv
}




