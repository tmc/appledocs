// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptCoercionHandler] class.
var (
	scriptCoercionHandlerClass     _ScriptCoercionHandlerClass
	scriptCoercionHandlerClassOnce sync.Once
)

func getScriptCoercionHandlerClass() _ScriptCoercionHandlerClass {
	scriptCoercionHandlerClassOnce.Do(func() {
		scriptCoercionHandlerClass = _ScriptCoercionHandlerClass{objc.GetClass("NSScriptCoercionHandler")}
	})
	return scriptCoercionHandlerClass
}

type _ScriptCoercionHandlerClass struct {
	class objc.Class
}

// An interface definition for the [ScriptCoercionHandler] class.
type IScriptCoercionHandler interface {
	objectivec.IObject
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

// Alloc allocates a new instance without initialization.
func (sc _ScriptCoercionHandlerClass) Alloc() ScriptCoercionHandler {
	rv := objc.Send[ScriptCoercionHandler](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScriptCoercionHandlerClass) New() ScriptCoercionHandler {
	rv := objc.Send[ScriptCoercionHandler](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptCoercionHandler) Init() ScriptCoercionHandler {
	rv := objc.Send[ScriptCoercionHandler](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptCoercionHandler) Autorelease() ScriptCoercionHandler {
	rv := objc.Send[ScriptCoercionHandler](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptCoercionHandler creates a new ScriptCoercionHandler instance.
func NewScriptCoercionHandler() ScriptCoercionHandler {
	return getScriptCoercionHandlerClass().New()
}




