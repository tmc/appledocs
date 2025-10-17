// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Invocation] class.
var InvocationClass objc.Class

func init() {
	InvocationClass = objc.GetClass("NSInvocation")
}

type Invocation struct {
	objc.ID
}

func InvocationFrom(ptr unsafe.Pointer) Invocation {
	return Invocation{
		ID: objc.ID(ptr),
	}
}




