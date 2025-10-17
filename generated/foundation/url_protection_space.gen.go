// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLProtectionSpace] class.
var uRLProtectionSpaceClass = _URLProtectionSpaceClass{objc.GetClass("NSURLProtectionSpace")}

type _URLProtectionSpaceClass struct {
	class objc.Class
}

// A server or an area on a server, commonly referred to as a realm, that requires authentication. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace

type URLProtectionSpace struct {
	objectivec.Object
}

// URLProtectionSpaceFrom constructs a [URLProtectionSpace] from an unsafe.Pointer.
//
// A server or an area on a server, commonly referred to as a realm, that requires authentication.
func URLProtectionSpaceFrom(ptr unsafe.Pointer) URLProtectionSpace {
	return URLProtectionSpace{objectivec.Object{objc.ID(ptr)}}
}



