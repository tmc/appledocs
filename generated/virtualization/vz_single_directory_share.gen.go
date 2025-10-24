// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZSingleDirectoryShare */


/* debug [class_header]: Header for VZSingleDirectoryShare */
// The class instance for the [VZSingleDirectoryShare] class.
var (
	VZSingleDirectoryShareClass     _VZSingleDirectoryShareClass
	VZSingleDirectoryShareClassOnce sync.Once
)

func getVZSingleDirectoryShareClass() _VZSingleDirectoryShareClass {
	VZSingleDirectoryShareClassOnce.Do(func() {
		VZSingleDirectoryShareClass = _VZSingleDirectoryShareClass{objc.GetClass("VZSingleDirectoryShare")}
	})
	return VZSingleDirectoryShareClass
}

type _VZSingleDirectoryShareClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZSingleDirectoryShare */
// An interface definition for the [VZSingleDirectoryShare] class.
type IVZSingleDirectoryShare interface {
	IVZDirectoryShare
	
/* debug [class_interface_properties]: Properties for VZSingleDirectoryShare */
	// properties:
	Directory() IVZSharedDirectory
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZSingleDirectoryShare */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZSingleDirectoryShare */
// Alloc allocates a new instance without initialization.
func (vc _VZSingleDirectoryShareClass) Alloc() VZSingleDirectoryShare {
	rv := objc.Send[VZSingleDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZSingleDirectoryShareClass) New() VZSingleDirectoryShare {
	rv := objc.Send[VZSingleDirectoryShare](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZSingleDirectoryShare) Init() VZSingleDirectoryShare {
	rv := objc.Send[VZSingleDirectoryShare](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZSingleDirectoryShare) Autorelease() VZSingleDirectoryShare {
	rv := objc.Send[VZSingleDirectoryShare](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSingleDirectoryShare creates a new VZSingleDirectoryShare instance.
func NewVZSingleDirectoryShare() VZSingleDirectoryShare {
	return getVZSingleDirectoryShareClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZSingleDirectoryShare */
// An object that defines the directory share for a single directory.
//
// This directory share exposes a single directory from the host file system to the guest.


// An object that defines the directory share for a single directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSingleDirectoryShare
type VZSingleDirectoryShare struct {
	VZDirectoryShare
}

// VZSingleDirectoryShareFrom constructs a [VZSingleDirectoryShare] from an unsafe.Pointer.
//
// An object that defines the directory share for a single directory.
func VZSingleDirectoryShareFrom(ptr unsafe.Pointer) VZSingleDirectoryShare {
	return VZSingleDirectoryShare{
		VZDirectoryShare: VZDirectoryShareFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZSingleDirectoryShare */

// Creates a directory share with a directory that you specify on the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSingleDirectoryShare/init(directory:)
func NewVZSingleDirectoryShareWithDirectory(directory IVZSharedDirectory) VZSingleDirectoryShare {
	instance := getVZSingleDirectoryShareClass().Alloc()
	rv := objc.Send[VZSingleDirectoryShare](instance.ID, objc.Sel("initWithDirectory:"), directory)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZSingleDirectoryShareWithDirectory */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZSingleDirectoryShare */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZSingleDirectoryShare */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZSingleDirectoryShare */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZSingleDirectoryShare */

// The directory on the host to share with the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSingleDirectoryShare/directory
func (v_ VZSingleDirectoryShare) Directory() IVZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](v_.ID, objc.Sel("directory"))
	return rv
}/* debug [instance_properties/getter]: directory */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZSingleDirectoryShare */


