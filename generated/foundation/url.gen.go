// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URL] class.
var URLClass objc.Class

func init() {
	URLClass = objc.GetClass("NSURL")
}

type URL struct {
	objc.ID
}

func URLFrom(ptr unsafe.Pointer) URL {
	return URL{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc URL) Alloc() URL {
	ret := objc.ID(URLClass).Send(objc.RegisterName("alloc"))
	return URL{ret}
}

// Init initializes the instance.
func (u_ URL) Init() URL {
	ret := u_.ID.Send(objc.RegisterName("init"))
	return URL{ret}
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(absoluteURLWithDataRepresentation:relativeTo:)
func NewURLAbsoluteURLWithDataRepresentationRelativeToURL(data unsafe.Pointer, baseURL unsafe.Pointer) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initAbsoluteURLWithDataRepresentation:relativeToURL:")
	ret := instance.ID.Send(sel, data, baseURL)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(dataRepresentation:relativeTo:)
func NewURLWithDataRepresentationRelativeToURL(data unsafe.Pointer, baseURL unsafe.Pointer) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initWithDataRepresentation:relativeToURL:")
	ret := instance.ID.Send(sel, data, baseURL)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a URL object with a C string representing a local file system path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(fileURLWithFileSystemRepresentation:isDirectory:relativeTo:)
func NewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path unsafe.Pointer, isDir bool, baseURL unsafe.Pointer) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initFileURLWithFileSystemRepresentation:isDirectory:relativeToURL:")
	ret := instance.ID.Send(sel, path, isDir, baseURL)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly created NSURL referencing the local file or directory at  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(fileURLWithPath:)
func NewURLFileURLWithPath(path string) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initFileURLWithPath:")
	ret := instance.ID.Send(sel, path)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly created NSURL referencing the local file or directory at  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:)
func NewURLFileURLWithPathIsDirectory(path string, isDir bool) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initFileURLWithPath:isDirectory:")
	ret := instance.ID.Send(sel, path, isDir)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:relativeTo:)
func NewURLFileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL unsafe.Pointer) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initFileURLWithPath:isDirectory:relativeToURL:")
	ret := instance.ID.Send(sel, path, isDir, baseURL)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(fileURLWithPath:relativeTo:)
func NewURLFileURLWithPathRelativeToURL(path string, baseURL unsafe.Pointer) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initFileURLWithPath:relativeToURL:")
	ret := instance.ID.Send(sel, path, baseURL)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly created NSURL that points to a location specified by resolving bookmark data. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(resolvingBookmarkData:options:relativeTo:bookmarkDataIsStale:)
func NewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData unsafe.Pointer, options unsafe.Pointer, relativeURL unsafe.Pointer, isStale unsafe.Pointer, error unsafe.Pointer) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:")
	ret := instance.ID.Send(sel, bookmarkData, options, relativeURL, isStale, error)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly created NSURL with a specified scheme, host, and path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(scheme:host:path:)
func NewURLWithSchemeHostPath(scheme string, host string, path string) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initWithScheme:host:path:")
	ret := instance.ID.Send(sel, scheme, host, path)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes an NSURL object with a provided URL string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(string:)
func NewURLWithString(URLString string) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initWithString:")
	ret := instance.ID.Send(sel, URLString)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Creates an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(string:encodingInvalidCharacters:)
func NewURLWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initWithString:encodingInvalidCharacters:")
	ret := instance.ID.Send(sel, URLString, encodingInvalidCharacters)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes an NSURL object with a base URL and a relative string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(string:relativeTo:)
