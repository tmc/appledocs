// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSessionTask] class.
var (
	uRLSessionTaskClass     _URLSessionTaskClass
	uRLSessionTaskClassOnce sync.Once
)

func getURLSessionTaskClass() _URLSessionTaskClass {
	uRLSessionTaskClassOnce.Do(func() {
		uRLSessionTaskClass = _URLSessionTaskClass{objc.GetClass("NSURLSessionTask")}
	})
	return uRLSessionTaskClass
}

type _URLSessionTaskClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionTask] class.
type IURLSessionTask interface {
	objectivec.IObject
	Cancel()
	Resume()
}

// A task, like downloading a specific resource, performed in a URL session.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask
type URLSessionTask struct {
	objectivec.Object
}

// URLSessionTaskFrom constructs a [URLSessionTask] from an unsafe.Pointer.
//
// A task, like downloading a specific resource, performed in a URL session.
func URLSessionTaskFrom(ptr unsafe.Pointer) URLSessionTask {
	return URLSessionTask{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLSessionTaskClass) Alloc() URLSessionTask {
	rv := objc.Send[URLSessionTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLSessionTaskClass) New() URLSessionTask {
	rv := objc.Send[URLSessionTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionTask) Init() URLSessionTask {
	rv := objc.Send[URLSessionTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionTask) Autorelease() URLSessionTask {
	rv := objc.Send[URLSessionTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionTask creates a new URLSessionTask instance.
func NewURLSessionTask() URLSessionTask {
	return getURLSessionTaskClass().New()
}


// Cancels the task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/cancel()
func (u_ URLSessionTask) Cancel() {
	objc.Send[objc.ID](u_.ID, objc.Sel("cancel"))
}
// Resumes the task, if it is suspended.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/resume()
func (u_ URLSessionTask) Resume() {
	objc.Send[objc.ID](u_.ID, objc.Sel("resume"))
}


