// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextLayer] class.
var textLayerClass = _TextLayerClass{objc.GetClass("CATextLayer")}

type _TextLayerClass struct {
	class objc.Class
}

// A layer that provides simple text layout and rendering of plain or attributed strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer

type TextLayer struct {
	Layer
}

// TextLayerFrom constructs a [TextLayer] from an unsafe.Pointer.
//
// A layer that provides simple text layout and rendering of plain or attributed strings.
func TextLayerFrom(ptr unsafe.Pointer) TextLayer {
	return TextLayer{
		Layer: LayerFrom(ptr),
	}
}



