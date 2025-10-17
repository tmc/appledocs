// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Unit] class.
var UnitClass = _UnitClass{objc.GetClass("NSUnit")}

type _UnitClass struct {
	class objc.Class
}

type Unit struct {
	objc.ID
}

func UnitFrom(ptr unsafe.Pointer) Unit {
	return Unit{
		ID: objc.ID(ptr),
	}
}




