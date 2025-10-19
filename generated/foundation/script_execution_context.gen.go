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
// Alloc allocates a new instance without initialization.
func (sc _ScriptExecutionContextClass) Alloc() ScriptExecutionContext {
	rv := objc.Send[ScriptExecutionContext](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _ScriptExecutionContextClass) New() ScriptExecutionContext {
	rv := objc.Send[ScriptExecutionContext](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptExecutionContext) Init() ScriptExecutionContext {
	rv := objc.Send[ScriptExecutionContext](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptExecutionContext) Autorelease() ScriptExecutionContext {
	rv := objc.Send[ScriptExecutionContext](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptExecutionContext creates a new ScriptExecutionContext instance.
func NewScriptExecutionContext() ScriptExecutionContext {
	return scriptExecutionContextClass.New()
}




