// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var CacheClass _CacheClass

func init() {
	CacheClass = _CacheClass{objc.GetClass("NSCache")}
}

type _CacheClass struct {
	class objc.Class
}

type Cache struct {
	objc.ID
}

func CacheFrom(ptr unsafe.Pointer) Cache {
	return Cache{
		ID: objc.ID(ptr),
	}
}




