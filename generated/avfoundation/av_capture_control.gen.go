// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureControl */


/* debug [class_header]: Header for AVCaptureControl */
// The class instance for the [CaptureControl] class.
var (
	CaptureControlClass     _CaptureControlClass
	CaptureControlClassOnce sync.Once
)

func getCaptureControlClass() _CaptureControlClass {
	CaptureControlClassOnce.Do(func() {
		CaptureControlClass = _CaptureControlClass{objc.GetClass("AVCaptureControl")}
	})
	return CaptureControlClass
}

type _CaptureControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureControl */
// An interface definition for the [CaptureControl] class.
type ICaptureControl interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureControl */
	// properties:
	Enabled() bool
	SetEnabled(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureControl */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureControl */
// Alloc allocates a new instance without initialization.
func (cc _CaptureControlClass) Alloc() CaptureControl {
	rv := objc.Send[CaptureControl](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureControlClass) New() CaptureControl {
	rv := objc.Send[CaptureControl](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureControl) Init() CaptureControl {
	rv := objc.Send[CaptureControl](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureControl) Autorelease() CaptureControl {
	rv := objc.Send[CaptureControl](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureControl creates a new CaptureControl instance.
func NewCaptureControl() CaptureControl {
	return getCaptureControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureControl */
// An abstract base class for controls that interact with the camera system.
//
// Capture controls provide the interface for interacting with the camera system from the Camera Control button on iPhone 16 devices. The framework provides several concrete subclasses of this class that allow apps to access built-in functionality and define custom controls.


// An abstract base class for controls that interact with the camera system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureControl
type CaptureControl struct {
	objectivec.Object
}

// CaptureControlFrom constructs a [CaptureControl] from an unsafe.Pointer.
//
// An abstract base class for controls that interact with the camera system.
func CaptureControlFrom(ptr unsafe.Pointer) CaptureControl {
	return CaptureControl{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureControl *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureControl */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureControl */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureControl */

// A Boolean value that indicates whether this control supports user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureControl/isEnabled
func (c_ CaptureControl) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether this control supports user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureControl/isEnabled
func (c_ CaptureControl) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A Boolean value that indicates whether this control supports user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturecontrol/isenabled
func (c_ CaptureControl) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether this control supports user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturecontrol/isenabled
func (c_ CaptureControl) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureControl */



