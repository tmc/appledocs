// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileWrapper] class.
var fileWrapperClass = _FileWrapperClass{objc.GetClass("NSFileWrapper")}

type _FileWrapperClass struct {
	class objc.Class
}

// A representation of a node (a file, directory, or symbolic link) in the file system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper

type FileWrapper struct {
	objectivec.Object
}

// FileWrapperFrom constructs a [FileWrapper] from an unsafe.Pointer.
//
// A representation of a node (a file, directory, or symbolic link) in the file system.
func FileWrapperFrom(ptr unsafe.Pointer) FileWrapper {
	return FileWrapper{objectivec.Object{objc.ID(ptr)}}
}



