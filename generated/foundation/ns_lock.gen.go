// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSLock */


/* debug [class_header]: Header for NSLock */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Lock */
// An interface definition for the [Lock] class.
type ILock interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Lock */
	// properties:
	Name() IString
	SetName(value IString)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Lock */
	// methods:
	LockBeforeDate(limit IDate) bool
	TryLock() bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Lock */
// Alloc allocates a new instance without initialization.
func (lc _LockClass) Alloc() Lock {
	rv := objc.Send[Lock](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Lock */
// An object that coordinates the operation of multiple threads of execution within the same application.
//
// An object can be used to mediate access to an application’s global data or to protect a critical section of code, allowing it to run atomically. You should not use this class to implement a recursive lock. Calling the method twice on the same thread will lock up your thread permanently. Use the class to implement recursive locks instead. Unlocking a lock that is not locked is considered a programmer error and should be fixed in your code. The class reports such errors by printing an error message to the console when they occur.


// An object that coordinates the operation of multiple threads of execution within the same application.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Lock *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Lock */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Lock */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Lock */

// Attempts to acquire a lock before a given time and returns a Boolean value indicating whether the attempt was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLock/lock(before:)
func (l_ Lock) LockBeforeDate(limit IDate) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("lockBeforeDate:"), limit)
	return rv
}/* debug [instance_methods/method]: LockBeforeDate */


// Attempts to acquire a lock and immediately returns a Boolean value that indicates whether the attempt was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLock/try()
func (l_ Lock) TryLock() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("tryLock"))
	return rv
}/* debug [instance_methods/method]: TryLock */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Lock */

// The name associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLock/name
func (l_ Lock) Name() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLock/name
func (l_ Lock) SetName(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSLock */



