// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DirectoryEnumerator] class.
var directoryEnumeratorClass = _DirectoryEnumeratorClass{objc.GetClass("NSDirectoryEnumerator")}

type _DirectoryEnumeratorClass struct {
	class objc.Class
}

// An interface definition for the [DirectoryEnumerator] class.
type IDirectoryEnumerator interface {
	IEnumerator
}

// An object that enumerates the contents of a directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator

type DirectoryEnumerator struct {
	Enumerator
}

// DirectoryEnumeratorFrom constructs a [DirectoryEnumerator] from an unsafe.Pointer.
//
// An object that enumerates the contents of a directory.
func DirectoryEnumeratorFrom(ptr unsafe.Pointer) DirectoryEnumerator {
	return DirectoryEnumerator{
		Enumerator: EnumeratorFrom(ptr),
	}
}



