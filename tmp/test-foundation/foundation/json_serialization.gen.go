// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var JSONSerializationClass _JSONSerializationClass

func init() {
	JSONSerializationClass = _JSONSerializationClass{objc.GetClass("NSJSONSerialization")}
}

type _JSONSerializationClass struct {
	class objc.Class
}

type JSONSerialization struct {
	objc.ID
}

func JSONSerializationFrom(ptr unsafe.Pointer) JSONSerialization {
	return JSONSerialization{
		ID: objc.ID(ptr),
	}
}




