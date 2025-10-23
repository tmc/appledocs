// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileManager] class.
var (
	FileManagerClass     _FileManagerClass
	FileManagerClassOnce sync.Once
)

func getFileManagerClass() _FileManagerClass {
	FileManagerClassOnce.Do(func() {
		FileManagerClass = _FileManagerClass{objc.GetClass("NSFileManager")}
	})
	return FileManagerClass
}

type _FileManagerClass struct {
	class objc.Class
}

// An interface definition for the [FileManager] class.
type IFileManager interface {
	objectivec.IObject
	// properties:
	CurrentDirectoryPath() string /* primitive/slice/pointer. */
	HomeDirectoryForCurrentUser() IURL
	UbiquityIdentityToken() objc.ID
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	TemporaryDirectory() IURL
	SetTemporaryDirectory(value IURL)
	NSFileManagerUnmountDissentingProcessIdentifierErrorKey() string /* primitive/slice/pointer. */
	NSFoundationVersionWithFileManagerResourceForkSupport() unsafe.Pointer
	SetNSFoundationVersionWithFileManagerResourceForkSupport(value unsafe.Pointer)
	// methods:
	EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url IURL, keys []string /* primitive/slice/pointer. */, mask DirectoryEnumerationOptions, handler unsafe.Pointer) unsafe.Pointer
	ContainerURLForSecurityApplicationGroupIdentifier(groupIdentifier string /* primitive/slice/pointer. */) IURL
	ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url IURL, keys []string /* primitive/slice/pointer. */, mask DirectoryEnumerationOptions, error_ IError) []URL /* primitive/slice/pointer. */
	ContentsOfDirectoryAtPathError(path string /* primitive/slice/pointer. */, error_ IError) []string /* primitive/slice/pointer. */
	CreateSymbolicLinkAtURLWithDestinationURLError(url IURL, destURL IURL, error_ IError) bool /* primitive/slice/pointer. */
	EnumeratorAtPath(path string /* primitive/slice/pointer. */) unsafe.Pointer
	EvictUbiquitousItemAtURLError(url IURL, error_ IError) bool /* primitive/slice/pointer. */
	FetchLatestRemoteVersionOfItemAtURLCompletionHandler(url IURL, completionHandler unsafe.Pointer)
	FileExistsAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	FileExistsAtPathIsDirectory(path string /* primitive/slice/pointer. */, isDirectory unsafe.Pointer) bool /* primitive/slice/pointer. */
	FileSystemRepresentationWithPath(path string /* primitive/slice/pointer. */) unsafe.Pointer
	HomeDirectoryForUser(userName string /* primitive/slice/pointer. */) IURL
	IsDeletableFileAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	IsExecutableFileAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	IsReadableFileAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	IsUbiquitousItemAtURL(url IURL) bool /* primitive/slice/pointer. */
	IsWritableFileAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys []string /* primitive/slice/pointer. */, options VolumeEnumerationOptions) []URL /* primitive/slice/pointer. */
	PauseSyncForUbiquitousItemAtURLCompletionHandler(url IURL, completionHandler unsafe.Pointer)
	RemoveItemAtURLError(URL IURL, error_ IError) bool /* primitive/slice/pointer. */
	RemoveItemAtPathError(path string /* primitive/slice/pointer. */, error_ IError) bool /* primitive/slice/pointer. */
	ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL IURL, newItemURL IURL, backupItemName string /* primitive/slice/pointer. */, options FileManagerItemReplacementOptions, resultingURL IURL, error_ IError) bool /* primitive/slice/pointer. */
	ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler(url IURL, behavior FileManagerResumeSyncBehavior, completionHandler unsafe.Pointer)
	SetAttributesOfItemAtPathError(attributes IDictionary /* already interface */, path string /* primitive/slice/pointer. */, error_ IError) bool /* primitive/slice/pointer. */
	SetUbiquitousItemAtURLDestinationURLError(flag bool /* primitive/slice/pointer. */, url IURL, destinationURL IURL, error_ IError) bool /* primitive/slice/pointer. */
	StartDownloadingUbiquitousItemAtURLError(url IURL, error_ IError) bool /* primitive/slice/pointer. */
	StringWithFileSystemRepresentationLength(str unsafe.Pointer, len_ uint /* primitive/slice/pointer. */) IString
	SubpathsAtPath(path string /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */
	SubpathsOfDirectoryAtPathError(path string /* primitive/slice/pointer. */, error_ IError) []string /* primitive/slice/pointer. */
	UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler(url IURL, conflictResolutionPolicy FileManagerUploadLocalVersionConflictPolicy, completionHandler unsafe.Pointer)
	URLForDirectoryInDomainAppropriateForURLCreateError(directory SearchPathDirectory, domain SearchPathDomainMask, url IURL, shouldCreate bool /* primitive/slice/pointer. */, error_ IError) IURL
	URLForPublishingUbiquitousItemAtURLExpirationDateError(url IURL, outDate IDate, error_ IError) IURL
	URLForUbiquityContainerIdentifier(containerIdentifier string /* primitive/slice/pointer. */) IURL
	URLsForDirectoryInDomains(directory SearchPathDirectory, domainMask SearchPathDomainMask) []URL /* primitive/slice/pointer. */
}

