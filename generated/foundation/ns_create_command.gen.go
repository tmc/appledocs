// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSCreateCommand */


/* debug [class_header]: Header for NSCreateCommand */
// The class instance for the [CreateCommand] class.
var (
	CreateCommandClass     _CreateCommandClass
	CreateCommandClassOnce sync.Once
)

func getCreateCommandClass() _CreateCommandClass {
	CreateCommandClassOnce.Do(func() {
		CreateCommandClass = _CreateCommandClass{objc.GetClass("NSCreateCommand")}
	})
	return CreateCommandClass
}

type _CreateCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CreateCommand */
// An interface definition for the [CreateCommand] class.
type ICreateCommand interface {
	IScriptCommand
	
/* debug [class_interface_properties]: Properties for CreateCommand */
	// properties:
	CreateClassDescription() IScriptClassDescription
	ResolvedKeyDictionary() IDictionary
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CreateCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CreateCommand */
// Alloc allocates a new instance without initialization.
func (cc _CreateCommandClass) Alloc() CreateCommand {
	rv := objc.Send[CreateCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CreateCommandClass) New() CreateCommand {
	rv := objc.Send[CreateCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CreateCommand) Init() CreateCommand {
	rv := objc.Send[CreateCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CreateCommand) Autorelease() CreateCommand {
	rv := objc.Send[CreateCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCreateCommand creates a new CreateCommand instance.
func NewCreateCommand() CreateCommand {
	return getCreateCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CreateCommand */
// A command that creates a scriptable object.
//
// An instance of creates the specified scriptable object (such as a document), optionally supplying the new object with the specified attributes. This command corresponds to AppleScript’s command. is part of Cocoa’s built-in scripting support. Most applications don’t need to subclass or invoke its methods. When an instance of is executed, it creates a new object using (where is the class of the object to be created), unless the command has a argument. In the latter case, the new object is created by invoking . Any properties specified by a argument are then set in the new object using . If an object with no argument corresponding to the parameter is executed (for example, ), and the receiver of the command (not necessarily the application object) has a to-many relationship to objects of the class to be instantiated, and the class description for the receiving class returns when sent an message, the object creates a new object and sends the receiver an message to place the new object in the container. This is part of Cocoa’s scripting support for inserting newly-created objects into containers without explicitly specifying a location.


// A command that creates a scriptable object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateCommand
type CreateCommand struct {
	ScriptCommand
}

// CreateCommandFrom constructs a [CreateCommand] from an unsafe.Pointer.
//
// A command that creates a scriptable object.
func CreateCommandFrom(ptr unsafe.Pointer) CreateCommand {
	return CreateCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CreateCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CreateCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CreateCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CreateCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CreateCommand */

// Returns the class description for the class that is to be created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateCommand/createClassDescription
func (c_ CreateCommand) CreateClassDescription() IScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](c_.ID, objc.Sel("createClassDescription"))
	return rv
}/* debug [instance_properties/getter]: createClassDescription */


// Returns a dictionary that contains the properties that were specified in the Apple event command that has been converted to this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateCommand/resolvedKeyDictionary
func (c_ CreateCommand) ResolvedKeyDictionary() IDictionary {
	rv := objc.Send[Dictionary](c_.ID, objc.Sel("resolvedKeyDictionary"))
	return rv
}/* debug [instance_properties/getter]: resolvedKeyDictionary */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCreateCommand */



