// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INShareFocusStatusIntentResponse] class.
var (
	INShareFocusStatusIntentResponseClass     _INShareFocusStatusIntentResponseClass
	INShareFocusStatusIntentResponseClassOnce sync.Once
)

func getINShareFocusStatusIntentResponseClass() _INShareFocusStatusIntentResponseClass {
	INShareFocusStatusIntentResponseClassOnce.Do(func() {
		INShareFocusStatusIntentResponseClass = _INShareFocusStatusIntentResponseClass{objc.GetClass("INShareFocusStatusIntentResponse")}
	})
	return INShareFocusStatusIntentResponseClass
}

type _INShareFocusStatusIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INShareFocusStatusIntentResponse] class.
type IINShareFocusStatusIntentResponse interface {
	IINIntentResponse
}

// Your app’s response to an intent that shares the user’s focus status.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntentResponse
type INShareFocusStatusIntentResponse struct {
	INIntentResponse
}

// INShareFocusStatusIntentResponseFrom constructs a [INShareFocusStatusIntentResponse] from an unsafe.Pointer.
//
// Your app’s response to an intent that shares the user’s focus status.
func INShareFocusStatusIntentResponseFrom(ptr unsafe.Pointer) INShareFocusStatusIntentResponse {
	return INShareFocusStatusIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INShareFocusStatusIntentResponseClass) Alloc() INShareFocusStatusIntentResponse {
	rv := objc.Send[INShareFocusStatusIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INShareFocusStatusIntentResponseClass) New() INShareFocusStatusIntentResponse {
	rv := objc.Send[INShareFocusStatusIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INShareFocusStatusIntentResponse) Init() INShareFocusStatusIntentResponse {
	rv := objc.Send[INShareFocusStatusIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INShareFocusStatusIntentResponse) Autorelease() INShareFocusStatusIntentResponse {
	rv := objc.Send[INShareFocusStatusIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINShareFocusStatusIntentResponse creates a new INShareFocusStatusIntentResponse instance.
func NewINShareFocusStatusIntentResponse() INShareFocusStatusIntentResponse {
	return getINShareFocusStatusIntentResponseClass().New()
}




// Creates a response with the specified response code and user activity.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntentResponse/init(code:userActivity:)
func NewINShareFocusStatusIntentResponseWithCodeUserActivity(code unsafe.Pointer, userActivity unsafe.Pointer) INShareFocusStatusIntentResponse {
	instance := getINShareFocusStatusIntentResponseClass().Alloc()
	rv := objc.Send[INShareFocusStatusIntentResponse](instance.ID, objc.Sel("initWithCode:userActivity:"), code, userActivity)
	rv.Autorelease()
	return rv
}


// Your app’s ability to handle an intent that shares the user’s focus status.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntentResponse/code
func (i_ INShareFocusStatusIntentResponse) Code() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("code"))
	return rv
}


