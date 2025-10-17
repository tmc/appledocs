// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ConditionLock] class.
var conditionLockClass = _ConditionLockClass{objc.GetClass("NSConditionLock")}

type _ConditionLockClass struct {
	class objc.Class
}

// An interface definition for the [ConditionLock] class.
type IConditionLock interface {
	objectivec.IObject
}

// A lock that can be associated with specific, user-defined conditions. [Full Topic]
//
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



