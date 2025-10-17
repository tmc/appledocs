// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AlignmentFeedbackFilter] class.
var AlignmentFeedbackFilterClass objc.Class

func init() {
	AlignmentFeedbackFilterClass = objc.GetClass("NSAlignmentFeedbackFilter")
}

type AlignmentFeedbackFilter struct {
	objc.ID
}

func AlignmentFeedbackFilterFrom(ptr unsafe.Pointer) AlignmentFeedbackFilter {
	return AlignmentFeedbackFilter{
		ID: objc.ID(ptr),
	}
}



