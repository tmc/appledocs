// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFileWrapper */


/* debug [class_header]: Header for NSFileWrapper */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileWrapper */
// An interface definition for the [FileWrapper] class.
type IFileWrapper interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileWrapper */
	// properties:
	FileAttributes() IDictionary
	SetFileAttributes(value IDictionary)
	FileWrappers() IDictionary
	Filename() IString
	SetFilename(value IString)
	PreferredFilename() IString
	SetPreferredFilename(value IString)
	RegularFileContents() IData
	SerializedRepresentation() IData
	SymbolicLinkDestinationURL() IURL
	Icon() objectivec.IObject
	SetIcon(value objectivec.IObject)
	IsDirectory() bool
	SetIsDirectory(value bool)
	IsRegularFile() bool
	SetIsRegularFile(value bool)
	IsSymbolicLink() bool
	SetIsSymbolicLink(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileWrapper */
	// methods:
	AddFileWrapper(child IFileWrapper) IString
	AddRegularFileWithContentsPreferredFilename(data IData, fileName IString) IString
	KeyForFileWrapper(child IFileWrapper) IString
	MatchesContentsOfURL(url IURL) bool
	ReadFromURLOptionsError(url IURL, options FileWrapperReadingOptions, outError IError) bool
	RemoveFileWrapper(child IFileWrapper)
	WriteToURLOptionsOriginalContentsURLError(url IURL, options FileWrapperWritingOptions, originalContentsURL IURL, outError IError) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileWrapper */
// Alloc allocates a new instance without initialization.
func (fc _FileWrapperClass) Alloc() FileWrapper {
	rv := objc.Send[FileWrapper](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileWrapper */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileWrapper */

// Initializes the receiver as a directory file wrapper, with a given file-wrapper list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/init(directoryWithFileWrappers:)
func NewFileWrapperDirectoryWithFileWrappers(childrenByPreferredName IDictionary) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[FileWrapper](instance.ID, objc.Sel("initDirectoryWithFileWrappers:"), childrenByPreferredName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileWrapperDirectoryWithFileWrappers */


// Initializes the receiver as a regular-file file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/init(regularFileWithContents:)
func NewFileWrapperRegularFileWithContents(contents IData) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[FileWrapper](instance.ID, objc.Sel("initRegularFileWithContents:"), contents)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileWrapperRegularFileWithContents */


// Initializes the receiver as a symbolic-link file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/init(symbolicLinkWithDestination:)
func NewFileWrapperSymbolicLinkWithDestination(path IString) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[FileWrapper](instance.ID, objc.Sel("initSymbolicLinkWithDestination:"), path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileWrapperSymbolicLinkWithDestination */


// Initializes the receiver as a symbolic-link file wrapper that links to a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/init(symbolicLinkWithDestinationURL:)
func NewFileWrapperSymbolicLinkWithDestinationURL(url IURL) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[FileWrapper](instance.ID, objc.Sel("initSymbolicLinkWithDestinationURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileWrapperSymbolicLinkWithDestinationURL */


// Initializes a file wrapper instance whose kind is determined by the type of file-system node located by the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/init(path:)
func NewFileWrapperWithPath(path IString) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[FileWrapper](instance.ID, objc.Sel("initWithPath:"), path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileWrapperWithPath */


// Initializes the receiver as a regular-file file wrapper from given serialized data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/init(serializedRepresentation:)
func NewFileWrapperWithSerializedRepresentation(serializeRepresentation IData) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[FileWrapper](instance.ID, objc.Sel("initWithSerializedRepresentation:"), serializeRepresentation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileWrapperWithSerializedRepresentation */


// Initializes a file wrapper instance whose kind is determined by the type of file-system node located by the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/init(url:options:)
func NewFileWrapperWithURLOptionsError(url IURL, options FileWrapperReadingOptions, outError IError) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[FileWrapper](instance.ID, objc.Sel("initWithURL:options:error:"), url, options, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileWrapperWithURLOptionsError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileWrapper */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileWrapper */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileWrapper */

// Adds a child file wrapper to the receiver, which must be a directory file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/addFileWrapper(_:)
func (f_ FileWrapper) AddFileWrapper(child IFileWrapper) IString {
	rv := objc.Send[String](f_.ID, objc.Sel("addFileWrapper:"), child)
	return rv
}/* debug [instance_methods/method]: AddFileWrapper */


// Creates a regular-file file wrapper with the given contents and adds it to the receiver, which must be a directory file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/addRegularFile(withContents:preferredFilename:)
func (f_ FileWrapper) AddRegularFileWithContentsPreferredFilename(data IData, fileName IString) IString {
	rv := objc.Send[String](f_.ID, objc.Sel("addRegularFileWithContents:preferredFilename:"), data, fileName)
	return rv
}/* debug [instance_methods/method]: AddRegularFileWithContentsPreferredFilename */


// Returns the dictionary key used by a directory to identify a given file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/keyForChildFileWrapper(_:)
func (f_ FileWrapper) KeyForFileWrapper(child IFileWrapper) IString {
	rv := objc.Send[String](f_.ID, objc.Sel("keyForFileWrapper:"), child)
	return rv
}/* debug [instance_methods/method]: KeyForFileWrapper */


// Indicates whether the contents of a file wrapper matches a directory, regular file, or symbolic link on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/matchesContents(of:)
func (f_ FileWrapper) MatchesContentsOfURL(url IURL) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("matchesContentsOfURL:"), url)
	return rv
}/* debug [instance_methods/method]: MatchesContentsOfURL */


// Recursively rereads the entire contents of a file wrapper from the specified location on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/read(from:options:)
func (f_ FileWrapper) ReadFromURLOptionsError(url IURL, options FileWrapperReadingOptions, outError IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("readFromURL:options:error:"), url, options, outError)
	return rv
}/* debug [instance_methods/method]: ReadFromURLOptionsError */


// Removes a child file wrapper from the receiver, which must be a directory file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/removeFileWrapper(_:)
func (f_ FileWrapper) RemoveFileWrapper(child IFileWrapper) {
	objc.Send[objc.ID](f_.ID, objc.Sel("removeFileWrapper:"), child)
}/* debug [instance_methods/method]: RemoveFileWrapper */


// Recursively writes the entire contents of a file wrapper to a given file-system URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/write(to:options:originalContentsURL:)
func (f_ FileWrapper) WriteToURLOptionsOriginalContentsURLError(url IURL, options FileWrapperWritingOptions, originalContentsURL IURL, outError IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("writeToURL:options:originalContentsURL:error:"), url, options, originalContentsURL, outError)
	return rv
}/* debug [instance_methods/method]: WriteToURLOptionsOriginalContentsURLError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileWrapper */

// A dictionary of file attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/fileAttributes
func (f_ FileWrapper) FileAttributes() IDictionary {
	rv := objc.Send[Dictionary](f_.ID, objc.Sel("fileAttributes"))
	return rv
}/* debug [instance_properties/getter]: fileAttributes */


// A dictionary of file attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/fileAttributes
func (f_ FileWrapper) SetFileAttributes(value IDictionary) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileAttributes:"), value)
}/* debug [instance_properties/setter]: fileAttributes */


// The file wrappers contained by a directory file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/fileWrappers
func (f_ FileWrapper) FileWrappers() IDictionary {
	rv := objc.Send[Dictionary](f_.ID, objc.Sel("fileWrappers"))
	return rv
}/* debug [instance_properties/getter]: fileWrappers */


// The filename of the file wrapper object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/filename
func (f_ FileWrapper) Filename() IString {
	rv := objc.Send[String](f_.ID, objc.Sel("filename"))
	return rv
}/* debug [instance_properties/getter]: filename */


// The filename of the file wrapper object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/filename
func (f_ FileWrapper) SetFilename(value IString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFilename:"), value)
}/* debug [instance_properties/setter]: filename */


// The preferred filename for the file wrapper object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/preferredFilename
func (f_ FileWrapper) PreferredFilename() IString {
	rv := objc.Send[String](f_.ID, objc.Sel("preferredFilename"))
	return rv
}/* debug [instance_properties/getter]: preferredFilename */


// The preferred filename for the file wrapper object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/preferredFilename
func (f_ FileWrapper) SetPreferredFilename(value IString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPreferredFilename:"), value)
}/* debug [instance_properties/setter]: preferredFilename */


// The contents of the file-system node associated with a regular-file file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/regularFileContents
func (f_ FileWrapper) RegularFileContents() IData {
	rv := objc.Send[Data](f_.ID, objc.Sel("regularFileContents"))
	return rv
}/* debug [instance_properties/getter]: regularFileContents */


// The contents of the file wrapper as an opaque data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/serializedRepresentation
func (f_ FileWrapper) SerializedRepresentation() IData {
	rv := objc.Send[Data](f_.ID, objc.Sel("serializedRepresentation"))
	return rv
}/* debug [instance_properties/getter]: serializedRepresentation */


// The URL referenced by the file wrapper object, which must be a symbolic-link file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileWrapper/symbolicLinkDestinationURL
func (f_ FileWrapper) SymbolicLinkDestinationURL() IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("symbolicLinkDestinationURL"))
	return rv
}/* debug [instance_properties/getter]: symbolicLinkDestinationURL */


// The icon that represents the file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/icon
func (f_ FileWrapper) Icon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("icon"))
	return rv
}/* debug [instance_properties/getter]: icon */


// The icon that represents the file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/icon
func (f_ FileWrapper) SetIcon(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIcon:"), value)
}/* debug [instance_properties/setter]: icon */


// This property contains a boolean value indicating whether the file wrapper is a directory file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/isdirectory
func (f_ FileWrapper) IsDirectory() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDirectory"))
	return rv
}/* debug [instance_properties/getter]: isDirectory */


// This property contains a boolean value indicating whether the file wrapper is a directory file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/isdirectory
func (f_ FileWrapper) SetIsDirectory(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsDirectory:"), value)
}/* debug [instance_properties/setter]: isDirectory */


// This property contains a boolean value that indicates whether the file wrapper object is a regular-file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/isregularfile
func (f_ FileWrapper) IsRegularFile() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isRegularFile"))
	return rv
}/* debug [instance_properties/getter]: isRegularFile */


// This property contains a boolean value that indicates whether the file wrapper object is a regular-file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/isregularfile
func (f_ FileWrapper) SetIsRegularFile(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsRegularFile:"), value)
}/* debug [instance_properties/setter]: isRegularFile */


// A boolean that indicates whether the file wrapper object is a symbolic-link file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/issymboliclink
func (f_ FileWrapper) IsSymbolicLink() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isSymbolicLink"))
	return rv
}/* debug [instance_properties/getter]: isSymbolicLink */


// A boolean that indicates whether the file wrapper object is a symbolic-link file wrapper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filewrapper/issymboliclink
func (f_ FileWrapper) SetIsSymbolicLink(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsSymbolicLink:"), value)
}/* debug [instance_properties/setter]: isSymbolicLink */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileWrapper */


