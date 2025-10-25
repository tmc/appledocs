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
	CurrentDirectoryPath() IString
	HomeDirectoryForCurrentUser() IURL
	TemporaryDirectory() IURL
	UbiquityIdentityToken() unsafe.Pointer
	NSFileManagerUnmountDissentingProcessIdentifierErrorKey() IString
	NSFoundationVersionWithFileManagerResourceForkSupport() objectivec.IObject
	SetNSFoundationVersionWithFileManagerResourceForkSupport(value objectivec.IObject)


	

	// methods:
	AttributesOfFileSystemForPathError(path IString, error_ IError) IDictionary
	AttributesOfItemAtPathError(path IString, error_ IError) IDictionary
	ChangeCurrentDirectoryPath(path IString) bool
	ComponentsToDisplayForPath(path IString) []string
	ContentsAtPath(path IString) IData
	ContentsEqualAtPathAndPath(path1 IString, path2 IString) bool
	ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url IURL, keys []string, mask DirectoryEnumerationOptions, error_ IError) []URL
	ContentsOfDirectoryAtPathError(path IString, error_ IError) []string
	CopyItemAtURLToURLError(srcURL IURL, dstURL IURL, error_ IError) bool
	CopyItemAtPathToPathError(srcPath IString, dstPath IString, error_ IError) bool
	CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError(url IURL, createIntermediates bool, attributes IDictionary, error_ IError) bool
	CreateDirectoryAtPathWithIntermediateDirectoriesAttributesError(path IString, createIntermediates bool, attributes IDictionary, error_ IError) bool
	CreateFileAtPathContentsAttributes(path IString, data IData, attr IDictionary) bool
	CreateSymbolicLinkAtURLWithDestinationURLError(url IURL, destURL IURL, error_ IError) bool
	CreateSymbolicLinkAtPathWithDestinationPathError(path IString, destPath IString, error_ IError) bool
	DestinationOfSymbolicLinkAtPathError(path IString, error_ IError) IString
	DisplayNameAtPath(path IString) IString
	EnumeratorAtPath(path IString) unsafe.Pointer
	EvictUbiquitousItemAtURLError(url IURL, error_ IError) bool
	FetchLatestRemoteVersionOfItemAtURLCompletionHandler(url IURL, completionHandler unsafe.Pointer)
	FileExistsAtPath(path IString) bool
	FileExistsAtPathIsDirectory(path IString, isDirectory objectivec.IObject) bool
	FileSystemRepresentationWithPath(path IString) objectivec.IObject
	GetFileProviderServicesForItemAtURLCompletionHandler(url IURL, completionHandler unsafe.Pointer)
	HomeDirectoryForUser(userName IString) IURL
	IsDeletableFileAtPath(path IString) bool
	IsExecutableFileAtPath(path IString) bool
	IsReadableFileAtPath(path IString) bool
	IsUbiquitousItemAtURL(url IURL) bool
	IsWritableFileAtPath(path IString) bool
	LinkItemAtURLToURLError(srcURL IURL, dstURL IURL, error_ IError) bool
	LinkItemAtPathToPathError(srcPath IString, dstPath IString, error_ IError) bool
	MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys []string, options VolumeEnumerationOptions) []URL
	MoveItemAtURLToURLError(srcURL IURL, dstURL IURL, error_ IError) bool
	MoveItemAtPathToPathError(srcPath IString, dstPath IString, error_ IError) bool
	PauseSyncForUbiquitousItemAtURLCompletionHandler(url IURL, completionHandler unsafe.Pointer)
	RemoveItemAtURLError(URL IURL, error_ IError) bool
	RemoveItemAtPathError(path IString, error_ IError) bool
	ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL IURL, newItemURL IURL, backupItemName IString, options FileManagerItemReplacementOptions, resultingURL IURL, error_ IError) bool
	ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler(url IURL, behavior FileManagerResumeSyncBehavior, completionHandler unsafe.Pointer)
	SetAttributesOfItemAtPathError(attributes IDictionary, path IString, error_ IError) bool
	SetUbiquitousItemAtURLDestinationURLError(flag bool, url IURL, destinationURL IURL, error_ IError) bool
	StartDownloadingUbiquitousItemAtURLError(url IURL, error_ IError) bool
	StringWithFileSystemRepresentationLength(str objectivec.IObject, len_ uint) IString
	SubpathsAtPath(path IString) []string
	SubpathsOfDirectoryAtPathError(path IString, error_ IError) []string
	TrashItemAtURLResultingItemURLError(url IURL, outResultingURL IURL, error_ IError) bool
	UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler(url IURL, conflictResolutionPolicy FileManagerUploadLocalVersionConflictPolicy, completionHandler unsafe.Pointer)
	URLForDirectoryInDomainAppropriateForURLCreateError(directory SearchPathDirectory, domain SearchPathDomainMask, url IURL, shouldCreate bool, error_ IError) IURL
	URLForPublishingUbiquitousItemAtURLExpirationDateError(url IURL, outDate IDate, error_ IError) IURL
	URLForUbiquityContainerIdentifier(containerIdentifier IString) IURL
	URLsForDirectoryInDomains(directory SearchPathDirectory, domainMask SearchPathDomainMask) []URL
	EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url IURL, keys []string, mask DirectoryEnumerationOptions, handler unsafe.Pointer) unsafe.Pointer


}





// Alloc allocates a new instance without initialization.
func (fc _FileManagerClass) Alloc() FileManager {
	rv := objc.Send[FileManager](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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




















// Returns a dictionary that describes the attributes of the mounted file system on which a given path resides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/attributesOfFileSystem(forPath:)
func (f_ FileManager) AttributesOfFileSystemForPathError(path IString, error_ IError) IDictionary {
	rv := objc.Send[Dictionary](f_.ID, objc.Sel("attributesOfFileSystemForPath:error:"), path, error_)
	return rv
}


// Returns the attributes of the item at a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/attributesOfItem(atPath:)
func (f_ FileManager) AttributesOfItemAtPathError(path IString, error_ IError) IDictionary {
	rv := objc.Send[Dictionary](f_.ID, objc.Sel("attributesOfItemAtPath:error:"), path, error_)
	return rv
}


// Changes the path of the current working directory to the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/changeCurrentDirectoryPath(_:)
func (f_ FileManager) ChangeCurrentDirectoryPath(path IString) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("changeCurrentDirectoryPath:"), path)
	return rv
}


// Returns an array of strings representing the user-visible components of a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/componentsToDisplay(forPath:)
func (f_ FileManager) ComponentsToDisplayForPath(path IString) []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("componentsToDisplayForPath:"), path)
	return rv
}


// Returns the contents of the file at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contents(atPath:)
func (f_ FileManager) ContentsAtPath(path IString) IData {
	rv := objc.Send[Data](f_.ID, objc.Sel("contentsAtPath:"), path)
	return rv
}


// Returns a Boolean value that indicates whether the files or directories in specified paths have the same contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contentsEqual(atPath:andPath:)
func (f_ FileManager) ContentsEqualAtPathAndPath(path1 IString, path2 IString) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("contentsEqualAtPath:andPath:"), path1, path2)
	return rv
}


// Performs a shallow search of the specified directory and returns URLs for the contained items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contentsOfDirectory(at:includingPropertiesForKeys:options:)
func (f_ FileManager) ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url IURL, keys []string, mask DirectoryEnumerationOptions, error_ IError) []URL {
	rv := objc.Send[[]URL](f_.ID, objc.Sel("contentsOfDirectoryAtURL:includingPropertiesForKeys:options:error:"), url, keys, mask, error_)
	return rv
}


// Performs a shallow search of the specified directory and returns the paths of any contained items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contentsOfDirectory(atPath:)
func (f_ FileManager) ContentsOfDirectoryAtPathError(path IString, error_ IError) []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("contentsOfDirectoryAtPath:error:"), path, error_)
	return rv
}


// Copies the file at the specified URL to a new location synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/copyItem(at:to:)
func (f_ FileManager) CopyItemAtURLToURLError(srcURL IURL, dstURL IURL, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("copyItemAtURL:toURL:error:"), srcURL, dstURL, error_)
	return rv
}


