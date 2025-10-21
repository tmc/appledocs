// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INCancelRideIntentResponse] class.
var (
	INCancelRideIntentResponseClass     _INCancelRideIntentResponseClass
	INCancelRideIntentResponseClassOnce sync.Once
)

func getINCancelRideIntentResponseClass() _INCancelRideIntentResponseClass {
	INCancelRideIntentResponseClassOnce.Do(func() {
		INCancelRideIntentResponseClass = _INCancelRideIntentResponseClass{objc.GetClass("INCancelRideIntentResponse")}
	})
	return INCancelRideIntentResponseClass
}

type _INCancelRideIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INCancelRideIntentResponse] class.
type IINCancelRideIntentResponse interface {
	IINIntentResponse
}

// Your app’s response to a cancel ride intent.
//
// An object contains your app’s response to the cancellation of a ride. After creating the response object, specify any cancellation-related fees using the properties of this object. Siri and Maps display your response information to the user during the confirmation phase. You create an object in the and methods of your handler object. For more information about implementing your handler object, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCancelRideIntentResponse
type INCancelRideIntentResponse struct {
	INIntentResponse
}

// INCancelRideIntentResponseFrom constructs a [INCancelRideIntentResponse] from an unsafe.Pointer.
//
// Your app’s response to a cancel ride intent.
func INCancelRideIntentResponseFrom(ptr unsafe.Pointer) INCancelRideIntentResponse {
	return INCancelRideIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INCancelRideIntentResponseClass) Alloc() INCancelRideIntentResponse {
	rv := objc.Send[INCancelRideIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCancelRideIntentResponseClass) New() INCancelRideIntentResponse {
	rv := objc.Send[INCancelRideIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCancelRideIntentResponse) Init() INCancelRideIntentResponse {
	rv := objc.Send[INCancelRideIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCancelRideIntentResponse) Autorelease() INCancelRideIntentResponse {
	rv := objc.Send[INCancelRideIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCancelRideIntentResponse creates a new INCancelRideIntentResponse instance.
func NewINCancelRideIntentResponse() INCancelRideIntentResponse {
	return getINCancelRideIntentResponseClass().New()
}




