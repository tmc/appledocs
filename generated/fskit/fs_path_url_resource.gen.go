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
	// properties:
	Writable() bool /* primitive/slice/pointer. */
	IsWritable() bool /* primitive/slice/pointer. */
	SetIsWritable(value bool /* primitive/slice/pointer. */)
	Url() foundation.objc.IObject /* cross-framework: URL */
	SetUrl(value foundation.objc.IObject /* cross-framework: URL */)
	// methods:
}

// A resource representing a path
//
// Represents a file path (possibly security scoped URL).


// A resource representing a path
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSPathURLResource/isWritable
func (f_ FSPathURLResource) Writable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("writable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fspathurlresource/iswritable
func (f_ FSPathURLResource) IsWritable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isWritable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fspathurlresource/iswritable
func (f_ FSPathURLResource) SetIsWritable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsWritable:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fspathurlresource/url
func (f_ FSPathURLResource) Url() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("url"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fspathurlresource/url
func (f_ FSPathURLResource) SetUrl(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUrl:"), value)
}



