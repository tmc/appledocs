// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScriptExecutionContext */


/* debug [class_header]: Header for NSScriptExecutionContext */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScriptExecutionContext */
// An interface definition for the [ScriptExecutionContext] class.
type IScriptExecutionContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScriptExecutionContext */
	// properties:
	ObjectBeingTested() objc.ID
	SetObjectBeingTested(value objc.ID)
	RangeContainerObject() objc.ID
	SetRangeContainerObject(value objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScriptExecutionContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScriptExecutionContext */
// Alloc allocates a new instance without initialization.
func (sc _ScriptExecutionContextClass) Alloc() ScriptExecutionContext {
	rv := objc.Send[ScriptExecutionContext](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScriptExecutionContext */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScriptExecutionContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScriptExecutionContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScriptExecutionContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScriptExecutionContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScriptExecutionContext */

// Sets the top-level container object currently being tested in a “whose” qualifier to a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptExecutionContext/objectBeingTested
func (s_ ScriptExecutionContext) ObjectBeingTested() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("objectBeingTested"))
	return rv
}/* debug [instance_properties/getter]: objectBeingTested */


// Sets the top-level container object currently being tested in a “whose” qualifier to a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptExecutionContext/objectBeingTested
func (s_ ScriptExecutionContext) SetObjectBeingTested(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setObjectBeingTested:"), value)
}/* debug [instance_properties/setter]: objectBeingTested */


// Sets the top-level container object for a range-specifier evaluation to a give object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptExecutionContext/rangeContainerObject
func (s_ ScriptExecutionContext) RangeContainerObject() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("rangeContainerObject"))
	return rv
}/* debug [instance_properties/getter]: rangeContainerObject */


// Sets the top-level container object for a range-specifier evaluation to a give object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptExecutionContext/rangeContainerObject
func (s_ ScriptExecutionContext) SetRangeContainerObject(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRangeContainerObject:"), value)
}/* debug [instance_properties/setter]: rangeContainerObject */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScriptExecutionContext */



