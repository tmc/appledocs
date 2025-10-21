// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZMacGraphicsDisplayConfiguration] class.
var (
	VZMacGraphicsDisplayConfigurationClass     _VZMacGraphicsDisplayConfigurationClass
	VZMacGraphicsDisplayConfigurationClassOnce sync.Once
)

func getVZMacGraphicsDisplayConfigurationClass() _VZMacGraphicsDisplayConfigurationClass {
	VZMacGraphicsDisplayConfigurationClassOnce.Do(func() {
		VZMacGraphicsDisplayConfigurationClass = _VZMacGraphicsDisplayConfigurationClass{objc.GetClass("VZMacGraphicsDisplayConfiguration")}
	})
	return VZMacGraphicsDisplayConfigurationClass
}

type _VZMacGraphicsDisplayConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZMacGraphicsDisplayConfiguration] class.
type IVZMacGraphicsDisplayConfiguration interface {
	IVZGraphicsDisplayConfiguration
}

// The configuration for a Mac graphics device.
//
// Use this device to attach a display that’s shown in a .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration
type VZMacGraphicsDisplayConfiguration struct {
	VZGraphicsDisplayConfiguration
}

// VZMacGraphicsDisplayConfigurationFrom constructs a [VZMacGraphicsDisplayConfiguration] from an unsafe.Pointer.
//
// The configuration for a Mac graphics device.
func VZMacGraphicsDisplayConfigurationFrom(ptr unsafe.Pointer) VZMacGraphicsDisplayConfiguration {
	return VZMacGraphicsDisplayConfiguration{
		VZGraphicsDisplayConfiguration: VZGraphicsDisplayConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMacGraphicsDisplayConfigurationClass) Alloc() VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMacGraphicsDisplayConfigurationClass) New() VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacGraphicsDisplayConfiguration) Init() VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacGraphicsDisplayConfiguration) Autorelease() VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDisplayConfiguration creates a new VZMacGraphicsDisplayConfiguration instance.
func NewVZMacGraphicsDisplayConfiguration() VZMacGraphicsDisplayConfiguration {
	return getVZMacGraphicsDisplayConfigurationClass().New()
}




// Create a display configuration suitable for showing on the specified screen.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/init(for:sizeInPoints:)
func NewVZMacGraphicsDisplayConfigurationForScreenSizeInPoints(screen appkit.IScreen, sizeInPoints foundation.ISize) VZMacGraphicsDisplayConfiguration {
	instance := getVZMacGraphicsDisplayConfigurationClass().Alloc()
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](instance.ID, objc.Sel("initForScreen:sizeInPoints:"), screen, sizeInPoints)
	rv.Autorelease()
	return rv
}



// Create a display configuration with the specified pixel dimensions and pixel density.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/init(widthInPixels:heightInPixels:pixelsPerInch:)
func NewVZMacGraphicsDisplayConfigurationWithWidthInPixelsHeightInPixelsPixelsPerInch(widthInPixels int, heightInPixels int, pixelsPerInch int) VZMacGraphicsDisplayConfiguration {
	instance := getVZMacGraphicsDisplayConfigurationClass().Alloc()
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](instance.ID, objc.Sel("initWithWidthInPixels:heightInPixels:pixelsPerInch:"), widthInPixels, heightInPixels, pixelsPerInch)
	rv.Autorelease()
	return rv
}


// The height of the display, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/heightInPixels
func (v_ VZMacGraphicsDisplayConfiguration) HeightInPixels() int {
	rv := objc.Send[int](v_.ID, objc.Sel("heightInPixels"))
	return rv
}


// SetHeightInPixels sets the value of the heightInPixels property.
// The height of the display, in pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/heightInPixels
func (v_ VZMacGraphicsDisplayConfiguration) SetHeightInPixels(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHeightInPixels:"), value)
}

// The pixel density in pixels per inch.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/pixelsPerInch
func (v_ VZMacGraphicsDisplayConfiguration) PixelsPerInch() int {
	rv := objc.Send[int](v_.ID, objc.Sel("pixelsPerInch"))
	return rv
}


// SetPixelsPerInch sets the value of the pixelsPerInch property.
// The pixel density in pixels per inch.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/pixelsPerInch
func (v_ VZMacGraphicsDisplayConfiguration) SetPixelsPerInch(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPixelsPerInch:"), value)
}

// The width of the display, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/widthInPixels
func (v_ VZMacGraphicsDisplayConfiguration) WidthInPixels() int {
	rv := objc.Send[int](v_.ID, objc.Sel("widthInPixels"))
	return rv
}


// SetWidthInPixels sets the value of the widthInPixels property.
// The width of the display, in pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/widthInPixels
func (v_ VZMacGraphicsDisplayConfiguration) SetWidthInPixels(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWidthInPixels:"), value)
}


