// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CloseCommand] class.
var CloseCommandClass objc.Class

func init() {
	CloseCommandClass = objc.GetClass("NSCloseCommand")
}

type CloseCommand struct {
	objc.ID
}

func CloseCommandFrom(ptr unsafe.Pointer) CloseCommand {
	return CloseCommand{
		ID: objc.ID(ptr),
	}
}



