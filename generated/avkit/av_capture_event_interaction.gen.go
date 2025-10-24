// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureEventInteraction */


/* debug [class_header]: Header for AVCaptureEventInteraction */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureEventInteraction */
// An interface definition for the [CaptureEventInteraction] class.
type ICaptureEventInteraction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureEventInteraction */
	// properties:
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureEventInteraction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureEventInteraction */
// Alloc allocates a new instance without initialization.
func (cc _CaptureEventInteractionClass) Alloc() CaptureEventInteraction {
	rv := objc.Send[CaptureEventInteraction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureEventInteraction */
// An object that registers handlers to respond to capture events from system hardware buttons.
//
// The system Camera app allows people to perform capture functions by pressing hardware buttons on their iOS device. UIKit apps can add similar functionality by using this type to register handlers that respond to interactions from device hardware. The following example shows how to add a handler that captures a photo when a user presses a hardware button on their device. The event handler queries the capture event to determine its phase, and when the interaction ends, captures a photo.


// An object that registers handlers to respond to capture events from system hardware buttons.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureEventInteraction */

// Creates a capture event interaction with a handler that responds to presses of hardware buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventInteraction/init(handler:)
func NewCaptureEventInteractionWithEventHandler(handler unsafe.Pointer) CaptureEventInteraction {
	instance := getCaptureEventInteractionClass().Alloc()
	rv := objc.Send[CaptureEventInteraction](instance.ID, objc.Sel("initWithEventHandler:"), handler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureEventInteractionWithEventHandler */


// Creates a capture event interaction with handlers that respond independently to presses of hardware buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventInteraction/init(primary:secondary:)
func NewCaptureEventInteractionWithPrimaryEventHandlerSecondaryEventHandler(primaryHandler unsafe.Pointer, secondaryHandler unsafe.Pointer) CaptureEventInteraction {
	instance := getCaptureEventInteractionClass().Alloc()
	rv := objc.Send[CaptureEventInteraction](instance.ID, objc.Sel("initWithPrimaryEventHandler:secondaryEventHandler:"), primaryHandler, secondaryHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureEventInteractionWithPrimaryEventHandlerSecondaryEventHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureEventInteraction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureEventInteraction */

// A Boolean value that indicates whether the default sound is in a disabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventInteraction/defaultCaptureSoundDisabled
func (cc _CaptureEventInteractionClass) DefaultCaptureSoundDisabled() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("defaultCaptureSoundDisabled"))
	return rv
}/* debug [class_properties_class/property]: defaultCaptureSoundDisabled */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureEventInteraction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureEventInteraction */

// A Boolean value that indicates whether this capture event interaction is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureeventinteraction/isenabled
func (c_ CaptureEventInteraction) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether this capture event interaction is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureeventinteraction/isenabled
func (c_ CaptureEventInteraction) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureEventInteraction */


