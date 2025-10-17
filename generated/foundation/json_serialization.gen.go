// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [JSONSerialization] class.
var jSONSerializationClass = _JSONSerializationClass{objc.GetClass("NSJSONSerialization")}

type _JSONSerializationClass struct {
	class objc.Class
}

// An object that converts between JSON and the equivalent Foundation objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization

type JSONSerialization struct {
	objectivec.Object
}

// JSONSerializationFrom constructs a [JSONSerialization] from an unsafe.Pointer.
//
// An object that converts between JSON and the equivalent Foundation objects.
func JSONSerializationFrom(ptr unsafe.Pointer) JSONSerialization {
	return JSONSerialization{objectivec.Object{objc.ID(ptr)}}
}



