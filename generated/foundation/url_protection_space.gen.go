// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLProtectionSpace] class.
var URLProtectionSpaceClass = _URLProtectionSpaceClass{objc.GetClass("NSURLProtectionSpace")}

type _URLProtectionSpaceClass struct {
	class objc.Class
}

type URLProtectionSpace struct {
	objc.ID
}

func URLProtectionSpaceFrom(ptr unsafe.Pointer) URLProtectionSpace {
	return URLProtectionSpace{
		ID: objc.ID(ptr),
	}
}




