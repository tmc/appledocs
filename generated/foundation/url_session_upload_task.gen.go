// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionUploadTask] class.
var (
	uRLSessionUploadTaskClass     _URLSessionUploadTaskClass
	uRLSessionUploadTaskClassOnce sync.Once
)

func getURLSessionUploadTaskClass() _URLSessionUploadTaskClass {
	uRLSessionUploadTaskClassOnce.Do(func() {
		uRLSessionUploadTaskClass = _URLSessionUploadTaskClass{objc.GetClass("NSURLSessionUploadTask")}
	})
	return uRLSessionUploadTaskClass
}

type _URLSessionUploadTaskClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionUploadTask] class.
type IURLSessionUploadTask interface {
	IURLSessionDataTask
}

// A URL session task that uploads data to the network in a request body. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionUploadTask
type URLSessionUploadTask struct {
	URLSessionDataTask
}

// URLSessionUploadTaskFrom constructs a [URLSessionUploadTask] from an unsafe.Pointer.
//
// A URL session task that uploads data to the network in a request body.
func URLSessionUploadTaskFrom(ptr unsafe.Pointer) URLSessionUploadTask {
	return URLSessionUploadTask{
		URLSessionDataTask: URLSessionDataTaskFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _URLSessionUploadTaskClass) Alloc() URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLSessionUploadTaskClass) New() URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionUploadTask) Init() URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionUploadTask) Autorelease() URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionUploadTask creates a new URLSessionUploadTask instance.
func NewURLSessionUploadTask() URLSessionUploadTask {
	return getURLSessionUploadTaskClass().New()
}




