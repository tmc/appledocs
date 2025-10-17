// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CharacterSet] class.
var characterSetClass = _CharacterSetClass{objc.GetClass("NSCharacterSet")}

type _CharacterSetClass struct {
	class objc.Class
}

// An object representing a fixed set of Unicode character values for use in search operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet

type CharacterSet struct {
	objectivec.Object
}

// CharacterSetFrom constructs a [CharacterSet] from an unsafe.Pointer.
//
// An object representing a fixed set of Unicode character values for use in search operations.
func CharacterSetFrom(ptr unsafe.Pointer) CharacterSet {
	return CharacterSet{objectivec.Object{objc.ID(ptr)}}
}



