// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DistributedLock] class.
var (
	distributedLockClass     _DistributedLockClass
	distributedLockClassOnce sync.Once
)

func getDistributedLockClass() _DistributedLockClass {
	distributedLockClassOnce.Do(func() {
		distributedLockClass = _DistributedLockClass{objc.GetClass("NSDistributedLock")}
	})
	return distributedLockClass
}

type _DistributedLockClass struct {
	class objc.Class
}

// An interface definition for the [DistributedLock] class.
type IDistributedLock interface {
	objectivec.IObject
}

// A lock that multiple applications on multiple hosts can use to restrict access to some shared resource, such as a file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistributedLock
type DistributedLock struct {
	objectivec.Object
}

// DistributedLockFrom constructs a [DistributedLock] from an unsafe.Pointer.
//
// A lock that multiple applications on multiple hosts can use to restrict access to some shared resource, such as a file.
func DistributedLockFrom(ptr unsafe.Pointer) DistributedLock {
	return DistributedLock{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DistributedLockClass) Alloc() DistributedLock {
	rv := objc.Send[DistributedLock](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DistributedLockClass) New() DistributedLock {
	rv := objc.Send[DistributedLock](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DistributedLock) Init() DistributedLock {
	rv := objc.Send[DistributedLock](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DistributedLock) Autorelease() DistributedLock {
	rv := objc.Send[DistributedLock](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDistributedLock creates a new DistributedLock instance.
func NewDistributedLock() DistributedLock {
	return getDistributedLockClass().New()
}




