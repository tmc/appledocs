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
	uRLClass     _URLClass
	uRLClassOnce sync.Once
)

func getURLClass() _URLClass {
	uRLClassOnce.Do(func() {
		uRLClass = _URLClass{objc.GetClass("NSURL")}
	})
	return uRLClass
}

type _URLClass struct {
	class objc.Class
}

// An interface definition for the [URL] class.
type IURL interface {
	objectivec.IObject
	URLHandleUsingCache(shouldUseCache bool) unsafe.Pointer
	URLByAppendingPathComponent(pathComponent string) unsafe.Pointer
	URLByAppendingPathComponentConformingToType(partialName string, contentType unsafe.Pointer) unsafe.Pointer
	URLByAppendingPathComponentIsDirectory(pathComponent string, isDirectory bool) unsafe.Pointer
	URLByAppendingPathExtension(pathExtension string) unsafe.Pointer
	URLByAppendingPathExtensionForType(contentType unsafe.Pointer) unsafe.Pointer
	BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options unsafe.Pointer, keys unsafe.Pointer, relativeURL unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	CheckPromisedItemIsReachableAndReturnError(error unsafe.Pointer) bool
	CheckResourceIsReachableAndReturnError(error unsafe.Pointer) bool
	FileReferenceURL() unsafe.Pointer
	GetFileSystemRepresentationMaxLength(buffer unsafe.Pointer, maxBufferLength uint) bool
	GetPromisedItemResourceValueForKeyError(value unsafe.Pointer, key unsafe.Pointer, error unsafe.Pointer) bool
	GetResourceValueForKeyError(value unsafe.Pointer, key unsafe.Pointer, error unsafe.Pointer) bool
	IsFileReferenceURL() bool
	LoadResourceDataNotifyingClientUsingCache(client objc.ID, shouldUseCache bool)
	PromisedItemResourceValuesForKeysError(keys unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	PropertyForKey(propertyKey string) objc.ID
	RemoveAllCachedResourceValues()
	RemoveCachedResourceValueForKey(key unsafe.Pointer)
	ResourceDataUsingCache(shouldUseCache bool) unsafe.Pointer
	ResourceValuesForKeysError(keys unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	SetPropertyForKey(property objc.ID, propertyKey string) bool
	SetResourceData(data unsafe.Pointer) bool
	SetResourceValueForKeyError(value objc.ID, key unsafe.Pointer, error unsafe.Pointer) bool
	SetResourceValuesError(keyedValues unsafe.Pointer, error unsafe.Pointer) bool
	SetTemporaryResourceValueForKey(value objc.ID, key unsafe.Pointer)
	StartAccessingSecurityScopedResource() bool
	StopAccessingSecurityScopedResource()
	WriteToPasteboard(pasteBoard unsafe.Pointer)
}

// An object that represents the location of a resource, such as an item on a remote server or the path to a local file.
//
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


// Reads an NSURL object off of the specified pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fromPasteboard:)
func NewURLFromPasteboard(pasteBoard unsafe.Pointer) URL {
	rv := objc.Send[URL](objc.ID(getURLClass().class), objc.Sel("URLFromPasteboard:"), pasteBoard)
	return rv
}
// Initializes a newly created NSURL that points to a location specified by resolving bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingBookmarkData:options:relativeTo:bookmarkDataIsStale:)
func NewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData unsafe.Pointer, options unsafe.Pointer, relativeURL unsafe.Pointer, isStale unsafe.Pointer, error unsafe.Pointer) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, error)
	rv.Autorelease()
	return rv
}
// Initializes a newly created NSURL with a specified scheme, host, and path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(scheme:host:path:)
func NewURLWithSchemeHostPath(scheme string, host string, path string) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithScheme:host:path:"), objc.String(scheme), objc.String(host), objc.String(path))
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(absoluteURLWithDataRepresentation:relativeTo:)
func NewURLAbsoluteURLWithDataRepresentationRelativeToURL(data unsafe.Pointer, baseURL unsafe.Pointer) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initAbsoluteURLWithDataRepresentation:relativeToURL:"), data, baseURL)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(dataRepresentation:relativeTo:)
func NewURLWithDataRepresentationRelativeToURL(data unsafe.Pointer, baseURL unsafe.Pointer) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithDataRepresentation:relativeToURL:"), data, baseURL)
	rv.Autorelease()
	return rv
}
// Returns a new URL made by resolving the alias file at .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingAliasFileAt:options:)
func NewURLByResolvingAliasFileAtURLOptionsError(url unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) URL {
	rv := objc.Send[URL](objc.ID(getURLClass().class), objc.Sel("URLByResolvingAliasFileAtURL:options:error:"), url, options, error)
	return rv
}
// Initializes an NSURL object with a provided URL string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:)
func NewURLWithString(URLString string) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:"), objc.String(URLString))
	rv.Autorelease()
	return rv
}
// Creates an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:encodingInvalidCharacters:)
func NewURLWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	rv.Autorelease()
	return rv
}
// Initializes an NSURL object with a base URL and a relative string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(string:relativeTo:)
func NewURLWithStringRelativeToURL(URLString string, baseURL unsafe.Pointer) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initWithString:relativeToURL:"), objc.String(URLString), baseURL)
	rv.Autorelease()
	return rv
}
// Initializes a URL object with a C string representing a local file system path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithFileSystemRepresentation:isDirectory:relativeTo:)
func NewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path unsafe.Pointer, isDir bool, baseURL unsafe.Pointer) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, baseURL)
	rv.Autorelease()
	return rv
}
// Initializes a newly created NSURL referencing the local file or directory at .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:)
func NewURLFileURLWithPath(path string) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:"), objc.String(path))
	rv.Autorelease()
	return rv
}
// Initializes a newly created NSURL referencing the local file or directory at .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:)
func NewURLFileURLWithPathIsDirectory(path string, isDir bool) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:"), objc.String(path), isDir)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:relativeTo:)
func NewURLFileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL unsafe.Pointer) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:relativeToURL:"), objc.String(path), isDir, baseURL)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:relativeTo:)
func NewURLFileURLWithPathRelativeToURL(path string, baseURL unsafe.Pointer) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:relativeToURL:"), objc.String(path), baseURL)
	rv.Autorelease()
	return rv
}


