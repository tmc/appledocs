// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var cloneCommandClass _CloneCommandClass

func init() {
	cloneCommandClass = _CloneCommandClass{objc.GetClass("NSCloneCommand")}
}

type _CloneCommandClass struct {
	class objc.Class
}

type CloneCommand struct {
	objc.ID
}

func CloneCommandFrom(ptr unsafe.Pointer) CloneCommand {
	return CloneCommand{
		ID: objc.ID(ptr),
	}
}




