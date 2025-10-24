// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScriptCommandDescription */


/* debug [class_header]: Header for NSScriptCommandDescription */
// The class instance for the [ScriptCommandDescription] class.
var (
	ScriptCommandDescriptionClass     _ScriptCommandDescriptionClass
	ScriptCommandDescriptionClassOnce sync.Once
)

func getScriptCommandDescriptionClass() _ScriptCommandDescriptionClass {
	ScriptCommandDescriptionClassOnce.Do(func() {
		ScriptCommandDescriptionClass = _ScriptCommandDescriptionClass{objc.GetClass("NSScriptCommandDescription")}
	})
	return ScriptCommandDescriptionClass
}

type _ScriptCommandDescriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScriptCommandDescription */
// An interface definition for the [ScriptCommandDescription] class.
type IScriptCommandDescription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScriptCommandDescription */
	// properties:
	CommandClassName() IString
	AppleEventClassCode() uint32 /* not a class type */
	SetAppleEventClassCode(value uint32 /* not a class type */)
	AppleEventCode() uint32 /* not a class type */
	SetAppleEventCode(value uint32 /* not a class type */)
	AppleEventCodeForReturnType() uint32 /* not a class type */
	SetAppleEventCodeForReturnType(value uint32 /* not a class type */)
	ArgumentNames() IString
	SetArgumentNames(value IString)
	CommandName() IString
	SetCommandName(value IString)
	ReturnType() IString
	SetReturnType(value IString)
	SuiteName() IString
	SetSuiteName(value IString)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScriptCommandDescription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScriptCommandDescription */
// Alloc allocates a new instance without initialization.
func (sc _ScriptCommandDescriptionClass) Alloc() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScriptCommandDescriptionClass) New() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptCommandDescription) Init() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptCommandDescription) Autorelease() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptCommandDescription creates a new ScriptCommandDescription instance.
func NewScriptCommandDescription() ScriptCommandDescription {
	return getScriptCommandDescriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScriptCommandDescription */
// A script command that a macOS app supports.
//
// A scriptable application provides scriptability information that describes the commands and objects scripters can use in scripts that target the application. An application’s scripting information is collected automatically by an instance of , which creates an for each command it finds, caches these objects in memory, and installs a command handler for each command. A script command instance stores the name, class, argument types, and return type of a command. For example, commands in AppleScript’s Core suite include , , , , , and . The public methods of are used primarily by Cocoa’s built-in scripting support in responding to Apple events that target the application. Although you can subclass the class, it is unlikely that you would need to do so, or to create instances of it.


// A script command that a macOS app supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription
type ScriptCommandDescription struct {
	objectivec.Object
}

// ScriptCommandDescriptionFrom constructs a [ScriptCommandDescription] from an unsafe.Pointer.
//
// A script command that a macOS app supports.
func ScriptCommandDescriptionFrom(ptr unsafe.Pointer) ScriptCommandDescription {
	return ScriptCommandDescription{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScriptCommandDescription *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScriptCommandDescription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScriptCommandDescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScriptCommandDescription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScriptCommandDescription */

// Returns the name of the class that will be instantiated to handle the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/commandClassName
func (s_ ScriptCommandDescription) CommandClassName() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("commandClassName"))
	return rv
}/* debug [instance_properties/getter]: commandClassName */


// Returns the four-character code for the Apple event class of the receiver’s command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventclasscode
func (s_ ScriptCommandDescription) AppleEventClassCode() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("appleEventClassCode"))
	return rv
}/* debug [instance_properties/getter]: appleEventClassCode */


// Returns the four-character code for the Apple event class of the receiver’s command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventclasscode
func (s_ ScriptCommandDescription) SetAppleEventClassCode(value uint32 /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAppleEventClassCode:"), value)
}/* debug [instance_properties/setter]: appleEventClassCode */


// Returns the four-character code for the Apple event ID of the receiver’s command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventcode
func (s_ ScriptCommandDescription) AppleEventCode() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("appleEventCode"))
	return rv
}/* debug [instance_properties/getter]: appleEventCode */


// Returns the four-character code for the Apple event ID of the receiver’s command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventcode
func (s_ ScriptCommandDescription) SetAppleEventCode(value uint32 /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAppleEventCode:"), value)
}/* debug [instance_properties/setter]: appleEventCode */


// Returns the Apple event code that identifies the command’s return type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventcodeforreturntype
func (s_ ScriptCommandDescription) AppleEventCodeForReturnType() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("appleEventCodeForReturnType"))
	return rv
}/* debug [instance_properties/getter]: appleEventCodeForReturnType */


// Returns the Apple event code that identifies the command’s return type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventcodeforreturntype
func (s_ ScriptCommandDescription) SetAppleEventCodeForReturnType(value uint32 /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAppleEventCodeForReturnType:"), value)
}/* debug [instance_properties/setter]: appleEventCodeForReturnType */


// Returns the names (or keys) for all arguments of the receiver’s command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/argumentnames
func (s_ ScriptCommandDescription) ArgumentNames() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("argumentNames"))
	return rv
}/* debug [instance_properties/getter]: argumentNames */


// Returns the names (or keys) for all arguments of the receiver’s command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/argumentnames
func (s_ ScriptCommandDescription) SetArgumentNames(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setArgumentNames:"), value)
}/* debug [instance_properties/setter]: argumentNames */


// Returns the name of the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/commandname
func (s_ ScriptCommandDescription) CommandName() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("commandName"))
	return rv
}/* debug [instance_properties/getter]: commandName */


// Returns the name of the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/commandname
func (s_ ScriptCommandDescription) SetCommandName(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCommandName:"), value)
}/* debug [instance_properties/setter]: commandName */


// Returns the return type of the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/returntype
func (s_ ScriptCommandDescription) ReturnType() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("returnType"))
	return rv
}/* debug [instance_properties/getter]: returnType */


// Returns the return type of the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/returntype
func (s_ ScriptCommandDescription) SetReturnType(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReturnType:"), value)
}/* debug [instance_properties/setter]: returnType */


// Returns the name of the suite that contains the command described by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/suitename
func (s_ ScriptCommandDescription) SuiteName() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("suiteName"))
	return rv
}/* debug [instance_properties/getter]: suiteName */


// Returns the name of the suite that contains the command described by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/suitename
func (s_ ScriptCommandDescription) SetSuiteName(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuiteName:"), value)
}/* debug [instance_properties/setter]: suiteName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScriptCommandDescription */



