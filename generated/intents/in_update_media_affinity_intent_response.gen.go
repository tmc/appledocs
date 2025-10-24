// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INUpdateMediaAffinityIntentResponse] class.
var (
	INUpdateMediaAffinityIntentResponseClass     _INUpdateMediaAffinityIntentResponseClass
	INUpdateMediaAffinityIntentResponseClassOnce sync.Once
)

func getINUpdateMediaAffinityIntentResponseClass() _INUpdateMediaAffinityIntentResponseClass {
	INUpdateMediaAffinityIntentResponseClassOnce.Do(func() {
		INUpdateMediaAffinityIntentResponseClass = _INUpdateMediaAffinityIntentResponseClass{objc.GetClass("INUpdateMediaAffinityIntentResponse")}
	})
	return INUpdateMediaAffinityIntentResponseClass
}

type _INUpdateMediaAffinityIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INUpdateMediaAffinityIntentResponse] class.
type IINUpdateMediaAffinityIntentResponse interface {
	IINIntentResponse
	Code() unsafe.Pointer
	SetCode(value unsafe.Pointer)
}

// An intents handler’s response to an update media affinity intent.
//
// Use an object to specify the results from trying to update the user’s affinity for a media item. After performing the add action using the criteria specified in the object, create an instance of this class with the results of the action. Siri communicates the response status to the user at appropriate times. You create an object in the and methods of your add media handler object. For more information about implementing your handler object, see .

// An intents handler’s response to an update media affinity intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INUpdateMediaAffinityIntentResponse
type INUpdateMediaAffinityIntentResponse struct {
	INIntentResponse
}

// INUpdateMediaAffinityIntentResponseFrom constructs a [INUpdateMediaAffinityIntentResponse] from an unsafe.Pointer.
//
// An intents handler’s response to an update media affinity intent.
func INUpdateMediaAffinityIntentResponseFrom(ptr unsafe.Pointer) INUpdateMediaAffinityIntentResponse {
	return INUpdateMediaAffinityIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INUpdateMediaAffinityIntentResponseClass) Alloc() INUpdateMediaAffinityIntentResponse {
	rv := objc.Send[INUpdateMediaAffinityIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INUpdateMediaAffinityIntentResponseClass) New() INUpdateMediaAffinityIntentResponse {
	rv := objc.Send[INUpdateMediaAffinityIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INUpdateMediaAffinityIntentResponse) Init() INUpdateMediaAffinityIntentResponse {
	rv := objc.Send[INUpdateMediaAffinityIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INUpdateMediaAffinityIntentResponse) Autorelease() INUpdateMediaAffinityIntentResponse {
	rv := objc.Send[INUpdateMediaAffinityIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINUpdateMediaAffinityIntentResponse creates a new INUpdateMediaAffinityIntentResponse instance.
func NewINUpdateMediaAffinityIntentResponse() INUpdateMediaAffinityIntentResponse {
	return getINUpdateMediaAffinityIntentResponseClass().New()
}

// The code that indicates whether the app successfully updated the user’s affinity for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inupdatemediaaffinityintentresponse/code
func (i_ INUpdateMediaAffinityIntentResponse) Code() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("code"))
	return rv
}

// The code that indicates whether the app successfully updated the user’s affinity for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inupdatemediaaffinityintentresponse/code
func (i_ INUpdateMediaAffinityIntentResponse) SetCode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCode:"), value)
}
