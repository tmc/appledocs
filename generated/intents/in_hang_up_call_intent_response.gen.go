// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INHangUpCallIntentResponse] class.
var (
	INHangUpCallIntentResponseClass     _INHangUpCallIntentResponseClass
	INHangUpCallIntentResponseClassOnce sync.Once
)

func getINHangUpCallIntentResponseClass() _INHangUpCallIntentResponseClass {
	INHangUpCallIntentResponseClassOnce.Do(func() {
		INHangUpCallIntentResponseClass = _INHangUpCallIntentResponseClass{objc.GetClass("INHangUpCallIntentResponse")}
	})
	return INHangUpCallIntentResponseClass
}

type _INHangUpCallIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INHangUpCallIntentResponse] class.
type IINHangUpCallIntentResponse interface {
	IINIntentResponse
}

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INHangUpCallIntentResponse
type INHangUpCallIntentResponse struct {
	INIntentResponse
}

// INHangUpCallIntentResponseFrom constructs a [INHangUpCallIntentResponse] from an unsafe.Pointer.
func INHangUpCallIntentResponseFrom(ptr unsafe.Pointer) INHangUpCallIntentResponse {
	return INHangUpCallIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INHangUpCallIntentResponseClass) Alloc() INHangUpCallIntentResponse {
	rv := objc.Send[INHangUpCallIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INHangUpCallIntentResponseClass) New() INHangUpCallIntentResponse {
	rv := objc.Send[INHangUpCallIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INHangUpCallIntentResponse) Init() INHangUpCallIntentResponse {
	rv := objc.Send[INHangUpCallIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INHangUpCallIntentResponse) Autorelease() INHangUpCallIntentResponse {
	rv := objc.Send[INHangUpCallIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINHangUpCallIntentResponse creates a new INHangUpCallIntentResponse instance.
func NewINHangUpCallIntentResponse() INHangUpCallIntentResponse {
	return getINHangUpCallIntentResponseClass().New()
}




