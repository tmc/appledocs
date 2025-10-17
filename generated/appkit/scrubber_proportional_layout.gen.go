// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberProportionalLayout] class.
var scrubberProportionalLayoutClass = _ScrubberProportionalLayoutClass{objc.GetClass("NSScrubberProportionalLayout")}

type _ScrubberProportionalLayoutClass struct {
	class objc.Class
}

// A concrete layout object that sizes each item to some fraction of the scrubber’s visible size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout

type ScrubberProportionalLayout struct {
	ScrubberLayout
}

// ScrubberProportionalLayoutFrom constructs a [ScrubberProportionalLayout] from an unsafe.Pointer.
//
// A concrete layout object that sizes each item to some fraction of the scrubber’s visible size.
func ScrubberProportionalLayoutFrom(ptr unsafe.Pointer) ScrubberProportionalLayout {
	return ScrubberProportionalLayout{
		ScrubberLayout: ScrubberLayoutFrom(ptr),
	}
}



