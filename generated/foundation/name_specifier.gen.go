// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NameSpecifier] class.
var nameSpecifierClass = _NameSpecifierClass{objc.GetClass("NSNameSpecifier")}

type _NameSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [NameSpecifier] class.
type INameSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier for an object in a collection (or container) by name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNameSpecifier

type NameSpecifier struct {
	ScriptObjectSpecifier
}

// NameSpecifierFrom constructs a [NameSpecifier] from an unsafe.Pointer.
//
// A specifier for an object in a collection (or container) by name.
func NameSpecifierFrom(ptr unsafe.Pointer) NameSpecifier {
	return NameSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}



