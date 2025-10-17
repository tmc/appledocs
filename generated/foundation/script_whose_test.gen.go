// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptWhoseTest] class.
var scriptWhoseTestClass = _ScriptWhoseTestClass{objc.GetClass("NSScriptWhoseTest")}

type _ScriptWhoseTestClass struct {
	class objc.Class
}

// An interface definition for the [ScriptWhoseTest] class.
type IScriptWhoseTest interface {
	objectivec.IObject
}

// An abstract class that provides the basis for testing specifiers one at a time or in groups. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptWhoseTest

type ScriptWhoseTest struct {
	objectivec.Object
}

// ScriptWhoseTestFrom constructs a [ScriptWhoseTest] from an unsafe.Pointer.
//
// An abstract class that provides the basis for testing specifiers one at a time or in groups.
func ScriptWhoseTestFrom(ptr unsafe.Pointer) ScriptWhoseTest {
	return ScriptWhoseTest{objectivec.Object{objc.ID(ptr)}}
}



