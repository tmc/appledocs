// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FilePromiseReceiver] class.
var FilePromiseReceiverClass objc.Class

func init() {
	FilePromiseReceiverClass = objc.GetClass("NSFilePromiseReceiver")
}

type FilePromiseReceiver struct {
	objc.ID
}

func FilePromiseReceiverFrom(ptr unsafe.Pointer) FilePromiseReceiver {
	return FilePromiseReceiver{
		ID: objc.ID(ptr),
	}
}




