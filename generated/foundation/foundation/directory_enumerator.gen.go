// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DirectoryEnumerator] class.
var DirectoryEnumeratorClass objc.Class

func init() {
	DirectoryEnumeratorClass = objc.GetClass("NSDirectoryEnumerator")
}

type DirectoryEnumerator struct {
	objc.ID
}

func DirectoryEnumeratorFrom(ptr unsafe.Pointer) DirectoryEnumerator {
	return DirectoryEnumerator{
		ID: objc.ID(ptr),
	}
}




