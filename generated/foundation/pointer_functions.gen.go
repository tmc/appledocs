// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PointerFunctions] class.
var PointerFunctionsClass = _PointerFunctionsClass{objc.GetClass("NSPointerFunctions")}

type _PointerFunctionsClass struct {
	class objc.Class
}

type PointerFunctions struct {
	objc.ID
}

func PointerFunctionsFrom(ptr unsafe.Pointer) PointerFunctions {
	return PointerFunctions{
		ID: objc.ID(ptr),
	}
}




