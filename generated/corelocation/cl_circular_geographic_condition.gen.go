// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CircularGeographicCondition] class.
var (
	CircularGeographicConditionClass     _CircularGeographicConditionClass
	CircularGeographicConditionClassOnce sync.Once
)

func getCircularGeographicConditionClass() _CircularGeographicConditionClass {
	CircularGeographicConditionClassOnce.Do(func() {
		CircularGeographicConditionClass = _CircularGeographicConditionClass{objc.GetClass("CLCircularGeographicCondition")}
	})
	return CircularGeographicConditionClass
}

type _CircularGeographicConditionClass struct {
	class objc.Class
}

// An interface definition for the [CircularGeographicCondition] class.
type ICircularGeographicCondition interface {
	ICondition
	Center() unsafe.Pointer
	Radius() unsafe.Pointer
}

// A circular geographic condition that a center point and radius define.
//
// Use to monitor events that occur in a circular geographic condition that you describe.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getCircularGeographicConditionClass().New()
}




// Creates a new circular geographic condition with the center point and radius you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLCircularGeographicCondition/initWithCenter:radius:
func NewCircularGeographicConditionWithCenterRadius(center unsafe.Pointer, radius unsafe.Pointer) CircularGeographicCondition {
	instance := getCircularGeographicConditionClass().Alloc()
	rv := objc.Send[CircularGeographicCondition](instance.ID, objc.Sel("initWithCenter:radius:"), center, radius)
	rv.Autorelease()
	return rv
}


// The center of the circular geographic condition.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLCircularGeographicCondition/center
func (c_ CircularGeographicCondition) Center() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("center"))
	return rv
}

// The radius of the circular geographic condition.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLCircularGeographicCondition/radius
func (c_ CircularGeographicCondition) Radius() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("radius"))
	return rv
}


