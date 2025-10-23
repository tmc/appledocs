// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [URL] class.
type IURL interface {
	objectivec.IObject
	// properties:
	AbsoluteString() string /* primitive/slice/pointer */
	AbsoluteURL() IURL
	BaseURL() IURL
	DataRepresentation() IData
	URLByDeletingLastPathComponent() IURL
	URLByDeletingPathExtension() IURL
	FilePathURL() IURL
	FileSystemRepresentation() unsafe.Pointer
	Fragment() string /* primitive/slice/pointer */
	HasDirectoryPath() bool /* primitive/slice/pointer */
	Host() string /* primitive/slice/pointer */
	FileURL() bool /* primitive/slice/pointer */
	LastPathComponent() string /* primitive/slice/pointer */
	ParameterString() string /* primitive/slice/pointer */
	Password() string /* primitive/slice/pointer */
	Path() string /* primitive/slice/pointer */
	PathComponents() []string /* primitive/slice/pointer */
	PathExtension() string /* primitive/slice/pointer */
	Port() Number /* foo */
	Query() string /* primitive/slice/pointer */
	RelativePath() string /* primitive/slice/pointer */
	RelativeString() string /* primitive/slice/pointer */
	URLByResolvingSymlinksInPath() IURL
	ResourceSpecifier() string /* primitive/slice/pointer */
	Scheme() string /* primitive/slice/pointer */
	StandardizedURL() IURL
	URLByStandardizingPath() IURL
	User() string /* primitive/slice/pointer */
	CustomPlaygroundQuickLook() unsafe.Pointer
	SetCustomPlaygroundQuickLook(value unsafe.Pointer)
	DeletingLastPathComponent() IURL
	SetDeletingLastPathComponent(value IURL)
	DeletingPathExtension() IURL
	SetDeletingPathExtension(value IURL)
	IsFileURL() bool /* primitive/slice/pointer */
	SetIsFileURL(value bool /* primitive/slice/pointer */)
	ResolvingSymlinksInPath() IURL
	SetResolvingSymlinksInPath(value IURL)
	Standardized() IURL
	SetStandardized(value IURL)
	StandardizingPath() IURL
	SetStandardizingPath(value IURL)
	// methods:
	URLByAppendingPathComponent(pathComponent string /* primitive/slice/pointer */) IURL
	URLByAppendingPathComponentConformingToType(partialName string /* primitive/slice/pointer */, contentType objectivec.IObject) IURL
	URLByAppendingPathComponentIsDirectory(pathComponent string /* primitive/slice/pointer */, isDirectory bool /* primitive/slice/pointer */) IURL
	URLByAppendingPathExtension(pathExtension string /* primitive/slice/pointer */) IURL
	URLByAppendingPathExtensionForType(contentType objectivec.IObject) IURL
	BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions, keys []string /* primitive/slice/pointer */, relativeURL IURL, error_ unsafe.Pointer) IData
	CheckPromisedItemIsReachableAndReturnError(error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	CheckResourceIsReachableAndReturnError(error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	FileReferenceURL() IURL
	GetFileSystemRepresentationMaxLength(buffer unsafe.Pointer, maxBufferLength uint /* primitive/slice/pointer */) bool /* primitive/slice/pointer */
	GetPromisedItemResourceValueForKeyError(value objectivec.IObject, key URLResourceKey /* foo */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	GetResourceValueForKeyError(value objectivec.IObject, key URLResourceKey /* foo */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	IsFileReferenceURL() bool /* primitive/slice/pointer */
	PromisedItemResourceValuesForKeysError(keys []string /* primitive/slice/pointer */, error_ unsafe.Pointer) IDictionary /* already interface */
	RemoveAllCachedResourceValues()
	RemoveCachedResourceValueForKey(key URLResourceKey /* foo */)
	ResourceValuesForKeysError(keys []string /* primitive/slice/pointer */, error_ unsafe.Pointer) IDictionary /* already interface */
	SetResourceValueForKeyError(value objectivec.IObject, key URLResourceKey /* foo */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	SetResourceValuesError(keyedValues IDictionary /* already interface */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	SetTemporaryResourceValueForKey(value objectivec.IObject, key URLResourceKey /* foo */)
	StartAccessingSecurityScopedResource() bool /* primitive/slice/pointer */
	StopAccessingSecurityScopedResource()
	WriteToPasteboard(pasteBoard Pasteboard /* foo */)
}

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

// Alloc allocates a new instance without initialization.
func (uc _URLClass) Alloc() URL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(absoluteURLWithDataRepresentation:relativeTo:)
func NewURLAbsoluteURLWithDataRepresentationRelativeToURL(data IData, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initAbsoluteURLWithDataRepresentation:relativeToURL:"), data, baseURL)
	rv.Autorelease()
	return rv
}


// Returns a new URL made by resolving the alias file at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingAliasFileAt:options:)
func NewURLByResolvingAliasFileAtURLOptionsError(url IURL, options URLBookmarkResolutionOptions, error_ unsafe.Pointer) URL {
	rv := objc.Send[URL](objc.ID(getURLClass().class), objc.Sel("URLByResolvingAliasFileAtURL:options:error:"), url, options, error_)
	return rv
}


// Initializes a newly created NSURL that points to a location specified by resolving bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingBookmarkData:options:relativeTo:bookmarkDataIsStale:)
func NewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData IData, options URLBookmarkResolutionOptions, relativeURL IURL, isStale unsafe.Pointer, error_ unsafe.Pointer) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, error_)
	rv.Autorelease()
	return rv
}


// Initializes a URL object with a C string representing a local file system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithFileSystemRepresentation:isDirectory:relativeTo:)
func NewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path unsafe.Pointer, isDir bool /* primitive/slice/pointer */, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, baseURL)
	rv.Autorelease()
	return rv
}


