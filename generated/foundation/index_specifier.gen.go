// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [IndexSpecifier] class.
var indexSpecifierClass = _IndexSpecifierClass{objc.GetClass("NSIndexSpecifier")}

type _IndexSpecifierClass struct {
	class objc.Class
}

// A specifier representing an object in a collection (or container) with an index number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier

type IndexSpecifier struct {
	ScriptObjectSpecifier
}

// IndexSpecifierFrom constructs a [IndexSpecifier] from an unsafe.Pointer.
//
// A specifier representing an object in a collection (or container) with an index number.
func IndexSpecifierFrom(ptr unsafe.Pointer) IndexSpecifier {
	return IndexSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}



