// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScriptCoercionHandler */


/* debug [class_header]: Header for NSScriptCoercionHandler */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScriptCoercionHandler */
// An interface definition for the [ScriptCoercionHandler] class.
type IScriptCoercionHandler interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScriptCoercionHandler */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScriptCoercionHandler */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScriptCoercionHandler */
// Alloc allocates a new instance without initialization.
func (sc _ScriptCoercionHandlerClass) Alloc() ScriptCoercionHandler {
	rv := objc.Send[ScriptCoercionHandler](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScriptCoercionHandler */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScriptCoercionHandler *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScriptCoercionHandler */

// Returns the shared for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCoercionHandler/shared()
func (sc _ScriptCoercionHandlerClass) SharedCoercionHandler() IScriptCoercionHandler {
	rv := objc.Send[ScriptCoercionHandler](objc.ID(sc.class), objc.Sel("sharedCoercionHandler"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedCoercionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScriptCoercionHandler */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScriptCoercionHandler */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScriptCoercionHandler */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScriptCoercionHandler */



