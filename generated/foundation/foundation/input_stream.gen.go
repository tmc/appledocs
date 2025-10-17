// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [InputStream] class.
var InputStreamClass objc.Class

func init() {
	InputStreamClass = objc.GetClass("NSInputStream")
}

type InputStream struct {
	objc.ID
}

func InputStreamFrom(ptr unsafe.Pointer) InputStream {
	return InputStream{
		ID: objc.ID(ptr),
	}
}




