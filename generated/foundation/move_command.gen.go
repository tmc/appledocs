// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MoveCommand] class.
var moveCommandClass = _MoveCommandClass{objc.GetClass("NSMoveCommand")}

type _MoveCommandClass struct {
	class objc.Class
}

// An interface definition for the [MoveCommand] class.
type IMoveCommand interface {
	IScriptCommand
}

// A command that moves one or more scriptable objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMoveCommand

type MoveCommand struct {
	ScriptCommand
}

// MoveCommandFrom constructs a [MoveCommand] from an unsafe.Pointer.
//
// A command that moves one or more scriptable objects.
func MoveCommandFrom(ptr unsafe.Pointer) MoveCommand {
	return MoveCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (mc _MoveCommandClass) Alloc() MoveCommand {
	rv := objc.Send[MoveCommand](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MoveCommandClass) New() MoveCommand {
	rv := objc.Send[MoveCommand](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MoveCommand) Init() MoveCommand {
	rv := objc.Send[MoveCommand](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MoveCommand) Autorelease() MoveCommand {
	rv := objc.Send[MoveCommand](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMoveCommand creates a new MoveCommand instance.
func NewMoveCommand() MoveCommand {
	return moveCommandClass.New()
}




