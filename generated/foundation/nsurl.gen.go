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
	URLHandleUsingCache(shouldUseCache bool) unsafe.Pointer
	URLByAppendingPathComponent(pathComponent string) URL
	URLByAppendingPathComponentConformingToType(partialName string, contentType unsafe.Pointer) URL
	URLByAppendingPathComponentIsDirectory(pathComponent string, isDirectory bool) URL
	URLByAppendingPathExtension(pathExtension string) URL
	URLByAppendingPathExtensionForType(contentType unsafe.Pointer) URL
	BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions, keys []string, relativeURL IURL, error_ IError) Data
	CheckPromisedItemIsReachableAndReturnError(error_ IError) bool
	CheckResourceIsReachableAndReturnError(error_ IError) bool
	FileReferenceURL() URL
	GetFileSystemRepresentationMaxLength(buffer unsafe.Pointer, maxBufferLength uint) bool
	GetPromisedItemResourceValueForKeyError(value objectivec.IObject, key URLResourceKey, error_ IError) bool
	GetResourceValueForKeyError(value objectivec.IObject, key URLResourceKey, error_ IError) bool
	IsFileReferenceURL() bool
	LoadResourceDataNotifyingClientUsingCache(client objectivec.IObject, shouldUseCache bool)
	PromisedItemResourceValuesForKeysError(keys []string, error_ IError) unsafe.Pointer
	PropertyForKey(propertyKey string) objc.ID
	RemoveAllCachedResourceValues()
	RemoveCachedResourceValueForKey(key URLResourceKey)
	ResourceDataUsingCache(shouldUseCache bool) Data
	ResourceValuesForKeysError(keys []string, error_ IError) unsafe.Pointer
	SetPropertyForKey(property objectivec.IObject, propertyKey string) bool
	SetResourceData(data IData) bool
	SetResourceValueForKeyError(value objectivec.IObject, key URLResourceKey, error_ IError) bool
	SetResourceValuesError(keyedValues unsafe.Pointer, error_ IError) bool
	SetTemporaryResourceValueForKey(value objectivec.IObject, key URLResourceKey)
	StartAccessingSecurityScopedResource() bool
	StopAccessingSecurityScopedResource()
	AbsoluteString() string
	AbsoluteURL() URL
	BaseURL() URL
	DataRepresentation() NSData
	URLByDeletingLastPathComponent() URL
	URLByDeletingPathExtension() URL
	FilePathURL() URL
	FileSystemRepresentation() unsafe.Pointer
	Fragment() string
	HasDirectoryPath() bool
	Host() string
	FileURL() bool
	LastPathComponent() string
	ParameterString() string
	Password() string
	Path() string
	PathComponents() []string
	PathExtension() string
	Port() Number
	Query() string
	RelativePath() string
	RelativeString() string
	URLByResolvingSymlinksInPath() URL
	ResourceSpecifier() string
	Scheme() string
	StandardizedURL() URL
	URLByStandardizingPath() URL
	User() string
	CustomPlaygroundQuickLook() unsafe.Pointer
	SetCustomPlaygroundQuickLook(value unsafe.Pointer)
	DeletingLastPathComponent() URL
	SetDeletingLastPathComponent(value IURL)
	DeletingPathExtension() URL
	SetDeletingPathExtension(value IURL)
	IsFileURL() bool
	SetIsFileURL(value bool)
	ResolvingSymlinksInPath() URL
	SetResolvingSymlinksInPath(value IURL)
	Standardized() URL
	SetStandardized(value IURL)
	StandardizingPath() URL
	SetStandardizingPath(value IURL)
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

func NewURLByResolvingAliasFileAtURLOptionsError(url IURL, options URLBookmarkResolutionOptions, error_ IError) URL {
	rv := objc.Send[URL](objc.ID(getURLClass().class), objc.Sel("URLByResolvingAliasFileAtURL:options:error:"), url, options, error_)
	return rv
}



// Initializes a newly created NSURL that points to a location specified by resolving bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingBookmarkData:options:relativeTo:bookmarkDataIsStale:)

func NewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData IData, options URLBookmarkResolutionOptions, relativeURL IURL, isStale unsafe.Pointer, error_ IError) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, error_)
	rv.Autorelease()
	return rv
}



// Initializes a URL object with a C string representing a local file system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithFileSystemRepresentation:isDirectory:relativeTo:)

func NewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path unsafe.Pointer, isDir bool, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, baseURL)
	rv.Autorelease()
	return rv
}



// Initializes a newly created NSURL referencing the local file or directory at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:)

func NewURLFileURLWithPath(path string) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:"), objc.String(path))
	rv.Autorelease()
	return rv
}



// Initializes a newly created NSURL referencing the local file or directory at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:)

func NewURLFileURLWithPathIsDirectory(path string, isDir bool) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:"), objc.String(path), isDir)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:relativeTo:)

func NewURLFileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:relativeToURL:"), objc.String(path), isDir, baseURL)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:relativeTo:)

func NewURLFileURLWithPathRelativeToURL(path string, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:relativeToURL:"), objc.String(path), baseURL)
	rv.Autorelease()
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

func NewURLWithSchemeHostPath(scheme string, host string, path string) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithScheme:host:path:"), objc.String(scheme), objc.String(host), objc.String(path))
	rv.Autorelease()
	return rv
}



// Initializes an NSURL object with a provided URL string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:)

func NewURLWithString(URLString string) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:"), objc.String(URLString))
	rv.Autorelease()
	return rv
}



// Creates an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:encodingInvalidCharacters:)

func NewURLWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	rv.Autorelease()
	return rv
}



// Initializes an NSURL object with a base URL and a relative string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:relativeTo:)

func NewURLWithStringRelativeToURL(URLString string, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:relativeToURL:"), objc.String(URLString), baseURL)
	rv.Autorelease()
	return rv
}



// Returns a new URL made by resolving bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:

func (uc _URLClass) URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData IData, options URLBookmarkResolutionOptions, relativeURL IURL, isStale unsafe.Pointer, error_ IError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithDataRepresentation:relativeToURL:

func (uc _URLClass) URLWithDataRepresentationRelativeToURL(data IData, baseURL IURL) URL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("URLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return rv
}


// Creates and returns an NSURL object initialized with a provided URL string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:

func (uc _URLClass) URLWithString(URLString string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLWithString:"), objc.String(URLString))
	return rv
}


// Creates and returns an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:encodingInvalidCharacters:

func (uc _URLClass) URLWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	return rv
}


// Creates and returns an NSURL object initialized with a base URL and a relative string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:relativeToURL:

func (uc _URLClass) URLWithStringRelativeToURL(URLString string, baseURL IURL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLWithString:relativeToURL:"), objc.String(URLString), baseURL)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/absoluteURL(withDataRepresentation:relativeTo:)

func (uc _URLClass) AbsoluteURLWithDataRepresentationRelativeToURL(data IData, baseURL IURL) URL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("absoluteURLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return rv
}


// Initializes and returns bookmark data derived from an alias file pointed to by a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/bookmarkData(withContentsOf:)

func (uc _URLClass) BookmarkDataWithContentsOfURLError(bookmarkFileURL IURL, error_ IError) Data {
	rv := objc.Send[Data](objc.ID(uc.class), objc.Sel("bookmarkDataWithContentsOfURL:error:"), bookmarkFileURL, error_)
	return rv
}


// Returns a new URL object initialized with a C string representing a local file system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withFileSystemRepresentation:isDirectory:relativeTo:)

func (uc _URLClass) FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path unsafe.Pointer, isDir bool, baseURL IURL) URL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, baseURL)
	return rv
}


// Initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:)

func (uc _URLClass) FileURLWithPath(path string) URL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:"), objc.String(path))
	return rv
}


// Initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:)

func (uc _URLClass) FileURLWithPathIsDirectory(path string, isDir bool) URL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:isDirectory:"), objc.String(path), isDir)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:relativeTo:)

func (uc _URLClass) FileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL IURL) URL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:isDirectory:relativeToURL:"), objc.String(path), isDir, baseURL)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:relativeTo:)

func (uc _URLClass) FileURLWithPathRelativeToURL(path string, baseURL IURL) URL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:relativeToURL:"), objc.String(path), baseURL)
	return rv
}


// Initializes and returns a newly created NSURL object as a file URL with specified path components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPathComponents:)

func (uc _URLClass) FileURLWithPathComponents(components []string) URL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPathComponents:"), components)
	return rv
}


// Returns a new URL made by resolving the alias file at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingAliasFileAt:options:)

func (uc _URLClass) URLByResolvingAliasFileAtURLOptionsError(url IURL, options URLBookmarkResolutionOptions, error_ IError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLByResolvingAliasFileAtURL:options:error:"), url, options, error_)
	return rv
}


// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceValues(forKeys:fromBookmarkData:)

func (uc _URLClass) ResourceValuesForKeysFromBookmarkData(keys []string, bookmarkData IData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("resourceValuesForKeys:fromBookmarkData:"), keys, bookmarkData)
	return rv
}


// Creates an alias file on disk at a specified location with specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/writeBookmarkData(_:to:options:)

func (uc _URLClass) WriteBookmarkDataToURLOptionsError(bookmarkData IData, bookmarkFileURL IURL, options URLBookmarkFileCreationOptions, error_ IError) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("writeBookmarkData:toURL:options:error:"), bookmarkData, bookmarkFileURL, options, error_)
	return rv
}



// Returns a URL handle to service the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLHandleUsingCache:

func (u_ URL) URLHandleUsingCache(shouldUseCache bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("URLHandleUsingCache:"), shouldUseCache)
	return rv
}



// Returns a new URL by appending a path component to the original URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:)

func (u_ URL) URLByAppendingPathComponent(pathComponent string) URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathComponent:"), objc.String(pathComponent))
	return rv
}



// Returns a URL by appending the specified path component with the file extension for a uniform type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:conformingTo:)

func (u_ URL) URLByAppendingPathComponentConformingToType(partialName string, contentType unsafe.Pointer) URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathComponent:conformingToType:"), objc.String(partialName), contentType)
	return rv
}



// Returns a new URL by appending a path component to the original URL, along with a trailing slash if the component is a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:isDirectory:)

func (u_ URL) URLByAppendingPathComponentIsDirectory(pathComponent string, isDirectory bool) URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathComponent:isDirectory:"), objc.String(pathComponent), isDirectory)
	return rv
}



// Returns a new URL by appending a path extension to the original URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathExtension(_:)

func (u_ URL) URLByAppendingPathExtension(pathExtension string) URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathExtension:"), objc.String(pathExtension))
	return rv
}



// Returns a URL by appending the path extension for a uniform type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathExtension(for:)

func (u_ URL) URLByAppendingPathExtensionForType(contentType unsafe.Pointer) URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathExtensionForType:"), contentType)
	return rv
}



// Returns a bookmark for the URL, created with specified options and resource values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/bookmarkData(options:includingResourceValuesForKeys:relativeTo:)

func (u_ URL) BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions, keys []string, relativeURL IURL, error_ IError) Data {
	rv := objc.Send[Data](u_.ID, objc.Sel("bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:"), options, keys, relativeURL, error_)
	return rv
}



// Returns whether the promised item can be reached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/checkPromisedItemIsReachableAndReturnError(_:)

func (u_ URL) CheckPromisedItemIsReachableAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("checkPromisedItemIsReachableAndReturnError:"), error_)
	return rv
}



// Returns whether the resource pointed to by a file URL can be reached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/checkResourceIsReachableAndReturnError(_:)

func (u_ URL) CheckResourceIsReachableAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("checkResourceIsReachableAndReturnError:"), error_)
	return rv
}



// Returns a new file reference URL that points to the same resource as the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileReferenceURL()

func (u_ URL) FileReferenceURL() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("fileReferenceURL"))
	return rv
}



// Fills the provided buffer with a C string representing a local file system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getFileSystemRepresentation(_:maxLength:)

func (u_ URL) GetFileSystemRepresentationMaxLength(buffer unsafe.Pointer, maxBufferLength uint) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("getFileSystemRepresentation:maxLength:"), buffer, maxBufferLength)
	return rv
}



// Returns the value of the resource property for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getPromisedItemResourceValue(_:forKey:)

