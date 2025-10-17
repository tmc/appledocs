// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SimpleCString] class.
var simpleCStringClass = _SimpleCStringClass{objc.GetClass("NSSimpleCString")}

type _SimpleCStringClass struct {
	class objc.Class
}

// An interface definition for the [SimpleCString] class.
type ISimpleCString interface {
	IString
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSimpleCString

type SimpleCString struct {
	String
}

// SimpleCStringFrom constructs a [SimpleCString] from an unsafe.Pointer.
func SimpleCStringFrom(ptr unsafe.Pointer) SimpleCString {
	return SimpleCString{
		String: StringFrom(ptr),
	}
}



