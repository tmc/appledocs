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
	AppleEventCodeForArgumentWithName(argumentName string) unsafe.Pointer
	CreateCommandInstance() ScriptCommand
	CreateCommandInstanceWithZone(zone unsafe.Pointer) ScriptCommand
	IsOptionalArgumentWithName(argumentName string) bool
	TypeForArgumentWithName(argumentName string) String
	AppleEventClassCode() unsafe.Pointer
	AppleEventCode() unsafe.Pointer
	AppleEventCodeForReturnType() unsafe.Pointer
	ArgumentNames() []string
	CommandClassName() string
	CommandName() string
	ReturnType() string
	SuiteName() string
}

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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/init(coder:)
func NewScriptCommandDescriptionWithCoder(inCoder ICoder) ScriptCommandDescription {
	instance := getScriptCommandDescriptionClass().Alloc()
	rv := objc.Send[ScriptCommandDescription](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	rv.Autorelease()
	return rv
}


// Initializes and returns a newly allocated instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/init(suiteName:commandName:dictionary:)
func NewScriptCommandDescriptionWithSuiteNameCommandNameDictionary(suiteName string, commandName string, commandDeclaration objectivec.IObject) ScriptCommandDescription {
	instance := getScriptCommandDescriptionClass().Alloc()
	rv := objc.Send[ScriptCommandDescription](instance.ID, objc.Sel("initWithSuiteName:commandName:dictionary:"), objc.String(suiteName), objc.String(commandName), commandDeclaration)
	rv.Autorelease()
	return rv
}



// Returns the Apple event code for the specified command argument of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/appleEventCodeForArgument(withName:)
func (s_ ScriptCommandDescription) AppleEventCodeForArgumentWithName(argumentName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("appleEventCodeForArgumentWithName:"), objc.String(argumentName))
	return rv
}


// Creates and returns an instance of the command object described by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/createCommandInstance()
func (s_ ScriptCommandDescription) CreateCommandInstance() ScriptCommand {
	rv := objc.Send[ScriptCommand](s_.ID, objc.Sel("createCommandInstance"))
	return rv
}


// Creates and returns an instance of the command object described by the receiver in the specified memory zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/createCommandInstance(with:)
func (s_ ScriptCommandDescription) CreateCommandInstanceWithZone(zone unsafe.Pointer) ScriptCommand {
	rv := objc.Send[ScriptCommand](s_.ID, objc.Sel("createCommandInstanceWithZone:"), zone)
	return rv
}


// Returns a Boolean value that indicates whether the command argument identified by the specified argument key is an optional argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/isOptionalArgument(withName:)
func (s_ ScriptCommandDescription) IsOptionalArgumentWithName(argumentName string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isOptionalArgumentWithName:"), objc.String(argumentName))
	return rv
}


// Returns the type of the command argument identified by the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/typeForArgument(withName:)
func (s_ ScriptCommandDescription) TypeForArgumentWithName(argumentName string) String {
	rv := objc.Send[String](s_.ID, objc.Sel("typeForArgumentWithName:"), objc.String(argumentName))
	return rv
}


// Returns the four-character code for the Apple event class of the receiver’s command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/appleEventClassCode
func (s_ ScriptCommandDescription) AppleEventClassCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("appleEventClassCode"))
	return rv
}


// Returns the four-character code for the Apple event ID of the receiver’s command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/appleEventCode
func (s_ ScriptCommandDescription) AppleEventCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("appleEventCode"))
	return rv
}


// Returns the Apple event code that identifies the command’s return type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/appleEventCodeForReturnType
func (s_ ScriptCommandDescription) AppleEventCodeForReturnType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("appleEventCodeForReturnType"))
	return rv
}


// Returns the names (or keys) for all arguments of the receiver’s command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/argumentNames
func (s_ ScriptCommandDescription) ArgumentNames() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("argumentNames"))
	return rv
}


// Returns the name of the class that will be instantiated to handle the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/commandClassName
func (s_ ScriptCommandDescription) CommandClassName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("commandClassName"))
	return rv
}


// Returns the name of the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/commandName
func (s_ ScriptCommandDescription) CommandName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("commandName"))
	return rv
}


// Returns the return type of the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/returnType
func (s_ ScriptCommandDescription) ReturnType() string {
	rv := objc.Send[string](s_.ID, objc.Sel("returnType"))
	return rv
}


// Returns the name of the suite that contains the command described by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription/suiteName
func (s_ ScriptCommandDescription) SuiteName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("suiteName"))
	return rv
}


