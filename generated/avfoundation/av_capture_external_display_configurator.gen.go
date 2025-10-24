// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureExternalDisplayConfigurator] class.
var (
	CaptureExternalDisplayConfiguratorClass     _CaptureExternalDisplayConfiguratorClass
	CaptureExternalDisplayConfiguratorClassOnce sync.Once
)

func getCaptureExternalDisplayConfiguratorClass() _CaptureExternalDisplayConfiguratorClass {
	CaptureExternalDisplayConfiguratorClassOnce.Do(func() {
		CaptureExternalDisplayConfiguratorClass = _CaptureExternalDisplayConfiguratorClass{objc.GetClass("AVCaptureExternalDisplayConfigurator")}
	})
	return CaptureExternalDisplayConfiguratorClass
}

type _CaptureExternalDisplayConfiguratorClass struct {
	class objc.Class
}





// An interface definition for the [CaptureExternalDisplayConfigurator] class.
type ICaptureExternalDisplayConfigurator interface {
	objectivec.IObject
	

	// properties:
	ActiveExternalDisplayFrameRate() float64
	Device() IAVCaptureDevice
	Active() bool
	PreviewLayer() objc.IObject /* cross-framework: Layer */
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
	IsActive() bool
	SetIsActive(value bool)


	

	// methods:
	Stop()


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureExternalDisplayConfiguratorClass) Alloc() CaptureExternalDisplayConfigurator {
	rv := objc.Send[CaptureExternalDisplayConfigurator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureExternalDisplayConfiguratorClass) New() CaptureExternalDisplayConfigurator {
	rv := objc.Send[CaptureExternalDisplayConfigurator](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureExternalDisplayConfigurator) Init() CaptureExternalDisplayConfigurator {
	rv := objc.Send[CaptureExternalDisplayConfigurator](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureExternalDisplayConfigurator) Autorelease() CaptureExternalDisplayConfigurator {
	rv := objc.Send[CaptureExternalDisplayConfigurator](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureExternalDisplayConfigurator creates a new CaptureExternalDisplayConfigurator instance.
func NewCaptureExternalDisplayConfigurator() CaptureExternalDisplayConfigurator {
	return getCaptureExternalDisplayConfiguratorClass().New()
}





// A configurator class allowing you to configure properties of an external display to match the camera’s active video format.
//
// An allows you to configure a connected external display to output a clean feed using a . Using the configurator, you can opt into automatic adjustment of the external display’s color space and / or frame rate to match your device’s capture configuration. These adjustments are only applied to the external display, not to the device.


// A configurator class allowing you to configure properties of an external display to match the camera’s active video format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator
type CaptureExternalDisplayConfigurator struct {
	objectivec.Object
}

// CaptureExternalDisplayConfiguratorFrom constructs a [CaptureExternalDisplayConfigurator] from an unsafe.Pointer.
//
// A configurator class allowing you to configure properties of an external display to match the camera’s active video format.
func CaptureExternalDisplayConfiguratorFrom(ptr unsafe.Pointer) CaptureExternalDisplayConfigurator {
	return CaptureExternalDisplayConfigurator{objectivec.Object{objc.ID(ptr)}}
}






// An external display configurator instance that attempts to synchronize the preview layer configuration with the device capture configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/init(device:previewLayer:configuration:)
func NewCaptureExternalDisplayConfiguratorWithDevicePreviewLayerConfiguration(device IAVCaptureDevice, previewLayer objc.IObject /* cross-framework: Layer */, configuration IAVCaptureExternalDisplayConfiguration) CaptureExternalDisplayConfigurator {
	instance := getCaptureExternalDisplayConfiguratorClass().Alloc()
	rv := objc.Send[CaptureExternalDisplayConfigurator](instance.ID, objc.Sel("initWithDevice:previewLayer:configuration:"), device, previewLayer, configuration)
	rv.Autorelease()
	return rv
}












// Whether the external display supports bypassing color space conversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/isBypassingColorSpaceConversionSupported
func (cc _CaptureExternalDisplayConfiguratorClass) SupportsBypassingColorSpaceConversion() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("supportsBypassingColorSpaceConversion"))
	return rv
}

// Whether the external display supports matching frame rate to a capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/isMatchingFrameRateSupported
func (cc _CaptureExternalDisplayConfiguratorClass) ShouldMatchFrameRateSupported() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("shouldMatchFrameRateSupported"))
	return rv
}

// Whether the external display supports configuration to your preferred resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/isPreferredResolutionSupported
func (cc _CaptureExternalDisplayConfiguratorClass) SupportsPreferredResolution() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("supportsPreferredResolution"))
	return rv
}






// Forces the external display configurator to asynchronously stop configuring the external display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/stop()
func (c_ CaptureExternalDisplayConfigurator) Stop() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stop"))
}







// The currently configured frame rate on the external display that’s displaying the preview layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/activeExternalDisplayFrameRate
func (c_ CaptureExternalDisplayConfigurator) ActiveExternalDisplayFrameRate() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("activeExternalDisplayFrameRate"))
	return rv
}


// The device for which the coordinator configures the preview layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/device
func (c_ CaptureExternalDisplayConfigurator) Device() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("device"))
	return rv
}


// This property tells you whether the configurator is actively configuring the external display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/isActive
func (c_ CaptureExternalDisplayConfigurator) Active() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("active"))
	return rv
}


// Whether the external display supports bypassing color space conversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/isBypassingColorSpaceConversionSupported
func (c_ CaptureExternalDisplayConfigurator) SupportsBypassingColorSpaceConversion() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsBypassingColorSpaceConversion"))
	return rv
}


// Whether the external display supports matching frame rate to a capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/isMatchingFrameRateSupported
func (c_ CaptureExternalDisplayConfigurator) ShouldMatchFrameRateSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldMatchFrameRateSupported"))
	return rv
}


// Whether the external display supports configuration to your preferred resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/isPreferredResolutionSupported
func (c_ CaptureExternalDisplayConfigurator) SupportsPreferredResolution() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsPreferredResolution"))
	return rv
}


// The layer for which the configurator adjusts display properties to match the device’s state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureExternalDisplayConfigurator/previewLayer
func (c_ CaptureExternalDisplayConfigurator) PreviewLayer() objc.IObject /* cross-framework: Layer */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("previewLayer"))
	return rv
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureExternalDisplayConfigurator) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeFormat"))
	return rv
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureExternalDisplayConfigurator) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}


// This property tells you whether the configurator is actively configuring the external display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureexternaldisplayconfigurator/isactive
func (c_ CaptureExternalDisplayConfigurator) IsActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isActive"))
	return rv
}


// This property tells you whether the configurator is actively configuring the external display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureexternaldisplayconfigurator/isactive
func (c_ CaptureExternalDisplayConfigurator) SetIsActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsActive:"), value)
}







