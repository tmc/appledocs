// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// A resource representing a path
//
// Represents a file path (possibly security scoped URL).
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource
type FSPathURLResource struct {
	FSResource
}

// FSPathURLResourceFrom constructs a [FSPathURLResource] from an unsafe.Pointer.
//
// A resource representing a path
func FSPathURLResourceFrom(ptr unsafe.Pointer) FSPathURLResource {
	return FSPathURLResource{
		FSResource: FSResourceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FSPathURLResourceClass) Alloc() FSPathURLResource {
	rv := objc.Send[FSPathURLResource](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource/init(url:writable:)
func NewFSPathURLResourceWithURLWritable(URL foundation.IURL, writable bool) FSPathURLResource {
	instance := getFSPathURLResourceClass().Alloc()
	rv := objc.Send[FSPathURLResource](instance.ID, objc.Sel("initWithURL:writable:"), URL, writable)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource/isWritable
func (f_ FSPathURLResource) Writable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("writable"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource/url
func (f_ FSPathURLResource) Url() foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("url"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fspathurlresource/iswritable
func (f_ FSPathURLResource) IsWritable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isWritable"))
	return rv
}


// SetIsWritable sets the value of the isWritable property.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fspathurlresource/iswritable
func (f_ FSPathURLResource) SetIsWritable(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsWritable:"), value)
}


