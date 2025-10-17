
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FilePromiseReceiver] class.
var FilePromiseReceiverClass _FilePromiseReceiverClass

func init() {
	FilePromiseReceiverClass = _FilePromiseReceiverClass{objc.GetClass("NSFilePromiseReceiver")}
}

type _FilePromiseReceiverClass struct {
	objc.Class
}

// An interface definition for the [FilePromiseReceiver] class.
type IFilePromiseReceiver interface {
	ID() objc.ID
}

type FilePromiseReceiver struct {
	id objc.ID
}

func FilePromiseReceiverFrom(ptr unsafe.Pointer) FilePromiseReceiver {
	return FilePromiseReceiver{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ FilePromiseReceiver) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _FilePromiseReceiverClass) Alloc() FilePromiseReceiver {
	rv := objc.Send[FilePromiseReceiver](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _FilePromiseReceiverClass) New() FilePromiseReceiver {
	rv := objc.Send[FilePromiseReceiver](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewFilePromiseReceiver creates and returns a new initialized instance.
func NewFilePromiseReceiver() FilePromiseReceiver {
	return FilePromiseReceiverClass.New()
}

// Init initializes the instance.
func (f_ FilePromiseReceiver) Init() FilePromiseReceiver {
	rv := objc.Send[FilePromiseReceiver](f_.ID(), selInit)
	return rv
}
