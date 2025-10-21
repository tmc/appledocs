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
	PerformDefaultImplementation() objc.ID
}

// A self-contained scripting statement.
//
// An instance of represents a scripting statement, such as , and contains the information needed to perform the operation specified by the statement. When an Apple event reaches a Cocoa application, Cocoa’s built-in scripting support transforms it into a script command (that is, an instance of or one of the subclasses provided by Cocoa scripting or by your application) and executes the command in the context of the application. Executing a command means either invoking the selector associated with the command on the object or objects designated to receive the command, or having the command perform its default implementation method ( ). Your application most likely calls methods of to extract the command arguments. You do this either in the method of a command subclass you have created, or in an object method designated as the selector to handle a particular command. As part of Cocoa’s standard scripting implementation, and its subclasses can handle the default command set for AppleScript’s Standard suite for most applications without any subclassing. The Standard suite includes commands such as , , , , , and , as well as common object classes such as , , and . For more information on working with script commands, see in .
//
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

// Overridden by subclasses to provide a default implementation for the command represented by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommand/performDefaultImplementation()
func (s_ ScriptCommand) PerformDefaultImplementation() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("performDefaultImplementation"))
	return rv
}

// Returns the object or objects to which the command is to be sent (called both the “receivers” or “targets” of script commands).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommand/evaluatedReceivers
func (s_ ScriptCommand) EvaluatedReceivers() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("evaluatedReceivers"))
	return rv
}

// Sets the object specifier to that, when evaluated, indicates the receiver or receivers of the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommand/receiversSpecifier
func (s_ ScriptCommand) ReceiversSpecifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("receiversSpecifier"))
	return rv
}

// SetReceiversSpecifier sets the value of the receiversSpecifier property.
// Sets the object specifier to that, when evaluated, indicates the receiver or receivers of the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommand/receiversSpecifier
func (s_ ScriptCommand) SetReceiversSpecifier(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReceiversSpecifier:"), value)
}
