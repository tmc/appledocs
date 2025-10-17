// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PropertyListSerialization] class.
var PropertyListSerializationClass objc.Class

func init() {
	PropertyListSerializationClass = objc.GetClass("NSPropertyListSerialization")
}

type PropertyListSerialization struct {
	objc.ID
}

func PropertyListSerializationFrom(ptr unsafe.Pointer) PropertyListSerialization {
	return PropertyListSerialization{
		ID: objc.ID(ptr),
	}
}



