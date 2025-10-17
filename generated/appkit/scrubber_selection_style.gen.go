// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScrubberSelectionStyle] class.
var scrubberSelectionStyleClass = _ScrubberSelectionStyleClass{objc.GetClass("NSScrubberSelectionStyle")}

type _ScrubberSelectionStyleClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberSelectionStyle] class.
type IScrubberSelectionStyle interface {
	objectivec.IObject
}

// An abstract class that provides decorative accessory views for selected and highlighted items within a scrubber control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle

type ScrubberSelectionStyle struct {
	objectivec.Object
}

// ScrubberSelectionStyleFrom constructs a [ScrubberSelectionStyle] from an unsafe.Pointer.
//
// An abstract class that provides decorative accessory views for selected and highlighted items within a scrubber control.
func ScrubberSelectionStyleFrom(ptr unsafe.Pointer) ScrubberSelectionStyle {
	return ScrubberSelectionStyle{objectivec.Object{objc.ID(ptr)}}
}



