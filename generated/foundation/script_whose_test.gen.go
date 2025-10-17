// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScriptWhoseTest] class.
var ScriptWhoseTestClass objc.Class

func init() {
	ScriptWhoseTestClass = objc.GetClass("NSScriptWhoseTest")
}

type ScriptWhoseTest struct {
	objc.ID
}

func ScriptWhoseTestFrom(ptr unsafe.Pointer) ScriptWhoseTest {
	return ScriptWhoseTest{
		ID: objc.ID(ptr),
	}
}



