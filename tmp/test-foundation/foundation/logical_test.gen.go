// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var LogicalTestClass _LogicalTestClass

func init() {
	LogicalTestClass = _LogicalTestClass{objc.GetClass("NSLogicalTest")}
}

type _LogicalTestClass struct {
	class objc.Class
}

type LogicalTest struct {
	objc.ID
}

func LogicalTestFrom(ptr unsafe.Pointer) LogicalTest {
	return LogicalTest{
		ID: objc.ID(ptr),
	}
}




