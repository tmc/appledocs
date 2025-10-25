// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURL */


/* debug [class_header]: Header for NSURL */
// The class instance for the [URL] class.
var (
	URLClass     _URLClass
	URLClassOnce sync.Once
)

func getURLClass() _URLClass {
	URLClassOnce.Do(func() {
		URLClass = _URLClass{objc.GetClass("NSURL")}
	})
	return URLClass
}

type _URLClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URL */
// An interface definition for the [URL] class.
type IURL interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URL */
	// properties:
	AbsoluteString() IString
	AbsoluteURL() IURL
	BaseURL() IURL
	DataRepresentation() IData
	URLByDeletingLastPathComponent() IURL
	URLByDeletingPathExtension() IURL
	FilePathURL() IURL
	FileSystemRepresentation() objectivec.IObject
	Fragment() IString
	HasDirectoryPath() bool
	Host() IString
	FileURL() bool
	LastPathComponent() IString
	ParameterString() IString
	Password() IString
	Path() IString
	PathComponents() []string
	PathExtension() IString
	Port() INumber
	Query() IString
	RelativePath() IString
	RelativeString() IString
	URLByResolvingSymlinksInPath() IURL
	ResourceSpecifier() IString
	Scheme() IString
	StandardizedURL() IURL
	URLByStandardizingPath() IURL
	User() IString
	CustomPlaygroundQuickLook() objectivec.IObject
	SetCustomPlaygroundQuickLook(value objectivec.IObject)
	DeletingLastPathComponent() IURL
	SetDeletingLastPathComponent(value IURL)
	DeletingPathExtension() IURL
	SetDeletingPathExtension(value IURL)
	IsFileURL() bool
	SetIsFileURL(value bool)
	ResolvingSymlinksInPath() IURL
	SetResolvingSymlinksInPath(value IURL)
	Standardized() IURL
	SetStandardized(value IURL)
	StandardizingPath() IURL
	SetStandardizingPath(value IURL)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URL */
	// methods:
	URLByAppendingPathComponent(pathComponent IString) IURL
	URLByAppendingPathComponentConformingToType(partialName IString, contentType objc.IObject) IURL
	URLByAppendingPathComponentIsDirectory(pathComponent IString, isDirectory bool) IURL
	URLByAppendingPathExtension(pathExtension IString) IURL
	URLByAppendingPathExtensionForType(contentType objc.IObject) IURL
	BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions, keys []string, relativeURL IURL, error_ IError) IData
	CheckPromisedItemIsReachableAndReturnError(error_ IError) bool
	CheckResourceIsReachableAndReturnError(error_ IError) bool
	FileReferenceURL() IURL
	GetFileSystemRepresentationMaxLength(buffer objectivec.IObject, maxBufferLength uint) bool
	GetPromisedItemResourceValueForKeyError(value objectivec.IObject, key URLResourceKey, error_ IError) bool
	GetResourceValueForKeyError(value objectivec.IObject, key URLResourceKey, error_ IError) bool
	IsFileReferenceURL() bool
	PromisedItemResourceValuesForKeysError(keys []string, error_ IError) IDictionary
	RemoveAllCachedResourceValues()
	RemoveCachedResourceValueForKey(key URLResourceKey)
	ResourceValuesForKeysError(keys []string, error_ IError) IDictionary
	SetResourceValueForKeyError(value objc.IObject, key URLResourceKey, error_ IError) bool
	SetResourceValuesError(keyedValues IDictionary, error_ IError) bool
	SetTemporaryResourceValueForKey(value objc.IObject, key URLResourceKey)
	StartAccessingSecurityScopedResource() bool
	StopAccessingSecurityScopedResource()
	WriteToPasteboard(pasteBoard objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URL */
// Alloc allocates a new instance without initialization.
func (uc _URLClass) Alloc() URL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _URLClass) New() URL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URL) Init() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URL) Autorelease() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURL creates a new URL instance.
func NewURL() URL {
	return getURLClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URL */
// An object that represents the location of a resource, such as an item on a remote server or the path to a local file.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. You can use URL objects to construct URLs and access their parts. For URLs that represent local files, you can also manipulate properties of those files directly, such as changing the file’s last modification date. Finally, you can pass URL objects to other APIs to retrieve the contents of those URLs. For example, you can use the , , and classes to access the contents of remote resources, as described in . URL objects are the preferred way to refer to local files. Most objects that read data from or write data to a file have methods that accept an object instead of a pathname as the file reference. For example, you can get the contents of a local file URL as an object using the initializer, or as an object using the initializer. You can also use URLs for interapplication communication. In macOS, the class provides the method to open a location specified by a URL. Similarly, in iOS, the class provides the method. Additionally, you can use URLs when working with pasteboards, as described in NSURL Additions Reference (part of the AppKit framework). The class is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.


// An object that represents the location of a resource, such as an item on a remote server or the path to a local file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL
type URL struct {
	objectivec.Object
}

// URLFrom constructs a [URL] from an unsafe.Pointer.
//
// An object that represents the location of a resource, such as an item on a remote server or the path to a local file.
func URLFrom(ptr unsafe.Pointer) URL {
	return URL{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URL */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(absoluteURLWithDataRepresentation:relativeTo:)
func NewURLAbsoluteURLWithDataRepresentationRelativeToURL(data IData, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initAbsoluteURLWithDataRepresentation:relativeToURL:"), data, baseURL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLAbsoluteURLWithDataRepresentationRelativeToURL */


// Returns a new URL made by resolving the alias file at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingAliasFileAt:options:)
func NewURLByResolvingAliasFileAtURLOptionsError(url IURL, options URLBookmarkResolutionOptions, error_ IError) URL {
	rv := objc.Send[URL](objc.ID(getURLClass().class), objc.Sel("URLByResolvingAliasFileAtURL:options:error:"), url, options, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewURLByResolvingAliasFileAtURLOptionsError */


// Initializes a newly created NSURL that points to a location specified by resolving bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingBookmarkData:options:relativeTo:bookmarkDataIsStale:)
func NewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData IData, options URLBookmarkResolutionOptions, relativeURL IURL, isStale objectivec.IObject, error_ IError) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError */


// Initializes a URL object with a C string representing a local file system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithFileSystemRepresentation:isDirectory:relativeTo:)
func NewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path objectivec.IObject, isDir bool, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, baseURL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL */


// Initializes a newly created NSURL referencing the local file or directory at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:)
func NewURLFileURLWithPath(path IString) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:"), path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLFileURLWithPath */


// Initializes a newly created NSURL referencing the local file or directory at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:)
func NewURLFileURLWithPathIsDirectory(path IString, isDir bool) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:"), path, isDir)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLFileURLWithPathIsDirectory */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:relativeTo:)
func NewURLFileURLWithPathIsDirectoryRelativeToURL(path IString, isDir bool, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:relativeToURL:"), path, isDir, baseURL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLFileURLWithPathIsDirectoryRelativeToURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:relativeTo:)
func NewURLFileURLWithPathRelativeToURL(path IString, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:relativeToURL:"), path, baseURL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLFileURLWithPathRelativeToURL */


// Reads an NSURL object off of the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fromPasteboard:)
func NewURLFromPasteboard(pasteBoard objectivec.IObject) URL {
	rv := objc.Send[URL](objc.ID(getURLClass().class), objc.Sel("URLFromPasteboard:"), pasteBoard)
	return rv
}/* debug [class_init_methods/constructor]: NewURLFromPasteboard */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(dataRepresentation:relativeTo:)
func NewURLWithDataRepresentationRelativeToURL(data IData, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithDataRepresentation:relativeToURL:"), data, baseURL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLWithDataRepresentationRelativeToURL */


// Initializes a newly created NSURL with a specified scheme, host, and path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(scheme:host:path:)
func NewURLWithSchemeHostPath(scheme IString, host IString, path IString) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithScheme:host:path:"), scheme, host, path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLWithSchemeHostPath */


// Initializes an NSURL object with a provided URL string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:)
func NewURLWithString(URLString IString) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:"), URLString)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLWithString */


// Creates an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:encodingInvalidCharacters:)
func NewURLWithStringEncodingInvalidCharacters(URLString IString, encodingInvalidCharacters bool) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:encodingInvalidCharacters:"), URLString, encodingInvalidCharacters)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLWithStringEncodingInvalidCharacters */


// Initializes an NSURL object with a base URL and a relative string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:relativeTo:)
func NewURLWithStringRelativeToURL(URLString IString, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:relativeToURL:"), URLString, baseURL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLWithStringRelativeToURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URL */

// Returns a new URL made by resolving bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:
func (uc _URLClass) URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData IData, options URLBookmarkResolutionOptions, relativeURL IURL, isStale bool, error_ IError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithDataRepresentation:relativeToURL:
func (uc _URLClass) URLWithDataRepresentationRelativeToURL(data IData, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("URLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=URLWithDataRepresentationRelativeToURL) */


// Creates and returns an NSURL object initialized with a provided URL string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:
func (uc _URLClass) URLWithString(URLString IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("URLWithString:"), URLString)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=URLWithString) */


// Creates and returns an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:encodingInvalidCharacters:
func (uc _URLClass) URLWithStringEncodingInvalidCharacters(URLString IString, encodingInvalidCharacters bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("URLWithString:encodingInvalidCharacters:"), URLString, encodingInvalidCharacters)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=URLWithStringEncodingInvalidCharacters) */


// Creates and returns an NSURL object initialized with a base URL and a relative string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:relativeToURL:
func (uc _URLClass) URLWithStringRelativeToURL(URLString IString, baseURL IURL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("URLWithString:relativeToURL:"), URLString, baseURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=URLWithStringRelativeToURL) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/absoluteURL(withDataRepresentation:relativeTo:)
func (uc _URLClass) AbsoluteURLWithDataRepresentationRelativeToURL(data IData, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("absoluteURLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AbsoluteURLWithDataRepresentationRelativeToURL) */


// Initializes and returns bookmark data derived from an alias file pointed to by a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/bookmarkData(withContentsOf:)
func (uc _URLClass) BookmarkDataWithContentsOfURLError(bookmarkFileURL IURL, error_ IError) IData {
	rv := objc.Send[Data](objc.ID(uc.class), objc.Sel("bookmarkDataWithContentsOfURL:error:"), bookmarkFileURL, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BookmarkDataWithContentsOfURLError) */


// Returns a new URL object initialized with a C string representing a local file system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withFileSystemRepresentation:isDirectory:relativeTo:)
func (uc _URLClass) FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path objectivec.IObject, isDir bool, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, baseURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL) */


// Initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:)
func (uc _URLClass) FileURLWithPath(path IString) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileURLWithPath) */


// Initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:)
func (uc _URLClass) FileURLWithPathIsDirectory(path IString, isDir bool) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:isDirectory:"), path, isDir)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileURLWithPathIsDirectory) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:relativeTo:)
func (uc _URLClass) FileURLWithPathIsDirectoryRelativeToURL(path IString, isDir bool, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:isDirectory:relativeToURL:"), path, isDir, baseURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileURLWithPathIsDirectoryRelativeToURL) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:relativeTo:)
func (uc _URLClass) FileURLWithPathRelativeToURL(path IString, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:relativeToURL:"), path, baseURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileURLWithPathRelativeToURL) */


// Initializes and returns a newly created NSURL object as a file URL with specified path components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPathComponents:)
func (uc _URLClass) FileURLWithPathComponents(components []string) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPathComponents:"), components)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileURLWithPathComponents) */


// Reads an NSURL object off of the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fromPasteboard:)
func (uc _URLClass) URLFromPasteboard(pasteBoard objectivec.IObject) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("URLFromPasteboard:"), pasteBoard)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=URLFromPasteboard) */


// Returns a new URL made by resolving the alias file at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingAliasFileAt:options:)
func (uc _URLClass) URLByResolvingAliasFileAtURLOptionsError(url IURL, options URLBookmarkResolutionOptions, error_ IError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("URLByResolvingAliasFileAtURL:options:error:"), url, options, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=URLByResolvingAliasFileAtURLOptionsError) */


// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceValues(forKeys:fromBookmarkData:)
func (uc _URLClass) ResourceValuesForKeysFromBookmarkData(keys []string, bookmarkData IData) IDictionary {
	rv := objc.Send[Dictionary](objc.ID(uc.class), objc.Sel("resourceValuesForKeys:fromBookmarkData:"), keys, bookmarkData)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ResourceValuesForKeysFromBookmarkData) */


// Creates an alias file on disk at a specified location with specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/writeBookmarkData(_:to:options:)
func (uc _URLClass) WriteBookmarkDataToURLOptionsError(bookmarkData IData, bookmarkFileURL IURL, options URLBookmarkFileCreationOptions, error_ IError) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("writeBookmarkData:toURL:options:error:"), bookmarkData, bookmarkFileURL, options, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WriteBookmarkDataToURLOptionsError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URL */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URL */

// Returns a new URL by appending a path component to the original URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:)
func (u_ URL) URLByAppendingPathComponent(pathComponent IString) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathComponent:"), pathComponent)
	return rv
}/* debug [instance_methods/method]: URLByAppendingPathComponent */


// Returns a URL by appending the specified path component with the file extension for a uniform type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:conformingTo:)
func (u_ URL) URLByAppendingPathComponentConformingToType(partialName IString, contentType objc.IObject) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathComponent:conformingToType:"), partialName, contentType)
	return rv
}/* debug [instance_methods/method]: URLByAppendingPathComponentConformingToType */


// Returns a new URL by appending a path component to the original URL, along with a trailing slash if the component is a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:isDirectory:)
func (u_ URL) URLByAppendingPathComponentIsDirectory(pathComponent IString, isDirectory bool) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathComponent:isDirectory:"), pathComponent, isDirectory)
	return rv
}/* debug [instance_methods/method]: URLByAppendingPathComponentIsDirectory */


// Returns a new URL by appending a path extension to the original URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathExtension(_:)
func (u_ URL) URLByAppendingPathExtension(pathExtension IString) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathExtension:"), pathExtension)
	return rv
}/* debug [instance_methods/method]: URLByAppendingPathExtension */


// Returns a URL by appending the path extension for a uniform type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathExtension(for:)
func (u_ URL) URLByAppendingPathExtensionForType(contentType objc.IObject) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathExtensionForType:"), contentType)
	return rv
}/* debug [instance_methods/method]: URLByAppendingPathExtensionForType */


// Returns a bookmark for the URL, created with specified options and resource values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/bookmarkData(options:includingResourceValuesForKeys:relativeTo:)
func (u_ URL) BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions, keys []string, relativeURL IURL, error_ IError) IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:"), options, keys, relativeURL, error_)
	return rv
}/* debug [instance_methods/method]: BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError */


// Returns whether the promised item can be reached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/checkPromisedItemIsReachableAndReturnError(_:)
func (u_ URL) CheckPromisedItemIsReachableAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("checkPromisedItemIsReachableAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: CheckPromisedItemIsReachableAndReturnError */


// Returns whether the resource pointed to by a file URL can be reached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/checkResourceIsReachableAndReturnError(_:)
func (u_ URL) CheckResourceIsReachableAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("checkResourceIsReachableAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: CheckResourceIsReachableAndReturnError */


// Returns a new file reference URL that points to the same resource as the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileReferenceURL()
func (u_ URL) FileReferenceURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("fileReferenceURL"))
	return rv
}/* debug [instance_methods/method]: FileReferenceURL */


// Fills the provided buffer with a C string representing a local file system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getFileSystemRepresentation(_:maxLength:)
func (u_ URL) GetFileSystemRepresentationMaxLength(buffer objectivec.IObject, maxBufferLength uint) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("getFileSystemRepresentation:maxLength:"), buffer, maxBufferLength)
	return rv
}/* debug [instance_methods/method]: GetFileSystemRepresentationMaxLength */


// Returns the value of the resource property for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getPromisedItemResourceValue(_:forKey:)
func (u_ URL) GetPromisedItemResourceValueForKeyError(value objectivec.IObject, key URLResourceKey, error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("getPromisedItemResourceValue:forKey:error:"), value, key, error_)
	return rv
}/* debug [instance_methods/method]: GetPromisedItemResourceValueForKeyError */


