// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MiddleSpecifier] class.
var middleSpecifierClass = _MiddleSpecifierClass{objc.GetClass("NSMiddleSpecifier")}

type _MiddleSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [MiddleSpecifier] class.
type IMiddleSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier indicating the middle object in a collection or, if not a one-to-many relationship, the sole object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMiddleSpecifier

type MiddleSpecifier struct {
	ScriptObjectSpecifier
}

// MiddleSpecifierFrom constructs a [MiddleSpecifier] from an unsafe.Pointer.
//
// A specifier indicating the middle object in a collection or, if not a one-to-many relationship, the sole object.
func MiddleSpecifierFrom(ptr unsafe.Pointer) MiddleSpecifier {
	return MiddleSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}



