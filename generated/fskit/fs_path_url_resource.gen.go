// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [FSPathURLResource] class.
var (
	FSPathURLResourceClass     _FSPathURLResourceClass
	FSPathURLResourceClassOnce sync.Once
)

func getFSPathURLResourceClass() _FSPathURLResourceClass {
	FSPathURLResourceClassOnce.Do(func() {
		FSPathURLResourceClass = _FSPathURLResourceClass{objc.GetClass("FSPathURLResource")}
	})
	return FSPathURLResourceClass
}

type _FSPathURLResourceClass struct {
	class objc.Class
}





// An interface definition for the [FSPathURLResource] class.
type IFSPathURLResource interface {
	IFSResource
	

	// properties:
	Writable() bool
	Url() foundation.foundation.INSURL
	IsWritable() bool
	SetIsWritable(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (fc _FSPathURLResourceClass) Alloc() FSPathURLResource {
	rv := objc.Send[FSPathURLResource](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSPathURLResourceClass) New() FSPathURLResource {
	rv := objc.Send[FSPathURLResource](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSPathURLResource) Init() FSPathURLResource {
	rv := objc.Send[FSPathURLResource](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSPathURLResource) Autorelease() FSPathURLResource {
	rv := objc.Send[FSPathURLResource](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSPathURLResource creates a new FSPathURLResource instance.
func NewFSPathURLResource() FSPathURLResource {
	return getFSPathURLResourceClass().New()
}





// A resource that represents a path in the system file space.
//
// The URL passed to may be a security-scoped URL. If the URL is a security-scoped URL, FSKit transports it intact from a client application to your extension.


// A resource that represents a path in the system file space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource
type FSPathURLResource struct {
	FSResource
}

// FSPathURLResourceFrom constructs a [FSPathURLResource] from an unsafe.Pointer.
//
// A resource that represents a path in the system file space.
func FSPathURLResourceFrom(ptr unsafe.Pointer) FSPathURLResource {
	return FSPathURLResource{
		FSResource: FSResourceFrom(ptr),
	}
}






// Creates a path URL resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource/init(url:writable:)
func NewFSPathURLResourceWithURLWritable(URL foundation.foundation.INSURL, writable bool) FSPathURLResource {
	instance := getFSPathURLResourceClass().Alloc()
	rv := objc.Send[FSPathURLResource](instance.ID, objc.Sel("initWithURL:writable:"), URL, writable)
	rv.Autorelease()
	return rv
}






















// A Boolean value that indicates whether the file system supports writing to the contents of the path URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource/isWritable
func (f_ FSPathURLResource) Writable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("writable"))
	return rv
}


// The URL represented by the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource/url
func (f_ FSPathURLResource) Url() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](f_.ID, objc.Sel("url"))
	return rv
}


// A Boolean value that indicates whether the file system supports writing to the contents of the path URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fspathurlresource/iswritable
func (f_ FSPathURLResource) IsWritable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isWritable"))
	return rv
}


// A Boolean value that indicates whether the file system supports writing to the contents of the path URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fspathurlresource/iswritable
func (f_ FSPathURLResource) SetIsWritable(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsWritable:"), value)
}







