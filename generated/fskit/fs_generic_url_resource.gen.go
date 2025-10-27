// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [FSGenericURLResource] class.
var (
	FSGenericURLResourceClass     _FSGenericURLResourceClass
	FSGenericURLResourceClassOnce sync.Once
)

func getFSGenericURLResourceClass() _FSGenericURLResourceClass {
	FSGenericURLResourceClassOnce.Do(func() {
		FSGenericURLResourceClass = _FSGenericURLResourceClass{objc.GetClass("FSGenericURLResource")}
	})
	return FSGenericURLResourceClass
}

type _FSGenericURLResourceClass struct {
	class objc.Class
}





// An interface definition for the [FSGenericURLResource] class.
type IFSGenericURLResource interface {
	IFSResource
	

	// properties:
	Url() foundation.foundation.INSURL


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (fc _FSGenericURLResourceClass) Alloc() FSGenericURLResource {
	rv := objc.Send[FSGenericURLResource](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSGenericURLResourceClass) New() FSGenericURLResource {
	rv := objc.Send[FSGenericURLResource](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSGenericURLResource) Init() FSGenericURLResource {
	rv := objc.Send[FSGenericURLResource](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSGenericURLResource) Autorelease() FSGenericURLResource {
	rv := objc.Send[FSGenericURLResource](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSGenericURLResource creates a new FSGenericURLResource instance.
func NewFSGenericURLResource() FSGenericURLResource {
	return getFSGenericURLResourceClass().New()
}





// A resource that represents an abstract URL.
//
// An is a completely abstract resource. The only reference to its contents is a single URL, the contents of which are arbitrary. This URL might represent a PCI locator string like , or some sort of network address for a remote file system. FSKit leaves interpretation of the URL and its contents entirely up to your implementation. Use the key to provide an array of case-insensitive URL schemes that your implementation supports. The following example shows how a hypothetical implementation declares support for the and URL schemes:


// A resource that represents an abstract URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSGenericURLResource
type FSGenericURLResource struct {
	FSResource
}

// FSGenericURLResourceFrom constructs a [FSGenericURLResource] from an unsafe.Pointer.
//
// A resource that represents an abstract URL.
func FSGenericURLResourceFrom(ptr unsafe.Pointer) FSGenericURLResource {
	return FSGenericURLResource{
		FSResource: FSResourceFrom(ptr),
	}
}






// Creates a generic URL resource with the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSGenericURLResource/init(url:)
func NewFSGenericURLResourceWithURL(url foundation.foundation.INSURL) FSGenericURLResource {
	instance := getFSGenericURLResourceClass().Alloc()
	rv := objc.Send[FSGenericURLResource](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}






















// The URL represented by the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSGenericURLResource/url
func (f_ FSGenericURLResource) Url() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](f_.ID, objc.Sel("url"))
	return rv
}







