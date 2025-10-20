// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var InvocationClass _InvocationClass

func init() {
	InvocationClass = _InvocationClass{objc.GetClass("NSInvocation")}
}

type _InvocationClass struct {
	class objc.Class
}

type Invocation struct {
	objc.ID
}

func InvocationFrom(ptr unsafe.Pointer) Invocation {
	return Invocation{
		ID: objc.ID(ptr),
	}
}




