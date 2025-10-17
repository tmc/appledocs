// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LayoutConstraint] class.
var layoutConstraintClass = _LayoutConstraintClass{objc.GetClass("NSLayoutConstraint")}

type _LayoutConstraintClass struct {
	class objc.Class
}

// An interface definition for the [LayoutConstraint] class.
type ILayoutConstraint interface {
	objectivec.IObject
}

// The relationship between two user interface objects that must be satisfied by the constraint-based layout system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint

type LayoutConstraint struct {
	objectivec.Object
}

// LayoutConstraintFrom constructs a [LayoutConstraint] from an unsafe.Pointer.
//
// The relationship between two user interface objects that must be satisfied by the constraint-based layout system.
func LayoutConstraintFrom(ptr unsafe.Pointer) LayoutConstraint {
	return LayoutConstraint{objectivec.Object{objc.ID(ptr)}}
}