// Initializes a newly created NSURL referencing the local file or directory at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:)
func NewURLFileURLWithPath(path string /* primitive/slice/pointer */) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:"), objc.String(path))
	rv.Autorelease()
	return rv
}


// Initializes a newly created NSURL referencing the local file or directory at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:)
func NewURLFileURLWithPathIsDirectory(path string /* primitive/slice/pointer */, isDir bool /* primitive/slice/pointer */) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:"), objc.String(path), isDir)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:relativeTo:)
func NewURLFileURLWithPathIsDirectoryRelativeToURL(path string /* primitive/slice/pointer */, isDir bool /* primitive/slice/pointer */, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:relativeToURL:"), objc.String(path), isDir, baseURL)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:relativeTo:)
func NewURLFileURLWithPathRelativeToURL(path string /* primitive/slice/pointer */, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:relativeToURL:"), objc.String(path), baseURL)
	rv.Autorelease()
	return rv
}


// Reads an NSURL object off of the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fromPasteboard:)
func NewURLFromPasteboard(pasteBoard Pasteboard /* foo */) URL {
	rv := objc.Send[URL](objc.ID(getURLClass().class), objc.Sel("URLFromPasteboard:"), pasteBoard)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(dataRepresentation:relativeTo:)
func NewURLWithDataRepresentationRelativeToURL(data IData, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithDataRepresentation:relativeToURL:"), data, baseURL)
	rv.Autorelease()
	return rv
}


// Initializes a newly created NSURL with a specified scheme, host, and path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(scheme:host:path:)
func NewURLWithSchemeHostPath(scheme string /* primitive/slice/pointer */, host string /* primitive/slice/pointer */, path string /* primitive/slice/pointer */) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithScheme:host:path:"), objc.String(scheme), objc.String(host), objc.String(path))
	rv.Autorelease()
	return rv
}


// Initializes an NSURL object with a provided URL string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:)
func NewURLWithString(URLString string /* primitive/slice/pointer */) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:"), objc.String(URLString))
	rv.Autorelease()
	return rv
}


// Creates an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:encodingInvalidCharacters:)
func NewURLWithStringEncodingInvalidCharacters(URLString string /* primitive/slice/pointer */, encodingInvalidCharacters bool /* primitive/slice/pointer */) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	rv.Autorelease()
	return rv
}


