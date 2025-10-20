// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var URLCacheClass _URLCacheClass

func init() {
	URLCacheClass = _URLCacheClass{objc.GetClass("NSURLCache")}
}

type _URLCacheClass struct {
	class objc.Class
}

type URLCache struct {
	objc.ID
}

func URLCacheFrom(ptr unsafe.Pointer) URLCache {
	return URLCache{
		ID: objc.ID(ptr),
	}
}




