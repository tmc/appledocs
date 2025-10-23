// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [VZSingleDirectoryShare] class.
type IVZSingleDirectoryShare interface {
	IVZDirectoryShare
	Directory() VZSharedDirectory
	SetDirectory(value VZSharedDirectory)
}

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

// Alloc allocates a new instance without initialization.
func (vc _VZSingleDirectoryShareClass) Alloc() VZSingleDirectoryShare {
	rv := objc.Send[VZSingleDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The directory on the host to share with the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzsingledirectoryshare/directory
func (v_ VZSingleDirectoryShare) Directory() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](v_.ID, objc.Sel("directory"))
	return rv
}


// The directory on the host to share with the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzsingledirectoryshare/directory
func (v_ VZSingleDirectoryShare) SetDirectory(value VZSharedDirectory) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDirectory:"), value)
}



