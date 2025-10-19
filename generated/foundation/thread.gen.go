// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Thread] class.
var (
	threadClass     _ThreadClass
	threadClassOnce sync.Once
)

func getThreadClass() _ThreadClass {
	threadClassOnce.Do(func() {
		threadClass = _ThreadClass{objc.GetClass("NSThread")}
	})
	return threadClass
}

type _ThreadClass struct {
	class objc.Class
}

// An interface definition for the [Thread] class.
type IThread interface {
	objectivec.IObject
}

// A thread of execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread
type Thread struct {
	objectivec.Object
}

// ThreadFrom constructs a [Thread] from an unsafe.Pointer.
//
// A thread of execution.
func ThreadFrom(ptr unsafe.Pointer) Thread {
	return Thread{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _ThreadClass) Alloc() Thread {
	rv := objc.Send[Thread](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ThreadClass) New() Thread {
	rv := objc.Send[Thread](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Thread) Init() Thread {
	rv := objc.Send[Thread](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Thread) Autorelease() Thread {
	rv := objc.Send[Thread](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewThread creates a new Thread instance.
func NewThread() Thread {
	return getThreadClass().New()
}




