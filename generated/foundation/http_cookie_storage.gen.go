// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [HTTPCookieStorage] class.
var HTTPCookieStorageClass objc.Class

func init() {
	HTTPCookieStorageClass = objc.GetClass("NSHTTPCookieStorage")
}

type HTTPCookieStorage struct {
	objc.ID
}

func HTTPCookieStorageFrom(ptr unsafe.Pointer) HTTPCookieStorage {
	return HTTPCookieStorage{
		ID: objc.ID(ptr),
	}
}



