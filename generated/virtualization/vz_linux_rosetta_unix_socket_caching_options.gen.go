// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZLinuxRosettaUnixSocketCachingOptions */


/* debug [class_header]: Header for VZLinuxRosettaUnixSocketCachingOptions */
// The class instance for the [VZLinuxRosettaUnixSocketCachingOptions] class.
var (
	VZLinuxRosettaUnixSocketCachingOptionsClass     _VZLinuxRosettaUnixSocketCachingOptionsClass
	VZLinuxRosettaUnixSocketCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaUnixSocketCachingOptionsClass() _VZLinuxRosettaUnixSocketCachingOptionsClass {
	VZLinuxRosettaUnixSocketCachingOptionsClassOnce.Do(func() {
		VZLinuxRosettaUnixSocketCachingOptionsClass = _VZLinuxRosettaUnixSocketCachingOptionsClass{objc.GetClass("VZLinuxRosettaUnixSocketCachingOptions")}
	})
	return VZLinuxRosettaUnixSocketCachingOptionsClass
}

type _VZLinuxRosettaUnixSocketCachingOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZLinuxRosettaUnixSocketCachingOptions */
// An interface definition for the [VZLinuxRosettaUnixSocketCachingOptions] class.
type IVZLinuxRosettaUnixSocketCachingOptions interface {
	IVZLinuxRosettaCachingOptions
	
/* debug [class_interface_properties]: Properties for VZLinuxRosettaUnixSocketCachingOptions */
	// properties:
	Path() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZLinuxRosettaUnixSocketCachingOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZLinuxRosettaUnixSocketCachingOptions */
// Alloc allocates a new instance without initialization.
func (vc _VZLinuxRosettaUnixSocketCachingOptionsClass) Alloc() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZLinuxRosettaUnixSocketCachingOptionsClass) New() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZLinuxRosettaUnixSocketCachingOptions) Init() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZLinuxRosettaUnixSocketCachingOptions) Autorelease() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaUnixSocketCachingOptions creates a new VZLinuxRosettaUnixSocketCachingOptions instance.
func NewVZLinuxRosettaUnixSocketCachingOptions() VZLinuxRosettaUnixSocketCachingOptions {
	return getVZLinuxRosettaUnixSocketCachingOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZLinuxRosettaUnixSocketCachingOptions */
// An object that represents caching options for a UNIX domain socket.
//
// This object configures Rosetta to communicate with the Rosetta daemon using a UNIX domain socket.


// An object that represents caching options for a UNIX domain socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions
type VZLinuxRosettaUnixSocketCachingOptions struct {
	VZLinuxRosettaCachingOptions
}

// VZLinuxRosettaUnixSocketCachingOptionsFrom constructs a [VZLinuxRosettaUnixSocketCachingOptions] from an unsafe.Pointer.
//
// An object that represents caching options for a UNIX domain socket.
func VZLinuxRosettaUnixSocketCachingOptionsFrom(ptr unsafe.Pointer) VZLinuxRosettaUnixSocketCachingOptions {
	return VZLinuxRosettaUnixSocketCachingOptions{
		VZLinuxRosettaCachingOptions: VZLinuxRosettaCachingOptionsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZLinuxRosettaUnixSocketCachingOptions */

// Creates a new Rosetta caching options object for a UNIX domain socket with the path you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/initWithPath:error:
func NewVZLinuxRosettaUnixSocketCachingOptionsWithPathError(path objc.IObject /* cross-framework: NSString */, error_ objectivec.IObject) VZLinuxRosettaUnixSocketCachingOptions {
	instance := getVZLinuxRosettaUnixSocketCachingOptionsClass().Alloc()
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](instance.ID, objc.Sel("initWithPath:error:"), path, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZLinuxRosettaUnixSocketCachingOptionsWithPathError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZLinuxRosettaUnixSocketCachingOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZLinuxRosettaUnixSocketCachingOptions */

// The maximum allowed length of the path to the UNIX domain socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/maximumPathLength
func (vc _VZLinuxRosettaUnixSocketCachingOptionsClass) MaximumPathLength() uint {
	rv := objc.Send[uint](objc.ID(vc.class), objc.Sel("maximumPathLength"))
	return rv
}/* debug [class_properties_class/property]: maximumPathLength */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZLinuxRosettaUnixSocketCachingOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZLinuxRosettaUnixSocketCachingOptions */

// The maximum allowed length of the path to the UNIX domain socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/maximumPathLength
func (v_ VZLinuxRosettaUnixSocketCachingOptions) MaximumPathLength() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("maximumPathLength"))
	return rv
}/* debug [instance_properties/getter]: maximumPathLength */


// The path to the UNIX domain socket that Rosetta uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/path
func (v_ VZLinuxRosettaUnixSocketCachingOptions) Path() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("path"))
	return rv
}/* debug [instance_properties/getter]: path */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZLinuxRosettaUnixSocketCachingOptions */


