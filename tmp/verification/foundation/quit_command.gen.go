// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var quitCommandClass _QuitCommandClass

func init() {
	quitCommandClass = _QuitCommandClass{objc.GetClass("NSQuitCommand")}
}

type _QuitCommandClass struct {
	class objc.Class
}

type QuitCommand struct {
	objc.ID
}

func QuitCommandFrom(ptr unsafe.Pointer) QuitCommand {
	return QuitCommand{
		ID: objc.ID(ptr),
	}
}




