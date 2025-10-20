// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ScriptSuiteRegistryClass _ScriptSuiteRegistryClass

func init() {
	ScriptSuiteRegistryClass = _ScriptSuiteRegistryClass{objc.GetClass("NSScriptSuiteRegistry")}
}

type _ScriptSuiteRegistryClass struct {
	class objc.Class
}

type ScriptSuiteRegistry struct {
	objc.ID
}

func ScriptSuiteRegistryFrom(ptr unsafe.Pointer) ScriptSuiteRegistry {
	return ScriptSuiteRegistry{
		ID: objc.ID(ptr),
	}
}




