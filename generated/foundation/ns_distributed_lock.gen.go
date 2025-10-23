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
	DistributedLockClass     _DistributedLockClass
	DistributedLockClassOnce sync.Once
)

func getDistributedLockClass() _DistributedLockClass {
	DistributedLockClassOnce.Do(func() {
		DistributedLockClass = _DistributedLockClass{objc.GetClass("NSDistributedLock")}
	})
	return DistributedLockClass
}

type _DistributedLockClass struct {
	class objc.Class
}

// An interface definition for the [DistributedLock] class.
type IDistributedLock interface {
	objectivec.IObject
	BreakLock()
	TryLock() bool
	Unlock()
	LockDate() IDate
}

// A lock that multiple applications on multiple hosts can use to restrict access to some shared resource, such as a file.
//
// The lock is implemented by an entry (such as a file or directory) in the file system. For multiple applications to use an object to coordinate their activities, the lock must be writable on a file system accessible to all hosts on which the applications might be running. Use the method to attempt to acquire a lock. You should generally use the method to release the lock rather than . doesn’t conform to the protocol, nor does it have a method. The protocol’s method is intended to block the execution of the thread until successful. For an object, this could mean polling the file system at some predetermined rate. A better solution is to provide the method and let you determine the polling frequency that makes sense for your application.


// A lock that multiple applications on multiple hosts can use to restrict access to some shared resource, such as a file.
//
// [Full Topic]
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



// Initializes an object to use as the lock the file-system entry specified by a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistributedLock/init(path:)
func NewDistributedLockWithPath(path string) DistributedLock {
	instance := getDistributedLockClass().Alloc()
	rv := objc.Send[DistributedLock](instance.ID, objc.Sel("initWithPath:"), objc.String(path))
	rv.Autorelease()
	return rv
}



// Returns an object initialized to use as the locking object the file-system entry specified by a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistributedLock/lockWithPath:
func (dc _DistributedLockClass) LockWithPath(path string) IDistributedLock {
	rv := objc.Send[DistributedLock](objc.ID(dc.class), objc.Sel("lockWithPath:"), objc.String(path))
	return rv
}


// Forces the lock to be relinquished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistributedLock/break()
func (d_ DistributedLock) BreakLock() {
	objc.Send[objc.ID](d_.ID, objc.Sel("breakLock"))
}


// Attempts to acquire the receiver and immediately returns a Boolean value that indicates whether the attempt was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistributedLock/try()
func (d_ DistributedLock) TryLock() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("tryLock"))
	return rv
}


// Relinquishes the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistributedLock/unlock()
func (d_ DistributedLock) Unlock() {
	objc.Send[objc.ID](d_.ID, objc.Sel("unlock"))
}


// Returns the time the receiver was acquired by any of the objects using the same path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistributedLock/lockDate
func (d_ DistributedLock) LockDate() IDate {
	rv := objc.Send[NSDate](d_.ID, objc.Sel("lockDate"))
	return rv
}


