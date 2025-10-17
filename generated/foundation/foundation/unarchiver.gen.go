// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Unarchiver] class.
var UnarchiverClass objc.Class

func init() {
	UnarchiverClass = objc.GetClass("NSUnarchiver")
}

type Unarchiver struct {
	objc.ID
}

func UnarchiverFrom(ptr unsafe.Pointer) Unarchiver {
	return Unarchiver{
		ID: objc.ID(ptr),
	}
}




