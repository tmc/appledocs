// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var scriptCommandClass _ScriptCommandClass

func init() {
	scriptCommandClass = _ScriptCommandClass{objc.GetClass("NSScriptCommand")}
}

type _ScriptCommandClass struct {
	class objc.Class
}

type ScriptCommand struct {
	objc.ID
}

func ScriptCommandFrom(ptr unsafe.Pointer) ScriptCommand {
	return ScriptCommand{
		ID: objc.ID(ptr),
	}
}




