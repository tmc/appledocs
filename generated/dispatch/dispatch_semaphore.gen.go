// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DispatchSemaphore] class.
var (
	DispatchSemaphoreClass     _DispatchSemaphoreClass
	DispatchSemaphoreClassOnce sync.Once
)

func getDispatchSemaphoreClass() _DispatchSemaphoreClass {
	DispatchSemaphoreClassOnce.Do(func() {
		DispatchSemaphoreClass = _DispatchSemaphoreClass{objc.GetClass("DispatchSemaphore")}
	})
	return DispatchSemaphoreClass
}

type _DispatchSemaphoreClass struct {
	class objc.Class
}

// An interface definition for the [DispatchSemaphore] class.
type IDispatchSemaphore interface {
	objectivec.IObject
}

// An object that controls access to a resource across multiple execution contexts through use of a traditional counting semaphore.
//
// A dispatch semaphore is an efficient implementation of a traditional counting semaphore. Dispatch semaphores call down to the kernel only when the calling thread needs to be blocked. If the calling semaphore does not need to block, no kernel call is made. You increment a semaphore count by calling the method, and decrement a semaphore count by calling or one of its variants that specifies a timeout.


// An object that controls access to a resource across multiple execution contexts through use of a traditional counting semaphore.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchSemaphore

type DispatchSemaphore struct {
	objectivec.Object
}

// DispatchSemaphoreFrom constructs a [DispatchSemaphore] from an unsafe.Pointer.
//
// An object that controls access to a resource across multiple execution contexts through use of a traditional counting semaphore.
func DispatchSemaphoreFrom(ptr unsafe.Pointer) DispatchSemaphore {
	return DispatchSemaphore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DispatchSemaphoreClass) Alloc() DispatchSemaphore {
	rv := objc.Send[DispatchSemaphore](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DispatchSemaphoreClass) New() DispatchSemaphore {
	rv := objc.Send[DispatchSemaphore](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchSemaphore) Init() DispatchSemaphore {
	rv := objc.Send[DispatchSemaphore](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchSemaphore) Autorelease() DispatchSemaphore {
	rv := objc.Send[DispatchSemaphore](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchSemaphore creates a new DispatchSemaphore instance.
func NewDispatchSemaphore() DispatchSemaphore {
	return getDispatchSemaphoreClass().New()
}




