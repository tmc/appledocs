// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OutputStream] class.
var OutputStreamClass objc.Class

func init() {
	OutputStreamClass = objc.GetClass("NSOutputStream")
}

type OutputStream struct {
	objc.ID
}

func OutputStreamFrom(ptr unsafe.Pointer) OutputStream {
	return OutputStream{
		ID: objc.ID(ptr),
	}
}




