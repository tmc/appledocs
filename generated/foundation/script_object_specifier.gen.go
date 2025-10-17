// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptObjectSpecifier] class.
var scriptObjectSpecifierClass = _ScriptObjectSpecifierClass{objc.GetClass("NSScriptObjectSpecifier")}

type _ScriptObjectSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [ScriptObjectSpecifier] class.
type IScriptObjectSpecifier interface {
	objectivec.IObject
}

// An abstract class used to represent natural language expressions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier

type ScriptObjectSpecifier struct {
	objectivec.Object
}

// ScriptObjectSpecifierFrom constructs a [ScriptObjectSpecifier] from an unsafe.Pointer.
//
// An abstract class used to represent natural language expressions.
func ScriptObjectSpecifierFrom(ptr unsafe.Pointer) ScriptObjectSpecifier {
	return ScriptObjectSpecifier{objectivec.Object{objc.ID(ptr)}}
}



