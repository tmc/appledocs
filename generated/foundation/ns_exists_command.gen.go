// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSExistsCommand */


/* debug [class_header]: Header for NSExistsCommand */
// The class instance for the [ExistsCommand] class.
var (
	ExistsCommandClass     _ExistsCommandClass
	ExistsCommandClassOnce sync.Once
)

func getExistsCommandClass() _ExistsCommandClass {
	ExistsCommandClassOnce.Do(func() {
		ExistsCommandClass = _ExistsCommandClass{objc.GetClass("NSExistsCommand")}
	})
	return ExistsCommandClass
}

type _ExistsCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ExistsCommand */
// An interface definition for the [ExistsCommand] class.
type IExistsCommand interface {
	IScriptCommand
	
/* debug [class_interface_properties]: Properties for ExistsCommand */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ExistsCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ExistsCommand */
// Alloc allocates a new instance without initialization.
func (ec _ExistsCommandClass) Alloc() ExistsCommand {
	rv := objc.Send[ExistsCommand](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _ExistsCommandClass) New() ExistsCommand {
	rv := objc.Send[ExistsCommand](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExistsCommand) Init() ExistsCommand {
	rv := objc.Send[ExistsCommand](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExistsCommand) Autorelease() ExistsCommand {
	rv := objc.Send[ExistsCommand](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExistsCommand creates a new ExistsCommand instance.
func NewExistsCommand() ExistsCommand {
	return getExistsCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ExistsCommand */
// A command that determines whether a scriptable object exists.
//
// An instance of determines whether a specified scriptable object, such as a word, paragraph, or image, exists. When an instance of is executed, it evaluates the receiver specifier for the command to determine if it specifies any objects. is part of Cocoa’s built-in scripting support. Most applications don’t need to subclass .


// A command that determines whether a scriptable object exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExistsCommand
type ExistsCommand struct {
	ScriptCommand
}

// ExistsCommandFrom constructs a [ExistsCommand] from an unsafe.Pointer.
//
// A command that determines whether a scriptable object exists.
func ExistsCommandFrom(ptr unsafe.Pointer) ExistsCommand {
	return ExistsCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ExistsCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ExistsCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ExistsCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ExistsCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ExistsCommand */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSExistsCommand */



