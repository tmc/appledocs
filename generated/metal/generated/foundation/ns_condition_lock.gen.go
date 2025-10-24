// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ConditionLock] class.
type IConditionLock interface {
	objectivec.IObject
	// properties:
	Condition() int /* primitive/slice/pointer. */
	SetCondition(value int /* primitive/slice/pointer. */)
	Name() IString
	SetName(value IString)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (cc _ConditionLockClass) Alloc() ConditionLock {
	rv := objc.Send[ConditionLock](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The condition associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/condition
func (c_ ConditionLock) Condition() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("condition"))
	return rv
}


// The condition associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/condition
func (c_ ConditionLock) SetCondition(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCondition:"), value)
}


// The name associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/name
func (c_ ConditionLock) Name() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("name"))
	return rv
}


// The name associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/name
func (c_ ConditionLock) SetName(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), value)
}



