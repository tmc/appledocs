// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SpecifierTest] class.
var specifierTestClass = _SpecifierTestClass{objc.GetClass("NSSpecifierTest")}

type _SpecifierTestClass struct {
	class objc.Class
}

// An interface definition for the [SpecifierTest] class.
type ISpecifierTest interface {
	IScriptWhoseTest
}

// A comparison between an object specifier and a test object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSpecifierTest

type SpecifierTest struct {
	ScriptWhoseTest
}

// SpecifierTestFrom constructs a [SpecifierTest] from an unsafe.Pointer.
//
// A comparison between an object specifier and a test object.
func SpecifierTestFrom(ptr unsafe.Pointer) SpecifierTest {
	return SpecifierTest{
		ScriptWhoseTest: ScriptWhoseTestFrom(ptr),
	}
}



