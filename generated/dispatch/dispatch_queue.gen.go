// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DispatchQueue] class.
var (
	DispatchQueueClass     _DispatchQueueClass
	DispatchQueueClassOnce sync.Once
)

func getDispatchQueueClass() _DispatchQueueClass {
	DispatchQueueClassOnce.Do(func() {
		DispatchQueueClass = _DispatchQueueClass{objc.GetClass("DispatchQueue")}
	})
	return DispatchQueueClass
}

type _DispatchQueueClass struct {
	class objc.Class
}

// An interface definition for the [DispatchQueue] class.
type IDispatchQueue interface {
	objectivec.IObject
}

// An object that manages the execution of tasks serially or concurrently on your app’s main thread or on a background thread.
//
// Dispatch queues are FIFO queues to which your application can submit tasks in the form of block objects. Dispatch queues execute tasks either serially or concurrently. Work submitted to dispatch queues executes on a pool of threads managed by the system. Except for the dispatch queue representing your app’s main thread, the system makes no guarantees about which thread it uses to execute a task. You schedule work items synchronously or asynchronously. When you schedule a work item synchronously, your code waits until that item finishes execution. When you schedule a work item asynchronously, your code continues executing while the work item runs elsewhere.


// An object that manages the execution of tasks serially or concurrently on your app’s main thread or on a background thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchQueue
type DispatchQueue struct {
	objectivec.Object
}

// DispatchQueueFrom constructs a [DispatchQueue] from an unsafe.Pointer.
//
// An object that manages the execution of tasks serially or concurrently on your app’s main thread or on a background thread.
func DispatchQueueFrom(ptr unsafe.Pointer) DispatchQueue {
	return DispatchQueue{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DispatchQueueClass) Alloc() DispatchQueue {
	rv := objc.Send[DispatchQueue](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DispatchQueueClass) New() DispatchQueue {
	rv := objc.Send[DispatchQueue](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchQueue) Init() DispatchQueue {
	rv := objc.Send[DispatchQueue](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchQueue) Autorelease() DispatchQueue {
	rv := objc.Send[DispatchQueue](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchQueue creates a new DispatchQueue instance.
func NewDispatchQueue() DispatchQueue {
	return getDispatchQueueClass().New()
}




