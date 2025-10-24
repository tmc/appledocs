//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for PictureInPictureController


// iOS-only properties

// A Boolean value that indicates whether Picture in Picture starts automatically when the controller embeds its content inline and the app transitions to the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/canStartPictureInPictureAutomaticallyFromInline
func (p_ PictureInPictureController) CanStartPictureInPictureAutomaticallyFromInline() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canStartPictureInPictureAutomaticallyFromInline"))
	return rv
}
func (p_ PictureInPictureController) SetCanStartPictureInPictureAutomaticallyFromInline(value bool) {
	p_.ID.Send(objc.RegisterName("setCanStartPictureInPictureAutomaticallyFromInline:"), value)
}

// A Boolean value that indicates whether Picture in Picture is active and is able to stop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/canStopPictureInPicture
func (p_ PictureInPictureController) CanStopPictureInPicture() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canStopPictureInPicture"))
	return rv
}