// Initializes an NSURL object with a base URL and a relative string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:relativeTo:)
func NewURLWithStringRelativeToURL(URLString string /* primitive/slice/pointer */, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:relativeToURL:"), objc.String(URLString), baseURL)
	rv.Autorelease()
	return rv
}



// Returns a new URL made by resolving bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:
func (uc _URLClass) URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData IData, options URLBookmarkResolutionOptions, relativeURL IURL, isStale bool /* primitive/slice/pointer */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithDataRepresentation:relativeToURL:
func (uc _URLClass) URLWithDataRepresentationRelativeToURL(data IData, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("URLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return rv
}


// Creates and returns an NSURL object initialized with a provided URL string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:
func (uc _URLClass) URLWithString(URLString string /* primitive/slice/pointer */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLWithString:"), objc.String(URLString))
	return rv
}


// Creates and returns an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:encodingInvalidCharacters:
func (uc _URLClass) URLWithStringEncodingInvalidCharacters(URLString string /* primitive/slice/pointer */, encodingInvalidCharacters bool /* primitive/slice/pointer */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	return rv
}


// Creates and returns an NSURL object initialized with a base URL and a relative string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:relativeToURL:
func (uc _URLClass) URLWithStringRelativeToURL(URLString string /* primitive/slice/pointer */, baseURL IURL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLWithString:relativeToURL:"), objc.String(URLString), baseURL)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/absoluteURL(withDataRepresentation:relativeTo:)
func (uc _URLClass) AbsoluteURLWithDataRepresentationRelativeToURL(data IData, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("absoluteURLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return rv
}


// Initializes and returns bookmark data derived from an alias file pointed to by a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/bookmarkData(withContentsOf:)
func (uc _URLClass) BookmarkDataWithContentsOfURLError(bookmarkFileURL IURL, error_ unsafe.Pointer) IData {
	rv := objc.Send[Data](objc.ID(uc.class), objc.Sel("bookmarkDataWithContentsOfURL:error:"), bookmarkFileURL, error_)
	return rv
}


// Returns a new URL object initialized with a C string representing a local file system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withFileSystemRepresentation:isDirectory:relativeTo:)
func (uc _URLClass) FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path unsafe.Pointer, isDir bool /* primitive/slice/pointer */, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, baseURL)
	return rv
}


// Initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:)
func (uc _URLClass) FileURLWithPath(path string /* primitive/slice/pointer */) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:"), objc.String(path))
	return rv
}


// Initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:)
func (uc _URLClass) FileURLWithPathIsDirectory(path string /* primitive/slice/pointer */, isDir bool /* primitive/slice/pointer */) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:isDirectory:"), objc.String(path), isDir)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:relativeTo:)
func (uc _URLClass) FileURLWithPathIsDirectoryRelativeToURL(path string /* primitive/slice/pointer */, isDir bool /* primitive/slice/pointer */, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:isDirectory:relativeToURL:"), objc.String(path), isDir, baseURL)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:relativeTo:)
func (uc _URLClass) FileURLWithPathRelativeToURL(path string /* primitive/slice/pointer */, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:relativeToURL:"), objc.String(path), baseURL)
	return rv
}


// Initializes and returns a newly created NSURL object as a file URL with specified path components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPathComponents:)
func (uc _URLClass) FileURLWithPathComponents(components []string /* primitive/slice/pointer */) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPathComponents:"), components)
	return rv
}


// Reads an NSURL object off of the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fromPasteboard:)
func (uc _URLClass) URLFromPasteboard(pasteBoard Pasteboard /* foo */) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("URLFromPasteboard:"), pasteBoard)
	return rv
}


// Returns a new URL made by resolving the alias file at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingAliasFileAt:options:)
func (uc _URLClass) URLByResolvingAliasFileAtURLOptionsError(url IURL, options URLBookmarkResolutionOptions, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLByResolvingAliasFileAtURL:options:error:"), url, options, error_)
	return rv
}


// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceValues(forKeys:fromBookmarkData:)
func (uc _URLClass) ResourceValuesForKeysFromBookmarkData(keys []string /* primitive/slice/pointer */, bookmarkData IData) IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](objc.ID(uc.class), objc.Sel("resourceValuesForKeys:fromBookmarkData:"), keys, bookmarkData)
	return rv
}


