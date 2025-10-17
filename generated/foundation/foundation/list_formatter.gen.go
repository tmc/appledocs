// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ListFormatter] class.
var ListFormatterClass objc.Class

func init() {
	ListFormatterClass = objc.GetClass("NSListFormatter")
}

type ListFormatter struct {
	objc.ID
}

func ListFormatterFrom(ptr unsafe.Pointer) ListFormatter {
	return ListFormatter{
		ID: objc.ID(ptr),
	}
}