// Returns a new URL made by resolving bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:
func (uc _URLClass) URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData unsafe.Pointer, options unsafe.Pointer, relativeURL unsafe.Pointer, isStale unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, error)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithDataRepresentation:relativeToURL:
func (uc _URLClass) URLWithDataRepresentationRelativeToURL(data unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return rv
}
// Creates and returns an NSURL object initialized with a provided URL string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:
func (uc _URLClass) URLWithString(URLString string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLWithString:"), objc.String(URLString))
	return rv
}
// Creates and returns an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:encodingInvalidCharacters:
func (uc _URLClass) URLWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	return rv
}
// Creates and returns an NSURL object initialized with a base URL and a relative string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLWithString:relativeToURL:
func (uc _URLClass) URLWithStringRelativeToURL(URLString string, baseURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLWithString:relativeToURL:"), objc.String(URLString), baseURL)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/absoluteURL(withDataRepresentation:relativeTo:)
func (uc _URLClass) AbsoluteURLWithDataRepresentationRelativeToURL(data unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("absoluteURLWithDataRepresentation:relativeToURL:"), data, baseURL)
	return rv
}
// Initializes and returns bookmark data derived from an alias file pointed to by a specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/bookmarkData(withContentsOf:)
func (uc _URLClass) BookmarkDataWithContentsOfURLError(bookmarkFileURL unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("bookmarkDataWithContentsOfURL:error:"), bookmarkFileURL, error)
	return rv
}
// Returns a new URL object initialized with a C string representing a local file system path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withFileSystemRepresentation:isDirectory:relativeTo:)
func (uc _URLClass) FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path unsafe.Pointer, isDir bool, baseURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("fileURLWithFileSystemRepresentation:isDirectory:relativeToURL:"), path, isDir, baseURL)
	return rv
}
// Initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:)
func (uc _URLClass) FileURLWithPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("fileURLWithPath:"), objc.String(path))
	return rv
}
// Initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:)
func (uc _URLClass) FileURLWithPathIsDirectory(path string, isDir bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("fileURLWithPath:isDirectory:"), objc.String(path), isDir)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:relativeTo:)
func (uc _URLClass) FileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("fileURLWithPath:isDirectory:relativeToURL:"), objc.String(path), isDir, baseURL)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:relativeTo:)
func (uc _URLClass) FileURLWithPathRelativeToURL(path string, baseURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("fileURLWithPath:relativeToURL:"), objc.String(path), baseURL)
	return rv
}
// Initializes and returns a newly created NSURL object as a file URL with specified path components.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPathComponents:)
func (uc _URLClass) FileURLWithPathComponents(components unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("fileURLWithPathComponents:"), components)
	return rv
}
// Reads an NSURL object off of the specified pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fromPasteboard:)
func (uc _URLClass) URLFromPasteboard(pasteBoard unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLFromPasteboard:"), pasteBoard)
	return rv
}
// Returns a new URL made by resolving the alias file at .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(resolvingAliasFileAt:options:)
func (uc _URLClass) URLByResolvingAliasFileAtURLOptionsError(url unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLByResolvingAliasFileAtURL:options:error:"), url, options, error)
	return rv
}
// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceValues(forKeys:fromBookmarkData:)
func (uc _URLClass) ResourceValuesForKeysFromBookmarkData(keys unsafe.Pointer, bookmarkData unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("resourceValuesForKeys:fromBookmarkData:"), keys, bookmarkData)
	return rv
}
// Creates an alias file on disk at a specified location with specified bookmark data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/writeBookmarkData(_:to:options:)
func (uc _URLClass) WriteBookmarkDataToURLOptionsError(bookmarkData unsafe.Pointer, bookmarkFileURL unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("writeBookmarkData:toURL:options:error:"), bookmarkData, bookmarkFileURL, options, error)
	return rv
}
// Returns a URL handle to service the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLHandleUsingCache:
func (u_ URL) URLHandleUsingCache(shouldUseCache bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("URLHandleUsingCache:"), shouldUseCache)
	return rv
}
// Returns a new URL by appending a path component to the original URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:)
func (u_ URL) URLByAppendingPathComponent(pathComponent string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("URLByAppendingPathComponent:"), objc.String(pathComponent))
	return rv
}
// Returns a URL by appending the specified path component with the file extension for a uniform type identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:conformingTo:)
func (u_ URL) URLByAppendingPathComponentConformingToType(partialName string, contentType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("URLByAppendingPathComponent:conformingToType:"), objc.String(partialName), contentType)
	return rv
}
// Returns a new URL by appending a path component to the original URL, along with a trailing slash if the component is a directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:isDirectory:)
func (u_ URL) URLByAppendingPathComponentIsDirectory(pathComponent string, isDirectory bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("URLByAppendingPathComponent:isDirectory:"), objc.String(pathComponent), isDirectory)
	return rv
}
// Returns a new URL by appending a path extension to the original URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathExtension(_:)
func (u_ URL) URLByAppendingPathExtension(pathExtension string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("URLByAppendingPathExtension:"), objc.String(pathExtension))
	return rv
}
// Returns a URL by appending the path extension for a uniform type identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathExtension(for:)
func (u_ URL) URLByAppendingPathExtensionForType(contentType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("URLByAppendingPathExtensionForType:"), contentType)
	return rv
}
// Returns a bookmark for the URL, created with specified options and resource values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/bookmarkData(options:includingResourceValuesForKeys:relativeTo:)
func (u_ URL) BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options unsafe.Pointer, keys unsafe.Pointer, relativeURL unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:"), options, keys, relativeURL, error)
	return rv
}
// Returns whether the promised item can be reached.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/checkPromisedItemIsReachableAndReturnError(_:)
func (u_ URL) CheckPromisedItemIsReachableAndReturnError(error unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("checkPromisedItemIsReachableAndReturnError:"), error)
	return rv
}
// Returns whether the resource pointed to by a file URL can be reached.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/checkResourceIsReachableAndReturnError(_:)
func (u_ URL) CheckResourceIsReachableAndReturnError(error unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("checkResourceIsReachableAndReturnError:"), error)
	return rv
}
// Returns a new file reference URL that points to the same resource as the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileReferenceURL()
func (u_ URL) FileReferenceURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("fileReferenceURL"))
	return rv
}
// Fills the provided buffer with a C string representing a local file system path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getFileSystemRepresentation(_:maxLength:)
func (u_ URL) GetFileSystemRepresentationMaxLength(buffer unsafe.Pointer, maxBufferLength uint) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("getFileSystemRepresentation:maxLength:"), buffer, maxBufferLength)
	return rv
}
// Returns the value of the resource property for the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getPromisedItemResourceValue(_:forKey:)
func (u_ URL) GetPromisedItemResourceValueForKeyError(value unsafe.Pointer, key unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("getPromisedItemResourceValue:forKey:error:"), value, key, error)
	return rv
}
// Returns the value of the resource property for the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getResourceValue(_:forKey:)
func (u_ URL) GetResourceValueForKeyError(value unsafe.Pointer, key unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("getResourceValue:forKey:error:"), value, key, error)
	return rv
}
// Returns whether the URL is a file reference URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/isFileReferenceURL()
func (u_ URL) IsFileReferenceURL() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isFileReferenceURL"))
	return rv
}
// Loads the receiver’s resource data in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/loadResourceDataNotifyingClient:usingCache:
func (u_ URL) LoadResourceDataNotifyingClientUsingCache(client objc.ID, shouldUseCache bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("loadResourceDataNotifyingClient:usingCache:"), client, shouldUseCache)
}
// Returns the resource values for the properties identified by specified array of keys.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/promisedItemResourceValues(forKeys:)
func (u_ URL) PromisedItemResourceValuesForKeysError(keys unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("promisedItemResourceValuesForKeys:error:"), keys, error)
	return rv
}
// Returns the specified property of the receiver’s resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/propertyForKey:
func (u_ URL) PropertyForKey(propertyKey string) objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("propertyForKey:"), objc.String(propertyKey))
	return rv
}
// Removes all cached resource values and temporary resource values from the URL object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/removeAllCachedResourceValues()
func (u_ URL) RemoveAllCachedResourceValues() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllCachedResourceValues"))
}
// Removes the cached resource value identified by a given key from the URL object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/removeCachedResourceValue(forKey:)
func (u_ URL) RemoveCachedResourceValueForKey(key unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeCachedResourceValueForKey:"), key)
}
// Returns the receiver’s resource data, loading it if necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceDataUsingCache:
func (u_ URL) ResourceDataUsingCache(shouldUseCache bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("resourceDataUsingCache:"), shouldUseCache)
	return rv
}
// Returns the resource values for the properties identified by specified array of keys.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceValues(forKeys:)
func (u_ URL) ResourceValuesForKeysError(keys unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("resourceValuesForKeys:error:"), keys, error)
	return rv
}
// Changes the specified property of the receiver’s resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setProperty:forKey:
func (u_ URL) SetPropertyForKey(property objc.ID, propertyKey string) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setProperty:forKey:"), property, objc.String(propertyKey))
	return rv
}
// Attempts to set the resource data for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceData:
func (u_ URL) SetResourceData(data unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceData:"), data)
	return rv
}
// Sets the URL’s resource property for a given key to a given value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValue(_:forKey:)
func (u_ URL) SetResourceValueForKeyError(value objc.ID, key unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceValue:forKey:error:"), value, key, error)
	return rv
}
// Sets the URL’s resource properties for a given set of keys to a given set of values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValues(_:)
func (u_ URL) SetResourceValuesError(keyedValues unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceValues:error:"), keyedValues, error)
	return rv
}
// Sets a temporary resource value on the URL object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setTemporaryResourceValue(_:forKey:)
func (u_ URL) SetTemporaryResourceValueForKey(value objc.ID, key unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTemporaryResourceValue:forKey:"), value, key)
}
// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/startAccessingSecurityScopedResource()
func (u_ URL) StartAccessingSecurityScopedResource() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("startAccessingSecurityScopedResource"))
	return rv
}
// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/stopAccessingSecurityScopedResource()
func (u_ URL) StopAccessingSecurityScopedResource() {
	objc.Send[objc.ID](u_.ID, objc.Sel("stopAccessingSecurityScopedResource"))
}
// Writes the URL to the specified pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/write(to:)
func (u_ URL) WriteToPasteboard(pasteBoard unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("writeToPasteboard:"), pasteBoard)
}

