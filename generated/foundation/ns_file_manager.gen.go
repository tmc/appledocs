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
	AttributesOfFileSystemForPathError(path string, error_ unsafe.Pointer) unsafe.Pointer
	AttributesOfItemAtPathError(path string, error_ unsafe.Pointer) unsafe.Pointer
	ChangeCurrentDirectoryPath(path string) bool
	ChangeFileAttributesAtPath(attributes objc.ID, path string) bool
	ComponentsToDisplayForPath(path string) []string
	ContainerURLForSecurityApplicationGroupIdentifier(groupIdentifier string) unsafe.Pointer
	ContentsAtPath(path string) unsafe.Pointer
	ContentsEqualAtPathAndPath(path1 string, path2 string) bool
	ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url unsafe.Pointer, keys unsafe.Pointer, mask unsafe.Pointer, error_ unsafe.Pointer) []URL
	ContentsOfDirectoryAtPathError(path string, error_ unsafe.Pointer) []string
	CopyItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error_ unsafe.Pointer) bool
	CopyItemAtPathToPathError(srcPath string, dstPath string, error_ unsafe.Pointer) bool
	CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError(url unsafe.Pointer, createIntermediates bool, attributes unsafe.Pointer, error_ unsafe.Pointer) bool
	CreateDirectoryAtPathAttributes(path string, attributes objc.ID) bool
	CreateDirectoryAtPathWithIntermediateDirectoriesAttributesError(path string, createIntermediates bool, attributes unsafe.Pointer, error_ unsafe.Pointer) bool
	CreateFileAtPathContentsAttributes(path string, data unsafe.Pointer, attr unsafe.Pointer) bool
	CreateSymbolicLinkAtURLWithDestinationURLError(url unsafe.Pointer, destURL unsafe.Pointer, error_ unsafe.Pointer) bool
	CreateSymbolicLinkAtPathPathContent(path string, otherpath string) bool
	CreateSymbolicLinkAtPathWithDestinationPathError(path string, destPath string, error_ unsafe.Pointer) bool
	DestinationOfSymbolicLinkAtPathError(path string, error_ unsafe.Pointer) string
	DirectoryContentsAtPath(path string) unsafe.Pointer
	DisplayNameAtPath(path string) string
	EnumeratorAtPath(path string) unsafe.Pointer
	EvictUbiquitousItemAtURLError(url unsafe.Pointer, error_ unsafe.Pointer) bool
	FetchLatestRemoteVersionOfItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer)
	FileAttributesAtPathTraverseLink(path string, yorn bool) unsafe.Pointer
	FileExistsAtPath(path string) bool
	FileExistsAtPathIsDirectory(path string, isDirectory unsafe.Pointer) bool
	FileSystemAttributesAtPath(path string) unsafe.Pointer
	FileSystemRepresentationWithPath(path string) unsafe.Pointer
	GetFileProviderServicesForItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer)
	GetRelationshipOfDirectoryInDomainToItemAtURLError(outRelationship unsafe.Pointer, directory unsafe.Pointer, domainMask unsafe.Pointer, url unsafe.Pointer, error_ unsafe.Pointer) bool
	GetRelationshipOfDirectoryAtURLToItemAtURLError(outRelationship unsafe.Pointer, directoryURL unsafe.Pointer, otherURL unsafe.Pointer, error_ unsafe.Pointer) bool
	HomeDirectoryForUser(userName string) unsafe.Pointer
	IsDeletableFileAtPath(path string) bool
	IsExecutableFileAtPath(path string) bool
	IsReadableFileAtPath(path string) bool
	IsUbiquitousItemAtURL(url unsafe.Pointer) bool
	IsWritableFileAtPath(path string) bool
	LinkItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error_ unsafe.Pointer) bool
	LinkItemAtPathToPathError(srcPath string, dstPath string, error_ unsafe.Pointer) bool
	MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys unsafe.Pointer, options unsafe.Pointer) []URL
	MoveItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error_ unsafe.Pointer) bool
	MoveItemAtPathToPathError(srcPath string, dstPath string, error_ unsafe.Pointer) bool
	PathContentOfSymbolicLinkAtPath(path string) string
	PauseSyncForUbiquitousItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer)
	RemoveItemAtURLError(URL unsafe.Pointer, error_ unsafe.Pointer) bool
	RemoveItemAtPathError(path string, error_ unsafe.Pointer) bool
	ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL unsafe.Pointer, newItemURL unsafe.Pointer, backupItemName string, options unsafe.Pointer, resultingURL unsafe.Pointer, error_ unsafe.Pointer) bool
	ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler(url unsafe.Pointer, behavior unsafe.Pointer, completionHandler unsafe.Pointer)
	SetAttributesOfItemAtPathError(attributes unsafe.Pointer, path string, error_ unsafe.Pointer) bool
	SetUbiquitousItemAtURLDestinationURLError(flag bool, url unsafe.Pointer, destinationURL unsafe.Pointer, error_ unsafe.Pointer) bool
	StartDownloadingUbiquitousItemAtURLError(url unsafe.Pointer, error_ unsafe.Pointer) bool
	StringWithFileSystemRepresentationLength(str unsafe.Pointer, len uint) string
	SubpathsAtPath(path string) []string
	SubpathsOfDirectoryAtPathError(path string, error_ unsafe.Pointer) []string
	TrashItemAtURLResultingItemURLError(url unsafe.Pointer, outResultingURL unsafe.Pointer, error_ unsafe.Pointer) bool
	UnmountVolumeAtURLOptionsCompletionHandler(url unsafe.Pointer, mask unsafe.Pointer, completionHandler unsafe.Pointer)
	UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler(url unsafe.Pointer, conflictResolutionPolicy unsafe.Pointer, completionHandler unsafe.Pointer)
	URLForDirectoryInDomainAppropriateForURLCreateError(directory unsafe.Pointer, domain unsafe.Pointer, url unsafe.Pointer, shouldCreate bool, error_ unsafe.Pointer) unsafe.Pointer
	URLForPublishingUbiquitousItemAtURLExpirationDateError(url unsafe.Pointer, outDate unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	URLForUbiquityContainerIdentifier(containerIdentifier string) unsafe.Pointer
	URLsForDirectoryInDomains(directory unsafe.Pointer, domainMask unsafe.Pointer) []URL
	CopyPathToPathHandler(src string, dest string, handler objc.ID) bool
	EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url unsafe.Pointer, keys unsafe.Pointer, mask unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer
	LinkPathToPathHandler(src string, dest string, handler objc.ID) bool
	MovePathToPathHandler(src string, dest string, handler objc.ID) bool
	RemoveFileAtPathHandler(path string, handler objc.ID) bool
}

// A convenient interface to the contents of the file system, and the primary means of interacting with it.
//
// A file manager object lets you examine the contents of the file system and make changes to it. The class provides convenient access to a shared file manager object that is suitable for most types of file-related manipulations. A file manager object is typically your primary mode of interaction with the file system. You use it to locate, create, copy, and move files and directories. You also use it to get information about a file or directory or change some of its attributes. When specifying the location of files, you can use either or objects. The use of the class is generally preferred for specifying file-system items because URLs can convert path information to a more efficient representation internally. You can also obtain a bookmark from an object, which is similar to an alias and offers a more sure way of locating the file or directory later. If you are moving, copying, linking, or removing files or directories, you can use a delegate in conjunction with a file manager object to manage those operations. The delegate’s role is to affirm the operation and to decide whether to proceed when errors occur. In macOS 10.7 and later, the delegate must conform to the protocol. In iOS 5.0 and later and in macOS 10.7 and later, includes methods for managing items stored in iCloud. Files and directories tagged for cloud storage are synced to iCloud so that they can be made available to the user’s iOS devices and Macintosh computers. Changes to an item in one location are propagated to all other locations to ensure the items stay in sync.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/init(authorization:)
func NewFileManagerWithAuthorization(authorization unsafe.Pointer) FileManager {
	rv := objc.Send[FileManager](objc.ID(getFileManagerClass().class), objc.Sel("fileManagerWithAuthorization:"), authorization)
	return rv
}


// Initializes a file manager object that is authorized to perform privileged file system operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/init(authorization:)
func (fc _FileManagerClass) FileManagerWithAuthorization(authorization unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fileManagerWithAuthorization:"), authorization)
	return rv
}

