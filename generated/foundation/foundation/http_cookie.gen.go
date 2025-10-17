// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [HTTPCookie] class.
var HTTPCookieClass objc.Class

func init() {
	HTTPCookieClass = objc.GetClass("NSHTTPCookie")
}

type HTTPCookie struct {
	objc.ID
}

func HTTPCookieFrom(ptr unsafe.Pointer) HTTPCookie {
	return HTTPCookie{
		ID: objc.ID(ptr),
	}
}




