// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FileHandle] class.
var FileHandleClass objc.Class

func init() {
	FileHandleClass = objc.GetClass("NSFileHandle")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileHandle/closeFile()
func (f_ FileHandle) CloseFile() {
	sel := objc.RegisterName("closeFile")
	f_.ID.Send(sel)
}

