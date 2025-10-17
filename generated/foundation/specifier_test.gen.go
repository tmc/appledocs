// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SpecifierTest] class.
var SpecifierTestClass objc.Class

func init() {
	SpecifierTestClass = objc.GetClass("NSSpecifierTest")
}

type SpecifierTest struct {
	objc.ID
}

func SpecifierTestFrom(ptr unsafe.Pointer) SpecifierTest {
	return SpecifierTest{
		ID: objc.ID(ptr),
	}
}



