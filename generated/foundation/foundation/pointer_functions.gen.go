// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PointerFunctions] class.
var PointerFunctionsClass objc.Class

func init() {
	PointerFunctionsClass = objc.GetClass("NSPointerFunctions")
}

type PointerFunctions struct {
	objc.ID
}

func PointerFunctionsFrom(ptr unsafe.Pointer) PointerFunctions {
	return PointerFunctions{
		ID: objc.ID(ptr),
	}
}




