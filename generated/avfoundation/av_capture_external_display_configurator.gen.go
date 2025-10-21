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




