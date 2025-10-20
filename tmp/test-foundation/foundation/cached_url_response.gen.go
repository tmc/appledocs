// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var CachedURLResponseClass _CachedURLResponseClass

func init() {
	CachedURLResponseClass = _CachedURLResponseClass{objc.GetClass("NSCachedURLResponse")}
}

type _CachedURLResponseClass struct {
	class objc.Class
}

type CachedURLResponse struct {
	objc.ID
}

func CachedURLResponseFrom(ptr unsafe.Pointer) CachedURLResponse {
	return CachedURLResponse{
		ID: objc.ID(ptr),
	}
}




