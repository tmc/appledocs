// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INIntentResponse] class.
var (
	INIntentResponseClass     _INIntentResponseClass
	INIntentResponseClassOnce sync.Once
)

func getINIntentResponseClass() _INIntentResponseClass {
	INIntentResponseClassOnce.Do(func() {
		INIntentResponseClass = _INIntentResponseClass{objc.GetClass("INIntentResponse")}
	})
	return INIntentResponseClass
}

type _INIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INIntentResponse] class.
type IINIntentResponse interface {
	objectivec.IObject
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.IUserActivity)
}

// Your response to an intent object.
//
// The class is abstract and provides common behavior for all response objects that your app creates. You do not create instances of this class directly. Instead, you create instances of the specific subclasses representing the intents that you are trying to handle. This class provides access to the custom user activity object that you provided when initializing your response object.


// Your response to an intent object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INIntentResponse
type INIntentResponse struct {
	objectivec.Object
}

// INIntentResponseFrom constructs a [INIntentResponse] from an unsafe.Pointer.
//
// Your response to an intent object.
func INIntentResponseFrom(ptr unsafe.Pointer) INIntentResponse {
	return INIntentResponse{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INIntentResponseClass) Alloc() INIntentResponse {
	rv := objc.Send[INIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INIntentResponseClass) New() INIntentResponse {
	rv := objc.Send[INIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INIntentResponse) Init() INIntentResponse {
	rv := objc.Send[INIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INIntentResponse) Autorelease() INIntentResponse {
	rv := objc.Send[INIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINIntentResponse creates a new INIntentResponse instance.
func NewINIntentResponse() INIntentResponse {
	return getINIntentResponseClass().New()
}



// The user activity object to use when launching the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inintentresponse/useractivity
func (i_ INIntentResponse) UserActivity() foundation.UserActivity {
	rv := objc.Send[foundation.UserActivity](i_.ID, objc.Sel("userActivity"))
	return rv
}


// The user activity object to use when launching the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inintentresponse/useractivity
func (i_ INIntentResponse) SetUserActivity(value foundation.IUserActivity) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUserActivity:"), value)
}



