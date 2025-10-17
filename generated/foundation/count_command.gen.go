// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CountCommand] class.
var CountCommandClass objc.Class

func init() {
	CountCommandClass = objc.GetClass("NSCountCommand")
}

type CountCommand struct {
	objc.ID
}

func CountCommandFrom(ptr unsafe.Pointer) CountCommand {
	return CountCommand{
		ID: objc.ID(ptr),
	}
}



