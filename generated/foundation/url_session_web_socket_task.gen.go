// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionWebSocketTask] class.
var (
	uRLSessionWebSocketTaskClass     _URLSessionWebSocketTaskClass
	uRLSessionWebSocketTaskClassOnce sync.Once
)

func getURLSessionWebSocketTaskClass() _URLSessionWebSocketTaskClass {
	uRLSessionWebSocketTaskClassOnce.Do(func() {
		uRLSessionWebSocketTaskClass = _URLSessionWebSocketTaskClass{objc.GetClass("NSURLSessionWebSocketTask")}
	})
	return uRLSessionWebSocketTaskClass
}

type _URLSessionWebSocketTaskClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionWebSocketTask] class.
type IURLSessionWebSocketTask interface {
	IURLSessionTask
}

// A URL session task that communicates over the WebSockets protocol standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask
type URLSessionWebSocketTask struct {
	URLSessionTask
}

// URLSessionWebSocketTaskFrom constructs a [URLSessionWebSocketTask] from an unsafe.Pointer.
//
// A URL session task that communicates over the WebSockets protocol standard.
func URLSessionWebSocketTaskFrom(ptr unsafe.Pointer) URLSessionWebSocketTask {
	return URLSessionWebSocketTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _URLSessionWebSocketTaskClass) Alloc() URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLSessionWebSocketTaskClass) New() URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionWebSocketTask) Init() URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionWebSocketTask) Autorelease() URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionWebSocketTask creates a new URLSessionWebSocketTask instance.
func NewURLSessionWebSocketTask() URLSessionWebSocketTask {
	return getURLSessionWebSocketTaskClass().New()
}




