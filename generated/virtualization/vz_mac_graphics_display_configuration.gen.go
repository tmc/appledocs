// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZMacGraphicsDisplayConfiguration */


/* debug [class_header]: Header for VZMacGraphicsDisplayConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZMacGraphicsDisplayConfiguration */
// An interface definition for the [VZMacGraphicsDisplayConfiguration] class.
type IVZMacGraphicsDisplayConfiguration interface {
	IVZGraphicsDisplayConfiguration
	
/* debug [class_interface_properties]: Properties for VZMacGraphicsDisplayConfiguration */
	// properties:
	HeightInPixels() int
	SetHeightInPixels(value int)
	PixelsPerInch() int
	SetPixelsPerInch(value int)
	WidthInPixels() int
	SetWidthInPixels(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZMacGraphicsDisplayConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZMacGraphicsDisplayConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZMacGraphicsDisplayConfigurationClass) Alloc() VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZMacGraphicsDisplayConfiguration */
// The configuration for a Mac graphics device.
//
// Use this device to attach a display that’s shown in a .


// The configuration for a Mac graphics device.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZMacGraphicsDisplayConfiguration */

// Create a display configuration suitable for showing on the specified screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/init(for:sizeInPoints:)
func NewVZMacGraphicsDisplayConfigurationForScreenSizeInPoints(screen appkit.Screen, sizeInPoints Size /* not a class type */) VZMacGraphicsDisplayConfiguration {
	instance := getVZMacGraphicsDisplayConfigurationClass().Alloc()
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](instance.ID, objc.Sel("initForScreen:sizeInPoints:"), screen, sizeInPoints)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZMacGraphicsDisplayConfigurationForScreenSizeInPoints */


// Create a display configuration with the specified pixel dimensions and pixel density.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/init(widthInPixels:heightInPixels:pixelsPerInch:)
func NewVZMacGraphicsDisplayConfigurationWithWidthInPixelsHeightInPixelsPixelsPerInch(widthInPixels int, heightInPixels int, pixelsPerInch int) VZMacGraphicsDisplayConfiguration {
	instance := getVZMacGraphicsDisplayConfigurationClass().Alloc()
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](instance.ID, objc.Sel("initWithWidthInPixels:heightInPixels:pixelsPerInch:"), widthInPixels, heightInPixels, pixelsPerInch)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZMacGraphicsDisplayConfigurationWithWidthInPixelsHeightInPixelsPixelsPerInch */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZMacGraphicsDisplayConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZMacGraphicsDisplayConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZMacGraphicsDisplayConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZMacGraphicsDisplayConfiguration */

// The height of the display, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/heightInPixels
func (v_ VZMacGraphicsDisplayConfiguration) HeightInPixels() int {
	rv := objc.Send[int](v_.ID, objc.Sel("heightInPixels"))
	return rv
}/* debug [instance_properties/getter]: heightInPixels */


// The height of the display, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/heightInPixels
func (v_ VZMacGraphicsDisplayConfiguration) SetHeightInPixels(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHeightInPixels:"), value)
}/* debug [instance_properties/setter]: heightInPixels */


// The pixel density in pixels per inch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/pixelsPerInch
func (v_ VZMacGraphicsDisplayConfiguration) PixelsPerInch() int {
	rv := objc.Send[int](v_.ID, objc.Sel("pixelsPerInch"))
	return rv
}/* debug [instance_properties/getter]: pixelsPerInch */


// The pixel density in pixels per inch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/pixelsPerInch
func (v_ VZMacGraphicsDisplayConfiguration) SetPixelsPerInch(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPixelsPerInch:"), value)
}/* debug [instance_properties/setter]: pixelsPerInch */


// The width of the display, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/widthInPixels
func (v_ VZMacGraphicsDisplayConfiguration) WidthInPixels() int {
	rv := objc.Send[int](v_.ID, objc.Sel("widthInPixels"))
	return rv
}/* debug [instance_properties/getter]: widthInPixels */


// The width of the display, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/widthInPixels
func (v_ VZMacGraphicsDisplayConfiguration) SetWidthInPixels(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWidthInPixels:"), value)
}/* debug [instance_properties/setter]: widthInPixels */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZMacGraphicsDisplayConfiguration */


