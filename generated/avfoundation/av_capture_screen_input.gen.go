// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CaptureScreenInput] class.
type ICaptureScreenInput interface {
	ICaptureInput
	

	// properties:
	CapturesCursor() bool
	SetCapturesCursor(value bool)
	CapturesMouseClicks() bool
	SetCapturesMouseClicks(value bool)
	CropRect() corefoundation.CGRect
	SetCropRect(value corefoundation.CGRect)
	MinFrameDuration() objectivec.IObject
	SetMinFrameDuration(value objectivec.IObject)
	RemovesDuplicateFrames() bool
	SetRemovesDuplicateFrames(value bool)
	ScaleFactor() float64
	SetScaleFactor(value float64)


	

	// methods:


}





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






// Initializes a capture screen input that provides media data from the specified display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/init(displayID:)
func NewCaptureScreenInputWithDisplayID(displayID DirectDisplayID /* not a class type */) CaptureScreenInput {
	instance := getCaptureScreenInputClass().Alloc()
	rv := objc.Send[CaptureScreenInput](instance.ID, objc.Sel("initWithDisplayID:"), displayID)
	rv.Autorelease()
	return rv
}






















// A Boolean value that specifies whether the mouse cursor appears in the captured output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/capturesCursor
func (c_ CaptureScreenInput) CapturesCursor() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("capturesCursor"))
	return rv
}


// A Boolean value that specifies whether the mouse cursor appears in the captured output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/capturesCursor
func (c_ CaptureScreenInput) SetCapturesCursor(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCapturesCursor:"), value)
}


// A Boolean value that specifies whether mouse clicks appear highlighted in the captured output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/capturesMouseClicks
func (c_ CaptureScreenInput) CapturesMouseClicks() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("capturesMouseClicks"))
	return rv
}


// A Boolean value that specifies whether mouse clicks appear highlighted in the captured output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/capturesMouseClicks
func (c_ CaptureScreenInput) SetCapturesMouseClicks(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCapturesMouseClicks:"), value)
}


// Indicates the bounding rectangle of the screen area to be captured, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/cropRect
func (c_ CaptureScreenInput) CropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("cropRect"))
	return rv
}


// Indicates the bounding rectangle of the screen area to be captured, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/cropRect
func (c_ CaptureScreenInput) SetCropRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCropRect:"), value)
}


// The screen input’s minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/minFrameDuration
func (c_ CaptureScreenInput) MinFrameDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("minFrameDuration"))
	return rv
}


// The screen input’s minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/minFrameDuration
func (c_ CaptureScreenInput) SetMinFrameDuration(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinFrameDuration:"), value)
}


// A Boolean value that specifies whether the capture input skips duplicate frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/removesDuplicateFrames
func (c_ CaptureScreenInput) RemovesDuplicateFrames() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("removesDuplicateFrames"))
	return rv
}


// A Boolean value that specifies whether the capture input skips duplicate frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/removesDuplicateFrames
func (c_ CaptureScreenInput) SetRemovesDuplicateFrames(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRemovesDuplicateFrames:"), value)
}


// Indicates the factor by which video buffers captured from the screen are to be scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/scaleFactor
func (c_ CaptureScreenInput) ScaleFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("scaleFactor"))
	return rv
}


// Indicates the factor by which video buffers captured from the screen are to be scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/scaleFactor
func (c_ CaptureScreenInput) SetScaleFactor(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactor:"), value)
}







