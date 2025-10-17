// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLProtectionSpace] class.
var URLProtectionSpaceClass objc.Class

func init() {
	URLProtectionSpaceClass = objc.GetClass("NSURLProtectionSpace")
}

type URLProtectionSpace struct {
	objc.ID
}

func URLProtectionSpaceFrom(ptr unsafe.Pointer) URLProtectionSpace {
	return URLProtectionSpace{
		ID: objc.ID(ptr),
	}
}




