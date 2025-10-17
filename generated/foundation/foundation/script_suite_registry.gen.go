// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScriptSuiteRegistry] class.
var ScriptSuiteRegistryClass objc.Class

func init() {
	ScriptSuiteRegistryClass = objc.GetClass("NSScriptSuiteRegistry")
}

type ScriptSuiteRegistry struct {
	objc.ID
}

func ScriptSuiteRegistryFrom(ptr unsafe.Pointer) ScriptSuiteRegistry {
	return ScriptSuiteRegistry{
		ID: objc.ID(ptr),
	}
}




