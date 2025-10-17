// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SetCommand] class.
var SetCommandClass objc.Class

func init() {
	SetCommandClass = objc.GetClass("NSSetCommand")
}

type SetCommand struct {
	objc.ID
}

func SetCommandFrom(ptr unsafe.Pointer) SetCommand {
	return SetCommand{
		ID: objc.ID(ptr),
	}
}



