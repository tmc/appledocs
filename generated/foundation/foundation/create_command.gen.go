// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CreateCommand] class.
var CreateCommandClass objc.Class

func init() {
	CreateCommandClass = objc.GetClass("NSCreateCommand")
}

type CreateCommand struct {
	objc.ID
}

func CreateCommandFrom(ptr unsafe.Pointer) CreateCommand {
	return CreateCommand{
		ID: objc.ID(ptr),
	}
}




