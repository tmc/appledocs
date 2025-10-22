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
	LockClass     _LockClass
	LockClassOnce sync.Once
)

func getLockClass() _LockClass {
	LockClassOnce.Do(func() {
		LockClass = _LockClass{objc.GetClass("NSLock")}
	})
	return LockClass
}

type _LockClass struct {
	class objc.Class
}

// An interface definition for the [Lock] class.
type ILock interface {
	objectivec.IObject
	LockBeforeDate(limit IDate) bool
	TryLock() bool
}

// An object that coordinates the operation of multiple threads of execution within the same application.
//
// An object can be used to mediate access to an application’s global data or to protect a critical section of code, allowing it to run atomically. You should not use this class to implement a recursive lock. Calling the method twice on the same thread will lock up your thread permanently. Use the class to implement recursive locks instead. Unlocking a lock that is not locked is considered a programmer error and should be fixed in your code. The class reports such errors by printing an error message to the console when they occur.
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


// Attempts to acquire a lock before a given time and returns a Boolean value indicating whether the attempt was successful.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLock/lock(before:)
func (l_ Lock) LockBeforeDate(limit IDate) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("lockBeforeDate:"), limit)
	return rv
}

// Attempts to acquire a lock and immediately returns a Boolean value that indicates whether the attempt was successful.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLock/try()
func (l_ Lock) TryLock() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("tryLock"))
	return rv
}

// The name associated with the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLock/name
func (l_ Lock) Name() string {
	rv := objc.Send[string](l_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name associated with the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLock/name
func (l_ Lock) SetName(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setName:"), objc.String(value))
}



