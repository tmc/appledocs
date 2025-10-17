// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GetCommand] class.
var GetCommandClass objc.Class

func init() {
	GetCommandClass = objc.GetClass("NSGetCommand")
}

type GetCommand struct {
	objc.ID
}

func GetCommandFrom(ptr unsafe.Pointer) GetCommand {
	return GetCommand{
		ID: objc.ID(ptr),
	}
}




