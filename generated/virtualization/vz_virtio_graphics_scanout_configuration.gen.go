// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioGraphicsScanoutConfiguration] class.
var (
	VZVirtioGraphicsScanoutConfigurationClass     _VZVirtioGraphicsScanoutConfigurationClass
	VZVirtioGraphicsScanoutConfigurationClassOnce sync.Once
)

func getVZVirtioGraphicsScanoutConfigurationClass() _VZVirtioGraphicsScanoutConfigurationClass {
	VZVirtioGraphicsScanoutConfigurationClassOnce.Do(func() {
		VZVirtioGraphicsScanoutConfigurationClass = _VZVirtioGraphicsScanoutConfigurationClass{objc.GetClass("VZVirtioGraphicsScanoutConfiguration")}
	})
	return VZVirtioGraphicsScanoutConfigurationClass
}

type _VZVirtioGraphicsScanoutConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioGraphicsScanoutConfiguration] class.
type IVZVirtioGraphicsScanoutConfiguration interface {
	IVZGraphicsDisplayConfiguration
	// properties:
	Scanouts() IVZVirtioGraphicsScanoutConfiguration
	SetScanouts(value IVZVirtioGraphicsScanoutConfiguration)
	HeightInPixels() int
	SetHeightInPixels(value int)
	WidthInPixels() int
	SetWidthInPixels(value int)
	// methods:
}

// The configuration for a Virtio graphics device that configures the dimensions of the graphics device for a Linux VM.
//
// Use a to configure the width and height of a Virtio graphics device.


// The configuration for a Virtio graphics device that configures the dimensions of the graphics device for a Linux VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanoutConfiguration
type VZVirtioGraphicsScanoutConfiguration struct {
	VZGraphicsDisplayConfiguration
}

// VZVirtioGraphicsScanoutConfigurationFrom constructs a [VZVirtioGraphicsScanoutConfiguration] from an unsafe.Pointer.
//
// The configuration for a Virtio graphics device that configures the dimensions of the graphics device for a Linux VM.
func VZVirtioGraphicsScanoutConfigurationFrom(ptr unsafe.Pointer) VZVirtioGraphicsScanoutConfiguration {
	return VZVirtioGraphicsScanoutConfiguration{
		VZGraphicsDisplayConfiguration: VZGraphicsDisplayConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioGraphicsScanoutConfigurationClass) Alloc() VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioGraphicsScanoutConfigurationClass) New() VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioGraphicsScanoutConfiguration) Init() VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioGraphicsScanoutConfiguration) Autorelease() VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsScanoutConfiguration creates a new VZVirtioGraphicsScanoutConfiguration instance.
func NewVZVirtioGraphicsScanoutConfiguration() VZVirtioGraphicsScanoutConfiguration {
	return getVZVirtioGraphicsScanoutConfigurationClass().New()
}



// The array of output devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiographicsdeviceconfiguration/scanouts
func (v_ VZVirtioGraphicsScanoutConfiguration) Scanouts() IVZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](v_.ID, objc.Sel("scanouts"))
	return rv
}


// The array of output devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiographicsdeviceconfiguration/scanouts
func (v_ VZVirtioGraphicsScanoutConfiguration) SetScanouts(value IVZVirtioGraphicsScanoutConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setScanouts:"), value)
}


// An integer value that describes the height of the graphics device in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiographicsscanoutconfiguration/heightinpixels
func (v_ VZVirtioGraphicsScanoutConfiguration) HeightInPixels() int {
	rv := objc.Send[int](v_.ID, objc.Sel("heightInPixels"))
	return rv
}


// An integer value that describes the height of the graphics device in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiographicsscanoutconfiguration/heightinpixels
func (v_ VZVirtioGraphicsScanoutConfiguration) SetHeightInPixels(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHeightInPixels:"), value)
}


// An integer value that describes the width of the graphics device in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiographicsscanoutconfiguration/widthinpixels
func (v_ VZVirtioGraphicsScanoutConfiguration) WidthInPixels() int {
	rv := objc.Send[int](v_.ID, objc.Sel("widthInPixels"))
	return rv
}


// An integer value that describes the width of the graphics device in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiographicsscanoutconfiguration/widthinpixels
func (v_ VZVirtioGraphicsScanoutConfiguration) SetWidthInPixels(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWidthInPixels:"), value)
}