// The shared file manager object for the process.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/default
func (fc _FileManagerClass) DefaultManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("defaultManager"))
	return rv
}
// Returns a dictionary that describes the attributes of the mounted file system on which a given path resides.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/attributesOfFileSystem(forPath:)
func (f_ FileManager) AttributesOfFileSystemForPathError(path string, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("attributesOfFileSystemForPath:error:"), objc.String(path), error_)
	return rv
}

// Returns the attributes of the item at a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/attributesOfItem(atPath:)
func (f_ FileManager) AttributesOfItemAtPathError(path string, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("attributesOfItemAtPath:error:"), objc.String(path), error_)
	return rv
}

// Changes the path of the current working directory to the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/changeCurrentDirectoryPath(_:)
func (f_ FileManager) ChangeCurrentDirectoryPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("changeCurrentDirectoryPath:"), objc.String(path))
	return rv
}

// Changes the attributes of a given file or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/changeFileAttributes(_:atPath:)
func (f_ FileManager) ChangeFileAttributesAtPath(attributes objc.ID, path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("changeFileAttributes:atPath:"), attributes, objc.String(path))
	return rv
}

// Returns an array of strings representing the user-visible components of a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/componentsToDisplay(forPath:)
func (f_ FileManager) ComponentsToDisplayForPath(path string) []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("componentsToDisplayForPath:"), objc.String(path))
	return rv
}

// Returns the container directory associated with the specified security application group identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/containerURL(forSecurityApplicationGroupIdentifier:)
func (f_ FileManager) ContainerURLForSecurityApplicationGroupIdentifier(groupIdentifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("containerURLForSecurityApplicationGroupIdentifier:"), objc.String(groupIdentifier))
	return rv
}

// Returns the contents of the file at the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contents(atPath:)
func (f_ FileManager) ContentsAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("contentsAtPath:"), objc.String(path))
	return rv
}

// Returns a Boolean value that indicates whether the files or directories in specified paths have the same contents.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contentsEqual(atPath:andPath:)
func (f_ FileManager) ContentsEqualAtPathAndPath(path1 string, path2 string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("contentsEqualAtPath:andPath:"), objc.String(path1), objc.String(path2))
	return rv
}

// Performs a shallow search of the specified directory and returns URLs for the contained items.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contentsOfDirectory(at:includingPropertiesForKeys:options:)
func (f_ FileManager) ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url unsafe.Pointer, keys unsafe.Pointer, mask unsafe.Pointer, error_ unsafe.Pointer) []URL {
	rv := objc.Send[[]URL](f_.ID, objc.Sel("contentsOfDirectoryAtURL:includingPropertiesForKeys:options:error:"), url, keys, mask, error_)
	return rv
}

// Performs a shallow search of the specified directory and returns the paths of any contained items.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contentsOfDirectory(atPath:)
func (f_ FileManager) ContentsOfDirectoryAtPathError(path string, error_ unsafe.Pointer) []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("contentsOfDirectoryAtPath:error:"), objc.String(path), error_)
	return rv
}

// Copies the file at the specified URL to a new location synchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/copyItem(at:to:)
func (f_ FileManager) CopyItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("copyItemAtURL:toURL:error:"), srcURL, dstURL, error_)
	return rv
}

// Copies the item at the specified path to a new location synchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/copyItem(atPath:toPath:)
func (f_ FileManager) CopyItemAtPathToPathError(srcPath string, dstPath string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("copyItemAtPath:toPath:error:"), objc.String(srcPath), objc.String(dstPath), error_)
	return rv
}

// Creates a directory with the given attributes at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createDirectory(at:withIntermediateDirectories:attributes:)
func (f_ FileManager) CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError(url unsafe.Pointer, createIntermediates bool, attributes unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createDirectoryAtURL:withIntermediateDirectories:attributes:error:"), url, createIntermediates, attributes, error_)
	return rv
}

// Creates a directory (without contents) at a given path with given attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createDirectory(atPath:attributes:)
func (f_ FileManager) CreateDirectoryAtPathAttributes(path string, attributes objc.ID) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createDirectoryAtPath:attributes:"), objc.String(path), attributes)
	return rv
}

// Creates a directory with given attributes at the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createDirectory(atPath:withIntermediateDirectories:attributes:)
func (f_ FileManager) CreateDirectoryAtPathWithIntermediateDirectoriesAttributesError(path string, createIntermediates bool, attributes unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createDirectoryAtPath:withIntermediateDirectories:attributes:error:"), objc.String(path), createIntermediates, attributes, error_)
	return rv
}

// Creates a file with the specified content and attributes at the given location.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createFile(atPath:contents:attributes:)
func (f_ FileManager) CreateFileAtPathContentsAttributes(path string, data unsafe.Pointer, attr unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createFileAtPath:contents:attributes:"), objc.String(path), data, attr)
	return rv
}

// Creates a symbolic link at the specified URL that points to an item at the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createSymbolicLink(at:withDestinationURL:)
func (f_ FileManager) CreateSymbolicLinkAtURLWithDestinationURLError(url unsafe.Pointer, destURL unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createSymbolicLinkAtURL:withDestinationURL:error:"), url, destURL, error_)
	return rv
}

// Creates a symbolic link identified by a given path that refers to a given location.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createSymbolicLink(atPath:pathContent:)
func (f_ FileManager) CreateSymbolicLinkAtPathPathContent(path string, otherpath string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createSymbolicLinkAtPath:pathContent:"), objc.String(path), objc.String(otherpath))
	return rv
}

// Creates a symbolic link that points to the specified destination.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createSymbolicLink(atPath:withDestinationPath:)
func (f_ FileManager) CreateSymbolicLinkAtPathWithDestinationPathError(path string, destPath string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createSymbolicLinkAtPath:withDestinationPath:error:"), objc.String(path), objc.String(destPath), error_)
	return rv
}

// Returns the path of the item pointed to by a symbolic link.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/destinationOfSymbolicLink(atPath:)
func (f_ FileManager) DestinationOfSymbolicLinkAtPathError(path string, error_ unsafe.Pointer) string {
	rv := objc.Send[string](f_.ID, objc.Sel("destinationOfSymbolicLinkAtPath:error:"), objc.String(path), error_)
	return rv
}

// Returns the directories and files (including symbolic links) contained in a given directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/directoryContents(atPath:)
func (f_ FileManager) DirectoryContentsAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("directoryContentsAtPath:"), objc.String(path))
	return rv
}

// Returns the display name of the file or directory at a specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/displayName(atPath:)
func (f_ FileManager) DisplayNameAtPath(path string) string {
	rv := objc.Send[string](f_.ID, objc.Sel("displayNameAtPath:"), objc.String(path))
	return rv
}

// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/enumerator(atPath:)
func (f_ FileManager) EnumeratorAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("enumeratorAtPath:"), objc.String(path))
	return rv
}

// Removes the local copy of the specified item that’s stored in iCloud.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/evictUbiquitousItem(at:)
func (f_ FileManager) EvictUbiquitousItemAtURLError(url unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("evictUbiquitousItemAtURL:error:"), url, error_)
	return rv
}

// Asynchronously fetches the latest remote version of a given item from the server.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fetchLatestRemoteVersionOfItem(at:completionHandler:)
func (f_ FileManager) FetchLatestRemoteVersionOfItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("fetchLatestRemoteVersionOfItemAtURL:completionHandler:"), url, completionHandler)
}

// Returns a dictionary that describes the POSIX attributes of the file specified at a given.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileAttributes(atPath:traverseLink:)
func (f_ FileManager) FileAttributesAtPathTraverseLink(path string, yorn bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fileAttributesAtPath:traverseLink:"), objc.String(path), yorn)
	return rv
}

// Returns a Boolean value that indicates whether a file or directory exists at a specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileExists(atPath:)
func (f_ FileManager) FileExistsAtPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("fileExistsAtPath:"), objc.String(path))
	return rv
}

// Returns a Boolean value that indicates whether a file or directory exists at a specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileExists(atPath:isDirectory:)
func (f_ FileManager) FileExistsAtPathIsDirectory(path string, isDirectory unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("fileExistsAtPath:isDirectory:"), objc.String(path), isDirectory)
	return rv
}

// Returns a dictionary that describes the attributes of the mounted file system on which a given path resides.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileSystemAttributes(atPath:)
func (f_ FileManager) FileSystemAttributesAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fileSystemAttributesAtPath:"), objc.String(path))
	return rv
}

// Returns a C-string representation of a given path that properly encodes Unicode strings for use by the file system.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileSystemRepresentation(withPath:)
func (f_ FileManager) FileSystemRepresentationWithPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fileSystemRepresentationWithPath:"), objc.String(path))
	return rv
}

