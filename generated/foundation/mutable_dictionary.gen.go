// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableDictionary] class.
var MutableDictionaryClass = _MutableDictionaryClass{objc.GetClass("NSMutableDictionary")}

type _MutableDictionaryClass struct {
	class objc.Class
}

type MutableDictionary struct {
	objc.ID
}

func MutableDictionaryFrom(ptr unsafe.Pointer) MutableDictionary {
	return MutableDictionary{
		ID: objc.ID(ptr),
	}
}




