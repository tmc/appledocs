// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSessionDataTask] class.
var (
	URLSessionDataTaskClass     _URLSessionDataTaskClass
	URLSessionDataTaskClassOnce sync.Once
)

func getURLSessionDataTaskClass() _URLSessionDataTaskClass {
	URLSessionDataTaskClassOnce.Do(func() {
		URLSessionDataTaskClass = _URLSessionDataTaskClass{objc.GetClass("NSURLSessionDataTask")}
	})
	return URLSessionDataTaskClass
}

type _URLSessionDataTaskClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionDataTask] class.
type IURLSessionDataTask interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other Foundation classes.


// A parent class referenced by other Foundation classes. [Full Topic]
type URLSessionDataTask struct {
	objectivec.Object
}

// URLSessionDataTaskFrom constructs a [URLSessionDataTask] from an unsafe.Pointer.
//
// A parent class referenced by other Foundation classes.
func URLSessionDataTaskFrom(ptr unsafe.Pointer) URLSessionDataTask {
	return URLSessionDataTask{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLSessionDataTaskClass) Alloc() URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLSessionDataTaskClass) New() URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionDataTask) Init() URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionDataTask) Autorelease() URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionDataTask creates a new URLSessionDataTask instance.
func NewURLSessionDataTask() URLSessionDataTask {
	return getURLSessionDataTaskClass().New()
}




