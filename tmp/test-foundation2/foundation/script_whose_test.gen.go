// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var scriptWhoseTestClass _ScriptWhoseTestClass

func init() {
	scriptWhoseTestClass = _ScriptWhoseTestClass{objc.GetClass("NSScriptWhoseTest")}
}

type _ScriptWhoseTestClass struct {
	class objc.Class
}

type ScriptWhoseTest struct {
	objc.ID
}

func ScriptWhoseTestFrom(ptr unsafe.Pointer) ScriptWhoseTest {
	return ScriptWhoseTest{
		ID: objc.ID(ptr),
	}
}




