// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptCommand] class.
var (
	ScriptCommandClass     _ScriptCommandClass
	ScriptCommandClassOnce sync.Once
)

func getScriptCommandClass() _ScriptCommandClass {
	ScriptCommandClassOnce.Do(func() {
		ScriptCommandClass = _ScriptCommandClass{objc.GetClass("NSScriptCommand")}
	})
	return ScriptCommandClass
}

type _ScriptCommandClass struct {
	class objc.Class
}

// An interface definition for the [ScriptCommand] class.
type IScriptCommand interface {
	objectivec.IObject
	// properties:
	AppleEvent() IAppleEventDescriptor
	SetAppleEvent(value IAppleEventDescriptor)
	Arguments() IString
	SetArguments(value IString)
	CommandDescription() IScriptCommandDescription
	SetCommandDescription(value IScriptCommandDescription)
	DirectParameter() unsafe.Pointer
	SetDirectParameter(value unsafe.Pointer)
	EvaluatedArguments() IString
	SetEvaluatedArguments(value IString)
	EvaluatedReceivers() unsafe.Pointer
	SetEvaluatedReceivers(value unsafe.Pointer)
	IsWellFormed() bool
	SetIsWellFormed(value bool)
	ReceiversSpecifier() IScriptObjectSpecifier
	SetReceiversSpecifier(value IScriptObjectSpecifier)
	ScriptErrorExpectedTypeDescriptor() IAppleEventDescriptor
	SetScriptErrorExpectedTypeDescriptor(value IAppleEventDescriptor)
	ScriptErrorNumber() int
	SetScriptErrorNumber(value int)
	ScriptErrorOffendingObjectDescriptor() IAppleEventDescriptor
	SetScriptErrorOffendingObjectDescriptor(value IAppleEventDescriptor)
	ScriptErrorString() IString
	SetScriptErrorString(value IString)
	// methods:
}

// A self-contained scripting statement.
//
// An instance of represents a scripting statement, such as , and contains the information needed to perform the operation specified by the statement. When an Apple event reaches a Cocoa application, Cocoa’s built-in scripting support transforms it into a script command (that is, an instance of or one of the subclasses provided by Cocoa scripting or by your application) and executes the command in the context of the application. Executing a command means either invoking the selector associated with the command on the object or objects designated to receive the command, or having the command perform its default implementation method ( ). Your application most likely calls methods of to extract the command arguments. You do this either in the method of a command subclass you have created, or in an object method designated as the selector to handle a particular command. As part of Cocoa’s standard scripting implementation, and its subclasses can handle the default command set for AppleScript’s Standard suite for most applications without any subclassing. The Standard suite includes commands such as , , , , , and , as well as common object classes such as , , and . For more information on working with script commands, see in .


// A self-contained scripting statement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommand
type ScriptCommand struct {
	objectivec.Object
}

// ScriptCommandFrom constructs a [ScriptCommand] from an unsafe.Pointer.
//
// A self-contained scripting statement.
func ScriptCommandFrom(ptr unsafe.Pointer) ScriptCommand {
	return ScriptCommand{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScriptCommandClass) Alloc() ScriptCommand {
	rv := objc.Send[ScriptCommand](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScriptCommandClass) New() ScriptCommand {
	rv := objc.Send[ScriptCommand](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptCommand) Init() ScriptCommand {
	rv := objc.Send[ScriptCommand](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptCommand) Autorelease() ScriptCommand {
	rv := objc.Send[ScriptCommand](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptCommand creates a new ScriptCommand instance.
func NewScriptCommand() ScriptCommand {
	return getScriptCommandClass().New()
}



// If the receiver was constructed by Cocoa scripting’s built-in Apple event handling, returns the Apple event descriptor from which it was constructed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/appleevent
func (s_ ScriptCommand) AppleEvent() IAppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](s_.ID, objc.Sel("appleEvent"))
	return rv
}


// If the receiver was constructed by Cocoa scripting’s built-in Apple event handling, returns the Apple event descriptor from which it was constructed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/appleevent
func (s_ ScriptCommand) SetAppleEvent(value IAppleEventDescriptor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAppleEvent:"), value)
}


// Sets the arguments of the command to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/arguments
func (s_ ScriptCommand) Arguments() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("arguments"))
	return rv
}


// Sets the arguments of the command to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/arguments
func (s_ ScriptCommand) SetArguments(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setArguments:"), value)
}


// Returns the command description for the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/commanddescription
func (s_ ScriptCommand) CommandDescription() IScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](s_.ID, objc.Sel("commandDescription"))
	return rv
}


