// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ScriptCoercionHandlerClass _ScriptCoercionHandlerClass

func init() {
	ScriptCoercionHandlerClass = _ScriptCoercionHandlerClass{objc.GetClass("NSScriptCoercionHandler")}
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




