// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Cache] class.
var (
	cacheClass     _CacheClass
	cacheClassOnce sync.Once
)

func getCacheClass() _CacheClass {
	cacheClassOnce.Do(func() {
		cacheClass = _CacheClass{objc.GetClass("NSCache")}
	})
	return cacheClass
}

type _CacheClass struct {
	class objc.Class
}

// An interface definition for the [Cache] class.
type ICache interface {
	objectivec.IObject
}

// A mutable collection you use to temporarily store transient key-value pairs that are subject to eviction when resources are low. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache
type Cache struct {
	objectivec.Object
}

// CacheFrom constructs a [Cache] from an unsafe.Pointer.
//
// A mutable collection you use to temporarily store transient key-value pairs that are subject to eviction when resources are low.
func CacheFrom(ptr unsafe.Pointer) Cache {
	return Cache{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CacheClass) Alloc() Cache {
	rv := objc.Send[Cache](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CacheClass) New() Cache {
	rv := objc.Send[Cache](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Cache) Init() Cache {
	rv := objc.Send[Cache](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Cache) Autorelease() Cache {
	rv := objc.Send[Cache](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCache creates a new Cache instance.
func NewCache() Cache {
	return getCacheClass().New()
}




