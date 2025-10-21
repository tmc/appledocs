// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [PHLivePhotoView] class.
var (
	PHLivePhotoViewClass     _PHLivePhotoViewClass
	PHLivePhotoViewClassOnce sync.Once
)

func getPHLivePhotoViewClass() _PHLivePhotoViewClass {
	PHLivePhotoViewClassOnce.Do(func() {
		PHLivePhotoViewClass = _PHLivePhotoViewClass{objc.GetClass("PHLivePhotoView")}
	})
	return PHLivePhotoViewClass
}

type _PHLivePhotoViewClass struct {
	class objc.Class
}

// An interface definition for the [PHLivePhotoView] class.
type IPHLivePhotoView interface {
	appkit.IView
	StartPlaybackWithStyle(playbackStyle unsafe.Pointer)
	StopPlayback()
	StopPlaybackAnimated(animated bool)
}

// A view that displays a Live Photo—a picture that also includes motion and sound from the moments just before and after its capture.
//
// Use a Live Photo view to display the photo and control playback of its motion and sound content. In iOS and tvOS, you can obtain Live Photo objects from the Photos library, using the or and classes, or by creating one from asset resources exported from a Photos library. In macOS, Live Photo objects are available only when editing Live Photo content in a photo editing extension that runs in the Photos app—see the class to access Live Photo content in an editing session. By default, a Live Photo view uses its own gesture recognizer to allow the user to play the motion and sound content of a Live Photo with the same interactions and visual effects seen in the Photos app. To customize this gesture recognizer—for example, to install it on a different view for proper event handling in your app’s view hierarchy—use the property. To animate the view briefly to hint that a picture is a Live Photo, use the method with the option.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView
type PHLivePhotoView struct {
	appkit.View
}

// PHLivePhotoViewFrom constructs a [PHLivePhotoView] from an unsafe.Pointer.
//
// A view that displays a Live Photo—a picture that also includes motion and sound from the moments just before and after its capture.
func PHLivePhotoViewFrom(ptr unsafe.Pointer) PHLivePhotoView {
	return PHLivePhotoView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHLivePhotoViewClass) Alloc() PHLivePhotoView {
	rv := objc.Send[PHLivePhotoView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHLivePhotoViewClass) New() PHLivePhotoView {
	rv := objc.Send[PHLivePhotoView](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHLivePhotoView) Init() PHLivePhotoView {
	rv := objc.Send[PHLivePhotoView](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHLivePhotoView) Autorelease() PHLivePhotoView {
	rv := objc.Send[PHLivePhotoView](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHLivePhotoView creates a new PHLivePhotoView instance.
func NewPHLivePhotoView() PHLivePhotoView {
	return getPHLivePhotoViewClass().New()
}


// Returns an icon image for the specified Live Photo semantic options.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/livePhotoBadgeImage(options:)
func (pc _PHLivePhotoViewClass) LivePhotoBadgeImageWithOptions(badgeOptions unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("livePhotoBadgeImageWithOptions:"), badgeOptions)
	return rv
}

// Begins playback of Live Photo content in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/startPlayback(with:)
func (p_ PHLivePhotoView) StartPlaybackWithStyle(playbackStyle unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startPlaybackWithStyle:"), playbackStyle)
}

// Stops playback of a Live Photo.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/stopPlayback()
func (p_ PHLivePhotoView) StopPlayback() {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopPlayback"))
}

// Stops playback of a Live Photo in an animated manner.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/stopPlayback(animated:)
func (p_ PHLivePhotoView) StopPlaybackAnimated(animated bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopPlaybackAnimated:"), animated)
}

// The audio gain to apply to the Live Photo’s movie content during playback.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/audioVolume
func (p_ PHLivePhotoView) AudioVolume() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("audioVolume"))
	return rv
}


// SetAudioVolume sets the value of the audioVolume property.
// The audio gain to apply to the Live Photo’s movie content during playback.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/audioVolume
func (p_ PHLivePhotoView) SetAudioVolume(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudioVolume:"), value)
}

// The mode in which the view displays its content.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/contentMode
func (p_ PHLivePhotoView) ContentMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contentMode"))
	return rv
}


// SetContentMode sets the value of the contentMode property.
// The mode in which the view displays its content.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/contentMode
func (p_ PHLivePhotoView) SetContentMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/contentsRect
func (p_ PHLivePhotoView) ContentsRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("contentsRect"))
	return rv
}


// SetContentsRect sets the value of the contentsRect property.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/contentsRect
func (p_ PHLivePhotoView) SetContentsRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentsRect:"), value)
}

// An object to be notified when Live Photo playback begins or ends.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/delegate
func (p_ PHLivePhotoView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// An object to be notified when Live Photo playback begins or ends.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/delegate
func (p_ PHLivePhotoView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that determines whether the view plays the audio content of its Live Photo.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/isMuted
func (p_ PHLivePhotoView) Muted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("muted"))
	return rv
}


// SetMuted sets the value of the muted property.
// A Boolean value that determines whether the view plays the audio content of its Live Photo.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/isMuted
func (p_ PHLivePhotoView) SetMuted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMuted:"), value)
}

// The Live Photo displayed in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/livePhoto
func (p_ PHLivePhotoView) LivePhoto() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("livePhoto"))
	return rv
}


// SetLivePhoto sets the value of the livePhoto property.
// The Live Photo displayed in the view.

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/livePhoto
func (p_ PHLivePhotoView) SetLivePhoto(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLivePhoto:"), value)
}

// A view for displaying Live Photo status.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/livePhotoBadgeView
func (p_ PHLivePhotoView) LivePhotoBadgeView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("livePhotoBadgeView"))
	return rv
}

// A gesture recognizer that controls playback of the Live Photo in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoView/playbackGestureRecognizer
func (p_ PHLivePhotoView) PlaybackGestureRecognizer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("playbackGestureRecognizer"))
	return rv
}



