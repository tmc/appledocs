// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [QuitCommand] class.
type IQuitCommand interface {
	IScriptCommand
	SaveOptions() NSSaveOptions
}

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

// Alloc allocates a new instance without initialization.
func (qc _QuitCommandClass) Alloc() QuitCommand {
	rv := objc.Send[QuitCommand](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns a constant indicating how to deal with closing any modified documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSQuitCommand/saveOptions
func (q_ QuitCommand) SaveOptions() NSSaveOptions {
	rv := objc.Send[SaveOptions](q_.ID, objc.Sel("saveOptions"))
	return rv
}



