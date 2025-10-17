// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SetCommand] class.
var SetCommandClass = _SetCommandClass{objc.GetClass("NSSetCommand")}

type _SetCommandClass struct {
	class objc.Class
}

type SetCommand struct {
	objc.ID
}

func SetCommandFrom(ptr unsafe.Pointer) SetCommand {
	return SetCommand{
		ID: objc.ID(ptr),
	}
}