// Returns the value of the resource property for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getResourceValue(_:forKey:)
func (u_ URL) GetResourceValueForKeyError(value objectivec.IObject, key URLResourceKey, error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("getResourceValue:forKey:error:"), value, key, error_)
	return rv
}/* debug [instance_methods/method]: GetResourceValueForKeyError */


// Returns whether the URL is a file reference URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/isFileReferenceURL()
func (u_ URL) IsFileReferenceURL() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isFileReferenceURL"))
	return rv
}/* debug [instance_methods/method]: IsFileReferenceURL */


// Returns the resource values for the properties identified by specified array of keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/promisedItemResourceValues(forKeys:)
func (u_ URL) PromisedItemResourceValuesForKeysError(keys []string, error_ IError) IDictionary {
	rv := objc.Send[Dictionary](u_.ID, objc.Sel("promisedItemResourceValuesForKeys:error:"), keys, error_)
	return rv
}/* debug [instance_methods/method]: PromisedItemResourceValuesForKeysError */


// Removes all cached resource values and temporary resource values from the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/removeAllCachedResourceValues()
func (u_ URL) RemoveAllCachedResourceValues() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllCachedResourceValues"))
}/* debug [instance_methods/method]: RemoveAllCachedResourceValues */


// Removes the cached resource value identified by a given key from the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/removeCachedResourceValue(forKey:)
func (u_ URL) RemoveCachedResourceValueForKey(key URLResourceKey) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeCachedResourceValueForKey:"), key)
}/* debug [instance_methods/method]: RemoveCachedResourceValueForKey */


