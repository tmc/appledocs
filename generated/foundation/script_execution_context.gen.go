// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScriptExecutionContext] class.
var ScriptExecutionContextClass = _ScriptExecutionContextClass{objc.GetClass("NSScriptExecutionContext")}

type _ScriptExecutionContextClass struct {
	class objc.Class
}

type ScriptExecutionContext struct {
	objc.ID
}

func ScriptExecutionContextFrom(ptr unsafe.Pointer) ScriptExecutionContext {
	return ScriptExecutionContext{
		ID: objc.ID(ptr),
	}
}