func NewURLWithStringRelativeToURL(URLString string, baseURL unsafe.Pointer) URL {
	instance := URL{}.Alloc()
	sel := objc.RegisterName("initWithString:relativeToURL:")
	ret := instance.ID.Send(sel, URLString, baseURL)
	instance = URL{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Returns a new URL made by resolving bookmark data. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:
func (uc URL) URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData unsafe.Pointer, options unsafe.Pointer, relativeURL unsafe.Pointer, isStale unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:")
	ret := objc.ID(URLClass).Send(sel, bookmarkData, options, relativeURL, isStale, error)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/URLWithDataRepresentation:relativeToURL:
func (uc URL) URLWithDataRepresentationRelativeToURL(data unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("URLWithDataRepresentation:relativeToURL:")
	ret := objc.ID(URLClass).Send(sel, data, baseURL)
	return unsafe.Pointer(ret)
}
// Creates and returns an NSURL object initialized with a provided URL string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/URLWithString:
func (uc URL) URLWithString(URLString string) unsafe.Pointer {
	sel := objc.RegisterName("URLWithString:")
	ret := objc.ID(URLClass).Send(sel, URLString)
	return unsafe.Pointer(ret)
}
// Creates and returns an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/URLWithString:encodingInvalidCharacters:
func (uc URL) URLWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) unsafe.Pointer {
	sel := objc.RegisterName("URLWithString:encodingInvalidCharacters:")
	ret := objc.ID(URLClass).Send(sel, URLString, encodingInvalidCharacters)
	return unsafe.Pointer(ret)
}
// Creates and returns an NSURL object initialized with a base URL and a relative string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/URLWithString:relativeToURL:
func (uc URL) URLWithStringRelativeToURL(URLString string, baseURL unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("URLWithString:relativeToURL:")
	ret := objc.ID(URLClass).Send(sel, URLString, baseURL)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/absoluteURL(withDataRepresentation:relativeTo:)
func (uc URL) AbsoluteURLWithDataRepresentationRelativeToURL(data unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("absoluteURLWithDataRepresentation:relativeToURL:")
	ret := objc.ID(URLClass).Send(sel, data, baseURL)
	return unsafe.Pointer(ret)
}
// Initializes and returns bookmark data derived from an alias file pointed to by a specified URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/bookmarkData(withContentsOf:)
func (uc URL) BookmarkDataWithContentsOfURLError(bookmarkFileURL unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("bookmarkDataWithContentsOfURL:error:")
	ret := objc.ID(URLClass).Send(sel, bookmarkFileURL, error)
	return unsafe.Pointer(ret)
}
// Returns a new URL object initialized with a C string representing a local file system path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/fileURL(withFileSystemRepresentation:isDirectory:relativeTo:)
func (uc URL) FileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(path unsafe.Pointer, isDir bool, baseURL unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("fileURLWithFileSystemRepresentation:isDirectory:relativeToURL:")
	ret := objc.ID(URLClass).Send(sel, path, isDir, baseURL)
	return unsafe.Pointer(ret)
}
// Initializes and returns a newly created NSURL object as a file URL with a specified path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/fileURL(withPath:)
func (uc URL) FileURLWithPath(path string) unsafe.Pointer {
	sel := objc.RegisterName("fileURLWithPath:")
	ret := objc.ID(URLClass).Send(sel, path)
	return unsafe.Pointer(ret)
}
// Initializes and returns a newly created NSURL object as a file URL with a specified path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:)
func (uc URL) FileURLWithPathIsDirectory(path string, isDir bool) unsafe.Pointer {
	sel := objc.RegisterName("fileURLWithPath:isDirectory:")
	ret := objc.ID(URLClass).Send(sel, path, isDir)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:relativeTo:)
func (uc URL) FileURLWithPathIsDirectoryRelativeToURL(path string, isDir bool, baseURL unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("fileURLWithPath:isDirectory:relativeToURL:")
	ret := objc.ID(URLClass).Send(sel, path, isDir, baseURL)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/fileURL(withPath:relativeTo:)
func (uc URL) FileURLWithPathRelativeToURL(path string, baseURL unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("fileURLWithPath:relativeToURL:")
	ret := objc.ID(URLClass).Send(sel, path, baseURL)
	return unsafe.Pointer(ret)
}
// Initializes and returns a newly created NSURL object as a file URL with specified path components. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/fileURL(withPathComponents:)
func (uc URL) FileURLWithPathComponents(components unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("fileURLWithPathComponents:")
	ret := objc.ID(URLClass).Send(sel, components)
	return unsafe.Pointer(ret)
}
// Reads an NSURL object off of the specified pasteboard. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(fromPasteboard:)
func (uc URL) URLFromPasteboard(pasteBoard unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("URLFromPasteboard:")
	ret := objc.ID(URLClass).Send(sel, pasteBoard)
	return unsafe.Pointer(ret)
}
// Returns a new URL made by resolving the alias file at  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/init(resolvingAliasFileAt:options:)
func (uc URL) URLByResolvingAliasFileAtURLOptionsError(url unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("URLByResolvingAliasFileAtURL:options:error:")
	ret := objc.ID(URLClass).Send(sel, url, options, error)
	return unsafe.Pointer(ret)
}
// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/resourceValues(forKeys:fromBookmarkData:)
func (uc URL) ResourceValuesForKeysFromBookmarkData(keys unsafe.Pointer, bookmarkData unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("resourceValuesForKeys:fromBookmarkData:")
	ret := objc.ID(URLClass).Send(sel, keys, bookmarkData)
	return unsafe.Pointer(ret)
}
// Creates an alias file on disk at a specified location with specified bookmark data. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/writeBookmarkData(_:to:options:)
func (uc URL) WriteBookmarkDataToURLOptionsError(bookmarkData unsafe.Pointer, bookmarkFileURL unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("writeBookmarkData:toURL:options:error:")
	ret := objc.ID(URLClass).Send(sel, bookmarkData, bookmarkFileURL, options, error)
	return bool(ret)
}
// Returns a URL handle to service the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/URLHandleUsingCache:
func (u_ URL) URLHandleUsingCache(shouldUseCache bool) unsafe.Pointer {
	sel := objc.RegisterName("URLHandleUsingCache:")
	ret := u_.ID.Send(sel, shouldUseCache)
	return unsafe.Pointer(ret)
}
// Returns a new URL by appending a path component to the original URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/appendingPathComponent(_:)
func (u_ URL) URLByAppendingPathComponent(pathComponent string) unsafe.Pointer {
	sel := objc.RegisterName("URLByAppendingPathComponent:")
	ret := u_.ID.Send(sel, pathComponent)
	return unsafe.Pointer(ret)
}
// Returns a URL by appending the specified path component with the file extension for a uniform type identifier. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/appendingPathComponent(_:conformingTo:)
func (u_ URL) URLByAppendingPathComponentConformingToType(partialName string, contentType unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("URLByAppendingPathComponent:conformingToType:")
	ret := u_.ID.Send(sel, partialName, contentType)
	return unsafe.Pointer(ret)
}
// Returns a new URL by appending a path component to the original URL, along with a trailing slash if the component is a directory. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/appendingPathComponent(_:isDirectory:)
func (u_ URL) URLByAppendingPathComponentIsDirectory(pathComponent string, isDirectory bool) unsafe.Pointer {
	sel := objc.RegisterName("URLByAppendingPathComponent:isDirectory:")
	ret := u_.ID.Send(sel, pathComponent, isDirectory)
	return unsafe.Pointer(ret)
}
// Returns a new URL by appending a path extension to the original URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/appendingPathExtension(_:)
func (u_ URL) URLByAppendingPathExtension(pathExtension string) unsafe.Pointer {
	sel := objc.RegisterName("URLByAppendingPathExtension:")
	ret := u_.ID.Send(sel, pathExtension)
	return unsafe.Pointer(ret)
}
// Returns a URL by appending the path extension for a uniform type identifier. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/appendingPathExtension(for:)
func (u_ URL) URLByAppendingPathExtensionForType(contentType unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("URLByAppendingPathExtensionForType:")
	ret := u_.ID.Send(sel, contentType)
	return unsafe.Pointer(ret)
}
// Returns a bookmark for the URL, created with specified options and resource values. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/bookmarkData(options:includingResourceValuesForKeys:relativeTo:)
func (u_ URL) BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options unsafe.Pointer, keys unsafe.Pointer, relativeURL unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:")
	ret := u_.ID.Send(sel, options, keys, relativeURL, error)
	return unsafe.Pointer(ret)
}
// Returns whether the promised item can be reached. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/checkPromisedItemIsReachableAndReturnError(_:)
func (u_ URL) CheckPromisedItemIsReachableAndReturnError(error unsafe.Pointer) bool {
	sel := objc.RegisterName("checkPromisedItemIsReachableAndReturnError:")
	ret := u_.ID.Send(sel, error)
	return ret != 0
}
// Returns whether the resource pointed to by a file URL can be reached. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/checkResourceIsReachableAndReturnError(_:)
func (u_ URL) CheckResourceIsReachableAndReturnError(error unsafe.Pointer) bool {
	sel := objc.RegisterName("checkResourceIsReachableAndReturnError:")
	ret := u_.ID.Send(sel, error)
	return ret != 0
}
// Returns a new file reference URL that points to the same resource as the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/fileReferenceURL()
func (u_ URL) FileReferenceURL() unsafe.Pointer {
	sel := objc.RegisterName("fileReferenceURL")
	ret := u_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Fills the provided buffer with a C string representing a local file system path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/getFileSystemRepresentation(_:maxLength:)
func (u_ URL) GetFileSystemRepresentationMaxLength(buffer unsafe.Pointer, maxBufferLength uint) bool {
	sel := objc.RegisterName("getFileSystemRepresentation:maxLength:")
	ret := u_.ID.Send(sel, buffer, maxBufferLength)
	return ret != 0
}
// Returns the value of the resource property for the specified key. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/getPromisedItemResourceValue(_:forKey:)
func (u_ URL) GetPromisedItemResourceValueForKeyError(value unsafe.Pointer, key unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("getPromisedItemResourceValue:forKey:error:")
	ret := u_.ID.Send(sel, value, key, error)
	return ret != 0
}
// Returns the value of the resource property for the specified key. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/getResourceValue(_:forKey:)
func (u_ URL) GetResourceValueForKeyError(value unsafe.Pointer, key unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("getResourceValue:forKey:error:")
	ret := u_.ID.Send(sel, value, key, error)
	return ret != 0
}
// Returns whether the URL is a file reference URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/isFileReferenceURL()
func (u_ URL) IsFileReferenceURL() bool {
	sel := objc.RegisterName("isFileReferenceURL")
	ret := u_.ID.Send(sel)
	return ret != 0
}
// Loads the receiver’s resource data in the background. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/loadResourceDataNotifyingClient:usingCache:
func (u_ URL) LoadResourceDataNotifyingClientUsingCache(client objc.ID, shouldUseCache bool) {
	sel := objc.RegisterName("loadResourceDataNotifyingClient:usingCache:")
	u_.ID.Send(sel, client, shouldUseCache)
}
// Returns the resource values for the properties identified by specified array of keys. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/promisedItemResourceValues(forKeys:)
func (u_ URL) PromisedItemResourceValuesForKeysError(keys unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("promisedItemResourceValuesForKeys:error:")
	ret := u_.ID.Send(sel, keys, error)
	return unsafe.Pointer(ret)
}
// Returns the specified property of the receiver’s resource. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/propertyForKey:
func (u_ URL) PropertyForKey(propertyKey string) objc.ID {
	sel := objc.RegisterName("propertyForKey:")
	ret := u_.ID.Send(sel, propertyKey)
	return ret
}
// Removes all cached resource values and temporary resource values from the URL object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/removeAllCachedResourceValues()
func (u_ URL) RemoveAllCachedResourceValues() {
	sel := objc.RegisterName("removeAllCachedResourceValues")
	u_.ID.Send(sel)
}
// Removes the cached resource value identified by a given key from the URL object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/removeCachedResourceValue(forKey:)
func (u_ URL) RemoveCachedResourceValueForKey(key unsafe.Pointer) {
	sel := objc.RegisterName("removeCachedResourceValueForKey:")
	u_.ID.Send(sel, key)
}
// Returns the receiver’s resource data, loading it if necessary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/resourceDataUsingCache:
func (u_ URL) ResourceDataUsingCache(shouldUseCache bool) unsafe.Pointer {
	sel := objc.RegisterName("resourceDataUsingCache:")
	ret := u_.ID.Send(sel, shouldUseCache)
	return unsafe.Pointer(ret)
}
// Returns the resource values for the properties identified by specified array of keys. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/resourceValues(forKeys:)
func (u_ URL) ResourceValuesForKeysError(keys unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("resourceValuesForKeys:error:")
	ret := u_.ID.Send(sel, keys, error)
	return unsafe.Pointer(ret)
}
// Changes the specified property of the receiver’s resource. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/setProperty:forKey:
func (u_ URL) SetPropertyForKey(property objc.ID, propertyKey string) bool {
	sel := objc.RegisterName("setProperty:forKey:")
	ret := u_.ID.Send(sel, property, propertyKey)
	return ret != 0
}
// Attempts to set the resource data for the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/setResourceData:
func (u_ URL) SetResourceData(data unsafe.Pointer) bool {
	sel := objc.RegisterName("setResourceData:")
	ret := u_.ID.Send(sel, data)
	return ret != 0
}
// Sets the URL’s resource property for a given key to a given value. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/setResourceValue(_:forKey:)
func (u_ URL) SetResourceValueForKeyError(value objc.ID, key unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("setResourceValue:forKey:error:")
	ret := u_.ID.Send(sel, value, key, error)
	return ret != 0
}
// Sets the URL’s resource properties for a given set of keys to a given set of values. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/setResourceValues(_:)
func (u_ URL) SetResourceValuesError(keyedValues unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("setResourceValues:error:")
	ret := u_.ID.Send(sel, keyedValues, error)
	return ret != 0
}
// Sets a temporary resource value on the URL object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/setTemporaryResourceValue(_:forKey:)
func (u_ URL) SetTemporaryResourceValueForKey(value objc.ID, key unsafe.Pointer) {
	sel := objc.RegisterName("setTemporaryResourceValue:forKey:")
	u_.ID.Send(sel, value, key)
}
// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/startAccessingSecurityScopedResource()
func (u_ URL) StartAccessingSecurityScopedResource() bool {
	sel := objc.RegisterName("startAccessingSecurityScopedResource")
	ret := u_.ID.Send(sel)
	return ret != 0
}
// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/stopAccessingSecurityScopedResource()
func (u_ URL) StopAccessingSecurityScopedResource() {
	sel := objc.RegisterName("stopAccessingSecurityScopedResource")
	u_.ID.Send(sel)
}
// Writes the URL to the specified pasteboard. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSURL/write(to:)
func (u_ URL) WriteToPasteboard(pasteBoard unsafe.Pointer) {
	sel := objc.RegisterName("writeToPasteboard:")
	u_.ID.Send(sel, pasteBoard)
}

