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
	ScriptCoercionHandlerClass     _ScriptCoercionHandlerClass
	ScriptCoercionHandlerClassOnce sync.Once
)

func getScriptCoercionHandlerClass() _ScriptCoercionHandlerClass {
	ScriptCoercionHandlerClassOnce.Do(func() {
		ScriptCoercionHandlerClass = _ScriptCoercionHandlerClass{objc.GetClass("NSScriptCoercionHandler")}
	})
	return ScriptCoercionHandlerClass
}

type _ScriptCoercionHandlerClass struct {
	class objc.Class
}

// An interface definition for the [ScriptCoercionHandler] class.
type IScriptCoercionHandler interface {
	objectivec.IObject
	RegisterCoercerSelectorToConvertFromClassToClass(coercer objectivec.IObject, selector objc.SEL, fromClass objc.Class, toClass objc.Class)
}

// A mechanism for converting one kind of scripting data to another.
//
// A shared instance of this class coerces (converts) object values to objects of another class using information supplied by classes that register with it. Coercions frequently are required during key-value coding.


// A mechanism for converting one kind of scripting data to another.
//
// [Full Topic]
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



// Registers a given object (typically a class) to handle coercions (conversions) from one given class to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCoercionHandler/registerCoercer(_:selector:toConvertFrom:to:)

func (s_ ScriptCoercionHandler) RegisterCoercerSelectorToConvertFromClassToClass(coercer objectivec.IObject, selector objc.SEL, fromClass objc.Class, toClass objc.Class) {
	objc.Send[objc.ID](s_.ID, objc.Sel("registerCoercer:selector:toConvertFromClass:toClass:"), coercer, selector, fromClass, toClass)
}



