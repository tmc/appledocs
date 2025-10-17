// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutConstraint] class.
var LayoutConstraintClass objc.Class

func init() {
	LayoutConstraintClass = objc.GetClass("NSLayoutConstraint")
}

type LayoutConstraint struct {
	objc.ID
}

func LayoutConstraintFrom(ptr unsafe.Pointer) LayoutConstraint {
	return LayoutConstraint{
		ID: objc.ID(ptr),
	}
}



