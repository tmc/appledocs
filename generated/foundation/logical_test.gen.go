// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LogicalTest] class.
var logicalTestClass = _LogicalTestClass{objc.GetClass("NSLogicalTest")}

type _LogicalTestClass struct {
	class objc.Class
}

// An interface definition for the [LogicalTest] class.
type ILogicalTest interface {
	IScriptWhoseTest
}

// The logical combination of one or more specifier tests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogicalTest

type LogicalTest struct {
	ScriptWhoseTest
}

// LogicalTestFrom constructs a [LogicalTest] from an unsafe.Pointer.
//
// The logical combination of one or more specifier tests.
func LogicalTestFrom(ptr unsafe.Pointer) LogicalTest {
	return LogicalTest{
		ScriptWhoseTest: ScriptWhoseTestFrom(ptr),
	}
}