// Creates an alias file on disk at a specified location with specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/writeBookmarkData(_:to:options:)
func (uc _URLClass) WriteBookmarkDataToURLOptionsError(bookmarkData IData, bookmarkFileURL IURL, options URLBookmarkFileCreationOptions /* foo */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("writeBookmarkData:toURL:options:error:"), bookmarkData, bookmarkFileURL, options, error_)
	return rv
}


// Returns a new URL by appending a path component to the original URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:)
func (u_ URL) URLByAppendingPathComponent(pathComponent string /* primitive/slice/pointer */) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathComponent:"), objc.String(pathComponent))
	return rv
}


// Returns a URL by appending the specified path component with the file extension for a uniform type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:conformingTo:)
func (u_ URL) URLByAppendingPathComponentConformingToType(partialName string /* primitive/slice/pointer */, contentType objectivec.IObject) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathComponent:conformingToType:"), objc.String(partialName), contentType)
	return rv
}


// Returns a new URL by appending a path component to the original URL, along with a trailing slash if the component is a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:isDirectory:)
func (u_ URL) URLByAppendingPathComponentIsDirectory(pathComponent string /* primitive/slice/pointer */, isDirectory bool /* primitive/slice/pointer */) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathComponent:isDirectory:"), objc.String(pathComponent), isDirectory)
	return rv
}


// Returns a new URL by appending a path extension to the original URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathExtension(_:)
func (u_ URL) URLByAppendingPathExtension(pathExtension string /* primitive/slice/pointer */) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathExtension:"), objc.String(pathExtension))
	return rv
}


// Returns a URL by appending the path extension for a uniform type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathExtension(for:)
func (u_ URL) URLByAppendingPathExtensionForType(contentType objectivec.IObject) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathExtensionForType:"), contentType)
	return rv
}


// Returns a bookmark for the URL, created with specified options and resource values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/bookmarkData(options:includingResourceValuesForKeys:relativeTo:)
func (u_ URL) BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions, keys []string /* primitive/slice/pointer */, relativeURL IURL, error_ unsafe.Pointer) IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:"), options, keys, relativeURL, error_)
	return rv
}


// Returns whether the promised item can be reached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/checkPromisedItemIsReachableAndReturnError(_:)
func (u_ URL) CheckPromisedItemIsReachableAndReturnError(error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("checkPromisedItemIsReachableAndReturnError:"), error_)
	return rv
}


// Returns whether the resource pointed to by a file URL can be reached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/checkResourceIsReachableAndReturnError(_:)
func (u_ URL) CheckResourceIsReachableAndReturnError(error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("checkResourceIsReachableAndReturnError:"), error_)
	return rv
}


// Returns a new file reference URL that points to the same resource as the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileReferenceURL()
func (u_ URL) FileReferenceURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("fileReferenceURL"))
	return rv
}


// Fills the provided buffer with a C string representing a local file system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getFileSystemRepresentation(_:maxLength:)
func (u_ URL) GetFileSystemRepresentationMaxLength(buffer unsafe.Pointer, maxBufferLength uint /* primitive/slice/pointer */) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("getFileSystemRepresentation:maxLength:"), buffer, maxBufferLength)
	return rv
}


// Returns the value of the resource property for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getPromisedItemResourceValue(_:forKey:)
func (u_ URL) GetPromisedItemResourceValueForKeyError(value objectivec.IObject, key URLResourceKey /* foo */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("getPromisedItemResourceValue:forKey:error:"), value, key, error_)
	return rv
}


// Returns the value of the resource property for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getResourceValue(_:forKey:)
func (u_ URL) GetResourceValueForKeyError(value objectivec.IObject, key URLResourceKey /* foo */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("getResourceValue:forKey:error:"), value, key, error_)
	return rv
}


// Returns whether the URL is a file reference URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/isFileReferenceURL()
func (u_ URL) IsFileReferenceURL() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("isFileReferenceURL"))
	return rv
}


// Returns the resource values for the properties identified by specified array of keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/promisedItemResourceValues(forKeys:)
func (u_ URL) PromisedItemResourceValuesForKeysError(keys []string /* primitive/slice/pointer */, error_ unsafe.Pointer) IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](u_.ID, objc.Sel("promisedItemResourceValuesForKeys:error:"), keys, error_)
	return rv
}


// Removes all cached resource values and temporary resource values from the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/removeAllCachedResourceValues()
func (u_ URL) RemoveAllCachedResourceValues() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllCachedResourceValues"))
}


// Removes the cached resource value identified by a given key from the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/removeCachedResourceValue(forKey:)
func (u_ URL) RemoveCachedResourceValueForKey(key URLResourceKey /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeCachedResourceValueForKey:"), key)
}


// Returns the resource values for the properties identified by specified array of keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceValues(forKeys:)
func (u_ URL) ResourceValuesForKeysError(keys []string /* primitive/slice/pointer */, error_ unsafe.Pointer) IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](u_.ID, objc.Sel("resourceValuesForKeys:error:"), keys, error_)
	return rv
}


// Sets the URL’s resource property for a given key to a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValue(_:forKey:)
func (u_ URL) SetResourceValueForKeyError(value objectivec.IObject, key URLResourceKey /* foo */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceValue:forKey:error:"), value, key, error_)
	return rv
}


// Sets the URL’s resource properties for a given set of keys to a given set of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValues(_:)
func (u_ URL) SetResourceValuesError(keyedValues IDictionary /* already interface */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceValues:error:"), keyedValues, error_)
	return rv
}


// Sets a temporary resource value on the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setTemporaryResourceValue(_:forKey:)
func (u_ URL) SetTemporaryResourceValueForKey(value objectivec.IObject, key URLResourceKey /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTemporaryResourceValue:forKey:"), value, key)
}


// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/startAccessingSecurityScopedResource()
func (u_ URL) StartAccessingSecurityScopedResource() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("startAccessingSecurityScopedResource"))
	return rv
}


// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/stopAccessingSecurityScopedResource()
func (u_ URL) StopAccessingSecurityScopedResource() {
	objc.Send[objc.ID](u_.ID, objc.Sel("stopAccessingSecurityScopedResource"))
}


// Writes the URL to the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/write(to:)
func (u_ URL) WriteToPasteboard(pasteBoard Pasteboard /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("writeToPasteboard:"), pasteBoard)
}


// The URL string for the receiver as an absolute URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/absoluteString
func (u_ URL) AbsoluteString() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("absoluteString"))
	return rv
}


// An absolute URL that refers to the same resource as the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/absoluteURL
func (u_ URL) AbsoluteURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("absoluteURL"))
	return rv
}


// The base URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/baseURL
func (u_ URL) BaseURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("baseURL"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/dataRepresentation
func (u_ URL) DataRepresentation() IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("dataRepresentation"))
	return rv
}


// A URL you create by removing the last path component from the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/deletingLastPathComponent
func (u_ URL) URLByDeletingLastPathComponent() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByDeletingLastPathComponent"))
	return rv
}


// A URL you create by removing the path extension from the receiver, if any. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/deletingPathExtension
func (u_ URL) URLByDeletingPathExtension() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByDeletingPathExtension"))
	return rv
}


// A file path URL that points to the same resource as the URL object. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/filePathURL
func (u_ URL) FilePathURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("filePathURL"))
	return rv
}


// A C string containing the URL’s file system path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileSystemRepresentation
func (u_ URL) FileSystemRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("fileSystemRepresentation"))
	return rv
}


// The fragment identifier, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fragment
func (u_ URL) Fragment() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("fragment"))
	return rv
}


// A Boolean value that indicates whether the URL string’s path represents a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/hasDirectoryPath
func (u_ URL) HasDirectoryPath() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasDirectoryPath"))
	return rv
}


// The host, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/host
func (u_ URL) Host() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("host"))
	return rv
}


// A boolean value that determines whether the receiver is a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/isFileURL
func (u_ URL) FileURL() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("fileURL"))
	return rv
}


