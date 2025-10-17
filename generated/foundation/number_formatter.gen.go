// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NumberFormatter] class.
var NumberFormatterClass objc.Class

func init() {
	NumberFormatterClass = objc.GetClass("NSNumberFormatter")
}

type NumberFormatter struct {
	objc.ID
}

func NumberFormatterFrom(ptr unsafe.Pointer) NumberFormatter {
	return NumberFormatter{
		ID: objc.ID(ptr),
	}
}



