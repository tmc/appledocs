// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SegmentedControl] class.
var segmentedControlClass = _SegmentedControlClass{objc.GetClass("NSSegmentedControl")}

type _SegmentedControlClass struct {
	class objc.Class
}

// An interface definition for the [SegmentedControl] class.
type ISegmentedControl interface {
	IControl
}

// Display one or more buttons in a single horizontal group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl

type SegmentedControl struct {
	Control
}

// SegmentedControlFrom constructs a [SegmentedControl] from an unsafe.Pointer.
//
// Display one or more buttons in a single horizontal group.
func SegmentedControlFrom(ptr unsafe.Pointer) SegmentedControl {
	return SegmentedControl{
		Control: ControlFrom(ptr),
	}
}



