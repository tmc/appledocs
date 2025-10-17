// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SocketPort] class.
var SocketPortClass objc.Class

func init() {
	SocketPortClass = objc.GetClass("NSSocketPort")
}

type SocketPort struct {
	objc.ID
}

func SocketPortFrom(ptr unsafe.Pointer) SocketPort {
	return SocketPort{
		ID: objc.ID(ptr),
	}
}




