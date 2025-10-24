// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSQuitCommand */


/* debug [class_header]: Header for NSQuitCommand */
// The class instance for the [QuitCommand] class.
var (
	QuitCommandClass     _QuitCommandClass
	QuitCommandClassOnce sync.Once
)

func getQuitCommandClass() _QuitCommandClass {
	QuitCommandClassOnce.Do(func() {
		QuitCommandClass = _QuitCommandClass{objc.GetClass("NSQuitCommand")}
	})
	return QuitCommandClass
}

type _QuitCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QuitCommand */
// An interface definition for the [QuitCommand] class.
type IQuitCommand interface {
	IScriptCommand
	
/* debug [class_interface_properties]: Properties for QuitCommand */
	// properties:
	SaveOptions() SaveOptions /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QuitCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QuitCommand */
// Alloc allocates a new instance without initialization.
func (qc _QuitCommandClass) Alloc() QuitCommand {
	rv := objc.Send[QuitCommand](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QuitCommandClass) New() QuitCommand {
	rv := objc.Send[QuitCommand](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuitCommand) Init() QuitCommand {
	rv := objc.Send[QuitCommand](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuitCommand) Autorelease() QuitCommand {
	rv := objc.Send[QuitCommand](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuitCommand creates a new QuitCommand instance.
func NewQuitCommand() QuitCommand {
	return getQuitCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QuitCommand */
// A command that quits the specified app.
//
// The quit command may optionally specify how to handle modified documents (automatically save changes, don’t save them, or ask the user). For details, see the description for the command in “Apple Events Sent By the Mac OS” in in . is part of Cocoa’s built-in scripting support. Most applications don’t need to subclass or call its methods.


// A command that quits the specified app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSQuitCommand
type QuitCommand struct {
	ScriptCommand
}

// QuitCommandFrom constructs a [QuitCommand] from an unsafe.Pointer.
//
// A command that quits the specified app.
func QuitCommandFrom(ptr unsafe.Pointer) QuitCommand {
	return QuitCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QuitCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QuitCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QuitCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QuitCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QuitCommand */

// Returns a constant indicating how to deal with closing any modified documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSQuitCommand/saveOptions
func (q_ QuitCommand) SaveOptions() SaveOptions /* not a class type */ {
	rv := objc.Send[SaveOptions](q_.ID, objc.Sel("saveOptions"))
	return rv
}/* debug [instance_properties/getter]: saveOptions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSQuitCommand */



