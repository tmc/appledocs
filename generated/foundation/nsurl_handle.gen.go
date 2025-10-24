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





// Alloc allocates a new instance without initialization.
func (uc _URLHandleClass) Alloc() URLHandle {
	rv := objc.Send[URLHandle](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Initializes a newly created URL handle with the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLHandle/initWithURL:cached:
func NewURLHandleWithURLCached(anURL IURL, willCache bool) URLHandle {
	instance := getURLHandleClass().Alloc()
	rv := objc.Send[URLHandle](instance.ID, objc.Sel("initWithURL:cached:"), anURL, willCache)
	rv.Autorelease()
	return rv
}







// Returns the class of the URL handle that will be used for a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLHandle/URLHandleClassForURL:
func (uc _URLHandleClass) URLHandleClassForURL(anURL IURL) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(uc.class), objc.Sel("URLHandleClassForURL:"), anURL)
	return rv
}


// Returns the URL handle from the cache that has serviced the specified URL or another identical URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLHandle/cachedHandleForURL:
func (uc _URLHandleClass) CachedHandleForURL(anURL IURL) IURLHandle {
	rv := objc.Send[URLHandle](objc.ID(uc.class), objc.Sel("cachedHandleForURL:"), anURL)
	return rv
}


// Returns whether a URL handle can be initialized with a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLHandle/canInitWithURL:
func (uc _URLHandleClass) CanInitWithURL(anURL IURL) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("canInitWithURL:"), anURL)
	return rv
}


// Registers a subclass of as an available subclass for handling URLs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLHandle/registerURLHandleClass:
func (uc _URLHandleClass) RegisterURLHandleClass(anURLHandleSubclass objc.Class) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("registerURLHandleClass:"), anURLHandleSubclass)
}






















