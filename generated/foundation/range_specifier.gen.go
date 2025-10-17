// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RangeSpecifier] class.
var rangeSpecifierClass = _RangeSpecifierClass{objc.GetClass("NSRangeSpecifier")}

type _RangeSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [RangeSpecifier] class.
type IRangeSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier for a range of objects in a container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier

type RangeSpecifier struct {
	ScriptObjectSpecifier
}

// RangeSpecifierFrom constructs a [RangeSpecifier] from an unsafe.Pointer.
//
// A specifier for a range of objects in a container.
func RangeSpecifierFrom(ptr unsafe.Pointer) RangeSpecifier {
	return RangeSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}



