// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class FSPathURLResource */


/* debug [class_header]: Header for FSPathURLResource */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSPathURLResource */
// An interface definition for the [FSPathURLResource] class.
type IFSPathURLResource interface {
	IFSResource
	
/* debug [class_interface_properties]: Properties for FSPathURLResource */
	// properties:
	Writable() bool
	Url() objc.IObject /* cross-framework: NSURL */
	IsWritable() bool
	SetIsWritable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSPathURLResource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSPathURLResource */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSPathURLResource */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSPathURLResource */

// Creates a path URL resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource/init(url:writable:)
func NewFSPathURLResourceWithURLWritable(URL objc.IObject /* cross-framework: NSURL */, writable bool) FSPathURLResource {
	instance := getFSPathURLResourceClass().Alloc()
	rv := objc.Send[FSPathURLResource](instance.ID, objc.Sel("initWithURL:writable:"), URL, writable)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFSPathURLResourceWithURLWritable */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSPathURLResource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSPathURLResource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSPathURLResource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSPathURLResource */

// A Boolean value that indicates whether the file system supports writing to the contents of the path URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource/isWritable
func (f_ FSPathURLResource) Writable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("writable"))
	return rv
}/* debug [instance_properties/getter]: writable */


// The URL represented by the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource/url
func (f_ FSPathURLResource) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](f_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// A Boolean value that indicates whether the file system supports writing to the contents of the path URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fspathurlresource/iswritable
func (f_ FSPathURLResource) IsWritable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isWritable"))
	return rv
}/* debug [instance_properties/getter]: isWritable */


// A Boolean value that indicates whether the file system supports writing to the contents of the path URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fspathurlresource/iswritable
func (f_ FSPathURLResource) SetIsWritable(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsWritable:"), value)
}/* debug [instance_properties/setter]: isWritable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSPathURLResource */


