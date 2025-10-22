// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ScriptCommandDescription] class.
type IScriptCommandDescription interface {
	objectivec.IObject
}

// A script command that a macOS app supports.
//
// A scriptable application provides scriptability information that describes the commands and objects scripters can use in scripts that target the application. An application’s scripting information is collected automatically by an instance of , which creates an for each command it finds, caches these objects in memory, and installs a command handler for each command. A script command instance stores the name, class, argument types, and return type of a command. For example, commands in AppleScript’s Core suite include , , , , , and . The public methods of are used primarily by Cocoa’s built-in scripting support in responding to Apple events that target the application. Although you can subclass the class, it is unlikely that you would need to do so, or to create instances of it.
//
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

// Alloc allocates a new instance without initialization.
func (sc _ScriptCommandDescriptionClass) Alloc() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the four-character code for the Apple event class of the receiver’s command.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventclasscode
func (s_ ScriptCommandDescription) AppleEventClassCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("appleEventClassCode"))
	return rv
}


// SetAppleEventClassCode sets the value of the appleEventClassCode property.
// Returns the four-character code for the Apple event class of the receiver’s command.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventclasscode
func (s_ ScriptCommandDescription) SetAppleEventClassCode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAppleEventClassCode:"), value)
}

// Returns the four-character code for the Apple event ID of the receiver’s command.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventcode
func (s_ ScriptCommandDescription) AppleEventCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("appleEventCode"))
	return rv
}


// SetAppleEventCode sets the value of the appleEventCode property.
// Returns the four-character code for the Apple event ID of the receiver’s command.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventcode
func (s_ ScriptCommandDescription) SetAppleEventCode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAppleEventCode:"), value)
}

// Returns the Apple event code that identifies the command’s return type.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventcodeforreturntype
func (s_ ScriptCommandDescription) AppleEventCodeForReturnType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("appleEventCodeForReturnType"))
	return rv
}


// SetAppleEventCodeForReturnType sets the value of the appleEventCodeForReturnType property.
// Returns the Apple event code that identifies the command’s return type.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/appleeventcodeforreturntype
func (s_ ScriptCommandDescription) SetAppleEventCodeForReturnType(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAppleEventCodeForReturnType:"), value)
}

// Returns the names (or keys) for all arguments of the receiver’s command.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/argumentnames
func (s_ ScriptCommandDescription) ArgumentNames() string {
	rv := objc.Send[string](s_.ID, objc.Sel("argumentNames"))
	return rv
}


// SetArgumentNames sets the value of the argumentNames property.
// Returns the names (or keys) for all arguments of the receiver’s command.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/argumentnames
func (s_ ScriptCommandDescription) SetArgumentNames(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setArgumentNames:"), objc.String(value))
}

// Returns the name of the class that will be instantiated to handle the command.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/commandclassname
func (s_ ScriptCommandDescription) CommandClassName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("commandClassName"))
	return rv
}


// SetCommandClassName sets the value of the commandClassName property.
// Returns the name of the class that will be instantiated to handle the command.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/commandclassname
func (s_ ScriptCommandDescription) SetCommandClassName(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCommandClassName:"), objc.String(value))
}

// Returns the name of the command.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/commandname
func (s_ ScriptCommandDescription) CommandName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("commandName"))
	return rv
}


// SetCommandName sets the value of the commandName property.
// Returns the name of the command.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/commandname
func (s_ ScriptCommandDescription) SetCommandName(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCommandName:"), objc.String(value))
}

// Returns the return type of the command.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/returntype
func (s_ ScriptCommandDescription) ReturnType() string {
	rv := objc.Send[string](s_.ID, objc.Sel("returnType"))
	return rv
}


// SetReturnType sets the value of the returnType property.
// Returns the return type of the command.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/returntype
func (s_ ScriptCommandDescription) SetReturnType(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReturnType:"), objc.String(value))
}

// Returns the name of the suite that contains the command described by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/suitename
func (s_ ScriptCommandDescription) SuiteName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("suiteName"))
	return rv
}


// SetSuiteName sets the value of the suiteName property.
// Returns the name of the suite that contains the command described by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/suitename
func (s_ ScriptCommandDescription) SetSuiteName(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuiteName:"), objc.String(value))
}



