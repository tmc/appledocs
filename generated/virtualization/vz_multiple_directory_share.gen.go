// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [VZMultipleDirectoryShare] class.
type IVZMultipleDirectoryShare interface {
	IVZDirectoryShare
}

// An object that describes a directory share for multiple directories.
//
// This directory share exposes multiple directories from the host file system to the guest VM.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZMultipleDirectoryShareClass) Alloc() VZMultipleDirectoryShare {
	rv := objc.Send[VZMultipleDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates the directory share with a set of directories on the host.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/init(directories:)
func NewVZMultipleDirectoryShareWithDirectories(directories unsafe.Pointer) VZMultipleDirectoryShare {
	instance := getVZMultipleDirectoryShareClass().Alloc()
	rv := objc.Send[VZMultipleDirectoryShare](instance.ID, objc.Sel("initWithDirectories:"), directories)
	rv.Autorelease()
	return rv
}


// Transforms a string to be a valid directory name.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/canonicalizedName(from:)
func (vc _VZMultipleDirectoryShareClass) CanonicalizedNameFromName(name string) string {
	rv := objc.Send[string](objc.ID(vc.class), objc.Sel("canonicalizedNameFromName:"), objc.String(name))
	return rv
}

// Check if a name is a valid directory name.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/validateName(_:)
func (vc _VZMultipleDirectoryShareClass) ValidateNameError(name string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("validateName:error:"), objc.String(name), error_)
	return rv
}

// The directories on the host to expose to the guest.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/directories
func (v_ VZMultipleDirectoryShare) Directories() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("directories"))
	return rv
}


