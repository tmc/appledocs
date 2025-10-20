// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var scriptCommandDescriptionClass _ScriptCommandDescriptionClass

func init() {
	scriptCommandDescriptionClass = _ScriptCommandDescriptionClass{objc.GetClass("NSScriptCommandDescription")}
}

type _ScriptCommandDescriptionClass struct {
	class objc.Class
}

type ScriptCommandDescription struct {
	objc.ID
}

func ScriptCommandDescriptionFrom(ptr unsafe.Pointer) ScriptCommandDescription {
	return ScriptCommandDescription{
		ID: objc.ID(ptr),
	}
}




