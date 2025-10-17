// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GetCommand] class.
var getCommandClass = _GetCommandClass{objc.GetClass("NSGetCommand")}

type _GetCommandClass struct {
	class objc.Class
}

// An interface definition for the [GetCommand] class.
type IGetCommand interface {
	IScriptCommand
}

// A command that retrieves a value or object from a scriptable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGetCommand

type GetCommand struct {
	ScriptCommand
}

// GetCommandFrom constructs a [GetCommand] from an unsafe.Pointer.
//
// A command that retrieves a value or object from a scriptable object.
func GetCommandFrom(ptr unsafe.Pointer) GetCommand {
	return GetCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}



