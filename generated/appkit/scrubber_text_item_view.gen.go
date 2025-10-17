// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberTextItemView] class.
var scrubberTextItemViewClass = _ScrubberTextItemViewClass{objc.GetClass("NSScrubberTextItemView")}

type _ScrubberTextItemViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberTextItemView] class.
type IScrubberTextItemView interface {
	IScrubberItemView
}

// A concrete view subclass for displaying text for an item in a scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberTextItemView

type ScrubberTextItemView struct {
	ScrubberItemView
}

// ScrubberTextItemViewFrom constructs a [ScrubberTextItemView] from an unsafe.Pointer.
//
// A concrete view subclass for displaying text for an item in a scrubber.
func ScrubberTextItemViewFrom(ptr unsafe.Pointer) ScrubberTextItemView {
	return ScrubberTextItemView{
		ScrubberItemView: ScrubberItemViewFrom(ptr),
	}
}



