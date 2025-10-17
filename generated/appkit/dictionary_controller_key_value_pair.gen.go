// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DictionaryControllerKeyValuePair] class.
var dictionaryControllerKeyValuePairClass = _DictionaryControllerKeyValuePairClass{objc.GetClass("NSDictionaryControllerKeyValuePair")}

type _DictionaryControllerKeyValuePairClass struct {
	class objc.Class
}

// An interface definition for the [DictionaryControllerKeyValuePair] class.
type IDictionaryControllerKeyValuePair interface {
	objectivec.IObject
}

// A set of methods implemented by arranged objects to give access to information about those objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair

type DictionaryControllerKeyValuePair struct {
	objectivec.Object
}

// DictionaryControllerKeyValuePairFrom constructs a [DictionaryControllerKeyValuePair] from an unsafe.Pointer.
//
// A set of methods implemented by arranged objects to give access to information about those objects.
func DictionaryControllerKeyValuePairFrom(ptr unsafe.Pointer) DictionaryControllerKeyValuePair {
	return DictionaryControllerKeyValuePair{objectivec.Object{objc.ID(ptr)}}
}



