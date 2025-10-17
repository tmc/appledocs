// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ConstantString] class.
var constantStringClass = _ConstantStringClass{objc.GetClass("NSConstantString")}

type _ConstantStringClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConstantString

type ConstantString struct {
	SimpleCString
}

// ConstantStringFrom constructs a [ConstantString] from an unsafe.Pointer.
func ConstantStringFrom(ptr unsafe.Pointer) ConstantString {
	return ConstantString{
		SimpleCString: SimpleCStringFrom(ptr),
	}
}



