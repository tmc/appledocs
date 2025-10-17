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



