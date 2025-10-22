// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [VZLinuxRosettaDirectoryShare] class.
type IVZLinuxRosettaDirectoryShare interface {
	IVZDirectoryShare
	Options() VZLinuxRosettaCachingOptions
	SetOptions(value VZLinuxRosettaCachingOptions)
	CachingOptions() unsafe.Pointer
	SetCachingOptions(value unsafe.Pointer)
}

// The Linux directory share for Rosetta.
//
// This directory share exposes the Rosetta directory from the host file system to the guest. The example below shows the process of creating a , and then associating the Rosetta directory share with the VM configuration. For complete instructions on installing Rosetta see , which includes additional information about checking for Rosetta availability, mounting the directory share, and registering the Rosetta runtime binary to run Intel binaries in a guest VM. For information on using a custom kernel to enhance Rosetta performance, see .
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZLinuxRosettaDirectoryShareClass) Alloc() VZLinuxRosettaDirectoryShare {
	rv := objc.Send[VZLinuxRosettaDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a new Rosetta directory share, or returns an error if Rosetta isn’t installed.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/init()
func NewVZLinuxRosettaDirectoryShareWithError(error_ unsafe.Pointer) VZLinuxRosettaDirectoryShare {
	instance := getVZLinuxRosettaDirectoryShareClass().Alloc()
	rv := objc.Send[VZLinuxRosettaDirectoryShare](instance.ID, objc.Sel("initWithError:"), error_)
	rv.Autorelease()
	return rv
}


// Starts the installation of Rosetta.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/installRosetta(completionHandler:)
func (vc _VZLinuxRosettaDirectoryShareClass) InstallRosettaWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("installRosettaWithCompletionHandler:"), completionHandler)
}

// A value that indicates the current state of Rosetta’s availability.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/availability
func (vc _VZLinuxRosettaDirectoryShareClass) Availability() VZLinuxRosettaAvailability {
	rv := objc.Send[VZLinuxRosettaAvailability](objc.ID(vc.class), objc.Sel("availability"))
	return rv
}
// A value that indicates the current state of Rosetta’s availability.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/availability
func (v_ VZLinuxRosettaDirectoryShare) Availability() VZLinuxRosettaAvailability {
	rv := objc.Send[VZLinuxRosettaAvailability](v_.ID, objc.Sel("availability"))
	return rv
}

// The value that enables translation caching and configures the socket communication type for Rosetta.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/options
func (v_ VZLinuxRosettaDirectoryShare) Options() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](v_.ID, objc.Sel("options"))
	return rv
}


// SetOptions sets the value of the options property.
// The value that enables translation caching and configures the socket communication type for Rosetta.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/options
func (v_ VZLinuxRosettaDirectoryShare) SetOptions(value VZLinuxRosettaCachingOptions) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setOptions:"), value)
}

// The value that enables translation caching and configures the socket communication type for Rosetta.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzlinuxrosettadirectoryshare/cachingoptions-swift.property
func (v_ VZLinuxRosettaDirectoryShare) CachingOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("cachingOptions"))
	return rv
}


// SetCachingOptions sets the value of the cachingOptions property.
// The value that enables translation caching and configures the socket communication type for Rosetta.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzlinuxrosettadirectoryshare/cachingoptions-swift.property
func (v_ VZLinuxRosettaDirectoryShare) SetCachingOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCachingOptions:"), value)
}


