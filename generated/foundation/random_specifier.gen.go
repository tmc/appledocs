// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RandomSpecifier] class.
var randomSpecifierClass = _RandomSpecifierClass{objc.GetClass("NSRandomSpecifier")}

type _RandomSpecifierClass struct {
	class objc.Class
}

// A specifier for an arbitrary object in a collection or, if not a one-to-many relationship, the sole object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRandomSpecifier

type RandomSpecifier struct {
	ScriptObjectSpecifier
}

// RandomSpecifierFrom constructs a [RandomSpecifier] from an unsafe.Pointer.
//
// A specifier for an arbitrary object in a collection or, if not a one-to-many relationship, the sole object.
func RandomSpecifierFrom(ptr unsafe.Pointer) RandomSpecifier {
	return RandomSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}



