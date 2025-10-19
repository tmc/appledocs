// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVDepthData] class.
var aVDepthDataClass = _AVDepthDataClass{objc.GetClass("AVDepthData")}

type _AVDepthDataClass struct {
	class objc.Class
}

// A container for per-pixel distance or disparity information captured by compatible camera devices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData

type AVDepthData struct {
	objectivec.Object
}

// AVDepthDataFrom constructs a [AVDepthData] from an unsafe.Pointer.
//
// A container for per-pixel distance or disparity information captured by compatible camera devices.
func AVDepthDataFrom(ptr unsafe.Pointer) AVDepthData {
	return AVDepthData{objectivec.Object{objc.ID(ptr)}}
}



