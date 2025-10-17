// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberItemView] class.
var scrubberItemViewClass = _ScrubberItemViewClass{objc.GetClass("NSScrubberItemView")}

type _ScrubberItemViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberItemView] class.
type IScrubberItemView interface {
	IScrubberArrangedView
}

// An item at a specific index position in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberItemView

type ScrubberItemView struct {
	ScrubberArrangedView
}

// ScrubberItemViewFrom constructs a [ScrubberItemView] from an unsafe.Pointer.
//
// An item at a specific index position in the scrubber.
func ScrubberItemViewFrom(ptr unsafe.Pointer) ScrubberItemView {
	return ScrubberItemView{
		ScrubberArrangedView: ScrubberArrangedViewFrom(ptr),
	}
}



