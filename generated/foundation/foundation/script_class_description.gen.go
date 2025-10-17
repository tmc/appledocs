// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScriptClassDescription] class.
var ScriptClassDescriptionClass objc.Class

func init() {
	ScriptClassDescriptionClass = objc.GetClass("NSScriptClassDescription")
}

type ScriptClassDescription struct {
	objc.ID
}

func ScriptClassDescriptionFrom(ptr unsafe.Pointer) ScriptClassDescription {
	return ScriptClassDescription{
		ID: objc.ID(ptr),
	}
}




