// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RenderInfo] class.
var renderInfoClass = _RenderInfoClass{objc.GetClass("CIRenderInfo")}

type _RenderInfoClass struct {
	class objc.Class
}

// An encapsulation of a render task’s timing, passes, and pixels processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderInfo

type RenderInfo struct {
	objectivec.Object
}

// RenderInfoFrom constructs a [RenderInfo] from an unsafe.Pointer.
//
// An encapsulation of a render task’s timing, passes, and pixels processed.
func RenderInfoFrom(ptr unsafe.Pointer) RenderInfo {
	return RenderInfo{objectivec.Object{objc.ID(ptr)}}
}



