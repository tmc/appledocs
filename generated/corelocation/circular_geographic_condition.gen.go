// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CircularGeographicCondition] class.
var circularGeographicConditionClass = _CircularGeographicConditionClass{objc.GetClass("CLCircularGeographicCondition")}

type _CircularGeographicConditionClass struct {
	class objc.Class
}

// A circular geographic condition that a center point and radius define. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLCircularGeographicCondition

type CircularGeographicCondition struct {
	Condition
}

// CircularGeographicConditionFrom constructs a [CircularGeographicCondition] from an unsafe.Pointer.
//
// A circular geographic condition that a center point and radius define.
func CircularGeographicConditionFrom(ptr unsafe.Pointer) CircularGeographicCondition {
	return CircularGeographicCondition{
		Condition: ConditionFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (cc _CircularGeographicConditionClass) Alloc() CircularGeographicCondition {
	rv := objc.Send[CircularGeographicCondition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _CircularGeographicConditionClass) New() CircularGeographicCondition {
	rv := objc.Send[CircularGeographicCondition](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CircularGeographicCondition) Init() CircularGeographicCondition {
	rv := objc.Send[CircularGeographicCondition](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CircularGeographicCondition) Autorelease() CircularGeographicCondition {
	rv := objc.Send[CircularGeographicCondition](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCircularGeographicCondition creates a new CircularGeographicCondition instance.
func NewCircularGeographicCondition() CircularGeographicCondition {
	return circularGeographicConditionClass.New()
}
// Creates a new circular geographic condition with the center point and radius you provide. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLCircularGeographicCondition/initWithCenter:radius:
func NewCircularGeographicConditionWithCenterRadius(center unsafe.Pointer, radius unsafe.Pointer) CircularGeographicCondition {
	instance := circularGeographicConditionClass.Alloc()
	rv := objc.Send[CircularGeographicCondition](instance.ID, objc.Sel("initWithCenter:radius:"), center, radius)
	rv.Autorelease()
	return rv
}



