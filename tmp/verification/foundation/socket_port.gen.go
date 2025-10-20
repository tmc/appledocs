// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var socketPortClass _SocketPortClass

func init() {
	socketPortClass = _SocketPortClass{objc.GetClass("NSSocketPort")}
}

type _SocketPortClass struct {
	class objc.Class
}

type SocketPort struct {
	objc.ID
}

func SocketPortFrom(ptr unsafe.Pointer) SocketPort {
	return SocketPort{
		ID: objc.ID(ptr),
	}
}




