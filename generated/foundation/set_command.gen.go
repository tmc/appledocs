// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SetCommand] class.
var setCommandClass = _SetCommandClass{objc.GetClass("NSSetCommand")}

type _SetCommandClass struct {
	class objc.Class
}

// A command that sets one or more attributes or relationships to one or more values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSetCommand

type SetCommand struct {
	ScriptCommand
}

// SetCommandFrom constructs a [SetCommand] from an unsafe.Pointer.
//
// A command that sets one or more attributes or relationships to one or more values.
func SetCommandFrom(ptr unsafe.Pointer) SetCommand {
	return SetCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}



