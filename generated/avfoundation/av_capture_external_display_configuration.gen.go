// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CaptureExternalDisplayConfiguration] class.
type ICaptureExternalDisplayConfiguration interface {
	objectivec.IObject
	

	// properties:
	BypassColorSpaceConversion() bool
	SetBypassColorSpaceConversion(value bool)
	PreferredResolution() objectivec.IObject
	SetPreferredResolution(value objectivec.IObject)
	ShouldMatchFrameRate() bool
	SetShouldMatchFrameRate(value bool)


	

	// methods:


}





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

























// A property indicating whether the color space of the configurator’s preview layer should be preserved on the output display by avoiding color space conversions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/bypassColorSpaceConversion
func (c_ CaptureExternalDisplayConfiguration) BypassColorSpaceConversion() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bypassColorSpaceConversion"))
	return rv
}


// A property indicating whether the color space of the configurator’s preview layer should be preserved on the output display by avoiding color space conversions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/bypassColorSpaceConversion
func (c_ CaptureExternalDisplayConfiguration) SetBypassColorSpaceConversion(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBypassColorSpaceConversion:"), value)
}


// Your preferred external display resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/preferredResolution
func (c_ CaptureExternalDisplayConfiguration) PreferredResolution() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("preferredResolution"))
	return rv
}


// Your preferred external display resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/preferredResolution
func (c_ CaptureExternalDisplayConfiguration) SetPreferredResolution(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredResolution:"), value)
}


// A property indicating whether the frame rate of the external display should be configured to match the camera’s frame rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/shouldMatchFrameRate
func (c_ CaptureExternalDisplayConfiguration) ShouldMatchFrameRate() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldMatchFrameRate"))
	return rv
}


// A property indicating whether the frame rate of the external display should be configured to match the camera’s frame rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfiguration/shouldMatchFrameRate
func (c_ CaptureExternalDisplayConfiguration) SetShouldMatchFrameRate(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldMatchFrameRate:"), value)
}








