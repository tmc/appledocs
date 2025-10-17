// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetalDisplayLink] class.
var metalDisplayLinkClass = _MetalDisplayLinkClass{objc.GetClass("CAMetalDisplayLink")}

type _MetalDisplayLinkClass struct {
	class objc.Class
}

// A class your Metal app uses to register for callbacks to synchronize its animations for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink

type MetalDisplayLink struct {
	objectivec.Object
}

// MetalDisplayLinkFrom constructs a [MetalDisplayLink] from an unsafe.Pointer.
//
// A class your Metal app uses to register for callbacks to synchronize its animations for a display.
func MetalDisplayLinkFrom(ptr unsafe.Pointer) MetalDisplayLink {
	return MetalDisplayLink{objectivec.Object{objc.ID(ptr)}}
}



