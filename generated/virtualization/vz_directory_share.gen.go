// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZDirectoryShare */


/* debug [class_header]: Header for VZDirectoryShare */
// The class instance for the [VZDirectoryShare] class.
var (
	VZDirectoryShareClass     _VZDirectoryShareClass
	VZDirectoryShareClassOnce sync.Once
)

func getVZDirectoryShareClass() _VZDirectoryShareClass {
	VZDirectoryShareClassOnce.Do(func() {
		VZDirectoryShareClass = _VZDirectoryShareClass{objc.GetClass("VZDirectoryShare")}
	})
	return VZDirectoryShareClass
}

type _VZDirectoryShareClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZDirectoryShare */
// An interface definition for the [VZDirectoryShare] class.
type IVZDirectoryShare interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZDirectoryShare */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZDirectoryShare */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZDirectoryShare */
// Alloc allocates a new instance without initialization.
func (vc _VZDirectoryShareClass) Alloc() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZDirectoryShareClass) New() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZDirectoryShare) Init() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZDirectoryShare) Autorelease() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectoryShare creates a new VZDirectoryShare instance.
func NewVZDirectoryShare() VZDirectoryShare {
	return getVZDirectoryShareClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZDirectoryShare */
// The base class for a directory share.
//
// A directory share defines how the system exposes host directories to a guest VM. Don’t instantiate directly, use one of its subclasses such as or instead.


// The base class for a directory share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare
type VZDirectoryShare struct {
	objectivec.Object
}

// VZDirectoryShareFrom constructs a [VZDirectoryShare] from an unsafe.Pointer.
//
// The base class for a directory share.
func VZDirectoryShareFrom(ptr unsafe.Pointer) VZDirectoryShare {
	return VZDirectoryShare{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZDirectoryShare *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZDirectoryShare */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZDirectoryShare */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZDirectoryShare */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZDirectoryShare */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZDirectoryShare */



