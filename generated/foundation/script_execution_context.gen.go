// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptExecutionContext] class.
var scriptExecutionContextClass = _ScriptExecutionContextClass{objc.GetClass("NSScriptExecutionContext")}

type _ScriptExecutionContextClass struct {
	class objc.Class
}

// An interface definition for the [ScriptExecutionContext] class.
type IScriptExecutionContext interface {
	objectivec.IObject
}

// The context in which the current script command is executed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptExecutionContext

type ScriptExecutionContext struct {
	objectivec.Object
}

// ScriptExecutionContextFrom constructs a [ScriptExecutionContext] from an unsafe.Pointer.
//
// The context in which the current script command is executed.
func ScriptExecutionContextFrom(ptr unsafe.Pointer) ScriptExecutionContext {
	return ScriptExecutionContext{objectivec.Object{objc.ID(ptr)}}
}



