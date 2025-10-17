// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HTTPCookie] class.
var HTTPCookieClass = _HTTPCookieClass{objc.GetClass("NSHTTPCookie")}

type _HTTPCookieClass struct {
	class objc.Class
}

type HTTPCookie struct {
	objc.ID
}

func HTTPCookieFrom(ptr unsafe.Pointer) HTTPCookie {
	return HTTPCookie{
		ID: objc.ID(ptr),
	}
}




