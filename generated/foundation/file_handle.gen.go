// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileHandle] class.
var fileHandleClass = _FileHandleClass{objc.GetClass("NSFileHandle")}

type _FileHandleClass struct {
	class objc.Class
}

// An interface definition for the [FileHandle] class.
type IFileHandle interface {
	objectivec.IObject
	CloseFile()
}

// An object-oriented wrapper for a file descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle

type FileHandle struct {
	objectivec.Object
}

// FileHandleFrom constructs a [FileHandle] from an unsafe.Pointer.
//
// An object-oriented wrapper for a file descriptor.
func FileHandleFrom(ptr unsafe.Pointer) FileHandle {
	return FileHandle{objectivec.Object{objc.ID(ptr)}}
}

// Disallows further access to the represented file or communications channel and signals end of file on communications channels that permit writing. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/closeFile()
func (f_ FileHandle) CloseFile() {
	objc.Send[objc.ID](f_.ID, objc.Sel("closeFile"))
}


