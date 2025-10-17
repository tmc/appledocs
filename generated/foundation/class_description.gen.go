// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ClassDescription] class.
var ClassDescriptionClass = _ClassDescriptionClass{objc.GetClass("NSClassDescription")}

type _ClassDescriptionClass struct {
	class objc.Class
}

type ClassDescription struct {
	objc.ID
}

func ClassDescriptionFrom(ptr unsafe.Pointer) ClassDescription {
	return ClassDescription{
		ID: objc.ID(ptr),
	}
}




