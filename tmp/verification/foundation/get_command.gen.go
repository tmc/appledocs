// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var getCommandClass _GetCommandClass

func init() {
	getCommandClass = _GetCommandClass{objc.GetClass("NSGetCommand")}
}

type _GetCommandClass struct {
	class objc.Class
}

type GetCommand struct {
	objc.ID
}

func GetCommandFrom(ptr unsafe.Pointer) GetCommand {
	return GetCommand{
		ID: objc.ID(ptr),
	}
}




