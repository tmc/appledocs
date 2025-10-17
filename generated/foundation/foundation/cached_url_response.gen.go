// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CachedURLResponse] class.
var CachedURLResponseClass objc.Class

func init() {
	CachedURLResponseClass = objc.GetClass("NSCachedURLResponse")
}

type CachedURLResponse struct {
	objc.ID
}

func CachedURLResponseFrom(ptr unsafe.Pointer) CachedURLResponse {
	return CachedURLResponse{
		ID: objc.ID(ptr),
	}
}




