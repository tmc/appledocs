// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LogicalTest] class.
var LogicalTestClass objc.Class

func init() {
	LogicalTestClass = objc.GetClass("NSLogicalTest")
}

type LogicalTest struct {
	objc.ID
}

func LogicalTestFrom(ptr unsafe.Pointer) LogicalTest {
	return LogicalTest{
		ID: objc.ID(ptr),
	}
}



