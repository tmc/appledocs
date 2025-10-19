// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVMutableComposition] class.
var aVMutableCompositionClass = _AVMutableCompositionClass{objc.GetClass("AVMutableComposition")}

type _AVMutableCompositionClass struct {
	class objc.Class
}

// An object that you use to create a new composition from existing assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition

type AVMutableComposition struct {
	AVComposition
}

// AVMutableCompositionFrom constructs a [AVMutableComposition] from an unsafe.Pointer.
//
// An object that you use to create a new composition from existing assets.
func AVMutableCompositionFrom(ptr unsafe.Pointer) AVMutableComposition {
	return AVMutableComposition{
		AVComposition: AVCompositionFrom(ptr),
	}
}



