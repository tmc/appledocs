// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var conditionClass _ConditionClass

func init() {
	conditionClass = _ConditionClass{objc.GetClass("NSCondition")}
}

type _ConditionClass struct {
	class objc.Class
}

type Condition struct {
	objc.ID
}

func ConditionFrom(ptr unsafe.Pointer) Condition {
	return Condition{
		ID: objc.ID(ptr),
	}
}




