// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAsyncCallbackWorkQueue] class.
var (
	MTRAsyncCallbackWorkQueueClass     _MTRAsyncCallbackWorkQueueClass
	MTRAsyncCallbackWorkQueueClassOnce sync.Once
)

func getMTRAsyncCallbackWorkQueueClass() _MTRAsyncCallbackWorkQueueClass {
	MTRAsyncCallbackWorkQueueClassOnce.Do(func() {
		MTRAsyncCallbackWorkQueueClass = _MTRAsyncCallbackWorkQueueClass{objc.GetClass("MTRAsyncCallbackWorkQueue")}
	})
	return MTRAsyncCallbackWorkQueueClass
}

type _MTRAsyncCallbackWorkQueueClass struct {
	class objc.Class
}

// An interface definition for the [MTRAsyncCallbackWorkQueue] class.
type IMTRAsyncCallbackWorkQueue interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAsyncCallbackWorkQueue
type MTRAsyncCallbackWorkQueue struct {
	objectivec.Object
}

// MTRAsyncCallbackWorkQueueFrom constructs a [MTRAsyncCallbackWorkQueue] from an unsafe.Pointer.
func MTRAsyncCallbackWorkQueueFrom(ptr unsafe.Pointer) MTRAsyncCallbackWorkQueue {
	return MTRAsyncCallbackWorkQueue{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAsyncCallbackWorkQueueClass) Alloc() MTRAsyncCallbackWorkQueue {
	rv := objc.Send[MTRAsyncCallbackWorkQueue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAsyncCallbackWorkQueueClass) New() MTRAsyncCallbackWorkQueue {
	rv := objc.Send[MTRAsyncCallbackWorkQueue](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAsyncCallbackWorkQueue) Init() MTRAsyncCallbackWorkQueue {
	rv := objc.Send[MTRAsyncCallbackWorkQueue](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAsyncCallbackWorkQueue) Autorelease() MTRAsyncCallbackWorkQueue {
	rv := objc.Send[MTRAsyncCallbackWorkQueue](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAsyncCallbackWorkQueue creates a new MTRAsyncCallbackWorkQueue instance.
func NewMTRAsyncCallbackWorkQueue() MTRAsyncCallbackWorkQueue {
	return getMTRAsyncCallbackWorkQueueClass().New()
}




