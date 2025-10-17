// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CharacterSet] class.
var CharacterSetClass objc.Class

func init() {
	CharacterSetClass = objc.GetClass("NSCharacterSet")
}

type CharacterSet struct {
	objc.ID
}

func CharacterSetFrom(ptr unsafe.Pointer) CharacterSet {
	return CharacterSet{
		ID: objc.ID(ptr),
	}
}