func (u_ URL) GetPromisedItemResourceValueForKeyError(value objectivec.IObject, key URLResourceKey, error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("getPromisedItemResourceValue:forKey:error:"), value, key, error_)
	return rv
}



// Returns the value of the resource property for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getResourceValue(_:forKey:)

func (u_ URL) GetResourceValueForKeyError(value objectivec.IObject, key URLResourceKey, error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("getResourceValue:forKey:error:"), value, key, error_)
	return rv
}



// Returns whether the URL is a file reference URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/isFileReferenceURL()

func (u_ URL) IsFileReferenceURL() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isFileReferenceURL"))
	return rv
}



// Loads the receiver’s resource data in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/loadResourceDataNotifyingClient:usingCache:

func (u_ URL) LoadResourceDataNotifyingClientUsingCache(client objectivec.IObject, shouldUseCache bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("loadResourceDataNotifyingClient:usingCache:"), client, shouldUseCache)
}



// Returns the resource values for the properties identified by specified array of keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/promisedItemResourceValues(forKeys:)

func (u_ URL) PromisedItemResourceValuesForKeysError(keys []string, error_ IError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("promisedItemResourceValuesForKeys:error:"), keys, error_)
	return rv
}



// Returns the specified property of the receiver’s resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/propertyForKey:

func (u_ URL) PropertyForKey(propertyKey string) objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("propertyForKey:"), objc.String(propertyKey))
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

func (u_ URL) RemoveCachedResourceValueForKey(key URLResourceKey) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeCachedResourceValueForKey:"), key)
}



// Returns the receiver’s resource data, loading it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceDataUsingCache:

func (u_ URL) ResourceDataUsingCache(shouldUseCache bool) Data {
	rv := objc.Send[Data](u_.ID, objc.Sel("resourceDataUsingCache:"), shouldUseCache)
	return rv
}



// Returns the resource values for the properties identified by specified array of keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceValues(forKeys:)

func (u_ URL) ResourceValuesForKeysError(keys []string, error_ IError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("resourceValuesForKeys:error:"), keys, error_)
	return rv
}



// Changes the specified property of the receiver’s resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setProperty:forKey:

func (u_ URL) SetPropertyForKey(property objectivec.IObject, propertyKey string) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setProperty:forKey:"), property, objc.String(propertyKey))
	return rv
}



// Attempts to set the resource data for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceData:

func (u_ URL) SetResourceData(data IData) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceData:"), data)
	return rv
}



// Sets the URL’s resource property for a given key to a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValue(_:forKey:)

func (u_ URL) SetResourceValueForKeyError(value objectivec.IObject, key URLResourceKey, error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceValue:forKey:error:"), value, key, error_)
	return rv
}



// Sets the URL’s resource properties for a given set of keys to a given set of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValues(_:)

func (u_ URL) SetResourceValuesError(keyedValues unsafe.Pointer, error_ IError) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceValues:error:"), keyedValues, error_)
	return rv
}



// Sets a temporary resource value on the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setTemporaryResourceValue(_:forKey:)

func (u_ URL) SetTemporaryResourceValueForKey(value objectivec.IObject, key URLResourceKey) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTemporaryResourceValue:forKey:"), value, key)
}



// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/startAccessingSecurityScopedResource()

func (u_ URL) StartAccessingSecurityScopedResource() bool {
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


// The URL string for the receiver as an absolute URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/absoluteString

func (u_ URL) AbsoluteString() string {
	rv := objc.Send[string](u_.ID, objc.Sel("absoluteString"))
	return rv
}


// An absolute URL that refers to the same resource as the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/absoluteURL

func (u_ URL) AbsoluteURL() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("absoluteURL"))
	return rv
}


// The base URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/baseURL

func (u_ URL) BaseURL() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("baseURL"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/dataRepresentation

func (u_ URL) DataRepresentation() NSData {
	rv := objc.Send[NSData](u_.ID, objc.Sel("dataRepresentation"))
	return rv
}


// A URL you create by removing the last path component from the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/deletingLastPathComponent

func (u_ URL) URLByDeletingLastPathComponent() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByDeletingLastPathComponent"))
	return rv
}


