// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var scriptClassDescriptionClass _ScriptClassDescriptionClass

func init() {
	scriptClassDescriptionClass = _ScriptClassDescriptionClass{objc.GetClass("NSScriptClassDescription")}
}

type _ScriptClassDescriptionClass struct {
	class objc.Class
}

type ScriptClassDescription struct {
	objc.ID
}

func ScriptClassDescriptionFrom(ptr unsafe.Pointer) ScriptClassDescription {
	return ScriptClassDescription{
		ID: objc.ID(ptr),
	}
}