// Returns the services provided by the File Provider extension that manages the item at the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/getFileProviderServicesForItem(at:completionHandler:)
func (f_ FileManager) GetFileProviderServicesForItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getFileProviderServicesForItemAtURL:completionHandler:"), url, completionHandler)
}

// Determines the type of relationship that exists between a system directory and the specified item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/getRelationship(_:of:in:toItemAt:)
func (f_ FileManager) GetRelationshipOfDirectoryInDomainToItemAtURLError(outRelationship unsafe.Pointer, directory unsafe.Pointer, domainMask unsafe.Pointer, url unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("getRelationship:ofDirectory:inDomain:toItemAtURL:error:"), outRelationship, directory, domainMask, url, error_)
	return rv
}

// Determines the type of relationship that exists between a directory and an item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/getRelationship(_:ofDirectoryAt:toItemAt:)
func (f_ FileManager) GetRelationshipOfDirectoryAtURLToItemAtURLError(outRelationship unsafe.Pointer, directoryURL unsafe.Pointer, otherURL unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("getRelationship:ofDirectoryAtURL:toItemAtURL:error:"), outRelationship, directoryURL, otherURL, error_)
	return rv
}

// Returns the home directory for the specified user.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/homeDirectory(forUser:)
func (f_ FileManager) HomeDirectoryForUser(userName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("homeDirectoryForUser:"), objc.String(userName))
	return rv
}

// Returns a Boolean value that indicates whether the invoking object appears able to delete a specified file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isDeletableFile(atPath:)
func (f_ FileManager) IsDeletableFileAtPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDeletableFileAtPath:"), objc.String(path))
	return rv
}

// Returns a Boolean value that indicates whether the operating system appears able to execute a specified file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isExecutableFile(atPath:)
func (f_ FileManager) IsExecutableFileAtPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isExecutableFileAtPath:"), objc.String(path))
	return rv
}

// Returns a Boolean value that indicates whether the invoking object appears able to read a specified file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isReadableFile(atPath:)
func (f_ FileManager) IsReadableFileAtPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isReadableFileAtPath:"), objc.String(path))
	return rv
}

// Returns a Boolean indicating whether the item is targeted for storage in iCloud.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isUbiquitousItem(at:)
func (f_ FileManager) IsUbiquitousItemAtURL(url unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isUbiquitousItemAtURL:"), url)
	return rv
}

// Returns a Boolean value that indicates whether the invoking object appears able to write to a specified file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isWritableFile(atPath:)
func (f_ FileManager) IsWritableFileAtPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isWritableFileAtPath:"), objc.String(path))
	return rv
}

// Creates a hard link between the items at the specified URLs.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/linkItem(at:to:)
func (f_ FileManager) LinkItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("linkItemAtURL:toURL:error:"), srcURL, dstURL, error_)
	return rv
}

// Creates a hard link between the items at the specified paths.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/linkItem(atPath:toPath:)
func (f_ FileManager) LinkItemAtPathToPathError(srcPath string, dstPath string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("linkItemAtPath:toPath:error:"), objc.String(srcPath), objc.String(dstPath), error_)
	return rv
}

// Returns an array of URLs that identify the mounted volumes available on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/mountedVolumeURLs(includingResourceValuesForKeys:options:)
func (f_ FileManager) MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys unsafe.Pointer, options unsafe.Pointer) []URL {
	rv := objc.Send[[]URL](f_.ID, objc.Sel("mountedVolumeURLsIncludingResourceValuesForKeys:options:"), propertyKeys, options)
	return rv
}

// Moves the file or directory at the specified URL to a new location synchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/moveItem(at:to:)
func (f_ FileManager) MoveItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("moveItemAtURL:toURL:error:"), srcURL, dstURL, error_)
	return rv
}

// Moves the file or directory at the specified path to a new location synchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/moveItem(atPath:toPath:)
func (f_ FileManager) MoveItemAtPathToPathError(srcPath string, dstPath string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("moveItemAtPath:toPath:error:"), objc.String(srcPath), objc.String(dstPath), error_)
	return rv
}

// Returns the path of the directory or file that a symbolic link at a given path refers to.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/pathContentOfSymbolicLink(atPath:)
func (f_ FileManager) PathContentOfSymbolicLinkAtPath(path string) string {
	rv := objc.Send[string](f_.ID, objc.Sel("pathContentOfSymbolicLinkAtPath:"), objc.String(path))
	return rv
}

