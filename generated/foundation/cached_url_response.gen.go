// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CachedURLResponse] class.
var cachedURLResponseClass = _CachedURLResponseClass{objc.GetClass("NSCachedURLResponse")}

type _CachedURLResponseClass struct {
	class objc.Class
}

// A cached response to a URL request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/CachedURLResponse

type CachedURLResponse struct {
	objectivec.Object
}

// CachedURLResponseFrom constructs a [CachedURLResponse] from an unsafe.Pointer.
//
// A cached response to a URL request.
func CachedURLResponseFrom(ptr unsafe.Pointer) CachedURLResponse {
	return CachedURLResponse{objectivec.Object{objc.ID(ptr)}}
}



