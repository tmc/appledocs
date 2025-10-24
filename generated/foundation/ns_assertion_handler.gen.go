// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAssertionHandler */


/* debug [class_header]: Header for NSAssertionHandler */
// The class instance for the [AssertionHandler] class.
var (
	AssertionHandlerClass     _AssertionHandlerClass
	AssertionHandlerClassOnce sync.Once
)

func getAssertionHandlerClass() _AssertionHandlerClass {
	AssertionHandlerClassOnce.Do(func() {
		AssertionHandlerClass = _AssertionHandlerClass{objc.GetClass("NSAssertionHandler")}
	})
	return AssertionHandlerClass
}

type _AssertionHandlerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssertionHandler */
// An interface definition for the [AssertionHandler] class.
type IAssertionHandler interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssertionHandler */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssertionHandler */
	// methods:
	HandleFailureInFunctionFileLineNumberDescription(functionName IString, fileName IString, line int, format IString)
	HandleFailureInMethodObjectFileLineNumberDescription(selector objc.SEL, object objc.IObject, fileName IString, line int, format IString)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssertionHandler */
// Alloc allocates a new instance without initialization.
func (ac _AssertionHandlerClass) Alloc() AssertionHandler {
	rv := objc.Send[AssertionHandler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssertionHandlerClass) New() AssertionHandler {
	rv := objc.Send[AssertionHandler](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssertionHandler) Init() AssertionHandler {
	rv := objc.Send[AssertionHandler](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssertionHandler) Autorelease() AssertionHandler {
	rv := objc.Send[AssertionHandler](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssertionHandler creates a new AssertionHandler instance.
func NewAssertionHandler() AssertionHandler {
	return getAssertionHandlerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssertionHandler */
// An object that logs an assertion to the console.
//
// objects are automatically created to handle false assertions. Assertion macros, such as and , are used to evaluate a condition, and if the condition evaluates to false, the macros pass a string to an object describing the failure. Each thread has its own object. When invoked, an assertion handler prints an error message that includes the method and class (or function) containing the assertion and raises an . You create assertions only using the assertion macros—you rarely need to invoke methods directly. The macros for use inside methods and functions send and messages respectively to the current assertion handler. The assertion handler for the current thread is obtained using the class method. See if you need to customize the behavior of .


// An object that logs an assertion to the console.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAssertionHandler
type AssertionHandler struct {
	objectivec.Object
}

// AssertionHandlerFrom constructs a [AssertionHandler] from an unsafe.Pointer.
//
// An object that logs an assertion to the console.
func AssertionHandlerFrom(ptr unsafe.Pointer) AssertionHandler {
	return AssertionHandler{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssertionHandler *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssertionHandler */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssertionHandler */

// Returns the object associated with the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAssertionHandler/current
func (ac _AssertionHandlerClass) CurrentHandler() AssertionHandler {
	rv := objc.Send[AssertionHandler](objc.ID(ac.class), objc.Sel("currentHandler"))
	return rv
}/* debug [class_properties_class/property]: currentHandler */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssertionHandler */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAssertionHandler/handleFailureInFunction:file:lineNumber:description:
func (a_ AssertionHandler) HandleFailureInFunctionFileLineNumberDescription(functionName IString, fileName IString, line int, format IString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("handleFailureInFunction:file:lineNumber:description:"), functionName, fileName, line, format)
}/* debug [instance_methods/method]: HandleFailureInFunctionFileLineNumberDescription */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAssertionHandler/handleFailureInMethod:object:file:lineNumber:description:
func (a_ AssertionHandler) HandleFailureInMethodObjectFileLineNumberDescription(selector objc.SEL, object objc.IObject, fileName IString, line int, format IString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("handleFailureInMethod:object:file:lineNumber:description:"), selector, object, fileName, line, format)
}/* debug [instance_methods/method]: HandleFailureInMethodObjectFileLineNumberDescription */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssertionHandler */

// Returns the object associated with the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAssertionHandler/current
func (a_ AssertionHandler) CurrentHandler() IAssertionHandler {
	rv := objc.Send[AssertionHandler](a_.ID, objc.Sel("currentHandler"))
	return rv
}/* debug [instance_properties/getter]: currentHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAssertionHandler */



