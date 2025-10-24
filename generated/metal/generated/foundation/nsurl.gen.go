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
	BaseURL() IURL
	RelativeString() IString
	AbsoluteString() IString
	SetAbsoluteString(value IString)
	AbsoluteURL() IURL
	SetAbsoluteURL(value IURL)
	CustomPlaygroundQuickLook() unsafe.Pointer
	SetCustomPlaygroundQuickLook(value unsafe.Pointer)
	DataRepresentation() IData
	SetDataRepresentation(value IData)
	DeletingLastPathComponent() IURL
	SetDeletingLastPathComponent(value IURL)
	DeletingPathExtension() IURL
	SetDeletingPathExtension(value IURL)
	FilePathURL() IURL
	SetFilePathURL(value IURL)
	FileSystemRepresentation() unsafe.Pointer
	SetFileSystemRepresentation(value unsafe.Pointer)
	Fragment() IString
	SetFragment(value IString)
	HasDirectoryPath() bool /* primitive/slice/pointer. */
	SetHasDirectoryPath(value bool /* primitive/slice/pointer. */)
	Host() IString
	SetHost(value IString)
	IsFileURL() bool /* primitive/slice/pointer. */
	SetIsFileURL(value bool /* primitive/slice/pointer. */)
	LastPathComponent() IString
	SetLastPathComponent(value IString)
	ParameterString() IString
	SetParameterString(value IString)
	Password() IString
	SetPassword(value IString)
	Path() IString
	SetPath(value IString)
	PathComponents() IString
	SetPathComponents(value IString)
	PathExtension() IString
	SetPathExtension(value IString)
	Port() INumber
	SetPort(value INumber)
	Query() IString
	SetQuery(value IString)
	RelativePath() IString
	SetRelativePath(value IString)
	ResolvingSymlinksInPath() IURL
	SetResolvingSymlinksInPath(value IURL)
	ResourceSpecifier() IString
	SetResourceSpecifier(value IString)
	Scheme() IString
	SetScheme(value IString)
	Standardized() IURL
	SetStandardized(value IURL)
	StandardizingPath() IURL
	SetStandardizingPath(value IURL)
	User() IString
	SetUser(value IString)
	// methods:
	URLByAppendingPathComponent(pathComponent IString) IURL
	URLByAppendingPathComponentIsDirectory(pathComponent IString, isDirectory bool /* primitive/slice/pointer. */) IURL
	BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions /* not a class type */, keys []string /* primitive/slice/pointer. */, relativeURL IURL, error_ IError) IData
	GetResourceValueForKeyError(value unsafe.Pointer, key objc.IObject /* cross-framework: URLResourceKey */, error_ IError) bool /* primitive/slice/pointer. */
	RemoveAllCachedResourceValues()
	RemoveCachedResourceValueForKey(key objc.IObject /* cross-framework: URLResourceKey */)
	ResourceValuesForKeysError(keys []string /* primitive/slice/pointer. */, error_ IError) IDictionary /* already interface */
	SetResourceValueForKeyError(value objectivec.IObject, key objc.IObject /* cross-framework: URLResourceKey */, error_ IError) bool /* primitive/slice/pointer. */
	SetResourceValuesError(keyedValues IDictionary /* already interface */, error_ IError) bool /* primitive/slice/pointer. */
	SetTemporaryResourceValueForKey(value objectivec.IObject, key objc.IObject /* cross-framework: URLResourceKey */)
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



// Initializes a newly created NSURL referencing the local file or directory at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:)
func NewURLFileURLWithPath(path IString) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:"), path)
	rv.Autorelease()
	return rv
}


// Initializes a newly created NSURL referencing the local file or directory at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:)
func NewURLFileURLWithPathIsDirectory(path IString, isDir bool /* primitive/slice/pointer. */) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:"), path, isDir)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:isDirectory:relativeTo:)
func NewURLFileURLWithPathIsDirectoryRelativeToURL(path IString, isDir bool /* primitive/slice/pointer. */, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:isDirectory:relativeToURL:"), path, isDir, baseURL)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/init(fileURLWithPath:relativeTo:)
func NewURLFileURLWithPathRelativeToURL(path IString, baseURL IURL) URL {
	instance := getURLClass().Alloc()
	rv := objc.Send[URL](instance.ID, objc.Sel("initFileURLWithPath:relativeToURL:"), path, baseURL)
	rv.Autorelease()
	return rv
}



// Returns a new URL made by resolving bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:
func (uc _URLClass) URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(bookmarkData IData, options URLBookmarkResolutionOptions /* not a class type */, relativeURL IURL, isStale bool /* primitive/slice/pointer. */, error_ IError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLByResolvingBookmarkData:options:relativeToURL:bookmarkDataIsStale:error:"), bookmarkData, options, relativeURL, isStale, error_)
	return rv
}


// Initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:)
func (uc _URLClass) FileURLWithPath(path IString) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:"), path)
	return rv
}


