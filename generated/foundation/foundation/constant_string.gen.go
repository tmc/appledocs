// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ConstantString] class.
var ConstantStringClass objc.Class

func init() {
	ConstantStringClass = objc.GetClass("NSConstantString")
}

type ConstantString struct {
	objc.ID
}

func ConstantStringFrom(ptr unsafe.Pointer) ConstantString {
	return ConstantString{
		ID: objc.ID(ptr),
	}
}




