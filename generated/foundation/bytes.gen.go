// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [bytes] class.
var bytesClass = _bytesClass{objc.GetClass("bytes")}

type _bytesClass struct {
	class objc.Class
}

type bytes struct {
	objc.ID
}

func bytesFrom(ptr unsafe.Pointer) bytes {
	return bytes{
		ID: objc.ID(ptr),
	}
}




