// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CustomImageRep] class.
var CustomImageRepClass objc.Class

func init() {
	CustomImageRepClass = objc.GetClass("NSCustomImageRep")
}

type CustomImageRep struct {
	objc.ID
}

func CustomImageRepFrom(ptr unsafe.Pointer) CustomImageRep {
	return CustomImageRep{
		ID: objc.ID(ptr),
	}
}




