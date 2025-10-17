// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Condition] class.
var ConditionClass objc.Class

func init() {
	ConditionClass = objc.GetClass("NSCondition")
}

type Condition struct {
	objc.ID
}

func ConditionFrom(ptr unsafe.Pointer) Condition {
	return Condition{
		ID: objc.ID(ptr),
	}
}




