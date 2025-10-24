// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZGraphicsDisplay */

/* debug [class_header]: Header for VZGraphicsDisplay */
// The class instance for the [VZGraphicsDisplay] class.
var (
	VZGraphicsDisplayClass     _VZGraphicsDisplayClass
	VZGraphicsDisplayClassOnce sync.Once
)

func getVZGraphicsDisplayClass() _VZGraphicsDisplayClass {
	VZGraphicsDisplayClassOnce.Do(func() {
		VZGraphicsDisplayClass = _VZGraphicsDisplayClass{objc.GetClass("VZGraphicsDisplay")}
	})
	return VZGraphicsDisplayClass
}

type _VZGraphicsDisplayClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZGraphicsDisplay */
// An interface definition for the [VZGraphicsDisplay] class.
type IVZGraphicsDisplay interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZGraphicsDisplay */
	// properties:
	SizeInPixels() corefoundation.CGSize
	Displays() IVZGraphicsDisplay
	SetDisplays(value IVZGraphicsDisplay)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZGraphicsDisplay */
	// methods:
	AddObserver(observer unsafe.Pointer)
	ReconfigureWithConfigurationError(configuration IVZGraphicsDisplayConfiguration, error_ unsafe.Pointer) bool
	ReconfigureWithSizeInPixelsError(sizeInPixels corefoundation.CGSize, error_ unsafe.Pointer) bool
	RemoveObserver(observer unsafe.Pointer)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZGraphicsDisplay */
// Alloc allocates a new instance without initialization.
func (vc _VZGraphicsDisplayClass) Alloc() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZGraphicsDisplayClass) New() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZGraphicsDisplay) Init() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZGraphicsDisplay) Autorelease() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDisplay creates a new VZGraphicsDisplay instance.
func NewVZGraphicsDisplay() VZGraphicsDisplay {
	return getVZGraphicsDisplayClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZGraphicsDisplay */
// A class that represents a graphics display in a VM.
//
// Don’t instantiate a directly. Graphics displays are first configured on a subclass. When you create a from the configuration, the displays are available through the property of the configuration’s .

// A class that represents a graphics display in a VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay
type VZGraphicsDisplay struct {
	objectivec.Object
}

// VZGraphicsDisplayFrom constructs a [VZGraphicsDisplay] from an unsafe.Pointer.
//
// A class that represents a graphics display in a VM.
func VZGraphicsDisplayFrom(ptr unsafe.Pointer) VZGraphicsDisplay {
	return VZGraphicsDisplay{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZGraphicsDisplay */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZGraphicsDisplay */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZGraphicsDisplay */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZGraphicsDisplay */

// Adds an observer to notify about display configuration changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/addObserver(_:)
func (v_ VZGraphicsDisplay) AddObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addObserver:"), observer)
} /* debug [instance_methods/method]: AddObserver */

// Reconfigure this display with the new display configuration you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/reconfigure(configuration:)
func (v_ VZGraphicsDisplay) ReconfigureWithConfigurationError(configuration IVZGraphicsDisplayConfiguration, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("reconfigureWithConfiguration:error:"), configuration, error_)
	return rv
} /* debug [instance_methods/method]: ReconfigureWithConfigurationError */

// Resize this display with the new dimensions you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/reconfigure(sizeInPixels:)
func (v_ VZGraphicsDisplay) ReconfigureWithSizeInPixelsError(sizeInPixels corefoundation.CGSize, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("reconfigureWithSizeInPixels:error:"), sizeInPixels, error_)
	return rv
} /* debug [instance_methods/method]: ReconfigureWithSizeInPixelsError */

// Removes a display configuration change observer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/removeObserver(_:)
func (v_ VZGraphicsDisplay) RemoveObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeObserver:"), observer)
} /* debug [instance_methods/method]: RemoveObserver */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZGraphicsDisplay */

// Returns the size of the display, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/sizeInPixels
func (v_ VZGraphicsDisplay) SizeInPixels() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v_.ID, objc.Sel("sizeInPixels"))
	return rv
} /* debug [instance_properties/getter]: sizeInPixels */

// The list of graphics displays configured for this graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgraphicsdevice/displays
func (v_ VZGraphicsDisplay) Displays() IVZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](v_.ID, objc.Sel("displays"))
	return rv
} /* debug [instance_properties/getter]: displays */

// The list of graphics displays configured for this graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgraphicsdevice/displays
func (v_ VZGraphicsDisplay) SetDisplays(value IVZGraphicsDisplay) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDisplays:"), value)
} /* debug [instance_properties/setter]: displays */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZGraphicsDisplay */
