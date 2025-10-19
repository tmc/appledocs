// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSessionWebSocketMessage] class.
var uRLSessionWebSocketMessageClass = _URLSessionWebSocketMessageClass{objc.GetClass("NSURLSessionWebSocketMessage")}

type _URLSessionWebSocketMessageClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionWebSocketMessage] class.
type IURLSessionWebSocketMessage interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage

type URLSessionWebSocketMessage struct {
	objectivec.Object
}

// URLSessionWebSocketMessageFrom constructs a [URLSessionWebSocketMessage] from an unsafe.Pointer.
func URLSessionWebSocketMessageFrom(ptr unsafe.Pointer) URLSessionWebSocketMessage {
	return URLSessionWebSocketMessage{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (uc _URLSessionWebSocketMessageClass) Alloc() URLSessionWebSocketMessage {
	rv := objc.Send[URLSessionWebSocketMessage](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (uc _URLSessionWebSocketMessageClass) New() URLSessionWebSocketMessage {
	rv := objc.Send[URLSessionWebSocketMessage](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionWebSocketMessage) Init() URLSessionWebSocketMessage {
	rv := objc.Send[URLSessionWebSocketMessage](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionWebSocketMessage) Autorelease() URLSessionWebSocketMessage {
	rv := objc.Send[URLSessionWebSocketMessage](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionWebSocketMessage creates a new URLSessionWebSocketMessage instance.
func NewURLSessionWebSocketMessage() URLSessionWebSocketMessage {
	return uRLSessionWebSocketMessageClass.New()
}




