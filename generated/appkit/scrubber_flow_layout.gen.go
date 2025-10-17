// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberFlowLayout] class.
var scrubberFlowLayoutClass = _ScrubberFlowLayoutClass{objc.GetClass("NSScrubberFlowLayout")}

type _ScrubberFlowLayoutClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberFlowLayout] class.
type IScrubberFlowLayout interface {
	IScrubberLayout
}

// A concrete layout object that arranges items end-to-end in a linear strip. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout

type ScrubberFlowLayout struct {
	ScrubberLayout
}

// ScrubberFlowLayoutFrom constructs a [ScrubberFlowLayout] from an unsafe.Pointer.
//
// A concrete layout object that arranges items end-to-end in a linear strip.
func ScrubberFlowLayoutFrom(ptr unsafe.Pointer) ScrubberFlowLayout {
	return ScrubberFlowLayout{
		ScrubberLayout: ScrubberLayoutFrom(ptr),
	}
}



