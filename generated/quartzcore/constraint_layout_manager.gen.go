// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ConstraintLayoutManager] class.
var constraintLayoutManagerClass = _ConstraintLayoutManagerClass{objc.GetClass("CAConstraintLayoutManager")}

type _ConstraintLayoutManagerClass struct {
	class objc.Class
}

// An object that provides a constraint-based layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintLayoutManager

type ConstraintLayoutManager struct {
	objectivec.Object
}

// ConstraintLayoutManagerFrom constructs a [ConstraintLayoutManager] from an unsafe.Pointer.
//
// An object that provides a constraint-based layout manager.
func ConstraintLayoutManagerFrom(ptr unsafe.Pointer) ConstraintLayoutManager {
	return ConstraintLayoutManager{objectivec.Object{objc.ID(ptr)}}
}

// Returns the shared layout manager object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintLayoutManager/layoutManager
func (cc _ConstraintLayoutManagerClass) LayoutManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layoutManager"))
	return rv
}


