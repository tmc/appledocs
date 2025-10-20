// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DirectoryEnumerator] class.
var (
	directoryEnumeratorClass     _DirectoryEnumeratorClass
	directoryEnumeratorClassOnce sync.Once
)

func getDirectoryEnumeratorClass() _DirectoryEnumeratorClass {
	directoryEnumeratorClassOnce.Do(func() {
		directoryEnumeratorClass = _DirectoryEnumeratorClass{objc.GetClass("NSDirectoryEnumerator")}
	})
	return directoryEnumeratorClass
}

type _DirectoryEnumeratorClass struct {
	class objc.Class
}

// An interface definition for the [DirectoryEnumerator] class.
type IDirectoryEnumerator interface {
	IEnumerator
}

// An object that enumerates the contents of a directory.
//
// You obtain a directory enumerator using ’s method. The enumeration provides the pathnames of all files and directories contained within that directory. These pathnames are relative to the directory. An enumeration is recursive, including the files of all subdirectories, and crosses device boundaries. An enumeration does not resolve symbolic links, or attempt to traverse symbolic links that point to directories.
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

// Alloc allocates a new instance without initialization.
func (dc _DirectoryEnumeratorClass) Alloc() DirectoryEnumerator {
	rv := objc.Send[DirectoryEnumerator](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DirectoryEnumeratorClass) New() DirectoryEnumerator {
	rv := objc.Send[DirectoryEnumerator](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DirectoryEnumerator) Init() DirectoryEnumerator {
	rv := objc.Send[DirectoryEnumerator](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DirectoryEnumerator) Autorelease() DirectoryEnumerator {
	rv := objc.Send[DirectoryEnumerator](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDirectoryEnumerator creates a new DirectoryEnumerator instance.
func NewDirectoryEnumerator() DirectoryEnumerator {
	return getDirectoryEnumeratorClass().New()
}




