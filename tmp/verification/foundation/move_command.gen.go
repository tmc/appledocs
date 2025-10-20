// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var moveCommandClass _MoveCommandClass

func init() {
	moveCommandClass = _MoveCommandClass{objc.GetClass("NSMoveCommand")}
}

type _MoveCommandClass struct {
	class objc.Class
}

type MoveCommand struct {
	objc.ID
}

func MoveCommandFrom(ptr unsafe.Pointer) MoveCommand {
	return MoveCommand{
		ID: objc.ID(ptr),
	}
}




