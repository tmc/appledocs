// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var setCommandClass _SetCommandClass

func init() {
	setCommandClass = _SetCommandClass{objc.GetClass("NSSetCommand")}
}

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




