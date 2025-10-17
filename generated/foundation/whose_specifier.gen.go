// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [WhoseSpecifier] class.
var whoseSpecifierClass = _WhoseSpecifierClass{objc.GetClass("NSWhoseSpecifier")}

type _WhoseSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [WhoseSpecifier] class.
type IWhoseSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier that indicates every object in a collection matching a condition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier

type WhoseSpecifier struct {
	ScriptObjectSpecifier
}

// WhoseSpecifierFrom constructs a [WhoseSpecifier] from an unsafe.Pointer.
//
// A specifier that indicates every object in a collection matching a condition.
func WhoseSpecifierFrom(ptr unsafe.Pointer) WhoseSpecifier {
	return WhoseSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}



