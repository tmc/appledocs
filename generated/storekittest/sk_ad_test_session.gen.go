// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/storekit"
)

// The class instance for the [AdTestSession] class.
var (
	AdTestSessionClass     _AdTestSessionClass
	AdTestSessionClassOnce sync.Once
)

func getAdTestSessionClass() _AdTestSessionClass {
	AdTestSessionClassOnce.Do(func() {
		AdTestSessionClass = _AdTestSessionClass{objc.GetClass("SKAdTestSession")}
	})
	return AdTestSessionClass
}

type _AdTestSessionClass struct {
	class objc.Class
}

// An interface definition for the [AdTestSession] class.
type IAdTestSession interface {
	objectivec.IObject
	FlushPostbacksWithResponses(responses unsafe.Pointer)
	SetPostbacksError(postbacks []AdTestPostback, error_ unsafe.Pointer) bool
	ValidateImpressionPublicKeyError(impression storekit.IAdImpression, publicKey appkit.string, error_ unsafe.Pointer) bool
	ValidateImpressionWithParametersPublicKeyError(parameters unsafe.Pointer, publicKey appkit.string, error_ unsafe.Pointer) bool
	ValidateWebAdImpressionPayloadPublicKeyError(impressionData foundation.IData, publicKey appkit.string, error_ unsafe.Pointer) bool
}

// The class you use to test ad impressions and postbacks in Xcode.
//
// Use the class to test your implementations of SKAdNetwork. Create one instance of this class to use in multiple test cases. The instance represents a test session, and holds a set of test postbacks. Use to create test postbacks. Call to add test postbacks to the test session. The test session deletes the postbacks from the instance after you call .
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession
type AdTestSession struct {
	objectivec.Object
}

// AdTestSessionFrom constructs a [AdTestSession] from an unsafe.Pointer.
//
// The class you use to test ad impressions and postbacks in Xcode.
func AdTestSessionFrom(ptr unsafe.Pointer) AdTestSession {
	return AdTestSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AdTestSessionClass) Alloc() AdTestSession {
	rv := objc.Send[AdTestSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AdTestSessionClass) New() AdTestSession {
	rv := objc.Send[AdTestSession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdTestSession) Init() AdTestSession {
	rv := objc.Send[AdTestSession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdTestSession) Autorelease() AdTestSession {
	rv := objc.Send[AdTestSession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdTestSession creates a new AdTestSession instance.
func NewAdTestSession() AdTestSession {
	return getAdTestSessionClass().New()
}



// Sends the test postbacks and handles the responses.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/flushPostbacks(responses:)
func (a_ AdTestSession) FlushPostbacksWithResponses(responses unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("flushPostbacksWithResponses:"), responses)
}

// Add test postbacks to the test session.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/setPostbacks(_:)
func (a_ AdTestSession) SetPostbacksError(postbacks []AdTestPostback, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPostbacks:error:"), postbacks, error_)
	return rv
}

// Validates an impression for a view-through ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/validate(_:publicKey:)
func (a_ AdTestSession) ValidateImpressionPublicKeyError(impression storekit.IAdImpression, publicKey appkit.string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("validateImpression:publicKey:error:"), impression, publicKey, error_)
	return rv
}

// Validates an impression for a StoreKit-rendered ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/validateImpression(parameters:publicKey:)
func (a_ AdTestSession) ValidateImpressionWithParametersPublicKeyError(parameters unsafe.Pointer, publicKey appkit.string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("validateImpressionWithParameters:publicKey:error:"), parameters, publicKey, error_)
	return rv
}

// Validates an impression for a web ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/validateWebAdImpressionPayload(_:publicKey:)
func (a_ AdTestSession) ValidateWebAdImpressionPayloadPublicKeyError(impressionData foundation.IData, publicKey appkit.string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("validateWebAdImpressionPayload:publicKey:error:"), impressionData, publicKey, error_)
	return rv
}

// The URL that SKAdNetwork computes to send copies of winning postbacks to the advertised app’s developer.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/developerPostbackURL
func (a_ AdTestSession) DeveloperPostbackURL() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("developerPostbackURL"))
	return rv
}

// An array of test postbacks you set in the testing environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession/postbacks
func (a_ AdTestSession) Postbacks() []AdTestPostback {
	rv := objc.Send[[]AdTestPostback](a_.ID, objc.Sel("postbacks"))
	return rv
}


