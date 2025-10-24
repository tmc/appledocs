// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureScreenInput */


/* debug [class_header]: Header for AVCaptureScreenInput */
// The class instance for the [CaptureScreenInput] class.
var (
	CaptureScreenInputClass     _CaptureScreenInputClass
	CaptureScreenInputClassOnce sync.Once
)

func getCaptureScreenInputClass() _CaptureScreenInputClass {
	CaptureScreenInputClassOnce.Do(func() {
		CaptureScreenInputClass = _CaptureScreenInputClass{objc.GetClass("AVCaptureScreenInput")}
	})
	return CaptureScreenInputClass
}

type _CaptureScreenInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureScreenInput */
// An interface definition for the [CaptureScreenInput] class.
type ICaptureScreenInput interface {
	ICaptureInput
	
/* debug [class_interface_properties]: Properties for CaptureScreenInput */
	// properties:
	CapturesCursor() bool
	SetCapturesCursor(value bool)
	CapturesMouseClicks() bool
	SetCapturesMouseClicks(value bool)
	CropRect() corefoundation.CGRect
	SetCropRect(value corefoundation.CGRect)
	MinFrameDuration() objc.IObject /* cross-framework: Time */
	SetMinFrameDuration(value objc.IObject /* cross-framework: Time */)
	RemovesDuplicateFrames() bool
	SetRemovesDuplicateFrames(value bool)
	ScaleFactor() float64
	SetScaleFactor(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureScreenInput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureScreenInput */
// Alloc allocates a new instance without initialization.
func (cc _CaptureScreenInputClass) Alloc() CaptureScreenInput {
	rv := objc.Send[CaptureScreenInput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureScreenInputClass) New() CaptureScreenInput {
	rv := objc.Send[CaptureScreenInput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureScreenInput) Init() CaptureScreenInput {
	rv := objc.Send[CaptureScreenInput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureScreenInput) Autorelease() CaptureScreenInput {
	rv := objc.Send[CaptureScreenInput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureScreenInput creates a new CaptureScreenInput instance.
func NewCaptureScreenInput() CaptureScreenInput {
	return getCaptureScreenInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureScreenInput */
// A capture input for recording from a screen in macOS.
//
// This class is a concrete capture input subclass that provides an interface to capture media from a screen or a portion of a screen. Use instances of this class as input sources for objects that provide media data from one of the screens connected to the system, represented by .


// A capture input for recording from a screen in macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput
type CaptureScreenInput struct {
	CaptureInput
}

// CaptureScreenInputFrom constructs a [CaptureScreenInput] from an unsafe.Pointer.
//
// A capture input for recording from a screen in macOS.
func CaptureScreenInputFrom(ptr unsafe.Pointer) CaptureScreenInput {
	return CaptureScreenInput{
		CaptureInput: CaptureInputFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureScreenInput */

// Initializes a capture screen input that provides media data from the specified display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/init(displayID:)
func NewCaptureScreenInputWithDisplayID(displayID DirectDisplayID /* not a class type */) CaptureScreenInput {
	instance := getCaptureScreenInputClass().Alloc()
	rv := objc.Send[CaptureScreenInput](instance.ID, objc.Sel("initWithDisplayID:"), displayID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureScreenInputWithDisplayID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureScreenInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureScreenInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureScreenInput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureScreenInput */

// A Boolean value that specifies whether the mouse cursor appears in the captured output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/capturesCursor
func (c_ CaptureScreenInput) CapturesCursor() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("capturesCursor"))
	return rv
}/* debug [instance_properties/getter]: capturesCursor */


// A Boolean value that specifies whether the mouse cursor appears in the captured output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/capturesCursor
func (c_ CaptureScreenInput) SetCapturesCursor(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCapturesCursor:"), value)
}/* debug [instance_properties/setter]: capturesCursor */


// A Boolean value that specifies whether mouse clicks appear highlighted in the captured output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/capturesMouseClicks
func (c_ CaptureScreenInput) CapturesMouseClicks() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("capturesMouseClicks"))
	return rv
}/* debug [instance_properties/getter]: capturesMouseClicks */


// A Boolean value that specifies whether mouse clicks appear highlighted in the captured output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/capturesMouseClicks
func (c_ CaptureScreenInput) SetCapturesMouseClicks(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCapturesMouseClicks:"), value)
}/* debug [instance_properties/setter]: capturesMouseClicks */


// Indicates the bounding rectangle of the screen area to be captured, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/cropRect
func (c_ CaptureScreenInput) CropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("cropRect"))
	return rv
}/* debug [instance_properties/getter]: cropRect */


// Indicates the bounding rectangle of the screen area to be captured, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/cropRect
func (c_ CaptureScreenInput) SetCropRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCropRect:"), value)
}/* debug [instance_properties/setter]: cropRect */


// The screen input’s minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/minFrameDuration
func (c_ CaptureScreenInput) MinFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("minFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: minFrameDuration */


// The screen input’s minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/minFrameDuration
func (c_ CaptureScreenInput) SetMinFrameDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinFrameDuration:"), value)
}/* debug [instance_properties/setter]: minFrameDuration */


// A Boolean value that specifies whether the capture input skips duplicate frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/removesDuplicateFrames
func (c_ CaptureScreenInput) RemovesDuplicateFrames() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("removesDuplicateFrames"))
	return rv
}/* debug [instance_properties/getter]: removesDuplicateFrames */


// A Boolean value that specifies whether the capture input skips duplicate frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/removesDuplicateFrames
func (c_ CaptureScreenInput) SetRemovesDuplicateFrames(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRemovesDuplicateFrames:"), value)
}/* debug [instance_properties/setter]: removesDuplicateFrames */


// Indicates the factor by which video buffers captured from the screen are to be scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/scaleFactor
func (c_ CaptureScreenInput) ScaleFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("scaleFactor"))
	return rv
}/* debug [instance_properties/getter]: scaleFactor */


// Indicates the factor by which video buffers captured from the screen are to be scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/scaleFactor
func (c_ CaptureScreenInput) SetScaleFactor(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactor:"), value)
}/* debug [instance_properties/setter]: scaleFactor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureScreenInput */


