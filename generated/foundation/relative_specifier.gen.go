// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RelativeSpecifier] class.
var relativeSpecifierClass = _RelativeSpecifierClass{objc.GetClass("NSRelativeSpecifier")}

type _RelativeSpecifierClass struct {
	class objc.Class
}

// A specifier that indicates an object in a collection by its position relative to another object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRelativeSpecifier

type RelativeSpecifier struct {
	ScriptObjectSpecifier
}

// RelativeSpecifierFrom constructs a [RelativeSpecifier] from an unsafe.Pointer.
//
// A specifier that indicates an object in a collection by its position relative to another object.
func RelativeSpecifierFrom(ptr unsafe.Pointer) RelativeSpecifier {
	return RelativeSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}



