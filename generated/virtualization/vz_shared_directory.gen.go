// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZSharedDirectory] class.
var (
	VZSharedDirectoryClass     _VZSharedDirectoryClass
	VZSharedDirectoryClassOnce sync.Once
)

func getVZSharedDirectoryClass() _VZSharedDirectoryClass {
	VZSharedDirectoryClassOnce.Do(func() {
		VZSharedDirectoryClass = _VZSharedDirectoryClass{objc.GetClass("VZSharedDirectory")}
	})
	return VZSharedDirectoryClass
}

type _VZSharedDirectoryClass struct {
	class objc.Class
}

// An interface definition for the [VZSharedDirectory] class.
type IVZSharedDirectory interface {
	objectivec.IObject
}

// A directory on the host that you can expose to a guest.
//
// This exposes a directory from the host file system to the guest.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory
type VZSharedDirectory struct {
	objectivec.Object
}

// VZSharedDirectoryFrom constructs a [VZSharedDirectory] from an unsafe.Pointer.
//
// A directory on the host that you can expose to a guest.
func VZSharedDirectoryFrom(ptr unsafe.Pointer) VZSharedDirectory {
	return VZSharedDirectory{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZSharedDirectoryClass) Alloc() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZSharedDirectoryClass) New() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZSharedDirectory) Init() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZSharedDirectory) Autorelease() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSharedDirectory creates a new VZSharedDirectory instance.
func NewVZSharedDirectory() VZSharedDirectory {
	return getVZSharedDirectoryClass().New()
}




// Initialize with a host directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory/init(url:readOnly:)
func NewVZSharedDirectoryWithURLReadOnly(url foundation.URL, readOnly bool) VZSharedDirectory {
	instance := getVZSharedDirectoryClass().Alloc()
	rv := objc.Send[VZSharedDirectory](instance.ID, objc.Sel("initWithURL:readOnly:"), url, readOnly)
	rv.Autorelease()
	return rv
}


// A Boolean value that indicates whether the directory is read-only to the guest.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzshareddirectory/isreadonly
func (v_ VZSharedDirectory) IsReadOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isReadOnly"))
	return rv
}


// SetIsReadOnly sets the value of the isReadOnly property.
// A Boolean value that indicates whether the directory is read-only to the guest.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzshareddirectory/isreadonly
func (v_ VZSharedDirectory) SetIsReadOnly(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsReadOnly:"), value)
}

// A Boolean value that indicates whether the directory is read-only to the guest.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory/isReadOnly
func (v_ VZSharedDirectory) ReadOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("readOnly"))
	return rv
}

// A file URL to a directory on the host system to expose to the guest.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory/url
func (v_ VZSharedDirectory) URL() foundation.URL {
	rv := objc.Send[foundation.URL](v_.ID, objc.Sel("URL"))
	return rv
}


