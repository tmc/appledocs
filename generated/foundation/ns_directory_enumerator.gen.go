// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DirectoryEnumerator] class.
var (
	DirectoryEnumeratorClass     _DirectoryEnumeratorClass
	DirectoryEnumeratorClassOnce sync.Once
)

func getDirectoryEnumeratorClass() _DirectoryEnumeratorClass {
	DirectoryEnumeratorClassOnce.Do(func() {
		DirectoryEnumeratorClass = _DirectoryEnumeratorClass{objc.GetClass("NSDirectoryEnumerator")}
	})
	return DirectoryEnumeratorClass
}

type _DirectoryEnumeratorClass struct {
	class objc.Class
}

// An interface definition for the [DirectoryEnumerator] class.
type IDirectoryEnumerator interface {
	IEnumerator
	DirectoryAttributes() FileAttributeKey
	SetDirectoryAttributes(value IFileAttributeKey)
	FileAttributes() FileAttributeKey
	SetFileAttributes(value IFileAttributeKey)
	IsEnumeratingDirectoryPostOrder() bool
	SetIsEnumeratingDirectoryPostOrder(value bool)
	Level() int
	SetLevel(value int)
}

// An object that enumerates the contents of a directory.
//
// You obtain a directory enumerator using ’s method. The enumeration provides the pathnames of all files and directories contained within that directory. These pathnames are relative to the directory. An enumeration is recursive, including the files of all subdirectories, and crosses device boundaries. An enumeration does not resolve symbolic links, or attempt to traverse symbolic links that point to directories.


// An object that enumerates the contents of a directory.
//
// [Full Topic]
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



// A dictionary with the attributes of the directory at which enumeration started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/directoryenumerator/directoryattributes

func (d_ DirectoryEnumerator) DirectoryAttributes() FileAttributeKey {
	rv := objc.Send[FileAttributeKey](d_.ID, objc.Sel("directoryAttributes"))
	return rv
}


// A dictionary with the attributes of the directory at which enumeration started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/directoryenumerator/directoryattributes

func (d_ DirectoryEnumerator) SetDirectoryAttributes(value IFileAttributeKey) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDirectoryAttributes:"), value)
}


// A dictionary with the attributes of the most recently returned file or subdirectory (as referenced by the pathname).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/directoryenumerator/fileattributes

func (d_ DirectoryEnumerator) FileAttributes() FileAttributeKey {
	rv := objc.Send[FileAttributeKey](d_.ID, objc.Sel("fileAttributes"))
	return rv
}


// A dictionary with the attributes of the most recently returned file or subdirectory (as referenced by the pathname).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/directoryenumerator/fileattributes

func (d_ DirectoryEnumerator) SetFileAttributes(value IFileAttributeKey) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileAttributes:"), value)
}


//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/directoryenumerator/isenumeratingdirectorypostorder

func (d_ DirectoryEnumerator) IsEnumeratingDirectoryPostOrder() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isEnumeratingDirectoryPostOrder"))
	return rv
}


//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/directoryenumerator/isenumeratingdirectorypostorder

func (d_ DirectoryEnumerator) SetIsEnumeratingDirectoryPostOrder(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsEnumeratingDirectoryPostOrder:"), value)
}


// The number of levels deep the current object is in the directory hierarchy being enumerated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/directoryenumerator/level

func (d_ DirectoryEnumerator) Level() int {
	rv := objc.Send[int](d_.ID, objc.Sel("level"))
	return rv
}


// The number of levels deep the current object is in the directory hierarchy being enumerated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/directoryenumerator/level

func (d_ DirectoryEnumerator) SetLevel(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLevel:"), value)
}



