// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MethodSignature] class.
var MethodSignatureClass objc.Class

func init() {
	MethodSignatureClass = objc.GetClass("NSMethodSignature")
}

type MethodSignature struct {
	objc.ID
}

func MethodSignatureFrom(ptr unsafe.Pointer) MethodSignature {
	return MethodSignature{
		ID: objc.ID(ptr),
	}
}



