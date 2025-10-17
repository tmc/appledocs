// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AlignmentFeedbackFilter] class.
var alignmentFeedbackFilterClass = _AlignmentFeedbackFilterClass{objc.GetClass("NSAlignmentFeedbackFilter")}

type _AlignmentFeedbackFilterClass struct {
	class objc.Class
}

// An object that can filter the movement of an object and provides haptic feedback when alignment occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter

type AlignmentFeedbackFilter struct {
	objectivec.Object
}

// AlignmentFeedbackFilterFrom constructs a [AlignmentFeedbackFilter] from an unsafe.Pointer.
//
// An object that can filter the movement of an object and provides haptic feedback when alignment occurs.
func AlignmentFeedbackFilterFrom(ptr unsafe.Pointer) AlignmentFeedbackFilter {
	return AlignmentFeedbackFilter{objectivec.Object{objc.ID(ptr)}}
}



