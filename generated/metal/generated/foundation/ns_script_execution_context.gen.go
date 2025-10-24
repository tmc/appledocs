// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptExecutionContext] class.
var (
	ScriptExecutionContextClass     _ScriptExecutionContextClass
	ScriptExecutionContextClassOnce sync.Once
)

func getScriptExecutionContextClass() _ScriptExecutionContextClass {
	ScriptExecutionContextClassOnce.Do(func() {
		ScriptExecutionContextClass = _ScriptExecutionContextClass{objc.GetClass("NSScriptExecutionContext")}
	})
	return ScriptExecutionContextClass
}

type _ScriptExecutionContextClass struct {
	class objc.Class
}

// An interface definition for the [ScriptExecutionContext] class.
type IScriptExecutionContext interface {
	objectivec.IObject
	// properties:
	ObjectBeingTested() unsafe.Pointer
	SetObjectBeingTested(value unsafe.Pointer)
	RangeContainerObject() unsafe.Pointer
	SetRangeContainerObject(value unsafe.Pointer)
	TopLevelObject() unsafe.Pointer
	SetTopLevelObject(value unsafe.Pointer)
	// methods:
}

// The context in which the current script command is executed.
//
// An object is a shared instance (there is only one instance of the class) that represents the context in which the current script command is executed. tracks global state relating to the command being executed, especially the top-level container object (that is, the container implied by a specifier object that specifies no container) used in an evaluation of an object. In most cases, the top-level container for a complete series of nested object specifiers is automatically set to the application object ( ), and you can get this object with the method. But you can also set this top-level container to something else (using ) if the situation warrants it. It is unlikely that you will need to subclass .


// The context in which the current script command is executed.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getScriptExecutionContextClass().New()
}



// Sets the top-level container object currently being tested in a “whose” qualifier to a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/objectbeingtested
func (s_ ScriptExecutionContext) ObjectBeingTested() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("objectBeingTested"))
	return rv
}


// Sets the top-level container object currently being tested in a “whose” qualifier to a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/objectbeingtested
func (s_ ScriptExecutionContext) SetObjectBeingTested(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setObjectBeingTested:"), value)
}


// Sets the top-level container object for a range-specifier evaluation to a give object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/rangecontainerobject
func (s_ ScriptExecutionContext) RangeContainerObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rangeContainerObject"))
	return rv
}


// Sets the top-level container object for a range-specifier evaluation to a give object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/rangecontainerobject
func (s_ ScriptExecutionContext) SetRangeContainerObject(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRangeContainerObject:"), value)
}


// Sets the top-level object for an object-specifier evaluation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/toplevelobject
func (s_ ScriptExecutionContext) TopLevelObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("topLevelObject"))
	return rv
}


// Sets the top-level object for an object-specifier evaluation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptexecutioncontext/toplevelobject
func (s_ ScriptExecutionContext) SetTopLevelObject(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTopLevelObject:"), value)
}



