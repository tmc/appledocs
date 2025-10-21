// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Condition] class.
var (
	ConditionClass     _ConditionClass
	ConditionClassOnce sync.Once
)

func getConditionClass() _ConditionClass {
	ConditionClassOnce.Do(func() {
		ConditionClass = _ConditionClass{objc.GetClass("NSCondition")}
	})
	return ConditionClass
}

type _ConditionClass struct {
	class objc.Class
}

// An interface definition for the [Condition] class.
type ICondition interface {
	objectivec.IObject
	Broadcast()
	Wait()
}

// A condition variable whose semantics follow those used for POSIX-style conditions.
//
// A condition object acts as both a lock and a checkpoint in a given thread. The lock protects your code while it tests the condition and performs the task triggered by the condition. The checkpoint behavior requires that the condition be true before the thread proceeds with its task. While the condition is not true, the thread blocks. It remains blocked until another thread signals the condition object. The semantics for using an object are as follows: Lock the condition object. Test a boolean predicate. (This predicate is a boolean flag or other variable in your code that indicates whether it is safe to perform the task protected by the condition.) If the boolean predicate is false, call the condition object’s or method to block the thread. Upon returning from these methods, go to step 2 to retest your boolean predicate. (Continue waiting and retesting the predicate until it is true.) If the boolean predicate is true, perform the task. Optionally update any predicates (or signal any conditions) affected by your task. When your task is done, unlock the condition object. The pseudocode for performing the preceding steps would therefore look something like the following: Whenever you use a condition object, the first step is to lock the condition. Locking the condition ensures that your predicate and task code are protected from interference by other threads using the same condition. Once you have completed your task, you can set other predicates or signal other conditions based on the needs of your code. You should always set predicates and signal conditions while holding the condition object’s lock. When a thread waits on a condition, the condition object unlocks its lock and blocks the thread. When the condition is signaled, the system wakes up the thread. The condition object then reacquires its lock before returning from the or method. Thus, from the point of view of the thread, it is as if it always held the lock. A boolean predicate is an important part of the semantics of using conditions because of the way signaling works. Signaling a condition does not guarantee that the condition itself is true. There are timing issues involved in signaling that may cause false signals to appear. Using a predicate ensures that these spurious signals do not cause you to perform work before it is safe to do so. The predicate itself is simply a flag or other variable in your code that you test in order to acquire a Boolean result. For more information on how to use conditions, see Using POSIX Thread Locks in .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCondition
type Condition struct {
	objectivec.Object
}

// ConditionFrom constructs a [Condition] from an unsafe.Pointer.
//
// A condition variable whose semantics follow those used for POSIX-style conditions.
func ConditionFrom(ptr unsafe.Pointer) Condition {
	return Condition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ConditionClass) Alloc() Condition {
	rv := objc.Send[Condition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ConditionClass) New() Condition {
	rv := objc.Send[Condition](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Condition) Init() Condition {
	rv := objc.Send[Condition](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Condition) Autorelease() Condition {
	rv := objc.Send[Condition](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCondition creates a new Condition instance.
func NewCondition() Condition {
	return getConditionClass().New()
}


// Signals the condition, waking up all threads waiting on it.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCondition/broadcast()
func (c_ Condition) Broadcast() {
	objc.Send[objc.ID](c_.ID, objc.Sel("broadcast"))
}

// Blocks the current thread until the condition is signaled.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCondition/wait()
func (c_ Condition) Wait() {
	objc.Send[objc.ID](c_.ID, objc.Sel("wait"))
}

// The name of the condition.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscondition/name
func (c_ Condition) Name() string {
	rv := objc.Send[string](c_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the condition.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscondition/name
func (c_ Condition) SetName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), objc.String(value))
}



