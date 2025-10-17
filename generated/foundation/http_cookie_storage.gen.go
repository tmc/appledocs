// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HTTPCookieStorage] class.
var HTTPCookieStorageClass = _HTTPCookieStorageClass{objc.GetClass("NSHTTPCookieStorage")}

type _HTTPCookieStorageClass struct {
	class objc.Class
}

type HTTPCookieStorage struct {
	objc.ID
}

func HTTPCookieStorageFrom(ptr unsafe.Pointer) HTTPCookieStorage {
	return HTTPCookieStorage{
		ID: objc.ID(ptr),
	}
}




