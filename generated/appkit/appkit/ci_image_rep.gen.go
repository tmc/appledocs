// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CIImageRep] class.
var CIImageRepClass objc.Class

func init() {
	CIImageRepClass = objc.GetClass("NSCIImageRep")
}

type CIImageRep struct {
	objc.ID
}

func CIImageRepFrom(ptr unsafe.Pointer) CIImageRep {
	return CIImageRep{
		ID: objc.ID(ptr),
	}
}



