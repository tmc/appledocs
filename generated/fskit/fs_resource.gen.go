// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSResource */


/* debug [class_header]: Header for FSResource */
// The class instance for the [FSResource] class.
var (
	FSResourceClass     _FSResourceClass
	FSResourceClassOnce sync.Once
)

func getFSResourceClass() _FSResourceClass {
	FSResourceClassOnce.Do(func() {
		FSResourceClass = _FSResourceClass{objc.GetClass("FSResource")}
	})
	return FSResourceClass
}

type _FSResourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSResource */
// An interface definition for the [FSResource] class.
type IFSResource interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSResource */
	// properties:
	Revoked() bool
	IsRevoked() bool
	SetIsRevoked(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSResource */
	// methods:
	MakeProxy() unsafe.Pointer
	Revoke()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSResource */
// Alloc allocates a new instance without initialization.
func (fc _FSResourceClass) Alloc() FSResource {
	rv := objc.Send[FSResource](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSResourceClass) New() FSResource {
	rv := objc.Send[FSResource](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSResource) Init() FSResource {
	rv := objc.Send[FSResource](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSResource) Autorelease() FSResource {
	rv := objc.Send[FSResource](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSResource creates a new FSResource instance.
func NewFSResource() FSResource {
	return getFSResourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSResource */
// An abstract resource a file system uses to provide data for a volume.
//
// is a base class to represent the various possible sources of data for a file system. These range from dedicated storage devices like hard drives and flash storage to network connections, and beyond. Subclasses define behavior specific to a given kind of resource, such as for disk partition (IOMedia) file systems. These file systems are typical disk file systems such as HFS, APFS, ExFAT, ext2fs, or NTFS. A resource’s type also determines its life cycle. Resources based on block storage devices come into being when the system probes the media underlying the volumes and container. Other kinds of resources, like those based on URLs, might have different life cycles. For example, a resource based on a URL might iniitalize when a person uses the “Connect to server” command in the macOS Finder.


// An abstract resource a file system uses to provide data for a volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSResource
type FSResource struct {
	objectivec.Object
}

// FSResourceFrom constructs a [FSResource] from an unsafe.Pointer.
//
// An abstract resource a file system uses to provide data for a volume.
func FSResourceFrom(ptr unsafe.Pointer) FSResource {
	return FSResource{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSResource *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSResource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSResource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSResource */

// Creates a proxy object of this resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSResource/makeProxy()
func (f_ FSResource) MakeProxy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("makeProxy"))
	return rv
}/* debug [instance_methods/method]: MakeProxy */


// Revokes the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSResource/revoke()
func (f_ FSResource) Revoke() {
	objc.Send[objc.ID](f_.ID, objc.Sel("revoke"))
}/* debug [instance_methods/method]: Revoke */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSResource */

// A Boolean value that indicates whether the resource is revoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSResource/isRevoked
func (f_ FSResource) Revoked() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("revoked"))
	return rv
}/* debug [instance_properties/getter]: revoked */


// A Boolean value that indicates whether the resource is revoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsresource/isrevoked
func (f_ FSResource) IsRevoked() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isRevoked"))
	return rv
}/* debug [instance_properties/getter]: isRevoked */


// A Boolean value that indicates whether the resource is revoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsresource/isrevoked
func (f_ FSResource) SetIsRevoked(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsRevoked:"), value)
}/* debug [instance_properties/setter]: isRevoked */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSResource */



