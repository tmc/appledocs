// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INStartWorkoutIntentResponse] class.
var (
	INStartWorkoutIntentResponseClass     _INStartWorkoutIntentResponseClass
	INStartWorkoutIntentResponseClassOnce sync.Once
)

func getINStartWorkoutIntentResponseClass() _INStartWorkoutIntentResponseClass {
	INStartWorkoutIntentResponseClassOnce.Do(func() {
		INStartWorkoutIntentResponseClass = _INStartWorkoutIntentResponseClass{objc.GetClass("INStartWorkoutIntentResponse")}
	})
	return INStartWorkoutIntentResponseClass
}

type _INStartWorkoutIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INStartWorkoutIntentResponse] class.
type IINStartWorkoutIntentResponse interface {
	IINIntentResponse
}

// Your app’s response to a start workout intent.
//
// Use an object to specify whether your app is able to start a workout. The response object contains only the response code that indicates whether to launch your app or whether there was a problem. You create an object in the and methods of your start workout handler object. For more information about implementing your handler object, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartWorkoutIntentResponse
type INStartWorkoutIntentResponse struct {
	INIntentResponse
}

// INStartWorkoutIntentResponseFrom constructs a [INStartWorkoutIntentResponse] from an unsafe.Pointer.
//
// Your app’s response to a start workout intent.
func INStartWorkoutIntentResponseFrom(ptr unsafe.Pointer) INStartWorkoutIntentResponse {
	return INStartWorkoutIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INStartWorkoutIntentResponseClass) Alloc() INStartWorkoutIntentResponse {
	rv := objc.Send[INStartWorkoutIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INStartWorkoutIntentResponseClass) New() INStartWorkoutIntentResponse {
	rv := objc.Send[INStartWorkoutIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INStartWorkoutIntentResponse) Init() INStartWorkoutIntentResponse {
	rv := objc.Send[INStartWorkoutIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INStartWorkoutIntentResponse) Autorelease() INStartWorkoutIntentResponse {
	rv := objc.Send[INStartWorkoutIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINStartWorkoutIntentResponse creates a new INStartWorkoutIntentResponse instance.
func NewINStartWorkoutIntentResponse() INStartWorkoutIntentResponse {
	return getINStartWorkoutIntentResponseClass().New()
}


// The code that indicates whether you successfully handled the intent.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartworkoutintentresponse/code
func (i_ INStartWorkoutIntentResponse) Code() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("code"))
	return rv
}


// SetCode sets the value of the code property.
// The code that indicates whether you successfully handled the intent.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartworkoutintentresponse/code
func (i_ INStartWorkoutIntentResponse) SetCode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCode:"), value)
}



