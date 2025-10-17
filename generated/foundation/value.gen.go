// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Value] class.
var ValueClass = _ValueClass{objc.GetClass("NSValue")}

type _ValueClass struct {
	class objc.Class
}

type Value struct {
	objc.ID
}

func ValueFrom(ptr unsafe.Pointer) Value {
	return Value{
		ID: objc.ID(ptr),
	}
}




