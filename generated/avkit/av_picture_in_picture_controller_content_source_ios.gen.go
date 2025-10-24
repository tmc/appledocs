//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for PictureInPictureControllerContentSource


// iOS-only properties

// The view controller that presents the video call content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class/activeVideoCallContentViewController
func (p_ PictureInPictureControllerContentSource) ActiveVideoCallContentViewController() IAVPictureInPictureVideoCallViewController {
	rv := objc.Send[PictureInPictureVideoCallViewController](p_.ID, objc.Sel("activeVideoCallContentViewController"))
	return rv
}

// The view that contains the video content of the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class/activeVideoCallSourceView
func (p_ PictureInPictureControllerContentSource) ActiveVideoCallSourceView() appkit.View {
	rv := objc.Send[appkit.View](p_.ID, objc.Sel("activeVideoCallSourceView"))
	return rv
}




