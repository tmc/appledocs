// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var uUIDClass _UUIDClass

func init() {
	uUIDClass = _UUIDClass{objc.GetClass("NSUUID")}
}

type _UUIDClass struct {
	class objc.Class
}

type UUID struct {
	objc.ID
}

func UUIDFrom(ptr unsafe.Pointer) UUID {
	return UUID{
		ID: objc.ID(ptr),
	}
}




