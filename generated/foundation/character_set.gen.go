// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CharacterSet] class.
var CharacterSetClass = _CharacterSetClass{objc.GetClass("NSCharacterSet")}

type _CharacterSetClass struct {
	class objc.Class
}

type CharacterSet struct {
	objc.ID
}

func CharacterSetFrom(ptr unsafe.Pointer) CharacterSet {
	return CharacterSet{
		ID: objc.ID(ptr),
	}
}




