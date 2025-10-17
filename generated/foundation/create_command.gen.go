// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CreateCommand] class.
var CreateCommandClass = _CreateCommandClass{objc.GetClass("NSCreateCommand")}

type _CreateCommandClass struct {
	class objc.Class
}

type CreateCommand struct {
	objc.ID
}

func CreateCommandFrom(ptr unsafe.Pointer) CreateCommand {
	return CreateCommand{
		ID: objc.ID(ptr),
	}
}




