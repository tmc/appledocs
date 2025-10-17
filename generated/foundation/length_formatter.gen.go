// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LengthFormatter] class.
var LengthFormatterClass objc.Class

func init() {
	LengthFormatterClass = objc.GetClass("NSLengthFormatter")
}

type LengthFormatter struct {
	objc.ID
}

func LengthFormatterFrom(ptr unsafe.Pointer) LengthFormatter {
	return LengthFormatter{
		ID: objc.ID(ptr),
	}
}