// A convenient interface to the contents of the file system, and the primary means of interacting with it.
//
// A file manager object lets you examine the contents of the file system and make changes to it. The class provides convenient access to a shared file manager object that is suitable for most types of file-related manipulations. A file manager object is typically your primary mode of interaction with the file system. You use it to locate, create, copy, and move files and directories. You also use it to get information about a file or directory or change some of its attributes. When specifying the location of files, you can use either or objects. The use of the class is generally preferred for specifying file-system items because URLs can convert path information to a more efficient representation internally. You can also obtain a bookmark from an object, which is similar to an alias and offers a more sure way of locating the file or directory later. If you are moving, copying, linking, or removing files or directories, you can use a delegate in conjunction with a file manager object to manage those operations. The delegate’s role is to affirm the operation and to decide whether to proceed when errors occur. In macOS 10.7 and later, the delegate must conform to the protocol. In iOS 5.0 and later and in macOS 10.7 and later, includes methods for managing items stored in iCloud. Files and directories tagged for cloud storage are synced to iCloud so that they can be made available to the user’s iOS devices and Macintosh computers. Changes to an item in one location are propagated to all other locations to ensure the items stay in sync.


// A convenient interface to the contents of the file system, and the primary means of interacting with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager
type FileManager struct {
	objectivec.Object
}

// FileManagerFrom constructs a [FileManager] from an unsafe.Pointer.
//
// A convenient interface to the contents of the file system, and the primary means of interacting with it.
func FileManagerFrom(ptr unsafe.Pointer) FileManager {
	return FileManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileManagerClass) Alloc() FileManager {
	rv := objc.Send[FileManager](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileManagerClass) New() FileManager {
	rv := objc.Send[FileManager](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileManager) Init() FileManager {
	rv := objc.Send[FileManager](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileManager) Autorelease() FileManager {
	rv := objc.Send[FileManager](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileManager creates a new FileManager instance.
func NewFileManager() FileManager {
	return getFileManagerClass().New()
}



// Initializes a file manager object that is authorized to perform privileged file system operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/init(authorization:)
func NewFileManagerWithAuthorization(authorization objc.IObject /* cross-framework WorkspaceAuthorization */) FileManager {
	rv := objc.Send[FileManager](objc.ID(getFileManagerClass().class), objc.Sel("fileManagerWithAuthorization:"), authorization)
	return rv
}



// Initializes a file manager object that is authorized to perform privileged file system operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/init(authorization:)
func (fc _FileManagerClass) FileManagerWithAuthorization(authorization objc.IObject /* cross-framework WorkspaceAuthorization */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fileManagerWithAuthorization:"), authorization)
	return rv
}


// The shared file manager object for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/default
func (fc _FileManagerClass) DefaultManager() FileManager {
	rv := objc.Send[FileManager](objc.ID(fc.class), objc.Sel("defaultManager"))
	return rv
}

// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:
func (f_ FileManager) EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url IURL, keys []string /* primitive/slice/pointer. */, mask DirectoryEnumerationOptions, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:"), url, keys, mask, handler)
	return rv
}


// Returns the container directory associated with the specified security application group identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/containerURL(forSecurityApplicationGroupIdentifier:)
func (f_ FileManager) ContainerURLForSecurityApplicationGroupIdentifier(groupIdentifier string /* primitive/slice/pointer. */) IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("containerURLForSecurityApplicationGroupIdentifier:"), objc.String(groupIdentifier))
	return rv
}


// Performs a shallow search of the specified directory and returns URLs for the contained items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contentsOfDirectory(at:includingPropertiesForKeys:options:)
func (f_ FileManager) ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url IURL, keys []string /* primitive/slice/pointer. */, mask DirectoryEnumerationOptions, error_ IError) []URL /* primitive/slice/pointer. */ {
	rv := objc.Send[[]URL](f_.ID, objc.Sel("contentsOfDirectoryAtURL:includingPropertiesForKeys:options:error:"), url, keys, mask, error_)
	return rv
}


// Performs a shallow search of the specified directory and returns the paths of any contained items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contentsOfDirectory(atPath:)
func (f_ FileManager) ContentsOfDirectoryAtPathError(path string /* primitive/slice/pointer. */, error_ IError) []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](f_.ID, objc.Sel("contentsOfDirectoryAtPath:error:"), objc.String(path), error_)
	return rv
}


