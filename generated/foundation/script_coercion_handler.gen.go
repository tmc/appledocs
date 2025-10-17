// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptCoercionHandler] class.
var scriptCoercionHandlerClass = _ScriptCoercionHandlerClass{objc.GetClass("NSScriptCoercionHandler")}

type _ScriptCoercionHandlerClass struct {
	class objc.Class
}

// A mechanism for converting one kind of scripting data to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCoercionHandler

type ScriptCoercionHandler struct {
	objectivec.Object
}

// ScriptCoercionHandlerFrom constructs a [ScriptCoercionHandler] from an unsafe.Pointer.
//
// A mechanism for converting one kind of scripting data to another.
func ScriptCoercionHandlerFrom(ptr unsafe.Pointer) ScriptCoercionHandler {
	return ScriptCoercionHandler{objectivec.Object{objc.ID(ptr)}}
}



