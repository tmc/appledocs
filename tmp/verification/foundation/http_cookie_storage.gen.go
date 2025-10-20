// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var hTTPCookieStorageClass _HTTPCookieStorageClass

func init() {
	hTTPCookieStorageClass = _HTTPCookieStorageClass{objc.GetClass("NSHTTPCookieStorage")}
}

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




