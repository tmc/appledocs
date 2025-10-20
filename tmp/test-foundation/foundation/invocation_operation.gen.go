// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var InvocationOperationClass _InvocationOperationClass

func init() {
	InvocationOperationClass = _InvocationOperationClass{objc.GetClass("NSInvocationOperation")}
}

type _InvocationOperationClass struct {
	class objc.Class
}

type InvocationOperation struct {
	objc.ID
}

func InvocationOperationFrom(ptr unsafe.Pointer) InvocationOperation {
	return InvocationOperation{
		ID: objc.ID(ptr),
	}
}




