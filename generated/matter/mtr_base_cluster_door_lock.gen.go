// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterDoorLock] class.
var (
	MTRBaseClusterDoorLockClass     _MTRBaseClusterDoorLockClass
	MTRBaseClusterDoorLockClassOnce sync.Once
)

func getMTRBaseClusterDoorLockClass() _MTRBaseClusterDoorLockClass {
	MTRBaseClusterDoorLockClassOnce.Do(func() {
		MTRBaseClusterDoorLockClass = _MTRBaseClusterDoorLockClass{objc.GetClass("MTRBaseClusterDoorLock")}
	})
	return MTRBaseClusterDoorLockClass
}

type _MTRBaseClusterDoorLockClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterDoorLock] class.
type IMTRBaseClusterDoorLock interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDoorLock
type MTRBaseClusterDoorLock struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterDoorLockFrom constructs a [MTRBaseClusterDoorLock] from an unsafe.Pointer.
func MTRBaseClusterDoorLockFrom(ptr unsafe.Pointer) MTRBaseClusterDoorLock {
	return MTRBaseClusterDoorLock{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterDoorLockClass) Alloc() MTRBaseClusterDoorLock {
	rv := objc.Send[MTRBaseClusterDoorLock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterDoorLockClass) New() MTRBaseClusterDoorLock {
	rv := objc.Send[MTRBaseClusterDoorLock](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterDoorLock) Init() MTRBaseClusterDoorLock {
	rv := objc.Send[MTRBaseClusterDoorLock](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterDoorLock) Autorelease() MTRBaseClusterDoorLock {
	rv := objc.Send[MTRBaseClusterDoorLock](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterDoorLock creates a new MTRBaseClusterDoorLock instance.
func NewMTRBaseClusterDoorLock() MTRBaseClusterDoorLock {
	return getMTRBaseClusterDoorLockClass().New()
}




