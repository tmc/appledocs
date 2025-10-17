// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScrubberLayoutAttributes] class.
var scrubberLayoutAttributesClass = _ScrubberLayoutAttributesClass{objc.GetClass("NSScrubberLayoutAttributes")}

type _ScrubberLayoutAttributesClass struct {
	class objc.Class
}

// The layout of a scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes

type ScrubberLayoutAttributes struct {
	objectivec.Object
}

// ScrubberLayoutAttributesFrom constructs a [ScrubberLayoutAttributes] from an unsafe.Pointer.
//
// The layout of a scrubber item.
func ScrubberLayoutAttributesFrom(ptr unsafe.Pointer) ScrubberLayoutAttributes {
	return ScrubberLayoutAttributes{objectivec.Object{objc.ID(ptr)}}
}



