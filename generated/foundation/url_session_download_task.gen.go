// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionDownloadTask] class.
var (
	uRLSessionDownloadTaskClass     _URLSessionDownloadTaskClass
	uRLSessionDownloadTaskClassOnce sync.Once
)

func getURLSessionDownloadTaskClass() _URLSessionDownloadTaskClass {
	uRLSessionDownloadTaskClassOnce.Do(func() {
		uRLSessionDownloadTaskClass = _URLSessionDownloadTaskClass{objc.GetClass("NSURLSessionDownloadTask")}
	})
	return uRLSessionDownloadTaskClass
}

type _URLSessionDownloadTaskClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionDownloadTask] class.
type IURLSessionDownloadTask interface {
	IURLSessionTask
}

// A URL session task that stores downloaded data to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionDownloadTask
type URLSessionDownloadTask struct {
	URLSessionTask
}

// URLSessionDownloadTaskFrom constructs a [URLSessionDownloadTask] from an unsafe.Pointer.
//
// A URL session task that stores downloaded data to a file.
func URLSessionDownloadTaskFrom(ptr unsafe.Pointer) URLSessionDownloadTask {
	return URLSessionDownloadTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _URLSessionDownloadTaskClass) Alloc() URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLSessionDownloadTaskClass) New() URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionDownloadTask) Init() URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionDownloadTask) Autorelease() URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionDownloadTask creates a new URLSessionDownloadTask instance.
func NewURLSessionDownloadTask() URLSessionDownloadTask {
	return getURLSessionDownloadTaskClass().New()
}




