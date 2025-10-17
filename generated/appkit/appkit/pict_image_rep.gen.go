// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PICTImageRep] class.
var PICTImageRepClass objc.Class

func init() {
	PICTImageRepClass = objc.GetClass("NSPICTImageRep")
}

type PICTImageRep struct {
	objc.ID
}

func PICTImageRepFrom(ptr unsafe.Pointer) PICTImageRep {
	return PICTImageRep{
		ID: objc.ID(ptr),
	}
}




