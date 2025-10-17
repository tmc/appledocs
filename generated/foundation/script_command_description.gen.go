// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScriptCommandDescription] class.
var ScriptCommandDescriptionClass objc.Class

func init() {
	ScriptCommandDescriptionClass = objc.GetClass("NSScriptCommandDescription")
}

type ScriptCommandDescription struct {
	objc.ID
}

func ScriptCommandDescriptionFrom(ptr unsafe.Pointer) ScriptCommandDescription {
	return ScriptCommandDescription{
		ID: objc.ID(ptr),
	}
}



