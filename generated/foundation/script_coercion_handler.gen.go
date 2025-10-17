// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScriptCoercionHandler] class.
var ScriptCoercionHandlerClass objc.Class

func init() {
	ScriptCoercionHandlerClass = objc.GetClass("NSScriptCoercionHandler")
}

type ScriptCoercionHandler struct {
	objc.ID
}

func ScriptCoercionHandlerFrom(ptr unsafe.Pointer) ScriptCoercionHandler {
	return ScriptCoercionHandler{
		ID: objc.ID(ptr),
	}
}



