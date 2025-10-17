// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MassFormatter] class.
var MassFormatterClass objc.Class

func init() {
	MassFormatterClass = objc.GetClass("NSMassFormatter")
}

type MassFormatter struct {
	objc.ID
}

func MassFormatterFrom(ptr unsafe.Pointer) MassFormatter {
	return MassFormatter{
		ID: objc.ID(ptr),
	}
}




