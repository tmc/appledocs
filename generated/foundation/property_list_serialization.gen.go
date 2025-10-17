// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PropertyListSerialization] class.
var PropertyListSerializationClass = _PropertyListSerializationClass{objc.GetClass("NSPropertyListSerialization")}

type _PropertyListSerializationClass struct {
	class objc.Class
}

type PropertyListSerialization struct {
	objc.ID
}

func PropertyListSerializationFrom(ptr unsafe.Pointer) PropertyListSerialization {
	return PropertyListSerialization{
		ID: objc.ID(ptr),
	}
}




