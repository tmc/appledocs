// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var scriptCoercionHandlerClass _ScriptCoercionHandlerClass

func init() {
	scriptCoercionHandlerClass = _ScriptCoercionHandlerClass{objc.GetClass("NSScriptCoercionHandler")}
}

type _ScriptCoercionHandlerClass struct {
	class objc.Class
}

type ScriptCoercionHandler struct {
	objc.ID
}

func ScriptCoercionHandlerFrom(ptr unsafe.Pointer) ScriptCoercionHandler {
	return ScriptCoercionHandler{
		ID: objc.ID(ptr),
	}
}




