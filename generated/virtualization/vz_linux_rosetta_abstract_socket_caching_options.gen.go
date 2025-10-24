// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZLinuxRosettaAbstractSocketCachingOptions */


/* debug [class_header]: Header for VZLinuxRosettaAbstractSocketCachingOptions */
// The class instance for the [VZLinuxRosettaAbstractSocketCachingOptions] class.
var (
	VZLinuxRosettaAbstractSocketCachingOptionsClass     _VZLinuxRosettaAbstractSocketCachingOptionsClass
	VZLinuxRosettaAbstractSocketCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaAbstractSocketCachingOptionsClass() _VZLinuxRosettaAbstractSocketCachingOptionsClass {
	VZLinuxRosettaAbstractSocketCachingOptionsClassOnce.Do(func() {
		VZLinuxRosettaAbstractSocketCachingOptionsClass = _VZLinuxRosettaAbstractSocketCachingOptionsClass{objc.GetClass("VZLinuxRosettaAbstractSocketCachingOptions")}
	})
	return VZLinuxRosettaAbstractSocketCachingOptionsClass
}

type _VZLinuxRosettaAbstractSocketCachingOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZLinuxRosettaAbstractSocketCachingOptions */
// An interface definition for the [VZLinuxRosettaAbstractSocketCachingOptions] class.
type IVZLinuxRosettaAbstractSocketCachingOptions interface {
	IVZLinuxRosettaCachingOptions
	
/* debug [class_interface_properties]: Properties for VZLinuxRosettaAbstractSocketCachingOptions */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZLinuxRosettaAbstractSocketCachingOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZLinuxRosettaAbstractSocketCachingOptions */
// Alloc allocates a new instance without initialization.
func (vc _VZLinuxRosettaAbstractSocketCachingOptionsClass) Alloc() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZLinuxRosettaAbstractSocketCachingOptionsClass) New() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZLinuxRosettaAbstractSocketCachingOptions) Init() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZLinuxRosettaAbstractSocketCachingOptions) Autorelease() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaAbstractSocketCachingOptions creates a new VZLinuxRosettaAbstractSocketCachingOptions instance.
func NewVZLinuxRosettaAbstractSocketCachingOptions() VZLinuxRosettaAbstractSocketCachingOptions {
	return getVZLinuxRosettaAbstractSocketCachingOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZLinuxRosettaAbstractSocketCachingOptions */
// Caching options for an abstract socket.
//
// Use this object to configure Rosetta to communicate with the Rosetta daemon using an abstract socket.


// Caching options for an abstract socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions
type VZLinuxRosettaAbstractSocketCachingOptions struct {
	VZLinuxRosettaCachingOptions
}

// VZLinuxRosettaAbstractSocketCachingOptionsFrom constructs a [VZLinuxRosettaAbstractSocketCachingOptions] from an unsafe.Pointer.
//
// Caching options for an abstract socket.
func VZLinuxRosettaAbstractSocketCachingOptionsFrom(ptr unsafe.Pointer) VZLinuxRosettaAbstractSocketCachingOptions {
	return VZLinuxRosettaAbstractSocketCachingOptions{
		VZLinuxRosettaCachingOptions: VZLinuxRosettaCachingOptionsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZLinuxRosettaAbstractSocketCachingOptions */

// Initialize options to set on a Rosetta directory share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions/initWithName:error:
func NewVZLinuxRosettaAbstractSocketCachingOptionsWithNameError(name objc.IObject /* cross-framework: NSString */, error_ objectivec.IObject) VZLinuxRosettaAbstractSocketCachingOptions {
	instance := getVZLinuxRosettaAbstractSocketCachingOptionsClass().Alloc()
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](instance.ID, objc.Sel("initWithName:error:"), name, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZLinuxRosettaAbstractSocketCachingOptionsWithNameError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZLinuxRosettaAbstractSocketCachingOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZLinuxRosettaAbstractSocketCachingOptions */

// The maximum length of name that the framework allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions/maximumNameLength
func (vc _VZLinuxRosettaAbstractSocketCachingOptionsClass) MaximumNameLength() uint {
	rv := objc.Send[uint](objc.ID(vc.class), objc.Sel("maximumNameLength"))
	return rv
}/* debug [class_properties_class/property]: maximumNameLength */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZLinuxRosettaAbstractSocketCachingOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZLinuxRosettaAbstractSocketCachingOptions */

// The maximum length of name that the framework allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions/maximumNameLength
func (v_ VZLinuxRosettaAbstractSocketCachingOptions) MaximumNameLength() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("maximumNameLength"))
	return rv
}/* debug [instance_properties/getter]: maximumNameLength */


// The name of the abstract socket that Rosetta uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions/name
func (v_ VZLinuxRosettaAbstractSocketCachingOptions) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZLinuxRosettaAbstractSocketCachingOptions */


