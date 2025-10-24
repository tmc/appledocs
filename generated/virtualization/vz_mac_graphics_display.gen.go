// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZMacGraphicsDisplay */


/* debug [class_header]: Header for VZMacGraphicsDisplay */
// The class instance for the [VZMacGraphicsDisplay] class.
var (
	VZMacGraphicsDisplayClass     _VZMacGraphicsDisplayClass
	VZMacGraphicsDisplayClassOnce sync.Once
)

func getVZMacGraphicsDisplayClass() _VZMacGraphicsDisplayClass {
	VZMacGraphicsDisplayClassOnce.Do(func() {
		VZMacGraphicsDisplayClass = _VZMacGraphicsDisplayClass{objc.GetClass("VZMacGraphicsDisplay")}
	})
	return VZMacGraphicsDisplayClass
}

type _VZMacGraphicsDisplayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZMacGraphicsDisplay */
// An interface definition for the [VZMacGraphicsDisplay] class.
type IVZMacGraphicsDisplay interface {
	IVZGraphicsDisplay
	
/* debug [class_interface_properties]: Properties for VZMacGraphicsDisplay */
	// properties:
	PixelsPerInch() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZMacGraphicsDisplay */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZMacGraphicsDisplay */
// Alloc allocates a new instance without initialization.
func (vc _VZMacGraphicsDisplayClass) Alloc() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacGraphicsDisplayClass) New() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacGraphicsDisplay) Init() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacGraphicsDisplay) Autorelease() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDisplay creates a new VZMacGraphicsDisplay instance.
func NewVZMacGraphicsDisplay() VZMacGraphicsDisplay {
	return getVZMacGraphicsDisplayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZMacGraphicsDisplay */
// An object that represents the graphics display on a Mac.


// An object that represents the graphics display on a Mac.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay
type VZMacGraphicsDisplay struct {
	VZGraphicsDisplay
}

// VZMacGraphicsDisplayFrom constructs a [VZMacGraphicsDisplay] from an unsafe.Pointer.
//
// An object that represents the graphics display on a Mac.
func VZMacGraphicsDisplayFrom(ptr unsafe.Pointer) VZMacGraphicsDisplay {
	return VZMacGraphicsDisplay{
		VZGraphicsDisplay: VZGraphicsDisplayFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZMacGraphicsDisplay *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZMacGraphicsDisplay */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZMacGraphicsDisplay */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZMacGraphicsDisplay */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZMacGraphicsDisplay */

// Returns the pixel density of the display in pixels per inch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay/pixelsPerInch
func (v_ VZMacGraphicsDisplay) PixelsPerInch() int {
	rv := objc.Send[int](v_.ID, objc.Sel("pixelsPerInch"))
	return rv
}/* debug [instance_properties/getter]: pixelsPerInch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZMacGraphicsDisplay */



