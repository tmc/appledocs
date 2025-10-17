// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLCache] class.
var URLCacheClass objc.Class

func init() {
	URLCacheClass = objc.GetClass("NSURLCache")
}

type URLCache struct {
	objc.ID
}

func URLCacheFrom(ptr unsafe.Pointer) URLCache {
	return URLCache{
		ID: objc.ID(ptr),
	}
}