// Copies the item at the specified path to a new location synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/copyItem(atPath:toPath:)
func (f_ FileManager) CopyItemAtPathToPathError(srcPath IString, dstPath IString, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("copyItemAtPath:toPath:error:"), srcPath, dstPath, error_)
	return rv
}


// Creates a directory with the given attributes at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createDirectory(at:withIntermediateDirectories:attributes:)
func (f_ FileManager) CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError(url IURL, createIntermediates bool, attributes IDictionary, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createDirectoryAtURL:withIntermediateDirectories:attributes:error:"), url, createIntermediates, attributes, error_)
	return rv
}


// Creates a directory with given attributes at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createDirectory(atPath:withIntermediateDirectories:attributes:)
func (f_ FileManager) CreateDirectoryAtPathWithIntermediateDirectoriesAttributesError(path IString, createIntermediates bool, attributes IDictionary, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createDirectoryAtPath:withIntermediateDirectories:attributes:error:"), path, createIntermediates, attributes, error_)
	return rv
}


// Creates a file with the specified content and attributes at the given location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createFile(atPath:contents:attributes:)
func (f_ FileManager) CreateFileAtPathContentsAttributes(path IString, data IData, attr IDictionary) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createFileAtPath:contents:attributes:"), path, data, attr)
	return rv
}


// Creates a symbolic link at the specified URL that points to an item at the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createSymbolicLink(at:withDestinationURL:)
func (f_ FileManager) CreateSymbolicLinkAtURLWithDestinationURLError(url IURL, destURL IURL, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createSymbolicLinkAtURL:withDestinationURL:error:"), url, destURL, error_)
	return rv
}


// Creates a symbolic link that points to the specified destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createSymbolicLink(atPath:withDestinationPath:)
func (f_ FileManager) CreateSymbolicLinkAtPathWithDestinationPathError(path IString, destPath IString, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createSymbolicLinkAtPath:withDestinationPath:error:"), path, destPath, error_)
	return rv
}


// Returns the path of the item pointed to by a symbolic link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/destinationOfSymbolicLink(atPath:)
func (f_ FileManager) DestinationOfSymbolicLinkAtPathError(path IString, error_ IError) IString {
	rv := objc.Send[String](f_.ID, objc.Sel("destinationOfSymbolicLinkAtPath:error:"), path, error_)
	return rv
}


// Returns the display name of the file or directory at a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/displayName(atPath:)
func (f_ FileManager) DisplayNameAtPath(path IString) IString {
	rv := objc.Send[String](f_.ID, objc.Sel("displayNameAtPath:"), path)
	return rv
}


// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/enumerator(atPath:)
func (f_ FileManager) EnumeratorAtPath(path IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("enumeratorAtPath:"), path)
	return rv
}


// Removes the local copy of the specified item that’s stored in iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/evictUbiquitousItem(at:)
func (f_ FileManager) EvictUbiquitousItemAtURLError(url IURL, error_ IError) bool {
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
func (f_ FileManager) FileExistsAtPath(path IString) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("fileExistsAtPath:"), path)
	return rv
}


// Returns a Boolean value that indicates whether a file or directory exists at a specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileExists(atPath:isDirectory:)
func (f_ FileManager) FileExistsAtPathIsDirectory(path IString, isDirectory objectivec.IObject) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("fileExistsAtPath:isDirectory:"), path, isDirectory)
	return rv
}


// Returns a C-string representation of a given path that properly encodes Unicode strings for use by the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileSystemRepresentation(withPath:)
func (f_ FileManager) FileSystemRepresentationWithPath(path IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("fileSystemRepresentationWithPath:"), path)
	return rv
}


// Returns the services provided by the File Provider extension that manages the item at the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/getFileProviderServicesForItem(at:completionHandler:)
func (f_ FileManager) GetFileProviderServicesForItemAtURLCompletionHandler(url IURL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getFileProviderServicesForItemAtURL:completionHandler:"), url, completionHandler)
}


