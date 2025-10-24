//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for Composition


// iOS-only properties

// A Boolean value that indicates whether you can write the composition to the Saved Photos album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/isCompatibleWithSavedPhotosAlbum
func (c_ Composition) IsCompatibleWithSavedPhotosAlbum() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isCompatibleWithSavedPhotosAlbum"))
	return rv
}
func (c_ Composition) SetIsCompatibleWithSavedPhotosAlbum(value objectivec.IObject) {
	c_.ID.Send(objc.RegisterName("setIsCompatibleWithSavedPhotosAlbum:"), value)
}

// The asset’s display mode preference for optimal playback of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredDisplayCriteria
func (c_ Composition) PreferredDisplayCriteria() IAVDisplayCriteria {
	rv := objc.Send[DisplayCriteria](c_.ID, objc.Sel("preferredDisplayCriteria"))
	return rv
}
func (c_ Composition) SetPreferredDisplayCriteria(value IAVDisplayCriteria) {
	c_.ID.Send(objc.RegisterName("setPreferredDisplayCriteria:"), value)
}





