// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Value] class.
var ValueClass objc.Class

func init() {
	ValueClass = objc.GetClass("NSValue")
}

type Value struct {
	objc.ID
}

func ValueFrom(ptr unsafe.Pointer) Value {
	return Value{
		ID: objc.ID(ptr),
	}
}