// Returns the home directory for the specified user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/homeDirectory(forUser:)
func (f_ FileManager) HomeDirectoryForUser(userName IString) IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("homeDirectoryForUser:"), userName)
	return rv
}


// Returns a Boolean value that indicates whether the invoking object appears able to delete a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isDeletableFile(atPath:)
func (f_ FileManager) IsDeletableFileAtPath(path IString) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDeletableFileAtPath:"), path)
	return rv
}


// Returns a Boolean value that indicates whether the operating system appears able to execute a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isExecutableFile(atPath:)
func (f_ FileManager) IsExecutableFileAtPath(path IString) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isExecutableFileAtPath:"), path)
	return rv
}


// Returns a Boolean value that indicates whether the invoking object appears able to read a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isReadableFile(atPath:)
func (f_ FileManager) IsReadableFileAtPath(path IString) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isReadableFileAtPath:"), path)
	return rv
}


// Returns a Boolean indicating whether the item is targeted for storage in iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isUbiquitousItem(at:)
func (f_ FileManager) IsUbiquitousItemAtURL(url IURL) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isUbiquitousItemAtURL:"), url)
	return rv
}


// Returns a Boolean value that indicates whether the invoking object appears able to write to a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isWritableFile(atPath:)
func (f_ FileManager) IsWritableFileAtPath(path IString) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isWritableFileAtPath:"), path)
	return rv
}


// Creates a hard link between the items at the specified URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/linkItem(at:to:)
func (f_ FileManager) LinkItemAtURLToURLError(srcURL IURL, dstURL IURL, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("linkItemAtURL:toURL:error:"), srcURL, dstURL, error_)
	return rv
}


// Creates a hard link between the items at the specified paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/linkItem(atPath:toPath:)
func (f_ FileManager) LinkItemAtPathToPathError(srcPath IString, dstPath IString, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("linkItemAtPath:toPath:error:"), srcPath, dstPath, error_)
	return rv
}


// Returns an array of URLs that identify the mounted volumes available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/mountedVolumeURLs(includingResourceValuesForKeys:options:)
func (f_ FileManager) MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys []string, options VolumeEnumerationOptions) []URL {
	rv := objc.Send[[]URL](f_.ID, objc.Sel("mountedVolumeURLsIncludingResourceValuesForKeys:options:"), propertyKeys, options)
	return rv
}


// Moves the file or directory at the specified URL to a new location synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/moveItem(at:to:)
func (f_ FileManager) MoveItemAtURLToURLError(srcURL IURL, dstURL IURL, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("moveItemAtURL:toURL:error:"), srcURL, dstURL, error_)
	return rv
}


// Moves the file or directory at the specified path to a new location synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/moveItem(atPath:toPath:)
func (f_ FileManager) MoveItemAtPathToPathError(srcPath IString, dstPath IString, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("moveItemAtPath:toPath:error:"), srcPath, dstPath, error_)
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
func (f_ FileManager) RemoveItemAtURLError(URL IURL, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("removeItemAtURL:error:"), URL, error_)
	return rv
}


// Removes the file or directory at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/removeItem(atPath:)
func (f_ FileManager) RemoveItemAtPathError(path IString, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("removeItemAtPath:error:"), path, error_)
	return rv
}


// Replaces the contents of the item at the specified URL in a manner that ensures no data loss occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/replaceItem(at:withItemAt:backupItemName:options:resultingItemURL:)
func (f_ FileManager) ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL IURL, newItemURL IURL, backupItemName IString, options FileManagerItemReplacementOptions, resultingURL IURL, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("replaceItemAtURL:withItemAtURL:backupItemName:options:resultingItemURL:error:"), originalItemURL, newItemURL, backupItemName, options, resultingURL, error_)
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
func (f_ FileManager) SetAttributesOfItemAtPathError(attributes IDictionary, path IString, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("setAttributes:ofItemAtPath:error:"), attributes, path, error_)
	return rv
}


// Indicates whether the item at the specified URL should be stored in iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/setUbiquitous(_:itemAt:destinationURL:)
func (f_ FileManager) SetUbiquitousItemAtURLDestinationURLError(flag bool, url IURL, destinationURL IURL, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("setUbiquitous:itemAtURL:destinationURL:error:"), flag, url, destinationURL, error_)
	return rv
}


