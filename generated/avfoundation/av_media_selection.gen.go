// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVMediaSelection] class.
var aVMediaSelectionClass = _AVMediaSelectionClass{objc.GetClass("AVMediaSelection")}

type _AVMediaSelectionClass struct {
	class objc.Class
}

// An object that represents a complete rendition of media selection options on an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection

type AVMediaSelection struct {
	objectivec.Object
}

// AVMediaSelectionFrom constructs a [AVMediaSelection] from an unsafe.Pointer.
//
// An object that represents a complete rendition of media selection options on an asset.
func AVMediaSelectionFrom(ptr unsafe.Pointer) AVMediaSelection {
	return AVMediaSelection{objectivec.Object{objc.ID(ptr)}}
}



