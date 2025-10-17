// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableData] class.
var MutableDataClass objc.Class

func init() {
	MutableDataClass = objc.GetClass("NSMutableData")
}

type MutableData struct {
	objc.ID
}

func MutableDataFrom(ptr unsafe.Pointer) MutableData {
	return MutableData{
		ID: objc.ID(ptr),
	}
}