// Creates a symbolic link at the specified URL that points to an item at the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createSymbolicLink(at:withDestinationURL:)
func (f_ FileManager) CreateSymbolicLinkAtURLWithDestinationURLError(url IURL, destURL IURL, error_ IError) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("createSymbolicLinkAtURL:withDestinationURL:error:"), url, destURL, error_)
	return rv
}


// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/enumerator(atPath:)
func (f_ FileManager) EnumeratorAtPath(path string /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("enumeratorAtPath:"), objc.String(path))
	return rv
}


// Removes the local copy of the specified item that’s stored in iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/evictUbiquitousItem(at:)
func (f_ FileManager) EvictUbiquitousItemAtURLError(url IURL, error_ IError) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("evictUbiquitousItemAtURL:error:"), url, error_)
	return rv
}


// Asynchronously fetches the latest remote version of a given item from the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fetchLatestRemoteVersionOfItem(at:completionHandler:)
func (f_ FileManager) FetchLatestRemoteVersionOfItemAtURLCompletionHandler(url IURL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("fetchLatestRemoteVersionOfItemAtURL:completionHandler:"), url, completionHandler)
}


// Returns a Boolean value that indicates whether a file or directory exists at a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileExists(atPath:)
func (f_ FileManager) FileExistsAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("fileExistsAtPath:"), objc.String(path))
	return rv
}


// Returns a Boolean value that indicates whether a file or directory exists at a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileExists(atPath:isDirectory:)
func (f_ FileManager) FileExistsAtPathIsDirectory(path string /* primitive/slice/pointer. */, isDirectory unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("fileExistsAtPath:isDirectory:"), objc.String(path), isDirectory)
	return rv
}


// Returns a C-string representation of a given path that properly encodes Unicode strings for use by the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileSystemRepresentation(withPath:)
func (f_ FileManager) FileSystemRepresentationWithPath(path string /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fileSystemRepresentationWithPath:"), objc.String(path))
	return rv
}


// Returns the home directory for the specified user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/homeDirectory(forUser:)
func (f_ FileManager) HomeDirectoryForUser(userName string /* primitive/slice/pointer. */) IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("homeDirectoryForUser:"), objc.String(userName))
	return rv
}


// Returns a Boolean value that indicates whether the invoking object appears able to delete a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isDeletableFile(atPath:)
func (f_ FileManager) IsDeletableFileAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDeletableFileAtPath:"), objc.String(path))
	return rv
}


// Returns a Boolean value that indicates whether the operating system appears able to execute a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isExecutableFile(atPath:)
func (f_ FileManager) IsExecutableFileAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isExecutableFileAtPath:"), objc.String(path))
	return rv
}


// Returns a Boolean value that indicates whether the invoking object appears able to read a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isReadableFile(atPath:)
func (f_ FileManager) IsReadableFileAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isReadableFileAtPath:"), objc.String(path))
	return rv
}


// Returns a Boolean indicating whether the item is targeted for storage in iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isUbiquitousItem(at:)
func (f_ FileManager) IsUbiquitousItemAtURL(url IURL) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isUbiquitousItemAtURL:"), url)
	return rv
}


// Returns a Boolean value that indicates whether the invoking object appears able to write to a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isWritableFile(atPath:)
func (f_ FileManager) IsWritableFileAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isWritableFileAtPath:"), objc.String(path))
	return rv
}


