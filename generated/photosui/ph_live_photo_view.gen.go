// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/photos"
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
	// properties:
	AudioVolume() float32
	SetAudioVolume(value float32)
	ContentMode() unsafe.Pointer
	SetContentMode(value unsafe.Pointer)
	ContentsRect() objc.IObject /* cross-framework: Rect */
	SetContentsRect(value objc.IObject /* cross-framework: Rect */)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	IntrinsicContentSize() objc.IObject /* cross-framework: Size */
	SetIntrinsicContentSize(value objc.IObject /* cross-framework: Size */)
	IsMuted() bool
	SetIsMuted(value bool)
	LivePhoto() objc.IObject /* cross-framework: PHLivePhoto */
	SetLivePhoto(value objc.IObject /* cross-framework: PHLivePhoto */)
	LivePhotoBadgeView() objc.IObject /* cross-framework: View */
	SetLivePhotoBadgeView(value objc.IObject /* cross-framework: View */)
	PlaybackGestureRecognizer() objc.IObject /* cross-framework: GestureRecognizer */
	SetPlaybackGestureRecognizer(value objc.IObject /* cross-framework: GestureRecognizer */)
	// methods:
}

// A view that displays a Live Photo—a picture that also includes motion and sound from the moments just before and after its capture.
//
// Use a Live Photo view to display the photo and control playback of its motion and sound content. In iOS and tvOS, you can obtain Live Photo objects from the Photos library, using the or and classes, or by creating one from asset resources exported from a Photos library. In macOS, Live Photo objects are available only when editing Live Photo content in a photo editing extension that runs in the Photos app—see the class to access Live Photo content in an editing session. By default, a Live Photo view uses its own gesture recognizer to allow the user to play the motion and sound content of a Live Photo with the same interactions and visual effects seen in the Photos app. To customize this gesture recognizer—for example, to install it on a different view for proper event handling in your app’s view hierarchy—use the property. To animate the view briefly to hint that a picture is a Live Photo, use the method with the option.


// A view that displays a Live Photo—a picture that also includes motion and sound from the moments just before and after its capture.
//
// [Full Topic]
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



// The audio gain to apply to the Live Photo’s movie content during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/audiovolume
func (p_ PHLivePhotoView) AudioVolume() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("audioVolume"))
	return rv
}


// The audio gain to apply to the Live Photo’s movie content during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/audiovolume
func (p_ PHLivePhotoView) SetAudioVolume(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudioVolume:"), value)
}


// The mode in which the view displays its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/contentmode
func (p_ PHLivePhotoView) ContentMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contentMode"))
	return rv
}


// The mode in which the view displays its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/contentmode
func (p_ PHLivePhotoView) SetContentMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/contentsrect
func (p_ PHLivePhotoView) ContentsRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](p_.ID, objc.Sel("contentsRect"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/contentsrect
func (p_ PHLivePhotoView) SetContentsRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentsRect:"), value)
}


// An object to be notified when Live Photo playback begins or ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/delegate
func (p_ PHLivePhotoView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}


// An object to be notified when Live Photo playback begins or ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/delegate
func (p_ PHLivePhotoView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/intrinsiccontentsize
func (p_ PHLivePhotoView) IntrinsicContentSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](p_.ID, objc.Sel("intrinsicContentSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/intrinsiccontentsize
func (p_ PHLivePhotoView) SetIntrinsicContentSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIntrinsicContentSize:"), value)
}


// A Boolean value that determines whether the view plays the audio content of its Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/ismuted
func (p_ PHLivePhotoView) IsMuted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isMuted"))
	return rv
}


// A Boolean value that determines whether the view plays the audio content of its Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/ismuted
func (p_ PHLivePhotoView) SetIsMuted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsMuted:"), value)
}


// The Live Photo displayed in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/livephoto
func (p_ PHLivePhotoView) LivePhoto() objc.IObject /* cross-framework: PHLivePhoto */ {
	rv := objc.Send[photos.PHLivePhoto](p_.ID, objc.Sel("livePhoto"))
	return rv
}


// The Live Photo displayed in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/livephoto
func (p_ PHLivePhotoView) SetLivePhoto(value objc.IObject /* cross-framework: PHLivePhoto */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLivePhoto:"), value)
}


// A view for displaying Live Photo status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/livephotobadgeview
func (p_ PHLivePhotoView) LivePhotoBadgeView() objc.IObject /* cross-framework: View */ {
	rv := objc.Send[appkit.View](p_.ID, objc.Sel("livePhotoBadgeView"))
	return rv
}


// A view for displaying Live Photo status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/livephotobadgeview
func (p_ PHLivePhotoView) SetLivePhotoBadgeView(value objc.IObject /* cross-framework: View */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLivePhotoBadgeView:"), value)
}


// A gesture recognizer that controls playback of the Live Photo in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/playbackgesturerecognizer
func (p_ PHLivePhotoView) PlaybackGestureRecognizer() objc.IObject /* cross-framework: GestureRecognizer */ {
	rv := objc.Send[appkit.GestureRecognizer](p_.ID, objc.Sel("playbackGestureRecognizer"))
	return rv
}


// A gesture recognizer that controls playback of the Live Photo in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phlivephotoview/playbackgesturerecognizer
func (p_ PHLivePhotoView) SetPlaybackGestureRecognizer(value objc.IObject /* cross-framework: GestureRecognizer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackGestureRecognizer:"), value)
}



