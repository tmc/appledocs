// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PictureInPictureController] class.
var (
	PictureInPictureControllerClass     _PictureInPictureControllerClass
	PictureInPictureControllerClassOnce sync.Once
)

func getPictureInPictureControllerClass() _PictureInPictureControllerClass {
	PictureInPictureControllerClassOnce.Do(func() {
		PictureInPictureControllerClass = _PictureInPictureControllerClass{objc.GetClass("AVPictureInPictureController")}
	})
	return PictureInPictureControllerClass
}

type _PictureInPictureControllerClass struct {
	class objc.Class
}

// An interface definition for the [PictureInPictureController] class.
type IPictureInPictureController interface {
	objectivec.IObject
	InvalidatePlaybackState()
	StartPictureInPicture()
	StopPictureInPicture()
}

// A controller that responds to user-initiated Picture in Picture playback of video in a floating, resizable window.
//
// To use Picture in Picture, you need to configure your app to support background audio playback. See for more details. Before presenting a user interface to start Picture in Picture, call the method to determine if the current device supports the feature, and check the property value to determine whether PiP is possible in the current context.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController
type PictureInPictureController struct {
	objectivec.Object
}

// PictureInPictureControllerFrom constructs a [PictureInPictureController] from an unsafe.Pointer.
//
// A controller that responds to user-initiated Picture in Picture playback of video in a floating, resizable window.
func PictureInPictureControllerFrom(ptr unsafe.Pointer) PictureInPictureController {
	return PictureInPictureController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PictureInPictureControllerClass) Alloc() PictureInPictureController {
	rv := objc.Send[PictureInPictureController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PictureInPictureControllerClass) New() PictureInPictureController {
	rv := objc.Send[PictureInPictureController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PictureInPictureController) Init() PictureInPictureController {
	rv := objc.Send[PictureInPictureController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PictureInPictureController) Autorelease() PictureInPictureController {
	rv := objc.Send[PictureInPictureController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPictureInPictureController creates a new PictureInPictureController instance.
func NewPictureInPictureController() PictureInPictureController {
	return getPictureInPictureControllerClass().New()
}




// Creates a Picture in Picture controller with a content source.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/init(contentSource:)
func NewPictureInPictureControllerWithContentSource(contentSource unsafe.Pointer) PictureInPictureController {
	instance := getPictureInPictureControllerClass().Alloc()
	rv := objc.Send[PictureInPictureController](instance.ID, objc.Sel("initWithContentSource:"), contentSource)
	rv.Autorelease()
	return rv
}



// Creates a Picture in Picture controller with a player layer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/init(playerLayer:)
func NewPictureInPictureControllerWithPlayerLayer(playerLayer unsafe.Pointer) PictureInPictureController {
	instance := getPictureInPictureControllerClass().Alloc()
	rv := objc.Send[PictureInPictureController](instance.ID, objc.Sel("initWithPlayerLayer:"), playerLayer)
	rv.Autorelease()
	return rv
}


// Returns a Boolean value that indicates whether the current device supports Picture in Picture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/isPictureInPictureSupported()
func (pc _PictureInPictureControllerClass) IsPictureInPictureSupported() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("isPictureInPictureSupported"))
	return rv
}

// Returns a system-default template image that’s compatible with a trait collection for the button that starts Picture in Picture in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/pictureInPictureButtonStartImage(compatibleWith:)
func (pc _PictureInPictureControllerClass) PictureInPictureButtonStartImageCompatibleWithTraitCollection(traitCollection unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pictureInPictureButtonStartImageCompatibleWithTraitCollection:"), traitCollection)
	return rv
}

// Returns a system-default template image that’s compatible with a trait collection for the button that stops Picture in Picture in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/pictureInPictureButtonStopImage(compatibleWith:)
func (pc _PictureInPictureControllerClass) PictureInPictureButtonStopImageCompatibleWithTraitCollection(traitCollection unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pictureInPictureButtonStopImageCompatibleWithTraitCollection:"), traitCollection)
	return rv
}

// A system-default template image for the button that starts Picture in Picture in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/pictureInPictureButtonStartImage
func (pc _PictureInPictureControllerClass) PictureInPictureButtonStartImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pictureInPictureButtonStartImage"))
	return rv
}
// A system-default template image for the button that stops Picture in Picture in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/pictureInPictureButtonStopImage
func (pc _PictureInPictureControllerClass) PictureInPictureButtonStopImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pictureInPictureButtonStopImage"))
	return rv
}
// Invalidates the controller’s current playback state and fetches the updated state from the sample buffer playback delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/invalidatePlaybackState()
func (p_ PictureInPictureController) InvalidatePlaybackState() {
	objc.Send[objc.ID](p_.ID, objc.Sel("invalidatePlaybackState"))
}

// Starts Picture in Picture, if possible.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/startPictureInPicture()
func (p_ PictureInPictureController) StartPictureInPicture() {
	objc.Send[objc.ID](p_.ID, objc.Sel("startPictureInPicture"))
}

// Stops Picture in Picture, if active.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/stopPictureInPicture()
func (p_ PictureInPictureController) StopPictureInPicture() {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopPictureInPicture"))
}

// A Boolean value that indicates whether the Picture in Picture window is onscreen.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/ispictureinpictureactive
func (p_ PictureInPictureController) IsPictureInPictureActive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPictureInPictureActive"))
	return rv
}


