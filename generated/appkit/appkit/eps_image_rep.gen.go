// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [EPSImageRep] class.
var EPSImageRepClass objc.Class

func init() {
	EPSImageRepClass = objc.GetClass("NSEPSImageRep")
}

type EPSImageRep struct {
	objc.ID
}

func EPSImageRepFrom(ptr unsafe.Pointer) EPSImageRep {
	return EPSImageRep{
		ID: objc.ID(ptr),
	}
}




