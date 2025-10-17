// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLCache] class.
var uRLCacheClass = _URLCacheClass{objc.GetClass("NSURLCache")}

type _URLCacheClass struct {
	class objc.Class
}

// An interface definition for the [URLCache] class.
type IURLCache interface {
	objectivec.IObject
}

// An object that maps URL requests to cached response objects. [Full Topic]
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



