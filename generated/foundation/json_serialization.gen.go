// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [JSONSerialization] class.
var JSONSerializationClass objc.Class

func init() {
	JSONSerializationClass = objc.GetClass("NSJSONSerialization")
}

type JSONSerialization struct {
	objc.ID
}

func JSONSerializationFrom(ptr unsafe.Pointer) JSONSerialization {
	return JSONSerialization{
		ID: objc.ID(ptr),
	}
}



