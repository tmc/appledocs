// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OS_dispatch_queue_global] class.
var (
	OS_dispatch_queue_globalClass     _OS_dispatch_queue_globalClass
	OS_dispatch_queue_globalClassOnce sync.Once
)

func getOS_dispatch_queue_globalClass() _OS_dispatch_queue_globalClass {
	OS_dispatch_queue_globalClassOnce.Do(func() {
		OS_dispatch_queue_globalClass = _OS_dispatch_queue_globalClass{objc.GetClass("OS_dispatch_queue_global")}
	})
	return OS_dispatch_queue_globalClass
}

type _OS_dispatch_queue_globalClass struct {
	class objc.Class
}

// An interface definition for the [OS_dispatch_queue_global] class.
type IOS_dispatch_queue_global interface {
	objectivec.IObject
}

// A system-provided dispatch queue that schedules tasks for concurrent execution.
//
// You do not create objects of this type directly. You receive a queue of the appropriate type when you create a new object.


// A system-provided dispatch queue that schedules tasks for concurrent execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/OS_dispatch_queue_global-swift.class

type OS_dispatch_queue_global struct {
	objectivec.Object
}

// OS_dispatch_queue_globalFrom constructs a [OS_dispatch_queue_global] from an unsafe.Pointer.
//
// A system-provided dispatch queue that schedules tasks for concurrent execution.
func OS_dispatch_queue_globalFrom(ptr unsafe.Pointer) OS_dispatch_queue_global {
	return OS_dispatch_queue_global{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OS_dispatch_queue_globalClass) Alloc() OS_dispatch_queue_global {
	rv := objc.Send[OS_dispatch_queue_global](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OS_dispatch_queue_globalClass) New() OS_dispatch_queue_global {
	rv := objc.Send[OS_dispatch_queue_global](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OS_dispatch_queue_global) Init() OS_dispatch_queue_global {
	rv := objc.Send[OS_dispatch_queue_global](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OS_dispatch_queue_global) Autorelease() OS_dispatch_queue_global {
	rv := objc.Send[OS_dispatch_queue_global](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOS_dispatch_queue_global creates a new OS_dispatch_queue_global instance.
func NewOS_dispatch_queue_global() OS_dispatch_queue_global {
	return getOS_dispatch_queue_globalClass().New()
}




