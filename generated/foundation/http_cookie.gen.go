// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HTTPCookie] class.
var hTTPCookieClass = _HTTPCookieClass{objc.GetClass("NSHTTPCookie")}

type _HTTPCookieClass struct {
	class objc.Class
}

// A representation of an HTTP cookie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie

type HTTPCookie struct {
	objectivec.Object
}

// HTTPCookieFrom constructs a [HTTPCookie] from an unsafe.Pointer.
//
// A representation of an HTTP cookie.
func HTTPCookieFrom(ptr unsafe.Pointer) HTTPCookie {
	return HTTPCookie{objectivec.Object{objc.ID(ptr)}}
}