// Returns the resource values for the properties identified by specified array of keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceValues(forKeys:)
func (u_ URL) ResourceValuesForKeysError(keys []string, error_ IError) IDictionary {
	rv := objc.Send[Dictionary](u_.ID, objc.Sel("resourceValuesForKeys:error:"), keys, error_)
	return rv
}/* debug [instance_methods/method]: ResourceValuesForKeysError */


// Sets the URL’s resource property for a given key to a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValue(_:forKey:)
func (u_ URL) SetResourceValueForKeyError(value objc.IObject, key URLResourceKey, error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceValue:forKey:error:"), value, key, error_)
	return rv
}/* debug [instance_methods/method]: SetResourceValueForKeyError */


// Sets the URL’s resource properties for a given set of keys to a given set of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValues(_:)
func (u_ URL) SetResourceValuesError(keyedValues IDictionary, error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceValues:error:"), keyedValues, error_)
	return rv
}/* debug [instance_methods/method]: SetResourceValuesError */


// Sets a temporary resource value on the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setTemporaryResourceValue(_:forKey:)
func (u_ URL) SetTemporaryResourceValueForKey(value objc.IObject, key URLResourceKey) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTemporaryResourceValue:forKey:"), value, key)
}/* debug [instance_methods/method]: SetTemporaryResourceValueForKey */


// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/startAccessingSecurityScopedResource()
func (u_ URL) StartAccessingSecurityScopedResource() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("startAccessingSecurityScopedResource"))
	return rv
}/* debug [instance_methods/method]: StartAccessingSecurityScopedResource */


// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/stopAccessingSecurityScopedResource()
func (u_ URL) StopAccessingSecurityScopedResource() {
	objc.Send[objc.ID](u_.ID, objc.Sel("stopAccessingSecurityScopedResource"))
}/* debug [instance_methods/method]: StopAccessingSecurityScopedResource */


// Writes the URL to the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/write(to:)
func (u_ URL) WriteToPasteboard(pasteBoard objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("writeToPasteboard:"), pasteBoard)
}/* debug [instance_methods/method]: WriteToPasteboard */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URL */

// The URL string for the receiver as an absolute URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/absoluteString
func (u_ URL) AbsoluteString() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("absoluteString"))
	return rv
}/* debug [instance_properties/getter]: absoluteString */


// An absolute URL that refers to the same resource as the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/absoluteURL
func (u_ URL) AbsoluteURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("absoluteURL"))
	return rv
}/* debug [instance_properties/getter]: absoluteURL */


// The base URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/baseURL
func (u_ URL) BaseURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("baseURL"))
	return rv
}/* debug [instance_properties/getter]: baseURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/dataRepresentation
func (u_ URL) DataRepresentation() IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("dataRepresentation"))
	return rv
}/* debug [instance_properties/getter]: dataRepresentation */


// A URL you create by removing the last path component from the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/deletingLastPathComponent
func (u_ URL) URLByDeletingLastPathComponent() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByDeletingLastPathComponent"))
	return rv
}/* debug [instance_properties/getter]: URLByDeletingLastPathComponent */


// A URL you create by removing the path extension from the receiver, if any. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/deletingPathExtension
func (u_ URL) URLByDeletingPathExtension() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByDeletingPathExtension"))
	return rv
}/* debug [instance_properties/getter]: URLByDeletingPathExtension */


// A file path URL that points to the same resource as the URL object. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/filePathURL
func (u_ URL) FilePathURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("filePathURL"))
	return rv
}/* debug [instance_properties/getter]: filePathURL */


// A C string containing the URL’s file system path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileSystemRepresentation
func (u_ URL) FileSystemRepresentation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("fileSystemRepresentation"))
	return rv
}/* debug [instance_properties/getter]: fileSystemRepresentation */


// The fragment identifier, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fragment
func (u_ URL) Fragment() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("fragment"))
	return rv
}/* debug [instance_properties/getter]: fragment */


// A Boolean value that indicates whether the URL string’s path represents a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/hasDirectoryPath
func (u_ URL) HasDirectoryPath() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasDirectoryPath"))
	return rv
}/* debug [instance_properties/getter]: hasDirectoryPath */


// The host, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/host
func (u_ URL) Host() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("host"))
	return rv
}/* debug [instance_properties/getter]: host */


// A boolean value that determines whether the receiver is a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/isFileURL
func (u_ URL) FileURL() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("fileURL"))
	return rv
}/* debug [instance_properties/getter]: fileURL */


// The last path component. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/lastPathComponent
func (u_ URL) LastPathComponent() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("lastPathComponent"))
	return rv
}/* debug [instance_properties/getter]: lastPathComponent */


// The parameter string conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/parameterString
func (u_ URL) ParameterString() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("parameterString"))
	return rv
}/* debug [instance_properties/getter]: parameterString */


// The password conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/password
func (u_ URL) Password() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("password"))
	return rv
}/* debug [instance_properties/getter]: password */


// The path, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/path
func (u_ URL) Path() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("path"))
	return rv
}/* debug [instance_properties/getter]: path */


// An array containing the path components. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/pathComponents
func (u_ URL) PathComponents() []string {
	rv := objc.Send[[]string](u_.ID, objc.Sel("pathComponents"))
	return rv
}/* debug [instance_properties/getter]: pathComponents */


// The path extension. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/pathExtension
func (u_ URL) PathExtension() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("pathExtension"))
	return rv
}/* debug [instance_properties/getter]: pathExtension */


// The port, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/port
func (u_ URL) Port() INumber {
	rv := objc.Send[Number](u_.ID, objc.Sel("port"))
	return rv
}/* debug [instance_properties/getter]: port */


// The query string, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/query
func (u_ URL) Query() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("query"))
	return rv
}/* debug [instance_properties/getter]: query */


// The relative path, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/relativePath
func (u_ URL) RelativePath() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("relativePath"))
	return rv
}/* debug [instance_properties/getter]: relativePath */


// A string representation of the relative portion of the URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/relativeString
func (u_ URL) RelativeString() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("relativeString"))
	return rv
}/* debug [instance_properties/getter]: relativeString */


// A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resolvingSymlinksInPath
func (u_ URL) URLByResolvingSymlinksInPath() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByResolvingSymlinksInPath"))
	return rv
}/* debug [instance_properties/getter]: URLByResolvingSymlinksInPath */


// The resource specifier. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceSpecifier
func (u_ URL) ResourceSpecifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("resourceSpecifier"))
	return rv
}/* debug [instance_properties/getter]: resourceSpecifier */


// The scheme. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/scheme
func (u_ URL) Scheme() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("scheme"))
	return rv
}/* debug [instance_properties/getter]: scheme */


// A copy of the URL with any instances of or removed from its path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/standardized
func (u_ URL) StandardizedURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("standardizedURL"))
	return rv
}/* debug [instance_properties/getter]: standardizedURL */


// A URL that points to the same resource as the original URL using an absolute path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/standardizingPath
func (u_ URL) URLByStandardizingPath() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByStandardizingPath"))
	return rv
}/* debug [instance_properties/getter]: URLByStandardizingPath */


// The user name, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/user
func (u_ URL) User() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("user"))
	return rv
}/* debug [instance_properties/getter]: user */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/customplaygroundquicklook
func (u_ URL) CustomPlaygroundQuickLook() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("customPlaygroundQuickLook"))
	return rv
}/* debug [instance_properties/getter]: customPlaygroundQuickLook */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/customplaygroundquicklook
func (u_ URL) SetCustomPlaygroundQuickLook(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCustomPlaygroundQuickLook:"), value)
}/* debug [instance_properties/setter]: customPlaygroundQuickLook */


// A URL you create by removing the last path component from the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/deletinglastpathcomponent
func (u_ URL) DeletingLastPathComponent() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("deletingLastPathComponent"))
	return rv
}/* debug [instance_properties/getter]: deletingLastPathComponent */


// A URL you create by removing the last path component from the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/deletinglastpathcomponent
func (u_ URL) SetDeletingLastPathComponent(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeletingLastPathComponent:"), value)
}/* debug [instance_properties/setter]: deletingLastPathComponent */


// A URL you create by removing the path extension from the receiver, if any. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/deletingpathextension
func (u_ URL) DeletingPathExtension() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("deletingPathExtension"))
	return rv
}/* debug [instance_properties/getter]: deletingPathExtension */


// A URL you create by removing the path extension from the receiver, if any. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/deletingpathextension
func (u_ URL) SetDeletingPathExtension(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeletingPathExtension:"), value)
}/* debug [instance_properties/setter]: deletingPathExtension */


// A boolean value that determines whether the receiver is a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/isfileurl
func (u_ URL) IsFileURL() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isFileURL"))
	return rv
}/* debug [instance_properties/getter]: isFileURL */


// A boolean value that determines whether the receiver is a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/isfileurl
func (u_ URL) SetIsFileURL(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsFileURL:"), value)
}/* debug [instance_properties/setter]: isFileURL */


// A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/resolvingsymlinksinpath
func (u_ URL) ResolvingSymlinksInPath() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("resolvingSymlinksInPath"))
	return rv
}/* debug [instance_properties/getter]: resolvingSymlinksInPath */


// A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/resolvingsymlinksinpath
func (u_ URL) SetResolvingSymlinksInPath(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResolvingSymlinksInPath:"), value)
}/* debug [instance_properties/setter]: resolvingSymlinksInPath */


// A copy of the URL with any instances of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/standardized
func (u_ URL) Standardized() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("standardized"))
	return rv
}/* debug [instance_properties/getter]: standardized */


// A copy of the URL with any instances of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/standardized
func (u_ URL) SetStandardized(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setStandardized:"), value)
}/* debug [instance_properties/setter]: standardized */


// A URL that points to the same resource as the original URL using an absolute path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/standardizingpath
func (u_ URL) StandardizingPath() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("standardizingPath"))
	return rv
}/* debug [instance_properties/getter]: standardizingPath */


// A URL that points to the same resource as the original URL using an absolute path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/standardizingpath
func (u_ URL) SetStandardizingPath(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setStandardizingPath:"), value)
}/* debug [instance_properties/setter]: standardizingPath */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURL */


