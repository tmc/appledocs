// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INStartAudioCallIntentResponse] class.
var (
	INStartAudioCallIntentResponseClass     _INStartAudioCallIntentResponseClass
	INStartAudioCallIntentResponseClassOnce sync.Once
)

func getINStartAudioCallIntentResponseClass() _INStartAudioCallIntentResponseClass {
	INStartAudioCallIntentResponseClassOnce.Do(func() {
		INStartAudioCallIntentResponseClass = _INStartAudioCallIntentResponseClass{objc.GetClass("INStartAudioCallIntentResponse")}
	})
	return INStartAudioCallIntentResponseClass
}

type _INStartAudioCallIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INStartAudioCallIntentResponse] class.
type IINStartAudioCallIntentResponse interface {
	IINIntentResponse
	Code() unsafe.Pointer
	SetCode(value unsafe.Pointer)
}

// An app’s response to an intent to start an audio call.
//
// Use an object to specify whether your app is able to initiate an audio-based call. You create instances of this class when confirming and handling an object. When it’s time to call the user, SiriKit launches your app and delivers the object contained in this object. Use that user activity object to specify any additional information that would assist your app in placing the call. You create an object in the and methods of your start audio call handler object. For more information about implementing your handler object, see .


// An app’s response to an intent to start an audio call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartAudioCallIntentResponse
type INStartAudioCallIntentResponse struct {
	INIntentResponse
}

// INStartAudioCallIntentResponseFrom constructs a [INStartAudioCallIntentResponse] from an unsafe.Pointer.
//
// An app’s response to an intent to start an audio call.
func INStartAudioCallIntentResponseFrom(ptr unsafe.Pointer) INStartAudioCallIntentResponse {
	return INStartAudioCallIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INStartAudioCallIntentResponseClass) Alloc() INStartAudioCallIntentResponse {
	rv := objc.Send[INStartAudioCallIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INStartAudioCallIntentResponseClass) New() INStartAudioCallIntentResponse {
	rv := objc.Send[INStartAudioCallIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INStartAudioCallIntentResponse) Init() INStartAudioCallIntentResponse {
	rv := objc.Send[INStartAudioCallIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INStartAudioCallIntentResponse) Autorelease() INStartAudioCallIntentResponse {
	rv := objc.Send[INStartAudioCallIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINStartAudioCallIntentResponse creates a new INStartAudioCallIntentResponse instance.
func NewINStartAudioCallIntentResponse() INStartAudioCallIntentResponse {
	return getINStartAudioCallIntentResponseClass().New()
}



// The code indicating whether you successfully handled the intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartaudiocallintentresponse/code
func (i_ INStartAudioCallIntentResponse) Code() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("code"))
	return rv
}


// The code indicating whether you successfully handled the intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartaudiocallintentresponse/code
func (i_ INStartAudioCallIntentResponse) SetCode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCode:"), value)
}



