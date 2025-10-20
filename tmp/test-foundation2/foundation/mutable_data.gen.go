// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var mutableDataClass _MutableDataClass

func init() {
	mutableDataClass = _MutableDataClass{objc.GetClass("NSMutableData")}
}

type _MutableDataClass struct {
	class objc.Class
}

type MutableData struct {
	objc.ID
}

func MutableDataFrom(ptr unsafe.Pointer) MutableData {
	return MutableData{
		ID: objc.ID(ptr),
	}
}