// Returns an array of URLs that identify the mounted volumes available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/mountedVolumeURLs(includingResourceValuesForKeys:options:)
func (f_ FileManager) MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys []string /* primitive/slice/pointer. */, options VolumeEnumerationOptions) []URL /* primitive/slice/pointer. */ {
	rv := objc.Send[[]URL](f_.ID, objc.Sel("mountedVolumeURLsIncludingResourceValuesForKeys:options:"), propertyKeys, options)
	return rv
}


// Asynchronously pauses sync of an item at the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/pauseSyncForUbiquitousItem(at:completionHandler:)
func (f_ FileManager) PauseSyncForUbiquitousItemAtURLCompletionHandler(url IURL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("pauseSyncForUbiquitousItemAtURL:completionHandler:"), url, completionHandler)
}


// Removes the file or directory at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/removeItem(at:)
func (f_ FileManager) RemoveItemAtURLError(URL IURL, error_ IError) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("removeItemAtURL:error:"), URL, error_)
	return rv
}


// Removes the file or directory at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/removeItem(atPath:)
func (f_ FileManager) RemoveItemAtPathError(path string /* primitive/slice/pointer. */, error_ IError) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("removeItemAtPath:error:"), objc.String(path), error_)
	return rv
}


// Replaces the contents of the item at the specified URL in a manner that ensures no data loss occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/replaceItem(at:withItemAt:backupItemName:options:resultingItemURL:)
func (f_ FileManager) ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL IURL, newItemURL IURL, backupItemName string /* primitive/slice/pointer. */, options FileManagerItemReplacementOptions, resultingURL IURL, error_ IError) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("replaceItemAtURL:withItemAtURL:backupItemName:options:resultingItemURL:error:"), originalItemURL, newItemURL, objc.String(backupItemName), options, resultingURL, error_)
	return rv
}


// Asynchronously resumes the sync on a paused item using the given resume behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/resumeSyncForUbiquitousItem(at:with:completionHandler:)
func (f_ FileManager) ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler(url IURL, behavior FileManagerResumeSyncBehavior, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("resumeSyncForUbiquitousItemAtURL:withBehavior:completionHandler:"), url, behavior, completionHandler)
}


// Sets the attributes of the specified file or directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/setAttributes(_:ofItemAtPath:)
func (f_ FileManager) SetAttributesOfItemAtPathError(attributes IDictionary /* already interface */, path string /* primitive/slice/pointer. */, error_ IError) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("setAttributes:ofItemAtPath:error:"), attributes, objc.String(path), error_)
	return rv
}


// Indicates whether the item at the specified URL should be stored in iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/setUbiquitous(_:itemAt:destinationURL:)
func (f_ FileManager) SetUbiquitousItemAtURLDestinationURLError(flag bool /* primitive/slice/pointer. */, url IURL, destinationURL IURL, error_ IError) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("setUbiquitous:itemAtURL:destinationURL:error:"), flag, url, destinationURL, error_)
	return rv
}


// Starts downloading (if necessary) the specified item to the local system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/startDownloadingUbiquitousItem(at:)
func (f_ FileManager) StartDownloadingUbiquitousItemAtURLError(url IURL, error_ IError) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("startDownloadingUbiquitousItemAtURL:error:"), url, error_)
	return rv
}


// Returns an object whose contents are derived from the specified C-string path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/string(withFileSystemRepresentation:length:)
func (f_ FileManager) StringWithFileSystemRepresentationLength(str unsafe.Pointer, len_ uint /* primitive/slice/pointer. */) IString {
	rv := objc.Send[String](f_.ID, objc.Sel("stringWithFileSystemRepresentation:length:"), str, len_)
	return rv
}


// Returns an array of strings identifying the paths for all items in the specified directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/subpaths(atPath:)
func (f_ FileManager) SubpathsAtPath(path string /* primitive/slice/pointer. */) []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](f_.ID, objc.Sel("subpathsAtPath:"), objc.String(path))
	return rv
}


// Performs a deep enumeration of the specified directory and returns the paths of all of the contained subdirectories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/subpathsOfDirectory(atPath:)
func (f_ FileManager) SubpathsOfDirectoryAtPathError(path string /* primitive/slice/pointer. */, error_ IError) []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](f_.ID, objc.Sel("subpathsOfDirectoryAtPath:error:"), objc.String(path), error_)
	return rv
}


