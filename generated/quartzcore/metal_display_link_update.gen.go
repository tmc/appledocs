// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetalDisplayLinkUpdate] class.
var metalDisplayLinkUpdateClass = _MetalDisplayLinkUpdateClass{objc.GetClass("CAMetalDisplayLinkUpdate")}

type _MetalDisplayLinkUpdateClass struct {
	class objc.Class
}

// Stores information about a single update from a Metal display link instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update

type MetalDisplayLinkUpdate struct {
	objectivec.Object
}

// MetalDisplayLinkUpdateFrom constructs a [MetalDisplayLinkUpdate] from an unsafe.Pointer.
//
// Stores information about a single update from a Metal display link instance.
func MetalDisplayLinkUpdateFrom(ptr unsafe.Pointer) MetalDisplayLinkUpdate {
	return MetalDisplayLinkUpdate{objectivec.Object{objc.ID(ptr)}}
}



