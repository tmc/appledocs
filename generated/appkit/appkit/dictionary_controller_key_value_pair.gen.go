// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DictionaryControllerKeyValuePair] class.
var DictionaryControllerKeyValuePairClass objc.Class

func init() {
	DictionaryControllerKeyValuePairClass = objc.GetClass("NSDictionaryControllerKeyValuePair")
}

type DictionaryControllerKeyValuePair struct {
	objc.ID
}

func DictionaryControllerKeyValuePairFrom(ptr unsafe.Pointer) DictionaryControllerKeyValuePair {
	return DictionaryControllerKeyValuePair{
		ID: objc.ID(ptr),
	}
}




