// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Condition] class.
var conditionClass = _ConditionClass{objc.GetClass("CLCondition")}

type _ConditionClass struct {
	class objc.Class
}

// The abstract base class that all other conditions derive from. [Full Topic]
//
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



