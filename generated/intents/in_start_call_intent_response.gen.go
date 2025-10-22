// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INStartCallIntentResponse] class.
var (
	INStartCallIntentResponseClass     _INStartCallIntentResponseClass
	INStartCallIntentResponseClassOnce sync.Once
)

func getINStartCallIntentResponseClass() _INStartCallIntentResponseClass {
	INStartCallIntentResponseClassOnce.Do(func() {
		INStartCallIntentResponseClass = _INStartCallIntentResponseClass{objc.GetClass("INStartCallIntentResponse")}
	})
	return INStartCallIntentResponseClass
}

type _INStartCallIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INStartCallIntentResponse] class.
type IINStartCallIntentResponse interface {
	IINIntentResponse
	Code() INStartCallIntentResponseCode
}

// Your app’s response to a start call intent.
//
// Use an object to specify whether your app is able to initiate an audio or video call. You create instances of this class when confirming and handling an object. When it’s time to call the user, SiriKit launches your app and delivers the object contained in this object. Use that user activity object to specify any additional information that assists your app in placing the call. You create an object in the and methods of your start call handler object. For more information about implementing your handler object, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntentResponse
type INStartCallIntentResponse struct {
	INIntentResponse
}

// INStartCallIntentResponseFrom constructs a [INStartCallIntentResponse] from an unsafe.Pointer.
//
// Your app’s response to a start call intent.
func INStartCallIntentResponseFrom(ptr unsafe.Pointer) INStartCallIntentResponse {
	return INStartCallIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INStartCallIntentResponseClass) Alloc() INStartCallIntentResponse {
	rv := objc.Send[INStartCallIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INStartCallIntentResponseClass) New() INStartCallIntentResponse {
	rv := objc.Send[INStartCallIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INStartCallIntentResponse) Init() INStartCallIntentResponse {
	rv := objc.Send[INStartCallIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INStartCallIntentResponse) Autorelease() INStartCallIntentResponse {
	rv := objc.Send[INStartCallIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINStartCallIntentResponse creates a new INStartCallIntentResponse instance.
func NewINStartCallIntentResponse() INStartCallIntentResponse {
	return getINStartCallIntentResponseClass().New()
}




// Initializes the response object with the specified code and user activity object.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntentResponse/init(code:userActivity:)
func NewINStartCallIntentResponseWithCodeUserActivity(code INStartCallIntentResponseCode, userActivity foundation.IUserActivity) INStartCallIntentResponse {
	instance := getINStartCallIntentResponseClass().Alloc()
	rv := objc.Send[INStartCallIntentResponse](instance.ID, objc.Sel("initWithCode:userActivity:"), code, userActivity)
	rv.Autorelease()
	return rv
}


// The code indicating whether you successfully handled the intent.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntentResponse/code
func (i_ INStartCallIntentResponse) Code() INStartCallIntentResponseCode {
	rv := objc.Send[INStartCallIntentResponseCode](i_.ID, objc.Sel("code"))
	return rv
}


