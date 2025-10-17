// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UUID] class.
var UUIDClass objc.Class

func init() {
	UUIDClass = objc.GetClass("NSUUID")
}

type UUID struct {
	objc.ID
}

func UUIDFrom(ptr unsafe.Pointer) UUID {
	return UUID{
		ID: objc.ID(ptr),
	}
}



