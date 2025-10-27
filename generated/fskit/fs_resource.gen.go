// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [FSResource] class.
type IFSResource interface {
	objectivec.IObject
	

	// properties:
	Revoked() bool
	IsRevoked() bool
	SetIsRevoked(value bool)


	

	// methods:
	MakeProxy() unsafe.Pointer
	Revoke()


}





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




















// Creates a proxy object of this resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSResource/makeProxy()
func (f_ FSResource) MakeProxy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("makeProxy"))
	return rv
}


// Revokes the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSResource/revoke()
func (f_ FSResource) Revoke() {
	objc.Send[objc.ID](f_.ID, objc.Sel("revoke"))
}







// A Boolean value that indicates whether the resource is revoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSResource/isRevoked
func (f_ FSResource) Revoked() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("revoked"))
	return rv
}


// A Boolean value that indicates whether the resource is revoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsresource/isrevoked
func (f_ FSResource) IsRevoked() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isRevoked"))
	return rv
}


// A Boolean value that indicates whether the resource is revoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsresource/isrevoked
func (f_ FSResource) SetIsRevoked(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsRevoked:"), value)
}








