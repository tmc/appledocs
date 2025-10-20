// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MoveCommand] class.
var (
	moveCommandClass     _MoveCommandClass
	moveCommandClassOnce sync.Once
)

func getMoveCommandClass() _MoveCommandClass {
	moveCommandClassOnce.Do(func() {
		moveCommandClass = _MoveCommandClass{objc.GetClass("NSMoveCommand")}
	})
	return moveCommandClass
}

type _MoveCommandClass struct {
	class objc.Class
}

// An interface definition for the [MoveCommand] class.
type IMoveCommand interface {
	IScriptCommand
}

// A command that moves one or more scriptable objects.
//
// An instance of moves the specified scriptable object or objects; for example, it may move words to a new location in a document or a file to a new directory. is part of Cocoa’s built-in scripting support. It works automatically to support the AppleScript command through key-value coding. Most applications don’t need to subclass or invoke its methods. However, for circumstances where you might choose to subclass this command, see “Modifying a Standard Command” in in . When an instance of is executed, it does not make copies of moved objects. It removes objects from the source container or containers, then inserts them into the destination container.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getMoveCommandClass().New()
}




