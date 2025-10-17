// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CloneCommand] class.
var cloneCommandClass = _CloneCommandClass{objc.GetClass("NSCloneCommand")}

type _CloneCommandClass struct {
	class objc.Class
}

// An interface definition for the [CloneCommand] class.
type ICloneCommand interface {
	IScriptCommand
}

// A command that clones one or more scriptable objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCloneCommand

type CloneCommand struct {
	ScriptCommand
}

// CloneCommandFrom constructs a [CloneCommand] from an unsafe.Pointer.
//
// A command that clones one or more scriptable objects.
func CloneCommandFrom(ptr unsafe.Pointer) CloneCommand {
	return CloneCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}



