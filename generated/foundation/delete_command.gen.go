// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DeleteCommand] class.
var DeleteCommandClass = _DeleteCommandClass{objc.GetClass("NSDeleteCommand")}

type _DeleteCommandClass struct {
	class objc.Class
}

type DeleteCommand struct {
	objc.ID
}

func DeleteCommandFrom(ptr unsafe.Pointer) DeleteCommand {
	return DeleteCommand{
		ID: objc.ID(ptr),
	}
}




