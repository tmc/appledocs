// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionStreamTask] class.
var uRLSessionStreamTaskClass = _URLSessionStreamTaskClass{objc.GetClass("NSURLSessionStreamTask")}

type _URLSessionStreamTaskClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionStreamTask] class.
type IURLSessionStreamTask interface {
	IURLSessionTask
	WriteDataTimeoutCompletionHandler(data unsafe.Pointer, timeout TimeInterval, completionHandler unsafe.Pointer)
}

// A URL session task that is stream-based. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionStreamTask

type URLSessionStreamTask struct {
	URLSessionTask
}

// URLSessionStreamTaskFrom constructs a [URLSessionStreamTask] from an unsafe.Pointer.
//
// A URL session task that is stream-based.
func URLSessionStreamTaskFrom(ptr unsafe.Pointer) URLSessionStreamTask {
	return URLSessionStreamTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (uc _URLSessionStreamTaskClass) Alloc() URLSessionStreamTask {
	rv := objc.Send[URLSessionStreamTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (uc _URLSessionStreamTaskClass) New() URLSessionStreamTask {
	rv := objc.Send[URLSessionStreamTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionStreamTask) Init() URLSessionStreamTask {
	rv := objc.Send[URLSessionStreamTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionStreamTask) Autorelease() URLSessionStreamTask {
	rv := objc.Send[URLSessionStreamTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionStreamTask creates a new URLSessionStreamTask instance.
func NewURLSessionStreamTask() URLSessionStreamTask {
	return uRLSessionStreamTaskClass.New()
}


// Asynchronously writes the specified data to the stream, and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionStreamTask/write(_:timeout:completionHandler:)
func (u_ URLSessionStreamTask) WriteDataTimeoutCompletionHandler(data unsafe.Pointer, timeout TimeInterval, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("writeData:timeout:completionHandler:"), data, timeout, completionHandler)
}


