// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLCache] class.
var (
	URLCacheClass     _URLCacheClass
	URLCacheClassOnce sync.Once
)

func getURLCacheClass() _URLCacheClass {
	URLCacheClassOnce.Do(func() {
		URLCacheClass = _URLCacheClass{objc.GetClass("NSURLCache")}
	})
	return URLCacheClass
}

type _URLCacheClass struct {
	class objc.Class
}

// An interface definition for the [URLCache] class.
type IURLCache interface {
	objectivec.IObject
}

// An object that maps URL requests to cached response objects.
//
// The class implements the caching of responses to URL load requests, by mapping objects to objects. It provides a composite in-memory and on-disk cache, and lets you manipulate the sizes of both the in-memory and on-disk portions. You can also control the path where cache data is persistently stored.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache
type URLCache struct {
	objectivec.Object
}

// URLCacheFrom constructs a [URLCache] from an unsafe.Pointer.
//
// An object that maps URL requests to cached response objects.
func URLCacheFrom(ptr unsafe.Pointer) URLCache {
	return URLCache{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLCacheClass) Alloc() URLCache {
	rv := objc.Send[URLCache](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLCacheClass) New() URLCache {
	rv := objc.Send[URLCache](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLCache) Init() URLCache {
	rv := objc.Send[URLCache](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLCache) Autorelease() URLCache {
	rv := objc.Send[URLCache](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLCache creates a new URLCache instance.
func NewURLCache() URLCache {
	return getURLCacheClass().New()
}
