// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [QuitCommand] class.
var QuitCommandClass objc.Class

func init() {
	QuitCommandClass = objc.GetClass("NSQuitCommand")
}

type QuitCommand struct {
	objc.ID
}

func QuitCommandFrom(ptr unsafe.Pointer) QuitCommand {
	return QuitCommand{
		ID: objc.ID(ptr),
	}
}



