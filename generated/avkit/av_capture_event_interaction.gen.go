// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureEventInteraction] class.
var (
	CaptureEventInteractionClass     _CaptureEventInteractionClass
	CaptureEventInteractionClassOnce sync.Once
)

func getCaptureEventInteractionClass() _CaptureEventInteractionClass {
	CaptureEventInteractionClassOnce.Do(func() {
		CaptureEventInteractionClass = _CaptureEventInteractionClass{objc.GetClass("AVCaptureEventInteraction")}
	})
	return CaptureEventInteractionClass
}

type _CaptureEventInteractionClass struct {
	class objc.Class
}

// An interface definition for the [CaptureEventInteraction] class.
type ICaptureEventInteraction interface {
	objectivec.IObject
}

// An object that registers handlers to respond to capture events from system hardware buttons.
//
// The system Camera app allows people to perform capture functions by pressing hardware buttons on their iOS device. UIKit apps can add similar functionality by using this type to register handlers that respond to interactions from device hardware. The following example shows how to add a handler that captures a photo when a user presses a hardware button on their device. The event handler queries the capture event to determine its phase, and when the interaction ends, captures a photo.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventInteraction
type CaptureEventInteraction struct {
	objectivec.Object
}

// CaptureEventInteractionFrom constructs a [CaptureEventInteraction] from an unsafe.Pointer.
//
// An object that registers handlers to respond to capture events from system hardware buttons.
func CaptureEventInteractionFrom(ptr unsafe.Pointer) CaptureEventInteraction {
	return CaptureEventInteraction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureEventInteractionClass) Alloc() CaptureEventInteraction {
	rv := objc.Send[CaptureEventInteraction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureEventInteractionClass) New() CaptureEventInteraction {
	rv := objc.Send[CaptureEventInteraction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureEventInteraction) Init() CaptureEventInteraction {
	rv := objc.Send[CaptureEventInteraction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureEventInteraction) Autorelease() CaptureEventInteraction {
	rv := objc.Send[CaptureEventInteraction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureEventInteraction creates a new CaptureEventInteraction instance.
func NewCaptureEventInteraction() CaptureEventInteraction {
	return getCaptureEventInteractionClass().New()
}


// Creates a capture event interaction with handlers that respond independently to presses of hardware buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventInteraction/init(primary:secondary:)
func NewCaptureEventInteractionWithPrimaryEventHandlerSecondaryEventHandler(primaryHandler unsafe.Pointer, secondaryHandler unsafe.Pointer) CaptureEventInteraction {
	instance := getCaptureEventInteractionClass().Alloc()
	rv := objc.Send[CaptureEventInteraction](instance.ID, objc.Sel("initWithPrimaryEventHandler:secondaryEventHandler:"), primaryHandler, secondaryHandler)
	rv.Autorelease()
	return rv
}

// Creates a capture event interaction with a handler that responds to presses of hardware buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventInteraction/init(handler:)
func NewCaptureEventInteractionWithEventHandler(handler unsafe.Pointer) CaptureEventInteraction {
	instance := getCaptureEventInteractionClass().Alloc()
	rv := objc.Send[CaptureEventInteraction](instance.ID, objc.Sel("initWithEventHandler:"), handler)
	rv.Autorelease()
	return rv
}


// A Boolean value that indicates whether this capture event interaction is in an enabled state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventInteraction/isEnabled
func (c_ CaptureEventInteraction) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value that indicates whether this capture event interaction is in an enabled state.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventInteraction/isEnabled
func (c_ CaptureEventInteraction) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}

