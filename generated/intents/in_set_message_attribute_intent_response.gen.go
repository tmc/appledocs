// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSetMessageAttributeIntentResponse] class.
var (
	INSetMessageAttributeIntentResponseClass     _INSetMessageAttributeIntentResponseClass
	INSetMessageAttributeIntentResponseClassOnce sync.Once
)

func getINSetMessageAttributeIntentResponseClass() _INSetMessageAttributeIntentResponseClass {
	INSetMessageAttributeIntentResponseClassOnce.Do(func() {
		INSetMessageAttributeIntentResponseClass = _INSetMessageAttributeIntentResponseClass{objc.GetClass("INSetMessageAttributeIntentResponse")}
	})
	return INSetMessageAttributeIntentResponseClass
}

type _INSetMessageAttributeIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INSetMessageAttributeIntentResponse] class.
type IINSetMessageAttributeIntentResponse interface {
	IINIntentResponse
	Code() unsafe.Pointer
	SetCode(value unsafe.Pointer)
}

// Your app’s response to a set message attribute intent.
//
// An object contains the status of modifying the specified messages. You create instances of this class when confirming or handling a set message attribute intent. You create an object in the and methods of your set message attribute handler object. For more information about implementing your handler object, see .


// Your app’s response to a set message attribute intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetMessageAttributeIntentResponse
type INSetMessageAttributeIntentResponse struct {
	INIntentResponse
}

// INSetMessageAttributeIntentResponseFrom constructs a [INSetMessageAttributeIntentResponse] from an unsafe.Pointer.
//
// Your app’s response to a set message attribute intent.
func INSetMessageAttributeIntentResponseFrom(ptr unsafe.Pointer) INSetMessageAttributeIntentResponse {
	return INSetMessageAttributeIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSetMessageAttributeIntentResponseClass) Alloc() INSetMessageAttributeIntentResponse {
	rv := objc.Send[INSetMessageAttributeIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSetMessageAttributeIntentResponseClass) New() INSetMessageAttributeIntentResponse {
	rv := objc.Send[INSetMessageAttributeIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSetMessageAttributeIntentResponse) Init() INSetMessageAttributeIntentResponse {
	rv := objc.Send[INSetMessageAttributeIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSetMessageAttributeIntentResponse) Autorelease() INSetMessageAttributeIntentResponse {
	rv := objc.Send[INSetMessageAttributeIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSetMessageAttributeIntentResponse creates a new INSetMessageAttributeIntentResponse instance.
func NewINSetMessageAttributeIntentResponse() INSetMessageAttributeIntentResponse {
	return getINSetMessageAttributeIntentResponseClass().New()
}



// The code indicating whether you successfully handled the intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetmessageattributeintentresponse/code
func (i_ INSetMessageAttributeIntentResponse) Code() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("code"))
	return rv
}


// The code indicating whether you successfully handled the intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetmessageattributeintentresponse/code
func (i_ INSetMessageAttributeIntentResponse) SetCode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCode:"), value)
}



