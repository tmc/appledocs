// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/avfoundation"
)

/* debug [class.gen.go]: Generating class AVCaptureView */


/* debug [class_header]: Header for AVCaptureView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureView */
// An interface definition for the [CaptureView] class.
type ICaptureView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for CaptureView */
	// properties:
	ControlsStyle() CaptureViewControlsStyle
	SetControlsStyle(value CaptureViewControlsStyle)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	FileOutput() avfoundation.CaptureFileOutput
	Session() avfoundation.CaptureSession
	VideoGravity() LayerVideoGravity /* not a class type */
	SetVideoGravity(value LayerVideoGravity /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureView */
	// methods:
	SetSessionShowVideoPreviewShowAudioPreview(session avfoundation.CaptureSession, showVideoPreview bool, showAudioPreview bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureView */
// Alloc allocates a new instance without initialization.
func (cc _CaptureViewClass) Alloc() CaptureView {
	rv := objc.Send[CaptureView](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureView */
// A view that displays standard user interface controls for capturing media data.


// A view that displays standard user interface controls for capturing media data.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureView */

// Sets the view’s capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/setSession(_:showVideoPreview:showAudioPreview:)
func (c_ CaptureView) SetSessionShowVideoPreviewShowAudioPreview(session avfoundation.CaptureSession, showVideoPreview bool, showAudioPreview bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSession:showVideoPreview:showAudioPreview:"), session, showVideoPreview, showAudioPreview)
}/* debug [instance_methods/method]: SetSessionShowVideoPreviewShowAudioPreview */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureView */

// The style of the capture controls presented by the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/controlsStyle
func (c_ CaptureView) ControlsStyle() CaptureViewControlsStyle {
	rv := objc.Send[CaptureViewControlsStyle](c_.ID, objc.Sel("controlsStyle"))
	return rv
}/* debug [instance_properties/getter]: controlsStyle */


// The style of the capture controls presented by the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/controlsStyle
func (c_ CaptureView) SetControlsStyle(value CaptureViewControlsStyle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlsStyle:"), value)
}/* debug [instance_properties/setter]: controlsStyle */


// The capture view’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/delegate
func (c_ CaptureView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The capture view’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/delegate
func (c_ CaptureView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The capture file output used to record media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/fileOutput
func (c_ CaptureView) FileOutput() avfoundation.CaptureFileOutput {
	rv := objc.Send[avfoundation.CaptureFileOutput](c_.ID, objc.Sel("fileOutput"))
	return rv
}/* debug [instance_properties/getter]: fileOutput */


// The view’s associated capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/session
func (c_ CaptureView) Session() avfoundation.CaptureSession {
	rv := objc.Send[avfoundation.CaptureSession](c_.ID, objc.Sel("session"))
	return rv
}/* debug [instance_properties/getter]: session */


// A string value that defines how the capture view displays video within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/videoGravity
func (c_ CaptureView) VideoGravity() LayerVideoGravity /* not a class type */ {
	rv := objc.Send[LayerVideoGravity](c_.ID, objc.Sel("videoGravity"))
	return rv
}/* debug [instance_properties/getter]: videoGravity */


// A string value that defines how the capture view displays video within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureView/videoGravity
func (c_ CaptureView) SetVideoGravity(value LayerVideoGravity /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoGravity:"), value)
}/* debug [instance_properties/setter]: videoGravity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureView */



