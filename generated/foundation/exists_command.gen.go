// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ExistsCommand] class.
var existsCommandClass = _ExistsCommandClass{objc.GetClass("NSExistsCommand")}

type _ExistsCommandClass struct {
	class objc.Class
}

// A command that determines whether a scriptable object exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExistsCommand

type ExistsCommand struct {
	ScriptCommand
}

// ExistsCommandFrom constructs a [ExistsCommand] from an unsafe.Pointer.
//
// A command that determines whether a scriptable object exists.
func ExistsCommandFrom(ptr unsafe.Pointer) ExistsCommand {
	return ExistsCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}



