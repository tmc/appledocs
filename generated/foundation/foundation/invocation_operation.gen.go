// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [InvocationOperation] class.
var InvocationOperationClass objc.Class

func init() {
	InvocationOperationClass = objc.GetClass("NSInvocationOperation")
}

type InvocationOperation struct {
	objc.ID
}

func InvocationOperationFrom(ptr unsafe.Pointer) InvocationOperation {
	return InvocationOperation{
		ID: objc.ID(ptr),
	}
}




