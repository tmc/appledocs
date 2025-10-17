// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableCharacterSet] class.
var MutableCharacterSetClass objc.Class

func init() {
	MutableCharacterSetClass = objc.GetClass("NSMutableCharacterSet")
}

type MutableCharacterSet struct {
	objc.ID
}

func MutableCharacterSetFrom(ptr unsafe.Pointer) MutableCharacterSet {
	return MutableCharacterSet{
		ID: objc.ID(ptr),
	}
}



