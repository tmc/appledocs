// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CachedImageRep] class.
var CachedImageRepClass objc.Class

func init() {
	CachedImageRepClass = objc.GetClass("NSCachedImageRep")
}

type CachedImageRep struct {
	objc.ID
}

func CachedImageRepFrom(ptr unsafe.Pointer) CachedImageRep {
	return CachedImageRep{
		ID: objc.ID(ptr),
	}
}



