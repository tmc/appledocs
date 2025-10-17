// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HTTPCookieStorage] class.
var hTTPCookieStorageClass = _HTTPCookieStorageClass{objc.GetClass("NSHTTPCookieStorage")}

type _HTTPCookieStorageClass struct {
	class objc.Class
}

// A container that manages the storage of cookies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage

type HTTPCookieStorage struct {
	objectivec.Object
}

// HTTPCookieStorageFrom constructs a [HTTPCookieStorage] from an unsafe.Pointer.
//
// A container that manages the storage of cookies.
func HTTPCookieStorageFrom(ptr unsafe.Pointer) HTTPCookieStorage {
	return HTTPCookieStorage{objectivec.Object{objc.ID(ptr)}}
}



