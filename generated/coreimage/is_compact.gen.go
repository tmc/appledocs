// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [isCompact] class.
var isCompactClass = _isCompactClass{objc.GetClass("isCompact")}

type _isCompactClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/isCompact-c.ivar

type isCompact struct {
	objectivec.Object
}

// isCompactFrom constructs a [isCompact] from an unsafe.Pointer.
func isCompactFrom(ptr unsafe.Pointer) isCompact {
	return isCompact{objectivec.Object{objc.ID(ptr)}}
}



