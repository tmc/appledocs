// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberImageItemView] class.
var scrubberImageItemViewClass = _ScrubberImageItemViewClass{objc.GetClass("NSScrubberImageItemView")}

type _ScrubberImageItemViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberImageItemView] class.
type IScrubberImageItemView interface {
	IScrubberItemView
}

// A concrete view subclass for displaying images in a scrubber items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView

type ScrubberImageItemView struct {
	ScrubberItemView
}

// ScrubberImageItemViewFrom constructs a [ScrubberImageItemView] from an unsafe.Pointer.
//
// A concrete view subclass for displaying images in a scrubber items.
func ScrubberImageItemViewFrom(ptr unsafe.Pointer) ScrubberImageItemView {
	return ScrubberImageItemView{
		ScrubberItemView: ScrubberItemViewFrom(ptr),
	}
}



