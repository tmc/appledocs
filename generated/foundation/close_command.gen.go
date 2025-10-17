// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CloseCommand] class.
var CloseCommandClass = _CloseCommandClass{objc.GetClass("NSCloseCommand")}

type _CloseCommandClass struct {
	class objc.Class
}

type CloseCommand struct {
	objc.ID
}

func CloseCommandFrom(ptr unsafe.Pointer) CloseCommand {
	return CloseCommand{
		ID: objc.ID(ptr),
	}
}




