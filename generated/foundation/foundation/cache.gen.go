// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Cache] class.
var CacheClass objc.Class

func init() {
	CacheClass = objc.GetClass("NSCache")
}

type Cache struct {
	objc.ID
}

func CacheFrom(ptr unsafe.Pointer) Cache {
	return Cache{
		ID: objc.ID(ptr),
	}
}




