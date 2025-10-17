// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DeleteCommand] class.
var DeleteCommandClass objc.Class

func init() {
	DeleteCommandClass = objc.GetClass("NSDeleteCommand")
}

type DeleteCommand struct {
	objc.ID
}

func DeleteCommandFrom(ptr unsafe.Pointer) DeleteCommand {
	return DeleteCommand{
		ID: objc.ID(ptr),
	}
}




