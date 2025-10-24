// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSendRideFeedbackIntentResponse] class.
var (
	INSendRideFeedbackIntentResponseClass     _INSendRideFeedbackIntentResponseClass
	INSendRideFeedbackIntentResponseClassOnce sync.Once
)

func getINSendRideFeedbackIntentResponseClass() _INSendRideFeedbackIntentResponseClass {
	INSendRideFeedbackIntentResponseClassOnce.Do(func() {
		INSendRideFeedbackIntentResponseClass = _INSendRideFeedbackIntentResponseClass{objc.GetClass("INSendRideFeedbackIntentResponse")}
	})
	return INSendRideFeedbackIntentResponseClass
}

type _INSendRideFeedbackIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INSendRideFeedbackIntentResponse] class.
type IINSendRideFeedbackIntentResponse interface {
	IINIntentResponse
	Code() unsafe.Pointer
	SetCode(value unsafe.Pointer)
}

// Your app’s response to a send ride feedback intent.
//
// An object contains your app’s response to a request for feedback about a ride. After creating the response object, specify any feedback using the properties of this object. Siri and Maps display your response information to the user during the confirmation phase. You create an object in the and methods of your handler object. For more information about implementing your handler object, see .

// Your app’s response to a send ride feedback intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendRideFeedbackIntentResponse
type INSendRideFeedbackIntentResponse struct {
	INIntentResponse
}

// INSendRideFeedbackIntentResponseFrom constructs a [INSendRideFeedbackIntentResponse] from an unsafe.Pointer.
//
// Your app’s response to a send ride feedback intent.
func INSendRideFeedbackIntentResponseFrom(ptr unsafe.Pointer) INSendRideFeedbackIntentResponse {
	return INSendRideFeedbackIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSendRideFeedbackIntentResponseClass) Alloc() INSendRideFeedbackIntentResponse {
	rv := objc.Send[INSendRideFeedbackIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSendRideFeedbackIntentResponseClass) New() INSendRideFeedbackIntentResponse {
	rv := objc.Send[INSendRideFeedbackIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSendRideFeedbackIntentResponse) Init() INSendRideFeedbackIntentResponse {
	rv := objc.Send[INSendRideFeedbackIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSendRideFeedbackIntentResponse) Autorelease() INSendRideFeedbackIntentResponse {
	rv := objc.Send[INSendRideFeedbackIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSendRideFeedbackIntentResponse creates a new INSendRideFeedbackIntentResponse instance.
func NewINSendRideFeedbackIntentResponse() INSendRideFeedbackIntentResponse {
	return getINSendRideFeedbackIntentResponseClass().New()
}

// The code indicating whether your app successfully handled the intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendridefeedbackintentresponse/code
func (i_ INSendRideFeedbackIntentResponse) Code() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("code"))
	return rv
}

// The code indicating whether your app successfully handled the intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insendridefeedbackintentresponse/code
func (i_ INSendRideFeedbackIntentResponse) SetCode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCode:"), value)
}