// Returns the command description for the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/commanddescription
func (s_ ScriptCommand) SetCommandDescription(value IScriptCommandDescription) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCommandDescription:"), value)
}


// Sets the object that corresponds to the direct parameter of the Apple event from which the receiver derives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/directparameter
func (s_ ScriptCommand) DirectParameter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("directParameter"))
	return rv
}


// Sets the object that corresponds to the direct parameter of the Apple event from which the receiver derives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/directparameter
func (s_ ScriptCommand) SetDirectParameter(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDirectParameter:"), value)
}


// Returns a dictionary containing the arguments of the command, evaluated from object specifiers to objects if necessary. The keys in the dictionary are the argument names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/evaluatedarguments
func (s_ ScriptCommand) EvaluatedArguments() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("evaluatedArguments"))
	return rv
}


// Returns a dictionary containing the arguments of the command, evaluated from object specifiers to objects if necessary. The keys in the dictionary are the argument names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/evaluatedarguments
func (s_ ScriptCommand) SetEvaluatedArguments(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEvaluatedArguments:"), value)
}


// Returns the object or objects to which the command is to be sent (called both the “receivers” or “targets” of script commands).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/evaluatedreceivers
func (s_ ScriptCommand) EvaluatedReceivers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("evaluatedReceivers"))
	return rv
}


// Returns the object or objects to which the command is to be sent (called both the “receivers” or “targets” of script commands).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/evaluatedreceivers
func (s_ ScriptCommand) SetEvaluatedReceivers(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEvaluatedReceivers:"), value)
}


// Returns a Boolean value indicating whether the receiver is well formed according to its command description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/iswellformed
func (s_ ScriptCommand) IsWellFormed() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isWellFormed"))
	return rv
}


// Returns a Boolean value indicating whether the receiver is well formed according to its command description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/iswellformed
func (s_ ScriptCommand) SetIsWellFormed(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsWellFormed:"), value)
}


// Sets the object specifier to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/receiversspecifier
func (s_ ScriptCommand) ReceiversSpecifier() IScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](s_.ID, objc.Sel("receiversSpecifier"))
	return rv
}


// Sets the object specifier to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/receiversspecifier
func (s_ ScriptCommand) SetReceiversSpecifier(value IScriptObjectSpecifier) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReceiversSpecifier:"), value)
}


// Sets a descriptor for the expected type that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/scripterrorexpectedtypedescriptor
func (s_ ScriptCommand) ScriptErrorExpectedTypeDescriptor() IAppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](s_.ID, objc.Sel("scriptErrorExpectedTypeDescriptor"))
	return rv
}


// Sets a descriptor for the expected type that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/scripterrorexpectedtypedescriptor
func (s_ ScriptCommand) SetScriptErrorExpectedTypeDescriptor(value IAppleEventDescriptor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScriptErrorExpectedTypeDescriptor:"), value)
}


// Sets a script error number that is associated with the execution of the command and is returned in the reply Apple event, if a reply was requested by the sender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/scripterrornumber
func (s_ ScriptCommand) ScriptErrorNumber() int {
	rv := objc.Send[int](s_.ID, objc.Sel("scriptErrorNumber"))
	return rv
}


// Sets a script error number that is associated with the execution of the command and is returned in the reply Apple event, if a reply was requested by the sender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/scripterrornumber
func (s_ ScriptCommand) SetScriptErrorNumber(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScriptErrorNumber:"), value)
}


// Sets a descriptor for an object that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/scripterroroffendingobjectdescriptor
func (s_ ScriptCommand) ScriptErrorOffendingObjectDescriptor() IAppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](s_.ID, objc.Sel("scriptErrorOffendingObjectDescriptor"))
	return rv
}


// Sets a descriptor for an object that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/scripterroroffendingobjectdescriptor
func (s_ ScriptCommand) SetScriptErrorOffendingObjectDescriptor(value IAppleEventDescriptor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScriptErrorOffendingObjectDescriptor:"), value)
}


// Sets a script error string that is associated with execution of the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/scripterrorstring
func (s_ ScriptCommand) ScriptErrorString() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("scriptErrorString"))
	return rv
}


// Sets a script error string that is associated with execution of the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/scripterrorstring
func (s_ ScriptCommand) SetScriptErrorString(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScriptErrorString:"), value)
}



