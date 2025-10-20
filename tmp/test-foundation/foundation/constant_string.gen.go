// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ConstantStringClass _ConstantStringClass

func init() {
	ConstantStringClass = _ConstantStringClass{objc.GetClass("NSConstantString")}
}

type _ConstantStringClass struct {
	class objc.Class
}

type ConstantString struct {
	objc.ID
}

func ConstantStringFrom(ptr unsafe.Pointer) ConstantString {
	return ConstantString{
		ID: objc.ID(ptr),
	}
}




