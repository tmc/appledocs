// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZLinuxRosettaDirectoryShare */

/* debug [class_header]: Header for VZLinuxRosettaDirectoryShare */
// The class instance for the [VZLinuxRosettaDirectoryShare] class.
var (
	VZLinuxRosettaDirectoryShareClass     _VZLinuxRosettaDirectoryShareClass
	VZLinuxRosettaDirectoryShareClassOnce sync.Once
)

func getVZLinuxRosettaDirectoryShareClass() _VZLinuxRosettaDirectoryShareClass {
	VZLinuxRosettaDirectoryShareClassOnce.Do(func() {
		VZLinuxRosettaDirectoryShareClass = _VZLinuxRosettaDirectoryShareClass{objc.GetClass("VZLinuxRosettaDirectoryShare")}
	})
	return VZLinuxRosettaDirectoryShareClass
}

type _VZLinuxRosettaDirectoryShareClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZLinuxRosettaDirectoryShare */
// An interface definition for the [VZLinuxRosettaDirectoryShare] class.
type IVZLinuxRosettaDirectoryShare interface {
	IVZDirectoryShare

	/* debug [class_interface_properties]: Properties for VZLinuxRosettaDirectoryShare */
	// properties:
	Options() IVZLinuxRosettaCachingOptions
	SetOptions(value IVZLinuxRosettaCachingOptions)
	CachingOptions() unsafe.Pointer
	SetCachingOptions(value unsafe.Pointer)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZLinuxRosettaDirectoryShare */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZLinuxRosettaDirectoryShare */
// Alloc allocates a new instance without initialization.
func (vc _VZLinuxRosettaDirectoryShareClass) Alloc() VZLinuxRosettaDirectoryShare {
	rv := objc.Send[VZLinuxRosettaDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZLinuxRosettaDirectoryShareClass) New() VZLinuxRosettaDirectoryShare {
	rv := objc.Send[VZLinuxRosettaDirectoryShare](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZLinuxRosettaDirectoryShare) Init() VZLinuxRosettaDirectoryShare {
	rv := objc.Send[VZLinuxRosettaDirectoryShare](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZLinuxRosettaDirectoryShare) Autorelease() VZLinuxRosettaDirectoryShare {
	rv := objc.Send[VZLinuxRosettaDirectoryShare](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaDirectoryShare creates a new VZLinuxRosettaDirectoryShare instance.
func NewVZLinuxRosettaDirectoryShare() VZLinuxRosettaDirectoryShare {
	return getVZLinuxRosettaDirectoryShareClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZLinuxRosettaDirectoryShare */
// The Linux directory share for Rosetta.
//
// This directory share exposes the Rosetta directory from the host file system to the guest. The example below shows the process of creating a , and then associating the Rosetta directory share with the VM configuration. For complete instructions on installing Rosetta see , which includes additional information about checking for Rosetta availability, mounting the directory share, and registering the Rosetta runtime binary to run Intel binaries in a guest VM. For information on using a custom kernel to enhance Rosetta performance, see .

// The Linux directory share for Rosetta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare
type VZLinuxRosettaDirectoryShare struct {
	VZDirectoryShare
}

// VZLinuxRosettaDirectoryShareFrom constructs a [VZLinuxRosettaDirectoryShare] from an unsafe.Pointer.
//
// The Linux directory share for Rosetta.
func VZLinuxRosettaDirectoryShareFrom(ptr unsafe.Pointer) VZLinuxRosettaDirectoryShare {
	return VZLinuxRosettaDirectoryShare{
		VZDirectoryShare: VZDirectoryShareFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZLinuxRosettaDirectoryShare */

// Creates a new Rosetta directory share, or returns an error if Rosetta isn’t installed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/init()
func NewVZLinuxRosettaDirectoryShareWithError(error_ unsafe.Pointer) VZLinuxRosettaDirectoryShare {
	instance := getVZLinuxRosettaDirectoryShareClass().Alloc()
	rv := objc.Send[VZLinuxRosettaDirectoryShare](instance.ID, objc.Sel("initWithError:"), error_)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZLinuxRosettaDirectoryShareWithError */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZLinuxRosettaDirectoryShare */

// Starts the installation of Rosetta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/installRosetta(completionHandler:)
func (vc _VZLinuxRosettaDirectoryShareClass) InstallRosettaWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("installRosettaWithCompletionHandler:"), completionHandler)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=InstallRosettaWithCompletionHandler) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZLinuxRosettaDirectoryShare */

// A value that indicates the current state of Rosetta’s availability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/availability
func (vc _VZLinuxRosettaDirectoryShareClass) Availability() VZLinuxRosettaAvailability {
	rv := objc.Send[VZLinuxRosettaAvailability](objc.ID(vc.class), objc.Sel("availability"))
	return rv
} /* debug [class_properties_class/property]: availability */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZLinuxRosettaDirectoryShare */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZLinuxRosettaDirectoryShare */

// A value that indicates the current state of Rosetta’s availability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/availability
func (v_ VZLinuxRosettaDirectoryShare) Availability() VZLinuxRosettaAvailability {
	rv := objc.Send[VZLinuxRosettaAvailability](v_.ID, objc.Sel("availability"))
	return rv
} /* debug [instance_properties/getter]: availability */

// The value that enables translation caching and configures the socket communication type for Rosetta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/options
func (v_ VZLinuxRosettaDirectoryShare) Options() IVZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](v_.ID, objc.Sel("options"))
	return rv
} /* debug [instance_properties/getter]: options */

// The value that enables translation caching and configures the socket communication type for Rosetta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/options
func (v_ VZLinuxRosettaDirectoryShare) SetOptions(value IVZLinuxRosettaCachingOptions) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setOptions:"), value)
} /* debug [instance_properties/setter]: options */

// The value that enables translation caching and configures the socket communication type for Rosetta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzlinuxrosettadirectoryshare/cachingoptions-swift.property
func (v_ VZLinuxRosettaDirectoryShare) CachingOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("cachingOptions"))
	return rv
} /* debug [instance_properties/getter]: cachingOptions */

// The value that enables translation caching and configures the socket communication type for Rosetta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzlinuxrosettadirectoryshare/cachingoptions-swift.property
func (v_ VZLinuxRosettaDirectoryShare) SetCachingOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCachingOptions:"), value)
} /* debug [instance_properties/setter]: cachingOptions */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZLinuxRosettaDirectoryShare */
