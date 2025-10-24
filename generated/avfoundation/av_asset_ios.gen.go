//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Asset


// iOS-only properties

// A Boolean value that indicates whether you can write the asset to the Saved Photos album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/isCompatibleWithSavedPhotosAlbum
func (a_ Asset) CompatibleWithSavedPhotosAlbum() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("compatibleWithSavedPhotosAlbum"))
	return rv
}

// The encoded or authored size of the visual portion of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/naturalSize
func (a_ Asset) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a_.ID, objc.Sel("naturalSize"))
	return rv
}

// The asset’s display mode preference for optimal playback of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/preferredDisplayCriteria
func (a_ Asset) PreferredDisplayCriteria() IAVDisplayCriteria {
	rv := objc.Send[DisplayCriteria](a_.ID, objc.Sel("preferredDisplayCriteria"))
	return rv
}




