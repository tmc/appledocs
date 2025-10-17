// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PersonNameComponents] class.
var PersonNameComponentsClass objc.Class

func init() {
	PersonNameComponentsClass = objc.GetClass("NSPersonNameComponents")
}

type PersonNameComponents struct {
	objc.ID
}

func PersonNameComponentsFrom(ptr unsafe.Pointer) PersonNameComponents {
	return PersonNameComponents{
		ID: objc.ID(ptr),
	}
}



