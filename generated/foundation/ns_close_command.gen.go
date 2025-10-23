// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CloseCommand] class.
var (
	CloseCommandClass     _CloseCommandClass
	CloseCommandClassOnce sync.Once
)

func getCloseCommandClass() _CloseCommandClass {
	CloseCommandClassOnce.Do(func() {
		CloseCommandClass = _CloseCommandClass{objc.GetClass("NSCloseCommand")}
	})
	return CloseCommandClass
}

type _CloseCommandClass struct {
	class objc.Class
}

// An interface definition for the [CloseCommand] class.
type ICloseCommand interface {
	IScriptCommand
	// properties:
	SaveOptions() SaveOptions
	// methods:
}

// A command that closes one or more scriptable objects.
//
// An instance of closes the specified scriptable object or objects—typically a document or window (and its associated document, if any). The command may optionally specify a location to save in and how to handle modified documents (by automatically saving changes, not saving them, or asking the user). is part of Cocoa’s built-in scripting support. It works automatically to support the command through key-value coding. Most applications don’t need to subclass or call its methods.


// A command that closes one or more scriptable objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCloseCommand
type CloseCommand struct {
	ScriptCommand
}

// CloseCommandFrom constructs a [CloseCommand] from an unsafe.Pointer.
//
// A command that closes one or more scriptable objects.
func CloseCommandFrom(ptr unsafe.Pointer) CloseCommand {
	return CloseCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CloseCommandClass) Alloc() CloseCommand {
	rv := objc.Send[CloseCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CloseCommandClass) New() CloseCommand {
	rv := objc.Send[CloseCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CloseCommand) Init() CloseCommand {
	rv := objc.Send[CloseCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CloseCommand) Autorelease() CloseCommand {
	rv := objc.Send[CloseCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCloseCommand creates a new CloseCommand instance.
func NewCloseCommand() CloseCommand {
	return getCloseCommandClass().New()
}



// Returns a constant indicating how to deal with closing any modified documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCloseCommand/saveOptions
func (c_ CloseCommand) SaveOptions() SaveOptions {
	rv := objc.Send[SaveOptions](c_.ID, objc.Sel("saveOptions"))
	return rv
}