// Initializes and returns a newly created NSURL object as a file URL with a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:)
func (uc _URLClass) FileURLWithPathIsDirectory(path IString, isDir bool /* primitive/slice/pointer. */) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:isDirectory:"), path, isDir)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:isDirectory:relativeTo:)
func (uc _URLClass) FileURLWithPathIsDirectoryRelativeToURL(path IString, isDir bool /* primitive/slice/pointer. */, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:isDirectory:relativeToURL:"), path, isDir, baseURL)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/fileURL(withPath:relativeTo:)
func (uc _URLClass) FileURLWithPathRelativeToURL(path IString, baseURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(uc.class), objc.Sel("fileURLWithPath:relativeToURL:"), path, baseURL)
	return rv
}


// Returns a new URL by appending a path component to the original URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:)
func (u_ URL) URLByAppendingPathComponent(pathComponent IString) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathComponent:"), pathComponent)
	return rv
}


// Returns a new URL by appending a path component to the original URL, along with a trailing slash if the component is a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/appendingPathComponent(_:isDirectory:)
func (u_ URL) URLByAppendingPathComponentIsDirectory(pathComponent IString, isDirectory bool /* primitive/slice/pointer. */) IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLByAppendingPathComponent:isDirectory:"), pathComponent, isDirectory)
	return rv
}


// Returns a bookmark for the URL, created with specified options and resource values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/bookmarkData(options:includingResourceValuesForKeys:relativeTo:)
func (u_ URL) BookmarkDataWithOptionsIncludingResourceValuesForKeysRelativeToURLError(options URLBookmarkCreationOptions /* not a class type */, keys []string /* primitive/slice/pointer. */, relativeURL IURL, error_ IError) IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("bookmarkDataWithOptions:includingResourceValuesForKeys:relativeToURL:error:"), options, keys, relativeURL, error_)
	return rv
}


// Returns the value of the resource property for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/getResourceValue(_:forKey:)
func (u_ URL) GetResourceValueForKeyError(value unsafe.Pointer, key objc.IObject /* cross-framework: URLResourceKey */, error_ IError) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("getResourceValue:forKey:error:"), value, key, error_)
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
func (u_ URL) RemoveCachedResourceValueForKey(key objc.IObject /* cross-framework: URLResourceKey */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeCachedResourceValueForKey:"), key)
}


// Returns the resource values for the properties identified by specified array of keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/resourceValues(forKeys:)
func (u_ URL) ResourceValuesForKeysError(keys []string /* primitive/slice/pointer. */, error_ IError) IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](u_.ID, objc.Sel("resourceValuesForKeys:error:"), keys, error_)
	return rv
}


// Sets the URL’s resource property for a given key to a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValue(_:forKey:)
func (u_ URL) SetResourceValueForKeyError(value objectivec.IObject, key objc.IObject /* cross-framework: URLResourceKey */, error_ IError) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceValue:forKey:error:"), value, key, error_)
	return rv
}


// Sets the URL’s resource properties for a given set of keys to a given set of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setResourceValues(_:)
func (u_ URL) SetResourceValuesError(keyedValues IDictionary /* already interface */, error_ IError) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("setResourceValues:error:"), keyedValues, error_)
	return rv
}


// Sets a temporary resource value on the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/setTemporaryResourceValue(_:forKey:)
func (u_ URL) SetTemporaryResourceValueForKey(value objectivec.IObject, key objc.IObject /* cross-framework: URLResourceKey */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTemporaryResourceValue:forKey:"), value, key)
}


// The base URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/baseURL
func (u_ URL) BaseURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("baseURL"))
	return rv
}


// A string representation of the relative portion of the URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/relativeString
func (u_ URL) RelativeString() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("relativeString"))
	return rv
}


// The URL string for the receiver as an absolute URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/absolutestring
func (u_ URL) AbsoluteString() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("absoluteString"))
	return rv
}


// The URL string for the receiver as an absolute URL. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/absolutestring
func (u_ URL) SetAbsoluteString(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAbsoluteString:"), value)
}


// An absolute URL that refers to the same resource as the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/absoluteurl
func (u_ URL) AbsoluteURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("absoluteURL"))
	return rv
}


// An absolute URL that refers to the same resource as the receiver. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/absoluteurl
func (u_ URL) SetAbsoluteURL(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAbsoluteURL:"), value)
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


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/datarepresentation
func (u_ URL) DataRepresentation() IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("dataRepresentation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/datarepresentation
func (u_ URL) SetDataRepresentation(value IData) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDataRepresentation:"), value)
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


// A file path URL that points to the same resource as the URL object. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/filepathurl
func (u_ URL) FilePathURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("filePathURL"))
	return rv
}


// A file path URL that points to the same resource as the URL object. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/filepathurl
func (u_ URL) SetFilePathURL(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setFilePathURL:"), value)
}


// A C string containing the URL’s file system path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/filesystemrepresentation
func (u_ URL) FileSystemRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("fileSystemRepresentation"))
	return rv
}


// A C string containing the URL’s file system path. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/filesystemrepresentation
func (u_ URL) SetFileSystemRepresentation(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setFileSystemRepresentation:"), value)
}


// The fragment identifier, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/fragment
func (u_ URL) Fragment() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("fragment"))
	return rv
}


// The fragment identifier, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/fragment
func (u_ URL) SetFragment(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setFragment:"), value)
}


// A Boolean value that indicates whether the URL string’s path represents a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/hasdirectorypath
func (u_ URL) HasDirectoryPath() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasDirectoryPath"))
	return rv
}


// A Boolean value that indicates whether the URL string’s path represents a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/hasdirectorypath
func (u_ URL) SetHasDirectoryPath(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHasDirectoryPath:"), value)
}


// The host, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/host
func (u_ URL) Host() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("host"))
	return rv
}


// The host, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/host
func (u_ URL) SetHost(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHost:"), value)
}


// A boolean value that determines whether the receiver is a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/isfileurl
func (u_ URL) IsFileURL() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("isFileURL"))
	return rv
}


// A boolean value that determines whether the receiver is a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/isfileurl
func (u_ URL) SetIsFileURL(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsFileURL:"), value)
}


// The last path component. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/lastpathcomponent
func (u_ URL) LastPathComponent() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("lastPathComponent"))
	return rv
}


// The last path component. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/lastpathcomponent
func (u_ URL) SetLastPathComponent(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setLastPathComponent:"), value)
}


// The parameter string conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/parameterstring
func (u_ URL) ParameterString() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("parameterString"))
	return rv
}


// The parameter string conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/parameterstring
func (u_ URL) SetParameterString(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setParameterString:"), value)
}


// The password conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/password
func (u_ URL) Password() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("password"))
	return rv
}


// The password conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/password
func (u_ URL) SetPassword(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPassword:"), value)
}


// The path, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/path
func (u_ URL) Path() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("path"))
	return rv
}


// The path, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/path
func (u_ URL) SetPath(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPath:"), value)
}


// An array containing the path components. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/pathcomponents
func (u_ URL) PathComponents() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("pathComponents"))
	return rv
}


// An array containing the path components. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/pathcomponents
func (u_ URL) SetPathComponents(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPathComponents:"), value)
}


// The path extension. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/pathextension
func (u_ URL) PathExtension() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("pathExtension"))
	return rv
}


// The path extension. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/pathextension
func (u_ URL) SetPathExtension(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPathExtension:"), value)
}


// The port, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/port
func (u_ URL) Port() INumber {
	rv := objc.Send[Number](u_.ID, objc.Sel("port"))
	return rv
}


// The port, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/port
func (u_ URL) SetPort(value INumber) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPort:"), value)
}


// The query string, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/query
func (u_ URL) Query() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("query"))
	return rv
}


// The query string, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/query
func (u_ URL) SetQuery(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setQuery:"), value)
}


// The relative path, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/relativepath
func (u_ URL) RelativePath() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("relativePath"))
	return rv
}


// The relative path, conforming to RFC 1808. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/relativepath
func (u_ URL) SetRelativePath(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRelativePath:"), value)
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


// The resource specifier. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/resourcespecifier
func (u_ URL) ResourceSpecifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("resourceSpecifier"))
	return rv
}


// The resource specifier. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/resourcespecifier
func (u_ URL) SetResourceSpecifier(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResourceSpecifier:"), value)
}


// The scheme. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/scheme
func (u_ URL) Scheme() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("scheme"))
	return rv
}


// The scheme. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/scheme
func (u_ URL) SetScheme(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setScheme:"), value)
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


// The user name, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/user
func (u_ URL) User() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("user"))
	return rv
}


// The user name, conforming to RFC 1808.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurl/user
func (u_ URL) SetUser(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUser:"), value)
}