// SetIsPictureInPictureActive sets the value of the isPictureInPictureActive property.
// A Boolean value that indicates whether the Picture in Picture window is onscreen.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/ispictureinpictureactive
func (p_ PictureInPictureController) SetIsPictureInPictureActive(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPictureInPictureActive:"), value)
}

// A Boolean value that indicates whether Picture in Picture playback is currently possible.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/ispictureinpicturepossible
func (p_ PictureInPictureController) IsPictureInPicturePossible() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPictureInPicturePossible"))
	return rv
}


// SetIsPictureInPicturePossible sets the value of the isPictureInPicturePossible property.
// A Boolean value that indicates whether Picture in Picture playback is currently possible.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/ispictureinpicturepossible
func (p_ PictureInPictureController) SetIsPictureInPicturePossible(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPictureInPicturePossible:"), value)
}

// A Boolean value that indicates whether the system suspends the controller’s Picture in Picture window.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/ispictureinpicturesuspended
func (p_ PictureInPictureController) IsPictureInPictureSuspended() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPictureInPictureSuspended"))
	return rv
}


// SetIsPictureInPictureSuspended sets the value of the isPictureInPictureSuspended property.
// A Boolean value that indicates whether the system suspends the controller’s Picture in Picture window.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/ispictureinpicturesuspended
func (p_ PictureInPictureController) SetIsPictureInPictureSuspended(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPictureInPictureSuspended:"), value)
}

// A Boolean value that indicates whether Picture in Picture starts automatically when the controller embeds its content inline and the app transitions to the background.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/canStartPictureInPictureAutomaticallyFromInline
func (p_ PictureInPictureController) CanStartPictureInPictureAutomaticallyFromInline() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canStartPictureInPictureAutomaticallyFromInline"))
	return rv
}


// SetCanStartPictureInPictureAutomaticallyFromInline sets the value of the canStartPictureInPictureAutomaticallyFromInline property.
// A Boolean value that indicates whether Picture in Picture starts automatically when the controller embeds its content inline and the app transitions to the background.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/canStartPictureInPictureAutomaticallyFromInline
func (p_ PictureInPictureController) SetCanStartPictureInPictureAutomaticallyFromInline(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanStartPictureInPictureAutomaticallyFromInline:"), value)
}

// A Boolean value that indicates whether Picture in Picture is active and is able to stop.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/canStopPictureInPicture
func (p_ PictureInPictureController) CanStopPictureInPicture() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canStopPictureInPicture"))
	return rv
}

// The source of the controller’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/contentSource-swift.property
func (p_ PictureInPictureController) ContentSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contentSource"))
	return rv
}


// SetContentSource sets the value of the contentSource property.
// The source of the controller’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/contentSource-swift.property
func (p_ PictureInPictureController) SetContentSource(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentSource:"), value)
}

// A delegate object for a Picture in Picture controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/delegate
func (p_ PictureInPictureController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate object for a Picture in Picture controller.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/delegate
func (p_ PictureInPictureController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the Picture in Picture window is onscreen.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/isPictureInPictureActive
func (p_ PictureInPictureController) PictureInPictureActive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pictureInPictureActive"))
	return rv
}

// A Boolean value that indicates whether Picture in Picture playback is currently possible.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/isPictureInPicturePossible
func (p_ PictureInPictureController) PictureInPicturePossible() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pictureInPicturePossible"))
	return rv
}

// A Boolean value that indicates whether the system suspends the controller’s Picture in Picture window.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/isPictureInPictureSuspended
func (p_ PictureInPictureController) PictureInPictureSuspended() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pictureInPictureSuspended"))
	return rv
}

// A system-default template image for the button that starts Picture in Picture in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/pictureInPictureButtonStartImage
func (p_ PictureInPictureController) PictureInPictureButtonStartImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pictureInPictureButtonStartImage"))
	return rv
}

// A system-default template image for the button that stops Picture in Picture in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/pictureInPictureButtonStopImage
func (p_ PictureInPictureController) PictureInPictureButtonStopImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pictureInPictureButtonStopImage"))
	return rv
}

// The layer that displays the video content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/playerLayer
func (p_ PictureInPictureController) PlayerLayer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("playerLayer"))
	return rv
}

// A Boolean value that determines whether the controller allows the user to skip media content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/requiresLinearPlayback
func (p_ PictureInPictureController) RequiresLinearPlayback() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("requiresLinearPlayback"))
	return rv
}


// SetRequiresLinearPlayback sets the value of the requiresLinearPlayback property.
// A Boolean value that determines whether the controller allows the user to skip media content.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/requiresLinearPlayback
func (p_ PictureInPictureController) SetRequiresLinearPlayback(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRequiresLinearPlayback:"), value)
}


