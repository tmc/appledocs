// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CreateCommand] class.
var createCommandClass = _CreateCommandClass{objc.GetClass("NSCreateCommand")}

type _CreateCommandClass struct {
	class objc.Class
}

// An interface definition for the [CreateCommand] class.
type ICreateCommand interface {
	IScriptCommand
}

// A command that creates a scriptable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateCommand

type CreateCommand struct {
	ScriptCommand
}

// CreateCommandFrom constructs a [CreateCommand] from an unsafe.Pointer.
//
// A command that creates a scriptable object.
func CreateCommandFrom(ptr unsafe.Pointer) CreateCommand {
	return CreateCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}



