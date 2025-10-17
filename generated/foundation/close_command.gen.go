// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CloseCommand] class.
var closeCommandClass = _CloseCommandClass{objc.GetClass("NSCloseCommand")}

type _CloseCommandClass struct {
	class objc.Class
}

// An interface definition for the [CloseCommand] class.
type ICloseCommand interface {
	IScriptCommand
}

// A command that closes one or more scriptable objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCloseCommand

type CloseCommand struct {
	ScriptCommand
}

// CloseCommandFrom constructs a [CloseCommand] from an unsafe.Pointer.
//
// A command that closes one or more scriptable objects.
func CloseCommandFrom(ptr unsafe.Pointer) CloseCommand {
	return CloseCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}



