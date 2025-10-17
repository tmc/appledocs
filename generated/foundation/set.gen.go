// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Set] class.
var SetClass objc.Class

func init() {
	SetClass = objc.GetClass("NSSet")
}

type Set struct {
	objc.ID
}

func SetFrom(ptr unsafe.Pointer) Set {
	return Set{
		ID: objc.ID(ptr),
	}
}



