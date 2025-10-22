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
	MatchesContentsOfURL(url IURL) bool
	NeedsToBeUpdatedFromPath(path string) bool
	SymbolicLinkDestination() String
	WriteToFileAtomicallyUpdateFilenames(path string, atomicFlag bool, updateFilenamesFlag bool) bool
	Filename() string
	SetFilename(value string)
	SerializedRepresentation() NSData
	FileAttributes() string
	SetFileAttributes(value string)
	FileWrappers() NSFileWrapper
	SetFileWrappers(value IFileWrapper)
	IsDirectory() bool
	SetIsDirectory(value bool)
	IsRegularFile() bool
	SetIsRegularFile(value bool)
	IsSymbolicLink() bool
	SetIsSymbolicLink(value bool)
	PreferredFilename() string
	SetPreferredFilename(value string)
	RegularFileContents() Data
	SetRegularFileContents(value IData)
	SymbolicLinkDestinationURL() URL
	SetSymbolicLinkDestinationURL(value IURL)
}

// A representation of a node (a file, directory, or symbolic link) in the file system.
//
// The class provides access to the attributes and contents of file system nodes. A file system node is a file, directory, or symbolic link. Instances of this class are known as file wrappers. File wrappers represent a file system node as an object that can be displayed as an image (and possibly edited in place), saved to the file system, or transmitted to another application. There are three types of file wrappers: Regular-file file wrapper: Represents a regular file. Directory file wrapper: Represents a directory. Symbolic-link file wrapper: Represents a symbolic link. A file wrapper has these attributes: Filename. Name of the file system node the file wrapper represents. file-system attributes. See for information on the contents of the dictionary. Regular-file contents. Applicable only to regular-file file wrappers. File wrappers. Applicable only to directory file wrappers. Destination node. Applicable only to symbolic-link file wrappers.


// A representation of a node (a file, directory, or symbolic link) in the file system.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/matchesContents(of:)

func (f_ FileWrapper) MatchesContentsOfURL(url IURL) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("matchesContentsOfURL:"), url)
	return rv
}



// Indicates whether the file wrapper needs to be updated to match a given file-system node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/needsToBeUpdated(fromPath:)

func (f_ FileWrapper) NeedsToBeUpdatedFromPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("needsToBeUpdatedFromPath:"), objc.String(path))
	return rv
}



// Provides the pathname referenced by the file wrapper object, which must be a symbolic-link file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/symbolicLinkDestination()

func (f_ FileWrapper) SymbolicLinkDestination() String {
	rv := objc.Send[String](f_.ID, objc.Sel("symbolicLinkDestination"))
	return rv
}



// Writes a file wrapper’s contents to a given file-system node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/write(toFile:atomically:updateFilenames:)

func (f_ FileWrapper) WriteToFileAtomicallyUpdateFilenames(path string, atomicFlag bool, updateFilenamesFlag bool) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("writeToFile:atomically:updateFilenames:"), objc.String(path), atomicFlag, updateFilenamesFlag)
	return rv
}


// The filename of the file wrapper object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/filename

func (f_ FileWrapper) Filename() string {
	rv := objc.Send[string](f_.ID, objc.Sel("filename"))
	return rv
}


// The filename of the file wrapper object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/filename

func (f_ FileWrapper) SetFilename(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFilename:"), objc.String(value))
}


// The contents of the file wrapper as an opaque data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/serializedRepresentation

func (f_ FileWrapper) SerializedRepresentation() NSData {
	rv := objc.Send[NSData](f_.ID, objc.Sel("serializedRepresentation"))
	return rv
}


// A dictionary of file attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/fileattributes

func (f_ FileWrapper) FileAttributes() string {
	rv := objc.Send[string](f_.ID, objc.Sel("fileAttributes"))
	return rv
}


// A dictionary of file attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/fileattributes

func (f_ FileWrapper) SetFileAttributes(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileAttributes:"), objc.String(value))
}


// The file wrappers contained by a directory file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/filewrappers

func (f_ FileWrapper) FileWrappers() NSFileWrapper {
	rv := objc.Send[NSFileWrapper](f_.ID, objc.Sel("fileWrappers"))
	return rv
}


// The file wrappers contained by a directory file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/filewrappers

func (f_ FileWrapper) SetFileWrappers(value IFileWrapper) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileWrappers:"), value)
}


// This property contains a boolean value indicating whether the file wrapper is a directory file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/isdirectory

func (f_ FileWrapper) IsDirectory() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDirectory"))
	return rv
}


// This property contains a boolean value indicating whether the file wrapper is a directory file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/isdirectory

func (f_ FileWrapper) SetIsDirectory(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsDirectory:"), value)
}


// This property contains a boolean value that indicates whether the file wrapper object is a regular-file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/isregularfile

func (f_ FileWrapper) IsRegularFile() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isRegularFile"))
	return rv
}


// This property contains a boolean value that indicates whether the file wrapper object is a regular-file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/isregularfile

func (f_ FileWrapper) SetIsRegularFile(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsRegularFile:"), value)
}


// A boolean that indicates whether the file wrapper object is a symbolic-link file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/issymboliclink

func (f_ FileWrapper) IsSymbolicLink() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isSymbolicLink"))
	return rv
}


// A boolean that indicates whether the file wrapper object is a symbolic-link file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/issymboliclink

func (f_ FileWrapper) SetIsSymbolicLink(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsSymbolicLink:"), value)
}


// The preferred filename for the file wrapper object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/preferredfilename

func (f_ FileWrapper) PreferredFilename() string {
	rv := objc.Send[string](f_.ID, objc.Sel("preferredFilename"))
	return rv
}


// The preferred filename for the file wrapper object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/preferredfilename

func (f_ FileWrapper) SetPreferredFilename(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPreferredFilename:"), objc.String(value))
}


// The contents of the file-system node associated with a regular-file file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/regularfilecontents

func (f_ FileWrapper) RegularFileContents() Data {
	rv := objc.Send[Data](f_.ID, objc.Sel("regularFileContents"))
	return rv
}


// The contents of the file-system node associated with a regular-file file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/regularfilecontents

func (f_ FileWrapper) SetRegularFileContents(value IData) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRegularFileContents:"), value)
}


// The URL referenced by the file wrapper object, which must be a symbolic-link file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/symboliclinkdestinationurl

func (f_ FileWrapper) SymbolicLinkDestinationURL() URL {
	rv := objc.Send[URL](f_.ID, objc.Sel("symbolicLinkDestinationURL"))
	return rv
}


// The URL referenced by the file wrapper object, which must be a symbolic-link file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/symboliclinkdestinationurl

func (f_ FileWrapper) SetSymbolicLinkDestinationURL(value IURL) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSymbolicLinkDestinationURL:"), value)
}



