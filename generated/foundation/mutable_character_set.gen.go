// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableCharacterSet] class.
var mutableCharacterSetClass = _MutableCharacterSetClass{objc.GetClass("NSMutableCharacterSet")}

type _MutableCharacterSetClass struct {
	class objc.Class
}

// An interface definition for the [MutableCharacterSet] class.
type IMutableCharacterSet interface {
	ICharacterSet
}

// An object representing a mutable set of Unicode character values for use in search operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet

type MutableCharacterSet struct {
	CharacterSet
}

// MutableCharacterSetFrom constructs a [MutableCharacterSet] from an unsafe.Pointer.
//
// An object representing a mutable set of Unicode character values for use in search operations.
func MutableCharacterSetFrom(ptr unsafe.Pointer) MutableCharacterSet {
	return MutableCharacterSet{
		CharacterSet: CharacterSetFrom(ptr),
	}
}



