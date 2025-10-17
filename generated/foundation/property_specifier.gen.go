// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PropertySpecifier] class.
var propertySpecifierClass = _PropertySpecifierClass{objc.GetClass("NSPropertySpecifier")}

type _PropertySpecifierClass struct {
	class objc.Class
}

// An interface definition for the [PropertySpecifier] class.
type IPropertySpecifier interface {
	IScriptObjectSpecifier
}

// A specifier for a simple attribute value, a one-to-one relationship, or all elements of a to-many relationship. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPropertySpecifier

type PropertySpecifier struct {
	ScriptObjectSpecifier
}

// PropertySpecifierFrom constructs a [PropertySpecifier] from an unsafe.Pointer.
//
// A specifier for a simple attribute value, a one-to-one relationship, or all elements of a to-many relationship.
func PropertySpecifierFrom(ptr unsafe.Pointer) PropertySpecifier {
	return PropertySpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}



