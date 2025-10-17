// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FilePromiseReceiver] class.
var filePromiseReceiverClass = _FilePromiseReceiverClass{objc.GetClass("NSFilePromiseReceiver")}

type _FilePromiseReceiverClass struct {
	class objc.Class
}

// An object that receives a file promise from the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver

type FilePromiseReceiver struct {
	objectivec.Object
}

// FilePromiseReceiverFrom constructs a [FilePromiseReceiver] from an unsafe.Pointer.
//
// An object that receives a file promise from the pasteboard.
func FilePromiseReceiverFrom(ptr unsafe.Pointer) FilePromiseReceiver {
	return FilePromiseReceiver{objectivec.Object{objc.ID(ptr)}}
}



