// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLHandle] class.
var (
	URLHandleClass     _URLHandleClass
	URLHandleClassOnce sync.Once
)

func getURLHandleClass() _URLHandleClass {
	URLHandleClassOnce.Do(func() {
		URLHandleClass = _URLHandleClass{objc.GetClass("NSURLHandle")}
	})
	return URLHandleClass
}

type _URLHandleClass struct {
	class objc.Class
}

// An interface definition for the [URLHandle] class.
type IURLHandle interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An object that accesses and manages resource data indicated by a URL.
//
// A single can service multiple equivalent objects, but only if these URLs map to the same resource.


// An object that accesses and manages resource data indicated by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLHandle
type URLHandle struct {
	objectivec.Object
}

// URLHandleFrom constructs a [URLHandle] from an unsafe.Pointer.
//
// An object that accesses and manages resource data indicated by a URL.
func URLHandleFrom(ptr unsafe.Pointer) URLHandle {
	return URLHandle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLHandleClass) Alloc() URLHandle {
	rv := objc.Send[URLHandle](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLHandleClass) New() URLHandle {
	rv := objc.Send[URLHandle](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLHandle) Init() URLHandle {
	rv := objc.Send[URLHandle](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLHandle) Autorelease() URLHandle {
	rv := objc.Send[URLHandle](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLHandle creates a new URLHandle instance.
func NewURLHandle() URLHandle {
	return getURLHandleClass().New()
}




