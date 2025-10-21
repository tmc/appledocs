// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/avfoundation"
)

// The class instance for the [CaptureView] class.
var (
	CaptureViewClass     _CaptureViewClass
	CaptureViewClassOnce sync.Once
)

func getCaptureViewClass() _CaptureViewClass {
	CaptureViewClassOnce.Do(func() {
		CaptureViewClass = _CaptureViewClass{objc.GetClass("AVCaptureView")}
	})
	return CaptureViewClass
}

type _CaptureViewClass struct {
	class objc.Class
}

// An interface definition for the [CaptureView] class.
type ICaptureView interface {
	appkit.IView
	SetSessionShowVideoPreviewShowAudioPreview(session avfoundation.ICaptureSession, showVideoPreview bool, showAudioPreview bool)
}

// A view that displays standard user interface controls for capturing media data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView
type CaptureView struct {
	appkit.View
}

// CaptureViewFrom constructs a [CaptureView] from an unsafe.Pointer.
//
// A view that displays standard user interface controls for capturing media data.
func CaptureViewFrom(ptr unsafe.Pointer) CaptureView {
	return CaptureView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureViewClass) Alloc() CaptureView {
	rv := objc.Send[CaptureView](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureViewClass) New() CaptureView {
	rv := objc.Send[CaptureView](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureView) Init() CaptureView {
	rv := objc.Send[CaptureView](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureView) Autorelease() CaptureView {
	rv := objc.Send[CaptureView](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureView creates a new CaptureView instance.
func NewCaptureView() CaptureView {
	return getCaptureViewClass().New()
}


// Sets the view’s capture session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/setSession(_:showVideoPreview:showAudioPreview:)
func (c_ CaptureView) SetSessionShowVideoPreviewShowAudioPreview(session avfoundation.ICaptureSession, showVideoPreview bool, showAudioPreview bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSession:showVideoPreview:showAudioPreview:"), session, showVideoPreview, showAudioPreview)
}

// The style of the capture controls presented by the view.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/controlsStyle
func (c_ CaptureView) ControlsStyle() CaptureViewControlsStyle {
	rv := objc.Send[CaptureViewControlsStyle](c_.ID, objc.Sel("controlsStyle"))
	return rv
}


// SetControlsStyle sets the value of the controlsStyle property.
// The style of the capture controls presented by the view.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/controlsStyle
func (c_ CaptureView) SetControlsStyle(value CaptureViewControlsStyle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlsStyle:"), value)
}

// The capture view’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/delegate
func (c_ CaptureView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The capture view’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/delegate
func (c_ CaptureView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}

// The capture file output used to record media data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/fileOutput
func (c_ CaptureView) FileOutput() avfoundation.CaptureFileOutput {
	rv := objc.Send[avfoundation.CaptureFileOutput](c_.ID, objc.Sel("fileOutput"))
	return rv
}

// The view’s associated capture session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/session
func (c_ CaptureView) Session() avfoundation.CaptureSession {
	rv := objc.Send[avfoundation.CaptureSession](c_.ID, objc.Sel("session"))
	return rv
}

// A string value that defines how the capture view displays video within its bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/videoGravity
func (c_ CaptureView) VideoGravity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoGravity"))
	return rv
}


// SetVideoGravity sets the value of the videoGravity property.
// A string value that defines how the capture view displays video within its bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/videoGravity
func (c_ CaptureView) SetVideoGravity(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoGravity:"), value)
}



