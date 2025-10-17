// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableCharacterSet] class.
var MutableCharacterSetClass = _MutableCharacterSetClass{objc.GetClass("NSMutableCharacterSet")}

type _MutableCharacterSetClass struct {
	class objc.Class
}

type MutableCharacterSet struct {
	objc.ID
}

func MutableCharacterSetFrom(ptr unsafe.Pointer) MutableCharacterSet {
	return MutableCharacterSet{
		ID: objc.ID(ptr),
	}
}




