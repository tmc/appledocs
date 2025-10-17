// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Cache] class.
var cacheClass = _CacheClass{objc.GetClass("NSCache")}

type _CacheClass struct {
	class objc.Class
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



