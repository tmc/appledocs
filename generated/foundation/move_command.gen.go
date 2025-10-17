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



