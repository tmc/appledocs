// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MessagePort] class.
var MessagePortClass objc.Class

func init() {
	MessagePortClass = objc.GetClass("NSMessagePort")
}

type MessagePort struct {
	objc.ID
}

func MessagePortFrom(ptr unsafe.Pointer) MessagePort {
	return MessagePort{
		ID: objc.ID(ptr),
	}
}