// A URL you create by removing the path extension from the receiver, if any. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/deletingPathExtension

func (u_ URL) URLByDeletingPathExtension() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByDeletingPathExtension"))
	return rv
}


// A file path URL that points to the same resource as the URL object. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/filePathURL

func (u_ URL) FilePathURL() URL {
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

func (u_ URL) Fragment() string {
	rv := objc.Send[string](u_.ID, objc.Sel("fragment"))
	return rv
}


// A Boolean value that indicates whether the URL string’s path represents a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/hasDirectoryPath

func (u_ URL) HasDirectoryPath() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasDirectoryPath"))
	return rv
}


// The host, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/host

func (u_ URL) Host() string {
	rv := objc.Send[string](u_.ID, objc.Sel("host"))
	return rv
}


// A boolean value that determines whether the receiver is a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/isFileURL

func (u_ URL) FileURL() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("fileURL"))
	return rv
}


// The last path component. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/lastPathComponent

func (u_ URL) LastPathComponent() string {
	rv := objc.Send[string](u_.ID, objc.Sel("lastPathComponent"))
	return rv
}


// The parameter string conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/parameterString

func (u_ URL) ParameterString() string {
	rv := objc.Send[string](u_.ID, objc.Sel("parameterString"))
	return rv
}


// The password conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/password

func (u_ URL) Password() string {
	rv := objc.Send[string](u_.ID, objc.Sel("password"))
	return rv
}


// The path, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/path

func (u_ URL) Path() string {
	rv := objc.Send[string](u_.ID, objc.Sel("path"))
	return rv
}


// An array containing the path components. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/pathComponents

func (u_ URL) PathComponents() []string {
	rv := objc.Send[[]string](u_.ID, objc.Sel("pathComponents"))
	return rv
}


// The path extension. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/pathExtension

func (u_ URL) PathExtension() string {
	rv := objc.Send[string](u_.ID, objc.Sel("pathExtension"))
	return rv
}


// The port, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/port

func (u_ URL) Port() Number {
	rv := objc.Send[Number](u_.ID, objc.Sel("port"))
	return rv
}


// The query string, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/query

func (u_ URL) Query() string {
	rv := objc.Send[string](u_.ID, objc.Sel("query"))
	return rv
}


// The relative path, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/relativePath

func (u_ URL) RelativePath() string {
	rv := objc.Send[string](u_.ID, objc.Sel("relativePath"))
	return rv
}


// A string representation of the relative portion of the URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/relativeString

func (u_ URL) RelativeString() string {
	rv := objc.Send[string](u_.ID, objc.Sel("relativeString"))
	return rv
}


// A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resolvingSymlinksInPath

func (u_ URL) URLByResolvingSymlinksInPath() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByResolvingSymlinksInPath"))
	return rv
}


// The resource specifier. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceSpecifier

func (u_ URL) ResourceSpecifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("resourceSpecifier"))
	return rv
}


// The scheme. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/scheme

func (u_ URL) Scheme() string {
	rv := objc.Send[string](u_.ID, objc.Sel("scheme"))
	return rv
}


// A copy of the URL with any instances of or removed from its path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/standardized

func (u_ URL) StandardizedURL() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("standardizedURL"))
	return rv
}


// A URL that points to the same resource as the original URL using an absolute path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/standardizingPath

func (u_ URL) URLByStandardizingPath() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByStandardizingPath"))
	return rv
}


// The user name, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/user

func (u_ URL) User() string {
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

func (u_ URL) DeletingLastPathComponent() URL {
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

func (u_ URL) DeletingPathExtension() URL {
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

func (u_ URL) IsFileURL() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isFileURL"))
	return rv
}


// A boolean value that determines whether the receiver is a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/isfileurl

func (u_ URL) SetIsFileURL(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsFileURL:"), value)
}


// A URL that points to the same resource as the receiver and includes no symbolic links. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/resolvingsymlinksinpath

func (u_ URL) ResolvingSymlinksInPath() URL {
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

func (u_ URL) Standardized() URL {
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

func (u_ URL) StandardizingPath() URL {
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


