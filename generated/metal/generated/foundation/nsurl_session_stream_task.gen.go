// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionStreamTask] class.
var (
	URLSessionStreamTaskClass     _URLSessionStreamTaskClass
	URLSessionStreamTaskClassOnce sync.Once
)

func getURLSessionStreamTaskClass() _URLSessionStreamTaskClass {
	URLSessionStreamTaskClassOnce.Do(func() {
		URLSessionStreamTaskClass = _URLSessionStreamTaskClass{objc.GetClass("NSURLSessionStreamTask")}
	})
	return URLSessionStreamTaskClass
}

type _URLSessionStreamTaskClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionStreamTask] class.
type IURLSessionStreamTask interface {
	IURLSessionTask
	// properties:
	HttpShouldUsePipelining() bool /* primitive/slice/pointer. */
	SetHttpShouldUsePipelining(value bool /* primitive/slice/pointer. */)
	// methods:
}

// A URL session task that is stream-based.
//
// is a concrete subclass of . Many of the methods in the class are documented in . The class provides an interface a TCP/IP connection created via . Tasks may be created from an using the and methods. They may also be created as a result of an being upgraded via the HTTP response header and appropriate use of the option of . A object performs asynchronous reads and writes, which are enqueued and executed serially, calling a handler upon completion being on the session delegate queue. If the task is canceled, all enqueued reads and writes will call their completion handlers with an appropriate error. When working with APIs that accept objects, you can create and objects from an object by calling the method.


// A URL session task that is stream-based.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getURLSessionStreamTaskClass().New()
}



// A Boolean value that determines whether the session should use HTTP pipelining.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpshouldusepipelining
func (u_ URLSessionStreamTask) HttpShouldUsePipelining() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("httpShouldUsePipelining"))
	return rv
}


// A Boolean value that determines whether the session should use HTTP pipelining.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpshouldusepipelining
func (u_ URLSessionStreamTask) SetHttpShouldUsePipelining(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpShouldUsePipelining:"), value)
}



