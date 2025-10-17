// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RelativeDateTimeFormatter] class.
var RelativeDateTimeFormatterClass objc.Class

func init() {
	RelativeDateTimeFormatterClass = objc.GetClass("NSRelativeDateTimeFormatter")
}

type RelativeDateTimeFormatter struct {
	objc.ID
}

func RelativeDateTimeFormatterFrom(ptr unsafe.Pointer) RelativeDateTimeFormatter {
	return RelativeDateTimeFormatter{
		ID: objc.ID(ptr),
	}
}