// Asynchronously uploads the local version of the item using the provided conflict resolution policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/uploadLocalVersionOfUbiquitousItem(at:withConflictResolutionPolicy:completionHandler:)
func (f_ FileManager) UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler(url IURL, conflictResolutionPolicy FileManagerUploadLocalVersionConflictPolicy, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("uploadLocalVersionOfUbiquitousItemAtURL:withConflictResolutionPolicy:completionHandler:"), url, conflictResolutionPolicy, completionHandler)
}


// Locates and optionally creates the specified common directory in a domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/url(for:in:appropriateFor:create:)
func (f_ FileManager) URLForDirectoryInDomainAppropriateForURLCreateError(directory SearchPathDirectory, domain SearchPathDomainMask, url IURL, shouldCreate bool /* primitive/slice/pointer. */, error_ IError) IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("URLForDirectory:inDomain:appropriateForURL:create:error:"), directory, domain, url, shouldCreate, error_)
	return rv
}


// Returns a URL that can be emailed to users to allow them to download a copy of a flat file item from iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/url(forPublishingUbiquitousItemAt:expiration:)
func (f_ FileManager) URLForPublishingUbiquitousItemAtURLExpirationDateError(url IURL, outDate IDate, error_ IError) IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("URLForPublishingUbiquitousItemAtURL:expirationDate:error:"), url, outDate, error_)
	return rv
}


// Returns the URL for the iCloud container associated with the specified identifier and establishes access to that container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/url(forUbiquityContainerIdentifier:)
func (f_ FileManager) URLForUbiquityContainerIdentifier(containerIdentifier string /* primitive/slice/pointer. */) IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("URLForUbiquityContainerIdentifier:"), objc.String(containerIdentifier))
	return rv
}


// Returns an array of URLs for the specified common directory in the requested domains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/urls(for:in:)
func (f_ FileManager) URLsForDirectoryInDomains(directory SearchPathDirectory, domainMask SearchPathDomainMask) []URL /* primitive/slice/pointer. */ {
	rv := objc.Send[[]URL](f_.ID, objc.Sel("URLsForDirectory:inDomains:"), directory, domainMask)
	return rv
}


// The path to the program’s current directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/currentDirectoryPath
func (f_ FileManager) CurrentDirectoryPath() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](f_.ID, objc.Sel("currentDirectoryPath"))
	return rv
}


// The shared file manager object for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/default
func (f_ FileManager) DefaultManager() IFileManager {
	rv := objc.Send[FileManager](f_.ID, objc.Sel("defaultManager"))
	return rv
}


// The home directory for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/homeDirectoryForCurrentUser
func (f_ FileManager) HomeDirectoryForCurrentUser() IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("homeDirectoryForCurrentUser"))
	return rv
}


// An opaque token that represents the current user’s iCloud Drive Documents identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ubiquityIdentityToken
func (f_ FileManager) UbiquityIdentityToken() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("ubiquityIdentityToken"))
	return rv
}


// The delegate of the file manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/delegate
func (f_ FileManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate of the file manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/delegate
func (f_ FileManager) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelegate:"), value)
}


// The temporary directory for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/temporarydirectory
func (f_ FileManager) TemporaryDirectory() IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("temporaryDirectory"))
	return rv
}


// The temporary directory for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/temporarydirectory
func (f_ FileManager) SetTemporaryDirectory(value IURL) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTemporaryDirectory:"), value)
}


// The process identifier of the process that prevented a volume from unmounting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerunmountdissentingprocessidentifiererrorkey
func (f_ FileManager) NSFileManagerUnmountDissentingProcessIdentifierErrorKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](f_.ID, objc.Sel("NSFileManagerUnmountDissentingProcessIdentifierErrorKey"))
	return rv
}


// The version of the Foundation framework in which
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfoundationversionwithfilemanagerresourceforksupport
func (f_ FileManager) NSFoundationVersionWithFileManagerResourceForkSupport() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("NSFoundationVersionWithFileManagerResourceForkSupport"))
	return rv
}


// The version of the Foundation framework in which
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfoundationversionwithfilemanagerresourceforksupport
func (f_ FileManager) SetNSFoundationVersionWithFileManagerResourceForkSupport(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNSFoundationVersionWithFileManagerResourceForkSupport:"), value)
}


