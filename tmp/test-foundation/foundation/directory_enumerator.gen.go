// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var DirectoryEnumeratorClass _DirectoryEnumeratorClass

func init() {
	DirectoryEnumeratorClass = _DirectoryEnumeratorClass{objc.GetClass("NSDirectoryEnumerator")}
}

type _DirectoryEnumeratorClass struct {
	class objc.Class
}

type DirectoryEnumerator struct {
	objc.ID
}

func DirectoryEnumeratorFrom(ptr unsafe.Pointer) DirectoryEnumerator {
	return DirectoryEnumerator{
		ID: objc.ID(ptr),
	}
}




