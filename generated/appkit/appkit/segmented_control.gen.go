// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SegmentedControl] class.
var SegmentedControlClass objc.Class

func init() {
	SegmentedControlClass = objc.GetClass("NSSegmentedControl")
}

type SegmentedControl struct {
	objc.ID
}

func SegmentedControlFrom(ptr unsafe.Pointer) SegmentedControl {
	return SegmentedControl{
		ID: objc.ID(ptr),
	}
}




