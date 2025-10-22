// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AssertionHandler] class.
type IAssertionHandler interface {
	objectivec.IObject
	HandleFailureInFunctionFileLineNumberDescription(functionName string, fileName string, line int, format string)
	HandleFailureInMethodObjectFileLineNumberDescription(selector objc.SEL, object objectivec.IObject, fileName string, line int, format string)
}

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

// Alloc allocates a new instance without initialization.
func (ac _AssertionHandlerClass) Alloc() AssertionHandler {
	rv := objc.Send[AssertionHandler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns the object associated with the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAssertionHandler/current

func (ac _AssertionHandlerClass) CurrentHandler() AssertionHandler {
	rv := objc.Send[NSAssertionHandler](objc.ID(ac.class), objc.Sel("currentHandler"))
	return rv
}

//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAssertionHandler/handleFailureInFunction:file:lineNumber:description:

func (a_ AssertionHandler) HandleFailureInFunctionFileLineNumberDescription(functionName string, fileName string, line int, format string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("handleFailureInFunction:file:lineNumber:description:"), objc.String(functionName), objc.String(fileName), line, objc.String(format))
}


//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAssertionHandler/handleFailureInMethod:object:file:lineNumber:description:

func (a_ AssertionHandler) HandleFailureInMethodObjectFileLineNumberDescription(selector objc.SEL, object objectivec.IObject, fileName string, line int, format string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("handleFailureInMethod:object:file:lineNumber:description:"), selector, object, objc.String(fileName), line, objc.String(format))
}


// Returns the object associated with the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAssertionHandler/current

func (a_ AssertionHandler) CurrentHandler() NSAssertionHandler {
	rv := objc.Send[NSAssertionHandler](a_.ID, objc.Sel("currentHandler"))
	return rv
}



