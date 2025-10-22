// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/quartzcore"
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
	ActiveFormat() unsafe.Pointer
	SetActiveFormat(value unsafe.Pointer)
	ActiveExternalDisplayFrameRate() float64
	SetActiveExternalDisplayFrameRate(value float64)
	Device() AVCaptureDevice
	SetDevice(value IAVCaptureDevice)
	IsActive() bool
	SetIsActive(value bool)
	PreviewLayer() quartzcore.Layer
	SetPreviewLayer(value quartzcore.ILayer)
}

// A configurator class allowing you to configure properties of an external display to match the camera’s active video format.
//
// An allows you to configure a connected external display to output a clean feed using a . Using the configurator, you can opt into automatic adjustment of the external display’s color space and / or frame rate to match your device’s capture configuration. These adjustments are only applied to the external display, not to the device.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CaptureExternalDisplayConfiguratorClass) Alloc() CaptureExternalDisplayConfigurator {
	rv := objc.Send[CaptureExternalDisplayConfigurator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The capture format in use by the device.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureExternalDisplayConfigurator) ActiveFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeFormat"))
	return rv
}


// SetActiveFormat sets the value of the activeFormat property.
// The capture format in use by the device.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureExternalDisplayConfigurator) SetActiveFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}

// The currently configured frame rate on the external display that’s displaying the preview layer.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureexternaldisplayconfigurator/activeexternaldisplayframerate
func (c_ CaptureExternalDisplayConfigurator) ActiveExternalDisplayFrameRate() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("activeExternalDisplayFrameRate"))
	return rv
}


// SetActiveExternalDisplayFrameRate sets the value of the activeExternalDisplayFrameRate property.
// The currently configured frame rate on the external display that’s displaying the preview layer.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureexternaldisplayconfigurator/activeexternaldisplayframerate
func (c_ CaptureExternalDisplayConfigurator) SetActiveExternalDisplayFrameRate(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveExternalDisplayFrameRate:"), value)
}

// The device for which the coordinator configures the preview layer.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureexternaldisplayconfigurator/device
func (c_ CaptureExternalDisplayConfigurator) Device() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("device"))
	return rv
}


// SetDevice sets the value of the device property.
// The device for which the coordinator configures the preview layer.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureexternaldisplayconfigurator/device
func (c_ CaptureExternalDisplayConfigurator) SetDevice(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDevice:"), value)
}

// This property tells you whether the configurator is actively configuring the external display.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureexternaldisplayconfigurator/isactive
func (c_ CaptureExternalDisplayConfigurator) IsActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isActive"))
	return rv
}


// SetIsActive sets the value of the isActive property.
// This property tells you whether the configurator is actively configuring the external display.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureexternaldisplayconfigurator/isactive
func (c_ CaptureExternalDisplayConfigurator) SetIsActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsActive:"), value)
}

// The layer for which the configurator adjusts display properties to match the device’s state.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureexternaldisplayconfigurator/previewlayer
func (c_ CaptureExternalDisplayConfigurator) PreviewLayer() quartzcore.Layer {
	rv := objc.Send[quartzcore.Layer](c_.ID, objc.Sel("previewLayer"))
	return rv
}


// SetPreviewLayer sets the value of the previewLayer property.
// The layer for which the configurator adjusts display properties to match the device’s state.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureexternaldisplayconfigurator/previewlayer
func (c_ CaptureExternalDisplayConfigurator) SetPreviewLayer(value quartzcore.ILayer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviewLayer:"), value)
}



