// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureExternalDisplayConfiguration */


/* debug [class_header]: Header for AVCaptureExternalDisplayConfiguration */
// The class instance for the [CaptureExternalDisplayConfiguration] class.
var (
	CaptureExternalDisplayConfigurationClass     _CaptureExternalDisplayConfigurationClass
	CaptureExternalDisplayConfigurationClassOnce sync.Once
)

func getCaptureExternalDisplayConfigurationClass() _CaptureExternalDisplayConfigurationClass {
	CaptureExternalDisplayConfigurationClassOnce.Do(func() {
		CaptureExternalDisplayConfigurationClass = _CaptureExternalDisplayConfigurationClass{objc.GetClass("AVCaptureExternalDisplayConfiguration")}
	})
	return CaptureExternalDisplayConfigurationClass
}

type _CaptureExternalDisplayConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureExternalDisplayConfiguration */
// An interface definition for the [CaptureExternalDisplayConfiguration] class.
type ICaptureExternalDisplayConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureExternalDisplayConfiguration */
	// properties:
	BypassColorSpaceConversion() bool
	SetBypassColorSpaceConversion(value bool)
	PreferredResolution() VideoDimensions /* not a class type */
	SetPreferredResolution(value VideoDimensions /* not a class type */)
	ShouldMatchFrameRate() bool
	SetShouldMatchFrameRate(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureExternalDisplayConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureExternalDisplayConfiguration */
// Alloc allocates a new instance without initialization.
func (cc _CaptureExternalDisplayConfigurationClass) Alloc() CaptureExternalDisplayConfiguration {
	rv := objc.Send[CaptureExternalDisplayConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureExternalDisplayConfigurationClass) New() CaptureExternalDisplayConfiguration {
	rv := objc.Send[CaptureExternalDisplayConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureExternalDisplayConfiguration) Init() CaptureExternalDisplayConfiguration {
	rv := objc.Send[CaptureExternalDisplayConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureExternalDisplayConfiguration) Autorelease() CaptureExternalDisplayConfiguration {
	rv := objc.Send[CaptureExternalDisplayConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureExternalDisplayConfiguration creates a new CaptureExternalDisplayConfiguration instance.
func NewCaptureExternalDisplayConfiguration() CaptureExternalDisplayConfiguration {
	return getCaptureExternalDisplayConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureExternalDisplayConfiguration */
// A class you use to specify a configuration to your external display configurator.
//
// Using an , you direct your how to configure an external display to match your device’s active video format.


// A class you use to specify a configuration to your external display configurator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration
type CaptureExternalDisplayConfiguration struct {
	objectivec.Object
}

// CaptureExternalDisplayConfigurationFrom constructs a [CaptureExternalDisplayConfiguration] from an unsafe.Pointer.
//
// A class you use to specify a configuration to your external display configurator.
func CaptureExternalDisplayConfigurationFrom(ptr unsafe.Pointer) CaptureExternalDisplayConfiguration {
	return CaptureExternalDisplayConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureExternalDisplayConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureExternalDisplayConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureExternalDisplayConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureExternalDisplayConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureExternalDisplayConfiguration */

// A property indicating whether the color space of the configurator’s preview layer should be preserved on the output display by avoiding color space conversions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/bypassColorSpaceConversion
func (c_ CaptureExternalDisplayConfiguration) BypassColorSpaceConversion() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bypassColorSpaceConversion"))
	return rv
}/* debug [instance_properties/getter]: bypassColorSpaceConversion */


// A property indicating whether the color space of the configurator’s preview layer should be preserved on the output display by avoiding color space conversions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/bypassColorSpaceConversion
func (c_ CaptureExternalDisplayConfiguration) SetBypassColorSpaceConversion(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBypassColorSpaceConversion:"), value)
}/* debug [instance_properties/setter]: bypassColorSpaceConversion */


// Your preferred external display resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/preferredResolution
func (c_ CaptureExternalDisplayConfiguration) PreferredResolution() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("preferredResolution"))
	return rv
}/* debug [instance_properties/getter]: preferredResolution */


// Your preferred external display resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/preferredResolution
func (c_ CaptureExternalDisplayConfiguration) SetPreferredResolution(value VideoDimensions /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredResolution:"), value)
}/* debug [instance_properties/setter]: preferredResolution */


// A property indicating whether the frame rate of the external display should be configured to match the camera’s frame rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/shouldMatchFrameRate
func (c_ CaptureExternalDisplayConfiguration) ShouldMatchFrameRate() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldMatchFrameRate"))
	return rv
}/* debug [instance_properties/getter]: shouldMatchFrameRate */


// A property indicating whether the frame rate of the external display should be configured to match the camera’s frame rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/shouldMatchFrameRate
func (c_ CaptureExternalDisplayConfiguration) SetShouldMatchFrameRate(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldMatchFrameRate:"), value)
}/* debug [instance_properties/setter]: shouldMatchFrameRate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureExternalDisplayConfiguration */



