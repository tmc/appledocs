// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var FileHandleClass _FileHandleClass

func init() {
	FileHandleClass = _FileHandleClass{objc.GetClass("NSFileHandle")}
}

type _FileHandleClass struct {
	class objc.Class
}

type FileHandle struct {
	objc.ID
}

func FileHandleFrom(ptr unsafe.Pointer) FileHandle {
	return FileHandle{
		ID: objc.ID(ptr),
	}
}


// Disallows further access to the represented file or communications channel and signals end of file on communications channels that permit writing. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/closeFile()
func (f_ FileHandle) CloseFile() {
	objc.Send[objc.ID](f_.ID, objc.Sel("closeFile"))
}


