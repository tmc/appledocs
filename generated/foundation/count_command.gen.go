// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CountCommand] class.
var countCommandClass = _CountCommandClass{objc.GetClass("NSCountCommand")}

type _CountCommandClass struct {
	class objc.Class
}

// A command that counts the number of objects of a specified class in the specified object container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountCommand

type CountCommand struct {
	ScriptCommand
}

// CountCommandFrom constructs a [CountCommand] from an unsafe.Pointer.
//
// A command that counts the number of objects of a specified class in the specified object container.
func CountCommandFrom(ptr unsafe.Pointer) CountCommand {
	return CountCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}



