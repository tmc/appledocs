// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Form] class.
var FormClass objc.Class

func init() {
	FormClass = objc.GetClass("NSForm")
}

type Form struct {
	objc.ID
}

func FormFrom(ptr unsafe.Pointer) Form {
	return Form{
		ID: objc.ID(ptr),
	}
}




