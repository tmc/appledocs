// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterDoorLock] class.
var (
	MTRClusterDoorLockClass     _MTRClusterDoorLockClass
	MTRClusterDoorLockClassOnce sync.Once
)

func getMTRClusterDoorLockClass() _MTRClusterDoorLockClass {
	MTRClusterDoorLockClassOnce.Do(func() {
		MTRClusterDoorLockClass = _MTRClusterDoorLockClass{objc.GetClass("MTRClusterDoorLock")}
	})
	return MTRClusterDoorLockClass
}

type _MTRClusterDoorLockClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterDoorLock] class.
type IMTRClusterDoorLock interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDoorLock
type MTRClusterDoorLock struct {
	MTRGenericCluster
}

// MTRClusterDoorLockFrom constructs a [MTRClusterDoorLock] from an unsafe.Pointer.
func MTRClusterDoorLockFrom(ptr unsafe.Pointer) MTRClusterDoorLock {
	return MTRClusterDoorLock{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterDoorLockClass) Alloc() MTRClusterDoorLock {
	rv := objc.Send[MTRClusterDoorLock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterDoorLockClass) New() MTRClusterDoorLock {
	rv := objc.Send[MTRClusterDoorLock](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterDoorLock) Init() MTRClusterDoorLock {
	rv := objc.Send[MTRClusterDoorLock](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterDoorLock) Autorelease() MTRClusterDoorLock {
	rv := objc.Send[MTRClusterDoorLock](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterDoorLock creates a new MTRClusterDoorLock instance.
func NewMTRClusterDoorLock() MTRClusterDoorLock {
	return getMTRClusterDoorLockClass().New()
}




