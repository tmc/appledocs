// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZSharedDirectory */

/* debug [class_header]: Header for VZSharedDirectory */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZSharedDirectory */
// An interface definition for the [VZSharedDirectory] class.
type IVZSharedDirectory interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZSharedDirectory */
	// properties:
	ReadOnly() bool
	URL() objc.IObject /* cross-framework: NSURL */
	IsReadOnly() bool
	SetIsReadOnly(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZSharedDirectory */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZSharedDirectory */
// Alloc allocates a new instance without initialization.
func (vc _VZSharedDirectoryClass) Alloc() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZSharedDirectory */
// A directory on the host that you can expose to a guest.
//
// This exposes a directory from the host file system to the guest.

// A directory on the host that you can expose to a guest.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZSharedDirectory */

// Initialize with a host directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory/init(url:readOnly:)
func NewVZSharedDirectoryWithURLReadOnly(url objc.IObject /* cross-framework: NSURL */, readOnly bool) VZSharedDirectory {
	instance := getVZSharedDirectoryClass().Alloc()
	rv := objc.Send[VZSharedDirectory](instance.ID, objc.Sel("initWithURL:readOnly:"), url, readOnly)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZSharedDirectoryWithURLReadOnly */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZSharedDirectory */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZSharedDirectory */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZSharedDirectory */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZSharedDirectory */

// A Boolean value that indicates whether the directory is read-only to the guest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory/isReadOnly
func (v_ VZSharedDirectory) ReadOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("readOnly"))
	return rv
} /* debug [instance_properties/getter]: readOnly */

// A file URL to a directory on the host system to expose to the guest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory/url
func (v_ VZSharedDirectory) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](v_.ID, objc.Sel("URL"))
	return rv
} /* debug [instance_properties/getter]: URL */

// A Boolean value that indicates whether the directory is read-only to the guest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzshareddirectory/isreadonly
func (v_ VZSharedDirectory) IsReadOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isReadOnly"))
	return rv
} /* debug [instance_properties/getter]: isReadOnly */

// A Boolean value that indicates whether the directory is read-only to the guest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzshareddirectory/isreadonly
func (v_ VZSharedDirectory) SetIsReadOnly(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsReadOnly:"), value)
} /* debug [instance_properties/setter]: isReadOnly */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZSharedDirectory */
