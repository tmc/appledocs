// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileWrapper] class.
var (
	FileWrapperClass     _FileWrapperClass
	FileWrapperClassOnce sync.Once
)

func getFileWrapperClass() _FileWrapperClass {
	FileWrapperClassOnce.Do(func() {
		FileWrapperClass = _FileWrapperClass{objc.GetClass("NSFileWrapper")}
	})
	return FileWrapperClass
}

type _FileWrapperClass struct {
	class objc.Class
}

// An interface definition for the [FileWrapper] class.
type IFileWrapper interface {
	objectivec.IObject
	MatchesContentsOfURL(url unsafe.Pointer) bool
	NeedsToBeUpdatedFromPath(path string) bool
	SymbolicLinkDestination() unsafe.Pointer
	WriteToFileAtomicallyUpdateFilenames(path string, atomicFlag bool, updateFilenamesFlag bool) bool
}

// A representation of a node (a file, directory, or symbolic link) in the file system.
//
// The class provides access to the attributes and contents of file system nodes. A file system node is a file, directory, or symbolic link. Instances of this class are known as file wrappers. File wrappers represent a file system node as an object that can be displayed as an image (and possibly edited in place), saved to the file system, or transmitted to another application. There are three types of file wrappers: Regular-file file wrapper: Represents a regular file. Directory file wrapper: Represents a directory. Symbolic-link file wrapper: Represents a symbolic link. A file wrapper has these attributes: Filename. Name of the file system node the file wrapper represents. file-system attributes. See for information on the contents of the dictionary. Regular-file contents. Applicable only to regular-file file wrappers. File wrappers. Applicable only to directory file wrappers. Destination node. Applicable only to symbolic-link file wrappers.
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

// Alloc allocates a new instance without initialization.
func (fc _FileWrapperClass) Alloc() FileWrapper {
	rv := objc.Send[FileWrapper](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileWrapperClass) New() FileWrapper {
	rv := objc.Send[FileWrapper](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileWrapper) Init() FileWrapper {
	rv := objc.Send[FileWrapper](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileWrapper) Autorelease() FileWrapper {
	rv := objc.Send[FileWrapper](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileWrapper creates a new FileWrapper instance.
func NewFileWrapper() FileWrapper {
	return getFileWrapperClass().New()
}


// Indicates whether the contents of a file wrapper matches a directory, regular file, or symbolic link on disk.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/matchesContents(of:)
func (f_ FileWrapper) MatchesContentsOfURL(url unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("matchesContentsOfURL:"), url)
	return rv
}

// Indicates whether the file wrapper needs to be updated to match a given file-system node.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/needsToBeUpdated(fromPath:)
func (f_ FileWrapper) NeedsToBeUpdatedFromPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("needsToBeUpdatedFromPath:"), objc.String(path))
	return rv
}

// Provides the pathname referenced by the file wrapper object, which must be a symbolic-link file wrapper.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/symbolicLinkDestination()
func (f_ FileWrapper) SymbolicLinkDestination() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("symbolicLinkDestination"))
	return rv
}

// Writes a file wrapper’s contents to a given file-system node.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/write(toFile:atomically:updateFilenames:)
func (f_ FileWrapper) WriteToFileAtomicallyUpdateFilenames(path string, atomicFlag bool, updateFilenamesFlag bool) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("writeToFile:atomically:updateFilenames:"), objc.String(path), atomicFlag, updateFilenamesFlag)
	return rv
}

// The filename of the file wrapper object
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/filename
func (f_ FileWrapper) Filename() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("filename"))
	return rv
}


// SetFilename sets the value of the filename property.
// The filename of the file wrapper object

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/filename
func (f_ FileWrapper) SetFilename(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFilename:"), value)
}
// The contents of the file wrapper as an opaque data object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/serializedRepresentation
func (f_ FileWrapper) SerializedRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("serializedRepresentation"))
	return rv
}



