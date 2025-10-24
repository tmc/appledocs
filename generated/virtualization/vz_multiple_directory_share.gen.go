// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZMultipleDirectoryShare */

/* debug [class_header]: Header for VZMultipleDirectoryShare */
// The class instance for the [VZMultipleDirectoryShare] class.
var (
	VZMultipleDirectoryShareClass     _VZMultipleDirectoryShareClass
	VZMultipleDirectoryShareClassOnce sync.Once
)

func getVZMultipleDirectoryShareClass() _VZMultipleDirectoryShareClass {
	VZMultipleDirectoryShareClassOnce.Do(func() {
		VZMultipleDirectoryShareClass = _VZMultipleDirectoryShareClass{objc.GetClass("VZMultipleDirectoryShare")}
	})
	return VZMultipleDirectoryShareClass
}

type _VZMultipleDirectoryShareClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZMultipleDirectoryShare */
// An interface definition for the [VZMultipleDirectoryShare] class.
type IVZMultipleDirectoryShare interface {
	IVZDirectoryShare

	/* debug [class_interface_properties]: Properties for VZMultipleDirectoryShare */
	// properties:
	Directories() foundation.IDictionary
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZMultipleDirectoryShare */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZMultipleDirectoryShare */
// Alloc allocates a new instance without initialization.
func (vc _VZMultipleDirectoryShareClass) Alloc() VZMultipleDirectoryShare {
	rv := objc.Send[VZMultipleDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMultipleDirectoryShareClass) New() VZMultipleDirectoryShare {
	rv := objc.Send[VZMultipleDirectoryShare](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMultipleDirectoryShare) Init() VZMultipleDirectoryShare {
	rv := objc.Send[VZMultipleDirectoryShare](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMultipleDirectoryShare) Autorelease() VZMultipleDirectoryShare {
	rv := objc.Send[VZMultipleDirectoryShare](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMultipleDirectoryShare creates a new VZMultipleDirectoryShare instance.
func NewVZMultipleDirectoryShare() VZMultipleDirectoryShare {
	return getVZMultipleDirectoryShareClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZMultipleDirectoryShare */
// An object that describes a directory share for multiple directories.
//
// This directory share exposes multiple directories from the host file system to the guest VM.

// An object that describes a directory share for multiple directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare
type VZMultipleDirectoryShare struct {
	VZDirectoryShare
}

// VZMultipleDirectoryShareFrom constructs a [VZMultipleDirectoryShare] from an unsafe.Pointer.
//
// An object that describes a directory share for multiple directories.
func VZMultipleDirectoryShareFrom(ptr unsafe.Pointer) VZMultipleDirectoryShare {
	return VZMultipleDirectoryShare{
		VZDirectoryShare: VZDirectoryShareFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZMultipleDirectoryShare */

// Creates the directory share with a set of directories on the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/init(directories:)
func NewVZMultipleDirectoryShareWithDirectories(directories foundation.IDictionary) VZMultipleDirectoryShare {
	instance := getVZMultipleDirectoryShareClass().Alloc()
	rv := objc.Send[VZMultipleDirectoryShare](instance.ID, objc.Sel("initWithDirectories:"), directories)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZMultipleDirectoryShareWithDirectories */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZMultipleDirectoryShare */

// Transforms a string to be a valid directory name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/canonicalizedName(from:)
func (vc _VZMultipleDirectoryShareClass) CanonicalizedNameFromName(name objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(vc.class), objc.Sel("canonicalizedNameFromName:"), name)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=CanonicalizedNameFromName) */

// Check if a name is a valid directory name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/validateName(_:)
func (vc _VZMultipleDirectoryShareClass) ValidateNameError(name objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("validateName:error:"), name, error_)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=ValidateNameError) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZMultipleDirectoryShare */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZMultipleDirectoryShare */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZMultipleDirectoryShare */

// The directories on the host to expose to the guest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/directories
func (v_ VZMultipleDirectoryShare) Directories() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](v_.ID, objc.Sel("directories"))
	return rv
} /* debug [instance_properties/getter]: directories */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZMultipleDirectoryShare */
