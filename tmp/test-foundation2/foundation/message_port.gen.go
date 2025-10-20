// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var messagePortClass _MessagePortClass

func init() {
	messagePortClass = _MessagePortClass{objc.GetClass("NSMessagePort")}
}

type _MessagePortClass struct {
	class objc.Class
}

type MessagePort struct {
	objc.ID
}

func MessagePortFrom(ptr unsafe.Pointer) MessagePort {
	return MessagePort{
		ID: objc.ID(ptr),
	}
}




