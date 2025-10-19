// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [QuitCommand] class.
var quitCommandClass = _QuitCommandClass{objc.GetClass("NSQuitCommand")}

type _QuitCommandClass struct {
	class objc.Class
}

// An interface definition for the [QuitCommand] class.
type IQuitCommand interface {
	IScriptCommand
}

// A command that quits the specified app. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return quitCommandClass.New()
}




