// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MassFormatter] class.
var MassFormatterClass = _MassFormatterClass{objc.GetClass("NSMassFormatter")}

type _MassFormatterClass struct {
	class objc.Class
}

type MassFormatter struct {
	objc.ID
}

func MassFormatterFrom(ptr unsafe.Pointer) MassFormatter {
	return MassFormatter{
		ID: objc.ID(ptr),
	}
}




