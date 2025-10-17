// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InputStream] class.
var InputStreamClass = _InputStreamClass{objc.GetClass("NSInputStream")}

type _InputStreamClass struct {
	class objc.Class
}

type InputStream struct {
	objc.ID
}

func InputStreamFrom(ptr unsafe.Pointer) InputStream {
	return InputStream{
		ID: objc.ID(ptr),
	}
}




