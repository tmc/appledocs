// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DeleteCommand] class.
var deleteCommandClass = _DeleteCommandClass{objc.GetClass("NSDeleteCommand")}

type _DeleteCommandClass struct {
	class objc.Class
}

// An interface definition for the [DeleteCommand] class.
type IDeleteCommand interface {
	IScriptCommand
}

// A command that deletes a scriptable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDeleteCommand

type DeleteCommand struct {
	ScriptCommand
}

// DeleteCommandFrom constructs a [DeleteCommand] from an unsafe.Pointer.
//
// A command that deletes a scriptable object.
func DeleteCommandFrom(ptr unsafe.Pointer) DeleteCommand {
	return DeleteCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}



