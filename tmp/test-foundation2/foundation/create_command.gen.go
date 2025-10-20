// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var createCommandClass _CreateCommandClass

func init() {
	createCommandClass = _CreateCommandClass{objc.GetClass("NSCreateCommand")}
}

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




