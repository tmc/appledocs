//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MutableMovie


// iOS-only properties

// A Boolean value that indicates whether you can write the composition to the Saved Photos album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isCompatibleWithSavedPhotosAlbum
func (m_ MutableMovie) IsCompatibleWithSavedPhotosAlbum() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("isCompatibleWithSavedPhotosAlbum"))
	return rv
}
func (m_ MutableMovie) SetIsCompatibleWithSavedPhotosAlbum(value objectivec.IObject) {
	m_.ID.Send(objc.RegisterName("setIsCompatibleWithSavedPhotosAlbum:"), value)
}