// Asynchronously pauses sync of an item at the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/pauseSyncForUbiquitousItem(at:completionHandler:)
func (f_ FileManager) PauseSyncForUbiquitousItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("pauseSyncForUbiquitousItemAtURL:completionHandler:"), url, completionHandler)
}

// Removes the file or directory at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/removeItem(at:)
func (f_ FileManager) RemoveItemAtURLError(URL unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("removeItemAtURL:error:"), URL, error_)
	return rv
}

// Removes the file or directory at the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/removeItem(atPath:)
func (f_ FileManager) RemoveItemAtPathError(path string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("removeItemAtPath:error:"), objc.String(path), error_)
	return rv
}

// Replaces the contents of the item at the specified URL in a manner that ensures no data loss occurs.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/replaceItem(at:withItemAt:backupItemName:options:resultingItemURL:)
func (f_ FileManager) ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL unsafe.Pointer, newItemURL unsafe.Pointer, backupItemName string, options unsafe.Pointer, resultingURL unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("replaceItemAtURL:withItemAtURL:backupItemName:options:resultingItemURL:error:"), originalItemURL, newItemURL, objc.String(backupItemName), options, resultingURL, error_)
	return rv
}

// Asynchronously resumes the sync on a paused item using the given resume behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/resumeSyncForUbiquitousItem(at:with:completionHandler:)
func (f_ FileManager) ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler(url unsafe.Pointer, behavior unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("resumeSyncForUbiquitousItemAtURL:withBehavior:completionHandler:"), url, behavior, completionHandler)
}

// Sets the attributes of the specified file or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/setAttributes(_:ofItemAtPath:)
func (f_ FileManager) SetAttributesOfItemAtPathError(attributes unsafe.Pointer, path string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("setAttributes:ofItemAtPath:error:"), attributes, objc.String(path), error_)
	return rv
}

// Indicates whether the item at the specified URL should be stored in iCloud.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/setUbiquitous(_:itemAt:destinationURL:)
func (f_ FileManager) SetUbiquitousItemAtURLDestinationURLError(flag bool, url unsafe.Pointer, destinationURL unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("setUbiquitous:itemAtURL:destinationURL:error:"), flag, url, destinationURL, error_)
	return rv
}

// Starts downloading (if necessary) the specified item to the local system.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/startDownloadingUbiquitousItem(at:)
func (f_ FileManager) StartDownloadingUbiquitousItemAtURLError(url unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("startDownloadingUbiquitousItemAtURL:error:"), url, error_)
	return rv
}

// Returns an object whose contents are derived from the specified C-string path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/string(withFileSystemRepresentation:length:)
func (f_ FileManager) StringWithFileSystemRepresentationLength(str unsafe.Pointer, len uint) string {
	rv := objc.Send[string](f_.ID, objc.Sel("stringWithFileSystemRepresentation:length:"), str, len)
	return rv
}

// Returns an array of strings identifying the paths for all items in the specified directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/subpaths(atPath:)
func (f_ FileManager) SubpathsAtPath(path string) []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("subpathsAtPath:"), objc.String(path))
	return rv
}

// Performs a deep enumeration of the specified directory and returns the paths of all of the contained subdirectories.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/subpathsOfDirectory(atPath:)
func (f_ FileManager) SubpathsOfDirectoryAtPathError(path string, error_ unsafe.Pointer) []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("subpathsOfDirectoryAtPath:error:"), objc.String(path), error_)
	return rv
}

// Moves an item to the trash.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/trashItem(at:resultingItemURL:)
func (f_ FileManager) TrashItemAtURLResultingItemURLError(url unsafe.Pointer, outResultingURL unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("trashItemAtURL:resultingItemURL:error:"), url, outResultingURL, error_)
	return rv
}

// Starts the process of unmounting the specified volume.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/unmountVolume(at:options:completionHandler:)
func (f_ FileManager) UnmountVolumeAtURLOptionsCompletionHandler(url unsafe.Pointer, mask unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("unmountVolumeAtURL:options:completionHandler:"), url, mask, completionHandler)
}

