// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableDictionary] class.
var mutableDictionaryClass = _MutableDictionaryClass{objc.GetClass("NSMutableDictionary")}

type _MutableDictionaryClass struct {
	class objc.Class
}

// An interface definition for the [MutableDictionary] class.
type IMutableDictionary interface {
	IDictionary
}

// A dynamic collection of objects associated with unique keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary

type MutableDictionary struct {
	Dictionary
}

// MutableDictionaryFrom constructs a [MutableDictionary] from an unsafe.Pointer.
//
// A dynamic collection of objects associated with unique keys.
func MutableDictionaryFrom(ptr unsafe.Pointer) MutableDictionary {
	return MutableDictionary{
		Dictionary: DictionaryFrom(ptr),
	}
}



