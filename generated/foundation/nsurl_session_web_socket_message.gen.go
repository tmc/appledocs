// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSessionWebSocketMessage] class.
var (
	URLSessionWebSocketMessageClass     _URLSessionWebSocketMessageClass
	URLSessionWebSocketMessageClassOnce sync.Once
)

func getURLSessionWebSocketMessageClass() _URLSessionWebSocketMessageClass {
	URLSessionWebSocketMessageClassOnce.Do(func() {
		URLSessionWebSocketMessageClass = _URLSessionWebSocketMessageClass{objc.GetClass("NSURLSessionWebSocketMessage")}
	})
	return URLSessionWebSocketMessageClass
}

type _URLSessionWebSocketMessageClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionWebSocketMessage] class.
type IURLSessionWebSocketMessage interface {
	objectivec.IObject
	Data() IData
	String() string
	Type() NSURLSessionWebSocketMessageType
}



// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getURLSessionWebSocketMessageClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/initWithData:
func NewURLSessionWebSocketMessageWithData(data IData) URLSessionWebSocketMessage {
	instance := getURLSessionWebSocketMessageClass().Alloc()
	rv := objc.Send[URLSessionWebSocketMessage](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/initWithString:
func NewURLSessionWebSocketMessageWithString(string_ string) URLSessionWebSocketMessage {
	instance := getURLSessionWebSocketMessageClass().Alloc()
	rv := objc.Send[URLSessionWebSocketMessage](instance.ID, objc.Sel("initWithString:"), objc.String(string_))
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/data
func (u_ URLSessionWebSocketMessage) Data() IData {
	rv := objc.Send[NSData](u_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/string
func (u_ URLSessionWebSocketMessage) String() string {
	rv := objc.Send[string](u_.ID, objc.Sel("string"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/type
func (u_ URLSessionWebSocketMessage) Type() NSURLSessionWebSocketMessageType {
	rv := objc.Send[URLSessionWebSocketMessageType](u_.ID, objc.Sel("type"))
	return rv
}


