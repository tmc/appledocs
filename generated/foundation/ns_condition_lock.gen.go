// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSConditionLock */


/* debug [class_header]: Header for NSConditionLock */
// The class instance for the [ConditionLock] class.
var (
	ConditionLockClass     _ConditionLockClass
	ConditionLockClassOnce sync.Once
)

func getConditionLockClass() _ConditionLockClass {
	ConditionLockClassOnce.Do(func() {
		ConditionLockClass = _ConditionLockClass{objc.GetClass("NSConditionLock")}
	})
	return ConditionLockClass
}

type _ConditionLockClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ConditionLock */
// An interface definition for the [ConditionLock] class.
type IConditionLock interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ConditionLock */
	// properties:
	Condition() int
	Name() IString
	SetName(value IString)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ConditionLock */
	// methods:
	LockBeforeDate(limit IDate) bool
	LockWhenCondition(condition int)
	LockWhenConditionBeforeDate(condition int, limit IDate) bool
	TryLock() bool
	TryLockWhenCondition(condition int) bool
	UnlockWithCondition(condition int)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ConditionLock */
// Alloc allocates a new instance without initialization.
func (cc _ConditionLockClass) Alloc() ConditionLock {
	rv := objc.Send[ConditionLock](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ConditionLockClass) New() ConditionLock {
	rv := objc.Send[ConditionLock](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConditionLock) Init() ConditionLock {
	rv := objc.Send[ConditionLock](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConditionLock) Autorelease() ConditionLock {
	rv := objc.Send[ConditionLock](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConditionLock creates a new ConditionLock instance.
func NewConditionLock() ConditionLock {
	return getConditionLockClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ConditionLock */
// A lock that can be associated with specific, user-defined conditions.
//
// Using an object, you can ensure that a thread can acquire a lock only if a certain condition is met. Once it has acquired the lock and executed the critical section of code, the thread can relinquish the lock and set the associated condition to something new. The conditions themselves are arbitrary: you define them as needed for your application.


// A lock that can be associated with specific, user-defined conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConditionLock
type ConditionLock struct {
	objectivec.Object
}

// ConditionLockFrom constructs a [ConditionLock] from an unsafe.Pointer.
//
// A lock that can be associated with specific, user-defined conditions.
func ConditionLockFrom(ptr unsafe.Pointer) ConditionLock {
	return ConditionLock{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ConditionLock */

// Initializes a newly allocated object and sets its condition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConditionLock/init(condition:)
func NewConditionLockWithCondition(condition int) ConditionLock {
	instance := getConditionLockClass().Alloc()
	rv := objc.Send[ConditionLock](instance.ID, objc.Sel("initWithCondition:"), condition)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewConditionLockWithCondition */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ConditionLock */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ConditionLock */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ConditionLock */

// Attempts to acquire a lock before a specified moment in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConditionLock/lock(before:)
func (c_ ConditionLock) LockBeforeDate(limit IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lockBeforeDate:"), limit)
	return rv
}/* debug [instance_methods/method]: LockBeforeDate */


// Attempts to acquire a lock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConditionLock/lock(whenCondition:)
func (c_ ConditionLock) LockWhenCondition(condition int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("lockWhenCondition:"), condition)
}/* debug [instance_methods/method]: LockWhenCondition */


// Attempts to acquire a lock before a specified moment in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConditionLock/lock(whenCondition:before:)
func (c_ ConditionLock) LockWhenConditionBeforeDate(condition int, limit IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lockWhenCondition:beforeDate:"), condition, limit)
	return rv
}/* debug [instance_methods/method]: LockWhenConditionBeforeDate */


// Attempts to acquire a lock without regard to the receiver’s condition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConditionLock/try()
func (c_ ConditionLock) TryLock() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("tryLock"))
	return rv
}/* debug [instance_methods/method]: TryLock */


// Attempts to acquire a lock if the receiver’s condition is equal to the specified condition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConditionLock/tryLock(whenCondition:)
func (c_ ConditionLock) TryLockWhenCondition(condition int) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("tryLockWhenCondition:"), condition)
	return rv
}/* debug [instance_methods/method]: TryLockWhenCondition */


// Relinquishes the lock and sets the receiver’s condition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConditionLock/unlock(withCondition:)
func (c_ ConditionLock) UnlockWithCondition(condition int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("unlockWithCondition:"), condition)
}/* debug [instance_methods/method]: UnlockWithCondition */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ConditionLock */

// The condition associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConditionLock/condition
func (c_ ConditionLock) Condition() int {
	rv := objc.Send[int](c_.ID, objc.Sel("condition"))
	return rv
}/* debug [instance_properties/getter]: condition */


// The name associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConditionLock/name
func (c_ ConditionLock) Name() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConditionLock/name
func (c_ ConditionLock) SetName(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSConditionLock */


