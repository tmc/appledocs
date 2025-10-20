// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var MethodSignatureClass _MethodSignatureClass

func init() {
	MethodSignatureClass = _MethodSignatureClass{objc.GetClass("NSMethodSignature")}
}

type _MethodSignatureClass struct {
	class objc.Class
}

type MethodSignature struct {
	objc.ID
}

func MethodSignatureFrom(ptr unsafe.Pointer) MethodSignature {
	return MethodSignature{
		ID: objc.ID(ptr),
	}
}




