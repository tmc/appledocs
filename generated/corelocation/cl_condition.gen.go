// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

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
		ConditionClass = _ConditionClass{objc.GetClass("CLCondition")}
	})
	return ConditionClass
}

type _ConditionClass struct {
	class objc.Class
}





// An interface definition for the [Condition] class.
type ICondition interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _ConditionClass) Alloc() Condition {
	rv := objc.Send[Condition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// The abstract base class that all other conditions derive from.


// The abstract base class that all other conditions derive from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLCondition-c.class
type Condition struct {
	objectivec.Object
}

// ConditionFrom constructs a [Condition] from an unsafe.Pointer.
//
// The abstract base class that all other conditions derive from.
func ConditionFrom(ptr unsafe.Pointer) Condition {
	return Condition{objectivec.Object{objc.ID(ptr)}}
}































