// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZVirtioGraphicsScanout */


/* debug [class_header]: Header for VZVirtioGraphicsScanout */
// The class instance for the [VZVirtioGraphicsScanout] class.
var (
	VZVirtioGraphicsScanoutClass     _VZVirtioGraphicsScanoutClass
	VZVirtioGraphicsScanoutClassOnce sync.Once
)

func getVZVirtioGraphicsScanoutClass() _VZVirtioGraphicsScanoutClass {
	VZVirtioGraphicsScanoutClassOnce.Do(func() {
		VZVirtioGraphicsScanoutClass = _VZVirtioGraphicsScanoutClass{objc.GetClass("VZVirtioGraphicsScanout")}
	})
	return VZVirtioGraphicsScanoutClass
}

type _VZVirtioGraphicsScanoutClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioGraphicsScanout */
// An interface definition for the [VZVirtioGraphicsScanout] class.
type IVZVirtioGraphicsScanout interface {
	IVZGraphicsDisplay
	
/* debug [class_interface_properties]: Properties for VZVirtioGraphicsScanout */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioGraphicsScanout */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioGraphicsScanout */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioGraphicsScanoutClass) Alloc() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioGraphicsScanoutClass) New() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioGraphicsScanout) Init() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioGraphicsScanout) Autorelease() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsScanout creates a new VZVirtioGraphicsScanout instance.
func NewVZVirtioGraphicsScanout() VZVirtioGraphicsScanout {
	return getVZVirtioGraphicsScanoutClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioGraphicsScanout */
// A Virtio graphics scanout that corresponds to a Virtio graphics scanout configuration.


// A Virtio graphics scanout that corresponds to a Virtio graphics scanout configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanout
type VZVirtioGraphicsScanout struct {
	VZGraphicsDisplay
}

// VZVirtioGraphicsScanoutFrom constructs a [VZVirtioGraphicsScanout] from an unsafe.Pointer.
//
// A Virtio graphics scanout that corresponds to a Virtio graphics scanout configuration.
func VZVirtioGraphicsScanoutFrom(ptr unsafe.Pointer) VZVirtioGraphicsScanout {
	return VZVirtioGraphicsScanout{
		VZGraphicsDisplay: VZGraphicsDisplayFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioGraphicsScanout *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioGraphicsScanout */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioGraphicsScanout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioGraphicsScanout */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioGraphicsScanout */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioGraphicsScanout */



