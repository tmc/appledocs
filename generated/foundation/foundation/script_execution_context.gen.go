// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScriptExecutionContext] class.
var ScriptExecutionContextClass objc.Class

func init() {
	ScriptExecutionContextClass = objc.GetClass("NSScriptExecutionContext")
}

type ScriptExecutionContext struct {
	objc.ID
}

func ScriptExecutionContextFrom(ptr unsafe.Pointer) ScriptExecutionContext {
	return ScriptExecutionContext{
		ID: objc.ID(ptr),
	}
}




