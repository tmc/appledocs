// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var RelativeDateTimeFormatterClass _RelativeDateTimeFormatterClass

func init() {
	RelativeDateTimeFormatterClass = _RelativeDateTimeFormatterClass{objc.GetClass("NSRelativeDateTimeFormatter")}
}

type _RelativeDateTimeFormatterClass struct {
	class objc.Class
}

type RelativeDateTimeFormatter struct {
	objc.ID
}

func RelativeDateTimeFormatterFrom(ptr unsafe.Pointer) RelativeDateTimeFormatter {
	return RelativeDateTimeFormatter{
		ID: objc.ID(ptr),
	}
}




