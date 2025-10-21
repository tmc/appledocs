// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RecursiveLock] class.
var (
	RecursiveLockClass     _RecursiveLockClass
	RecursiveLockClassOnce sync.Once
)

func getRecursiveLockClass() _RecursiveLockClass {
	RecursiveLockClassOnce.Do(func() {
		RecursiveLockClass = _RecursiveLockClass{objc.GetClass("NSRecursiveLock")}
	})
	return RecursiveLockClass
}

type _RecursiveLockClass struct {
	class objc.Class
}

// An interface definition for the [RecursiveLock] class.
type IRecursiveLock interface {
	objectivec.IObject
	TryLock() bool
}

// A lock that may be acquired multiple times by the same thread without causing a deadlock.
//
// defines a lock that may be acquired multiple times by the same thread without causing a deadlock, a situation where a thread is permanently blocked waiting for itself to relinquish a lock. While the locking thread has one or more locks, all other threads are prevented from accessing the code protected by the lock.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecursiveLock
type RecursiveLock struct {
	objectivec.Object
}

// RecursiveLockFrom constructs a [RecursiveLock] from an unsafe.Pointer.
//
// A lock that may be acquired multiple times by the same thread without causing a deadlock.
func RecursiveLockFrom(ptr unsafe.Pointer) RecursiveLock {
	return RecursiveLock{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RecursiveLockClass) Alloc() RecursiveLock {
	rv := objc.Send[RecursiveLock](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RecursiveLockClass) New() RecursiveLock {
	rv := objc.Send[RecursiveLock](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecursiveLock) Init() RecursiveLock {
	rv := objc.Send[RecursiveLock](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecursiveLock) Autorelease() RecursiveLock {
	rv := objc.Send[RecursiveLock](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecursiveLock creates a new RecursiveLock instance.
func NewRecursiveLock() RecursiveLock {
	return getRecursiveLockClass().New()
}


// Attempts to acquire a lock, and immediately returns a Boolean value that indicates whether the attempt was successful.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecursiveLock/try()
func (r_ RecursiveLock) TryLock() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("tryLock"))
	return rv
}

// The name associated with the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecursiveLock/name
func (r_ RecursiveLock) Name() appkit.string {
	rv := objc.Send[appkit.string](r_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name associated with the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecursiveLock/name
func (r_ RecursiveLock) SetName(value appkit.string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setName:"), value)
}



