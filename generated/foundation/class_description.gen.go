// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ClassDescription] class.
var ClassDescriptionClass objc.Class

func init() {
	ClassDescriptionClass = objc.GetClass("NSClassDescription")
}

type ClassDescription struct {
	objc.ID
}

func ClassDescriptionFrom(ptr unsafe.Pointer) ClassDescription {
	return ClassDescription{
		ID: objc.ID(ptr),
	}
}



