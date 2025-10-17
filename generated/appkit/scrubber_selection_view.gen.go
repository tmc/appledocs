// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberSelectionView] class.
var scrubberSelectionViewClass = _ScrubberSelectionViewClass{objc.GetClass("NSScrubberSelectionView")}

type _ScrubberSelectionViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberSelectionView] class.
type IScrubberSelectionView interface {
	IScrubberArrangedView
}

// An abstract base class for specifying the appearance of a highlighted or selected item in a scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionView

type ScrubberSelectionView struct {
	ScrubberArrangedView
}

// ScrubberSelectionViewFrom constructs a [ScrubberSelectionView] from an unsafe.Pointer.
//
// An abstract base class for specifying the appearance of a highlighted or selected item in a scrubber.
func ScrubberSelectionViewFrom(ptr unsafe.Pointer) ScrubberSelectionView {
	return ScrubberSelectionView{
		ScrubberArrangedView: ScrubberArrangedViewFrom(ptr),
	}
}



