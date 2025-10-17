// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MoveCommand] class.
var MoveCommandClass objc.Class

func init() {
	MoveCommandClass = objc.GetClass("NSMoveCommand")
}

type MoveCommand struct {
	objc.ID
}

func MoveCommandFrom(ptr unsafe.Pointer) MoveCommand {
	return MoveCommand{
		ID: objc.ID(ptr),
	}
}



