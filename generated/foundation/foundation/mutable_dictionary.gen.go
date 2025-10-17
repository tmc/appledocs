// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableDictionary] class.
var MutableDictionaryClass objc.Class

func init() {
	MutableDictionaryClass = objc.GetClass("NSMutableDictionary")
}

type MutableDictionary struct {
	objc.ID
}

func MutableDictionaryFrom(ptr unsafe.Pointer) MutableDictionary {
	return MutableDictionary{
		ID: objc.ID(ptr),
	}
}




