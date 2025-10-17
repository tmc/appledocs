// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ISO8601DateFormatter] class.
var ISO8601DateFormatterClass = _ISO8601DateFormatterClass{objc.GetClass("NSISO8601DateFormatter")}

type _ISO8601DateFormatterClass struct {
	class objc.Class
}

type ISO8601DateFormatter struct {
	objc.ID
}

func ISO8601DateFormatterFrom(ptr unsafe.Pointer) ISO8601DateFormatter {
	return ISO8601DateFormatter{
		ID: objc.ID(ptr),
	}
}




