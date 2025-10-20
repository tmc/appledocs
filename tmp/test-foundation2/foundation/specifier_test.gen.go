// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var specifierTestClass _SpecifierTestClass

func init() {
	specifierTestClass = _SpecifierTestClass{objc.GetClass("NSSpecifierTest")}
}

type _SpecifierTestClass struct {
	class objc.Class
}

type SpecifierTest struct {
	objc.ID
}

func SpecifierTestFrom(ptr unsafe.Pointer) SpecifierTest {
	return SpecifierTest{
		ID: objc.ID(ptr),
	}
}




