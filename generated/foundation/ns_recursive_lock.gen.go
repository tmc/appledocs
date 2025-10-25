// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSRecursiveLock */


/* debug [class_header]: Header for NSRecursiveLock */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecursiveLock */
// An interface definition for the [RecursiveLock] class.
type IRecursiveLock interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RecursiveLock */
	// properties:
	Name() IString
	SetName(value IString)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecursiveLock */
	// methods:
	LockBeforeDate(limit IDate) bool
	TryLock() bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecursiveLock */
// Alloc allocates a new instance without initialization.
func (rc _RecursiveLockClass) Alloc() RecursiveLock {
	rv := objc.Send[RecursiveLock](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecursiveLock */
// A lock that may be acquired multiple times by the same thread without causing a deadlock.
//
// defines a lock that may be acquired multiple times by the same thread without causing a deadlock, a situation where a thread is permanently blocked waiting for itself to relinquish a lock. While the locking thread has one or more locks, all other threads are prevented from accessing the code protected by the lock.


// A lock that may be acquired multiple times by the same thread without causing a deadlock.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecursiveLock *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecursiveLock */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecursiveLock */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecursiveLock */

// Attempts to acquire a lock before a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecursiveLock/lock(before:)
func (r_ RecursiveLock) LockBeforeDate(limit IDate) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("lockBeforeDate:"), limit)
	return rv
}/* debug [instance_methods/method]: LockBeforeDate */


// Attempts to acquire a lock, and immediately returns a Boolean value that indicates whether the attempt was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecursiveLock/try()
func (r_ RecursiveLock) TryLock() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("tryLock"))
	return rv
}/* debug [instance_methods/method]: TryLock */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecursiveLock */

// The name associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecursiveLock/name
func (r_ RecursiveLock) Name() IString {
	rv := objc.Send[String](r_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecursiveLock/name
func (r_ RecursiveLock) SetName(value IString) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSRecursiveLock */



