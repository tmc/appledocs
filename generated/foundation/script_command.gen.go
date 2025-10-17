// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScriptCommand] class.
var ScriptCommandClass objc.Class

func init() {
	ScriptCommandClass = objc.GetClass("NSScriptCommand")
}

type ScriptCommand struct {
	objc.ID
}

func ScriptCommandFrom(ptr unsafe.Pointer) ScriptCommand {
	return ScriptCommand{
		ID: objc.ID(ptr),
	}
}



