// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OS_dispatch_queue_main] class.
var (
	OS_dispatch_queue_mainClass     _OS_dispatch_queue_mainClass
	OS_dispatch_queue_mainClassOnce sync.Once
)

func getOS_dispatch_queue_mainClass() _OS_dispatch_queue_mainClass {
	OS_dispatch_queue_mainClassOnce.Do(func() {
		OS_dispatch_queue_mainClass = _OS_dispatch_queue_mainClass{objc.GetClass("OS_dispatch_queue_main")}
	})
	return OS_dispatch_queue_mainClass
}

type _OS_dispatch_queue_mainClass struct {
	class objc.Class
}

// An interface definition for the [OS_dispatch_queue_main] class.
type IOS_dispatch_queue_main interface {
	objectivec.IObject
}

// A system-provided dispatch queue that schedules tasks for serial execution on the app’s main thread.
//
// You do not create objects of this type directly. You receive a queue of the appropriate type when you create a new object.


// A system-provided dispatch queue that schedules tasks for serial execution on the app’s main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/OS_dispatch_queue_main-swift.class

type OS_dispatch_queue_main struct {
	objectivec.Object
}

// OS_dispatch_queue_mainFrom constructs a [OS_dispatch_queue_main] from an unsafe.Pointer.
//
// A system-provided dispatch queue that schedules tasks for serial execution on the app’s main thread.
func OS_dispatch_queue_mainFrom(ptr unsafe.Pointer) OS_dispatch_queue_main {
	return OS_dispatch_queue_main{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OS_dispatch_queue_mainClass) Alloc() OS_dispatch_queue_main {
	rv := objc.Send[OS_dispatch_queue_main](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OS_dispatch_queue_mainClass) New() OS_dispatch_queue_main {
	rv := objc.Send[OS_dispatch_queue_main](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OS_dispatch_queue_main) Init() OS_dispatch_queue_main {
	rv := objc.Send[OS_dispatch_queue_main](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OS_dispatch_queue_main) Autorelease() OS_dispatch_queue_main {
	rv := objc.Send[OS_dispatch_queue_main](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOS_dispatch_queue_main creates a new OS_dispatch_queue_main instance.
func NewOS_dispatch_queue_main() OS_dispatch_queue_main {
	return getOS_dispatch_queue_mainClass().New()
}





