// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Lock] class.
var (
	lockClass     _LockClass
	lockClassOnce sync.Once
)

func getLockClass() _LockClass {
	lockClassOnce.Do(func() {
		lockClass = _LockClass{objc.GetClass("NSLock")}
	})
	return lockClass
}

type _LockClass struct {
	class objc.Class
}

// An interface definition for the [Lock] class.
type ILock interface {
	objectivec.IObject
}

// An object that coordinates the operation of multiple threads of execution within the same application.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLock
type Lock struct {
	objectivec.Object
}

// LockFrom constructs a [Lock] from an unsafe.Pointer.
//
// An object that coordinates the operation of multiple threads of execution within the same application.
func LockFrom(ptr unsafe.Pointer) Lock {
	return Lock{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LockClass) Alloc() Lock {
	rv := objc.Send[Lock](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LockClass) New() Lock {
	rv := objc.Send[Lock](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ Lock) Init() Lock {
	rv := objc.Send[Lock](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ Lock) Autorelease() Lock {
	rv := objc.Send[Lock](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLock creates a new Lock instance.
func NewLock() Lock {
	return getLockClass().New()
}