// The last path component. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/lastPathComponent
func (u_ URL) LastPathComponent() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("lastPathComponent"))
	return rv
}


// The parameter string conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/parameterString
func (u_ URL) ParameterString() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("parameterString"))
	return rv
}


// The password conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/password
func (u_ URL) Password() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("password"))
	return rv
}


// The path, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/path
func (u_ URL) Path() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("path"))
	return rv
}


// An array containing the path components. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/pathComponents
func (u_ URL) PathComponents() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](u_.ID, objc.Sel("pathComponents"))
	return rv
}


// The path extension. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/pathExtension
func (u_ URL) PathExtension() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("pathExtension"))
	return rv
}


// The port, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/port
func (u_ URL) Port() Number /* foo */ {
	rv := objc.Send[Number](u_.ID, objc.Sel("port"))
	return rv
}


// The query string, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/query
func (u_ URL) Query() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("query"))
	return rv
}


// The relative path, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/relativePath
func (u_ URL) RelativePath() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("relativePath"))
	return rv
}


// A string representation of the relative portion of the URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/relativeString
func (u_ URL) RelativeString() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("relativeString"))
	return rv
}


// A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resolvingSymlinksInPath
func (u_ URL) URLByResolvingSymlinksInPath() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByResolvingSymlinksInPath"))
	return rv
}


// The resource specifier. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceSpecifier
func (u_ URL) ResourceSpecifier() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("resourceSpecifier"))
	return rv
}


// The scheme. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/scheme
func (u_ URL) Scheme() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("scheme"))
	return rv
}


// A copy of the URL with any instances of or removed from its path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/standardized
func (u_ URL) StandardizedURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("standardizedURL"))
	return rv
}


// A URL that points to the same resource as the original URL using an absolute path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/standardizingPath
func (u_ URL) URLByStandardizingPath() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByStandardizingPath"))
	return rv
}


// The user name, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/user
func (u_ URL) User() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("user"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/customplaygroundquicklook
func (u_ URL) CustomPlaygroundQuickLook() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("customPlaygroundQuickLook"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/customplaygroundquicklook
func (u_ URL) SetCustomPlaygroundQuickLook(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCustomPlaygroundQuickLook:"), value)
}


// A URL you create by removing the last path component from the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/deletinglastpathcomponent
func (u_ URL) DeletingLastPathComponent() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("deletingLastPathComponent"))
	return rv
}


// A URL you create by removing the last path component from the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/deletinglastpathcomponent
func (u_ URL) SetDeletingLastPathComponent(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeletingLastPathComponent:"), value)
}


// A URL you create by removing the path extension from the receiver, if any. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/deletingpathextension
func (u_ URL) DeletingPathExtension() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("deletingPathExtension"))
	return rv
}


// A URL you create by removing the path extension from the receiver, if any. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/deletingpathextension
func (u_ URL) SetDeletingPathExtension(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeletingPathExtension:"), value)
}


// A boolean value that determines whether the receiver is a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/isfileurl
func (u_ URL) IsFileURL() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("isFileURL"))
	return rv
}


// A boolean value that determines whether the receiver is a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/isfileurl
func (u_ URL) SetIsFileURL(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsFileURL:"), value)
}


// A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/resolvingsymlinksinpath
func (u_ URL) ResolvingSymlinksInPath() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("resolvingSymlinksInPath"))
	return rv
}


// A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/resolvingsymlinksinpath
func (u_ URL) SetResolvingSymlinksInPath(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResolvingSymlinksInPath:"), value)
}


// A copy of the URL with any instances of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/standardized
func (u_ URL) Standardized() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("standardized"))
	return rv
}


// A copy of the URL with any instances of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/standardized
func (u_ URL) SetStandardized(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setStandardized:"), value)
}


// A URL that points to the same resource as the original URL using an absolute path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/standardizingpath
func (u_ URL) StandardizingPath() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("standardizingPath"))
	return rv
}


// A URL that points to the same resource as the original URL using an absolute path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/standardizingpath
func (u_ URL) SetStandardizingPath(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setStandardizingPath:"), value)
}


