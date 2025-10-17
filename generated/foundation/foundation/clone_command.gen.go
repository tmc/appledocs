// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CloneCommand] class.
var CloneCommandClass objc.Class

func init() {
	CloneCommandClass = objc.GetClass("NSCloneCommand")
}

type CloneCommand struct {
	objc.ID
}

func CloneCommandFrom(ptr unsafe.Pointer) CloneCommand {
	return CloneCommand{
		ID: objc.ID(ptr),
	}
}




