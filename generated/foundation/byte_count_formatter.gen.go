// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ByteCountFormatter] class.
var ByteCountFormatterClass objc.Class

func init() {
	ByteCountFormatterClass = objc.GetClass("NSByteCountFormatter")
}

type ByteCountFormatter struct {
	objc.ID
}

func ByteCountFormatterFrom(ptr unsafe.Pointer) ByteCountFormatter {
	return ByteCountFormatter{
		ID: objc.ID(ptr),
	}
}



