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



