// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Condition] class.
var conditionClass = _ConditionClass{objc.GetClass("NSCondition")}

type _ConditionClass struct {
	class objc.Class
}

// An interface definition for the [Condition] class.
type ICondition interface {
	objectivec.IObject
}

// A condition variable whose semantics follow those used for POSIX-style conditions. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return conditionClass.New()
}




