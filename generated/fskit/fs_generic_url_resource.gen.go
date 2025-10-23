// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	Url() foundation.URL
	SetUrl(value foundation.URL)
}

// A resource representing an abstract URL


// A resource representing an abstract URL
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSGenericURLResource
type FSGenericURLResource struct {
	FSResource
}

// FSGenericURLResourceFrom constructs a [FSGenericURLResource] from an unsafe.Pointer.
//
// A resource representing an abstract URL
func FSGenericURLResourceFrom(ptr unsafe.Pointer) FSGenericURLResource {
	return FSGenericURLResource{
		FSResource: FSResourceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FSGenericURLResourceClass) Alloc() FSGenericURLResource {
	rv := objc.Send[FSGenericURLResource](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSGenericURLResource/init(url:)
func NewFSGenericURLResourceWithURL(url foundation.URL) FSGenericURLResource {
	instance := getFSGenericURLResourceClass().Alloc()
	rv := objc.Send[FSGenericURLResource](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsgenericurlresource/url
func (f_ FSGenericURLResource) Url() foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("url"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsgenericurlresource/url
func (f_ FSGenericURLResource) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUrl:"), value)
}


