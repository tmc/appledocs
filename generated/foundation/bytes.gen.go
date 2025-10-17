// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [bytes] class.
var bytesClass objc.Class

func init() {
	bytesClass = objc.GetClass("bytes")
}

type bytes struct {
	objc.ID
}

func bytesFrom(ptr unsafe.Pointer) bytes {
	return bytes{
		ID: objc.ID(ptr),
	}
}