// Asynchronously uploads the local version of the item using the provided conflict resolution policy.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/uploadLocalVersionOfUbiquitousItem(at:withConflictResolutionPolicy:completionHandler:)
func (f_ FileManager) UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler(url unsafe.Pointer, conflictResolutionPolicy unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("uploadLocalVersionOfUbiquitousItemAtURL:withConflictResolutionPolicy:completionHandler:"), url, conflictResolutionPolicy, completionHandler)
}

// Locates and optionally creates the specified common directory in a domain.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/url(for:in:appropriateFor:create:)
func (f_ FileManager) URLForDirectoryInDomainAppropriateForURLCreateError(directory unsafe.Pointer, domain unsafe.Pointer, url unsafe.Pointer, shouldCreate bool, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("URLForDirectory:inDomain:appropriateForURL:create:error:"), directory, domain, url, shouldCreate, error_)
	return rv
}

// Returns a URL that can be emailed to users to allow them to download a copy of a flat file item from iCloud.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/url(forPublishingUbiquitousItemAt:expiration:)
func (f_ FileManager) URLForPublishingUbiquitousItemAtURLExpirationDateError(url unsafe.Pointer, outDate unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("URLForPublishingUbiquitousItemAtURL:expirationDate:error:"), url, outDate, error_)
	return rv
}

// Returns the URL for the iCloud container associated with the specified identifier and establishes access to that container.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/url(forUbiquityContainerIdentifier:)
func (f_ FileManager) URLForUbiquityContainerIdentifier(containerIdentifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("URLForUbiquityContainerIdentifier:"), objc.String(containerIdentifier))
	return rv
}

// Returns an array of URLs for the specified common directory in the requested domains.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/urls(for:in:)
func (f_ FileManager) URLsForDirectoryInDomains(directory unsafe.Pointer, domainMask unsafe.Pointer) []URL {
	rv := objc.Send[[]URL](f_.ID, objc.Sel("URLsForDirectory:inDomains:"), directory, domainMask)
	return rv
}

// Copies the directory or file specified in a given path to a different location in the file system identified by another path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/copyPath:toPath:handler:
func (f_ FileManager) CopyPathToPathHandler(src string, dest string, handler objc.ID) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("copyPath:toPath:handler:"), objc.String(src), objc.String(dest), handler)
	return rv
}

// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:
func (f_ FileManager) EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url unsafe.Pointer, keys unsafe.Pointer, mask unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:"), url, keys, mask, handler)
	return rv
}

// Creates a link from a source to a destination.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/linkPath:toPath:handler:
func (f_ FileManager) LinkPathToPathHandler(src string, dest string, handler objc.ID) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("linkPath:toPath:handler:"), objc.String(src), objc.String(dest), handler)
	return rv
}

// Moves the directory or file specified by a given path to a different location in the file system identified by another path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/movePath:toPath:handler:
func (f_ FileManager) MovePathToPathHandler(src string, dest string, handler objc.ID) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("movePath:toPath:handler:"), objc.String(src), objc.String(dest), handler)
	return rv
}

// Deletes the file, link, or directory (including, recursively, all subdirectories, files, and links in the directory) identified by a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/removeFileAtPath:handler:
func (f_ FileManager) RemoveFileAtPathHandler(path string, handler objc.ID) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("removeFileAtPath:handler:"), objc.String(path), handler)
	return rv
}

// The path to the program’s current directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/currentDirectoryPath
func (f_ FileManager) CurrentDirectoryPath() string {
	rv := objc.Send[string](f_.ID, objc.Sel("currentDirectoryPath"))
	return rv
}

// The shared file manager object for the process.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/default
func (f_ FileManager) DefaultManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("defaultManager"))
	return rv
}

// The delegate of the file manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/delegate
func (f_ FileManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate of the file manager object.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/delegate
func (f_ FileManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelegate:"), value)
}

// The home directory for the current user.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/homeDirectoryForCurrentUser
func (f_ FileManager) HomeDirectoryForCurrentUser() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("homeDirectoryForCurrentUser"))
	return rv
}

// The temporary directory for the current user.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/temporaryDirectory
func (f_ FileManager) TemporaryDirectory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("temporaryDirectory"))
	return rv
}

// An opaque token that represents the current user’s iCloud Drive Documents identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/ubiquityIdentityToken
func (f_ FileManager) UbiquityIdentityToken() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("ubiquityIdentityToken"))
	return rv
}


