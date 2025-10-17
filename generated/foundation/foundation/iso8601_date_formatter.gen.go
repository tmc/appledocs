// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ISO8601DateFormatter] class.
var ISO8601DateFormatterClass objc.Class

func init() {
	ISO8601DateFormatterClass = objc.GetClass("NSISO8601DateFormatter")
}

type ISO8601DateFormatter struct {
	objc.ID
}

func ISO8601DateFormatterFrom(ptr unsafe.Pointer) ISO8601DateFormatter {
	return ISO8601DateFormatter{
		ID: objc.ID(ptr),
	}
}