// Starts downloading (if necessary) the specified item to the local system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/startDownloadingUbiquitousItem(at:)
func (f_ FileManager) StartDownloadingUbiquitousItemAtURLError(url IURL, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("startDownloadingUbiquitousItemAtURL:error:"), url, error_)
	return rv
}


// Returns an object whose contents are derived from the specified C-string path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/string(withFileSystemRepresentation:length:)
func (f_ FileManager) StringWithFileSystemRepresentationLength(str objectivec.IObject, len_ uint) IString {
	rv := objc.Send[String](f_.ID, objc.Sel("stringWithFileSystemRepresentation:length:"), str, len_)
	return rv
}


// Returns an array of strings identifying the paths for all items in the specified directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/subpaths(atPath:)
func (f_ FileManager) SubpathsAtPath(path IString) []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("subpathsAtPath:"), path)
	return rv
}


// Performs a deep enumeration of the specified directory and returns the paths of all of the contained subdirectories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/subpathsOfDirectory(atPath:)
func (f_ FileManager) SubpathsOfDirectoryAtPathError(path IString, error_ IError) []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("subpathsOfDirectoryAtPath:error:"), path, error_)
	return rv
}


// Moves an item to the trash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/trashItem(at:resultingItemURL:)
func (f_ FileManager) TrashItemAtURLResultingItemURLError(url IURL, outResultingURL IURL, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("trashItemAtURL:resultingItemURL:error:"), url, outResultingURL, error_)
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
func (f_ FileManager) URLForDirectoryInDomainAppropriateForURLCreateError(directory SearchPathDirectory, domain SearchPathDomainMask, url IURL, shouldCreate bool, error_ IError) IURL {
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
func (f_ FileManager) URLForUbiquityContainerIdentifier(containerIdentifier IString) IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("URLForUbiquityContainerIdentifier:"), containerIdentifier)
	return rv
}


// Returns an array of URLs for the specified common directory in the requested domains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/urls(for:in:)
func (f_ FileManager) URLsForDirectoryInDomains(directory SearchPathDirectory, domainMask SearchPathDomainMask) []URL {
	rv := objc.Send[[]URL](f_.ID, objc.Sel("URLsForDirectory:inDomains:"), directory, domainMask)
	return rv
}


// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:
func (f_ FileManager) EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url IURL, keys []string, mask DirectoryEnumerationOptions, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:"), url, keys, mask, handler)
	return rv
}







// The path to the program’s current directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/currentDirectoryPath
func (f_ FileManager) CurrentDirectoryPath() IString {
	rv := objc.Send[String](f_.ID, objc.Sel("currentDirectoryPath"))
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


// The temporary directory for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/temporaryDirectory
func (f_ FileManager) TemporaryDirectory() IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("temporaryDirectory"))
	return rv
}


// An opaque token that represents the current user’s iCloud Drive Documents identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ubiquityIdentityToken
func (f_ FileManager) UbiquityIdentityToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("ubiquityIdentityToken"))
	return rv
}


// The process identifier of the process that prevented a volume from unmounting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerunmountdissentingprocessidentifiererrorkey
func (f_ FileManager) NSFileManagerUnmountDissentingProcessIdentifierErrorKey() IString {
	rv := objc.Send[String](f_.ID, objc.Sel("NSFileManagerUnmountDissentingProcessIdentifierErrorKey"))
	return rv
}


// The version of the Foundation framework in which
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfoundationversionwithfilemanagerresourceforksupport
func (f_ FileManager) NSFoundationVersionWithFileManagerResourceForkSupport() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("NSFoundationVersionWithFileManagerResourceForkSupport"))
	return rv
}


// The version of the Foundation framework in which
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfoundationversionwithfilemanagerresourceforksupport
func (f_ FileManager) SetNSFoundationVersionWithFileManagerResourceForkSupport(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNSFoundationVersionWithFileManagerResourceForkSupport:"), value)
}







