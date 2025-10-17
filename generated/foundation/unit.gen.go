// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Unit] class.
var UnitClass objc.Class

func init() {
	UnitClass = objc.GetClass("NSUnit")
}

type Unit struct {
	objc.ID
}

func UnitFrom(ptr unsafe.Pointer) Unit {
	return Unit{
		ID: objc.ID(ptr),
	}
}



