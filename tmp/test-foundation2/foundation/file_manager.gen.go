// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var fileManagerClass _FileManagerClass

func init() {
	fileManagerClass = _FileManagerClass{objc.GetClass("NSFileManager")}
}

type _FileManagerClass struct {
	class objc.Class
}

type FileManager struct {
	objc.ID
}

func FileManagerFrom(ptr unsafe.Pointer) FileManager {
	return FileManager{
		ID: objc.ID(ptr),
	}
}


// Initializes a file manager object that is authorized to perform privileged file system operations. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/init(authorization:)
func (fc _FileManagerClass) FileManagerWithAuthorization(authorization unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fileManagerWithAuthorization:"), authorization)
	return rv
}
// Returns a dictionary that describes the attributes of the mounted file system on which a given path resides. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/attributesOfFileSystem(forPath:)
func (f_ FileManager) AttributesOfFileSystemForPathError(path string, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("attributesOfFileSystemForPath:error:"), path, error)
	return rv
}
// Returns the attributes of the item at a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/attributesOfItem(atPath:)
func (f_ FileManager) AttributesOfItemAtPathError(path string, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("attributesOfItemAtPath:error:"), path, error)
	return rv
}
// Changes the path of the current working directory to the specified path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/changeCurrentDirectoryPath(_:)
func (f_ FileManager) ChangeCurrentDirectoryPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("changeCurrentDirectoryPath:"), path)
	return rv
}
// Changes the attributes of a given file or directory. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/changeFileAttributes(_:atPath:)
func (f_ FileManager) ChangeFileAttributesAtPath(attributes unsafe.Pointer, path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("changeFileAttributes:atPath:"), attributes, path)
	return rv
}
// Returns an array of strings representing the user-visible components of a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/componentsToDisplay(forPath:)
func (f_ FileManager) ComponentsToDisplayForPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("componentsToDisplayForPath:"), path)
	return rv
}
// Returns the container directory associated with the specified security application group identifier. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/containerURL(forSecurityApplicationGroupIdentifier:)
func (f_ FileManager) ContainerURLForSecurityApplicationGroupIdentifier(groupIdentifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("containerURLForSecurityApplicationGroupIdentifier:"), groupIdentifier)
	return rv
}
// Returns the contents of the file at the specified path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contents(atPath:)
func (f_ FileManager) ContentsAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("contentsAtPath:"), path)
	return rv
}
// Returns a Boolean value that indicates whether the files or directories in specified paths have the same contents. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contentsEqual(atPath:andPath:)
func (f_ FileManager) ContentsEqualAtPathAndPath(path1 string, path2 string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("contentsEqualAtPath:andPath:"), path1, path2)
	return rv
}
// Performs a shallow search of the specified directory and returns URLs for the contained items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contentsOfDirectory(at:includingPropertiesForKeys:options:)
func (f_ FileManager) ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url unsafe.Pointer, keys unsafe.Pointer, mask unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("contentsOfDirectoryAtURL:includingPropertiesForKeys:options:error:"), url, keys, mask, error)
	return rv
}
// Performs a shallow search of the specified directory and returns the paths of any contained items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/contentsOfDirectory(atPath:)
func (f_ FileManager) ContentsOfDirectoryAtPathError(path string, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("contentsOfDirectoryAtPath:error:"), path, error)
	return rv
}
// Copies the file at the specified URL to a new location synchronously. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/copyItem(at:to:)
func (f_ FileManager) CopyItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("copyItemAtURL:toURL:error:"), srcURL, dstURL, error)
	return rv
}
// Copies the item at the specified path to a new location synchronously. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/copyItem(atPath:toPath:)
func (f_ FileManager) CopyItemAtPathToPathError(srcPath string, dstPath string, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("copyItemAtPath:toPath:error:"), srcPath, dstPath, error)
	return rv
}
// Creates a directory with the given attributes at the specified URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createDirectory(at:withIntermediateDirectories:attributes:)
func (f_ FileManager) CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError(url unsafe.Pointer, createIntermediates bool, attributes unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createDirectoryAtURL:withIntermediateDirectories:attributes:error:"), url, createIntermediates, attributes, error)
	return rv
}
// Creates a directory (without contents) at a given path with given attributes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createDirectory(atPath:attributes:)
func (f_ FileManager) CreateDirectoryAtPathAttributes(path string, attributes unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createDirectoryAtPath:attributes:"), path, attributes)
	return rv
}
// Creates a directory with given attributes at the specified path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createDirectory(atPath:withIntermediateDirectories:attributes:)
func (f_ FileManager) CreateDirectoryAtPathWithIntermediateDirectoriesAttributesError(path string, createIntermediates bool, attributes unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createDirectoryAtPath:withIntermediateDirectories:attributes:error:"), path, createIntermediates, attributes, error)
	return rv
}
// Creates a file with the specified content and attributes at the given location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createFile(atPath:contents:attributes:)
func (f_ FileManager) CreateFileAtPathContentsAttributes(path string, data unsafe.Pointer, attr unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createFileAtPath:contents:attributes:"), path, data, attr)
	return rv
}
// Creates a symbolic link at the specified URL that points to an item at the given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createSymbolicLink(at:withDestinationURL:)
func (f_ FileManager) CreateSymbolicLinkAtURLWithDestinationURLError(url unsafe.Pointer, destURL unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createSymbolicLinkAtURL:withDestinationURL:error:"), url, destURL, error)
	return rv
}
// Creates a symbolic link identified by a given path that refers to a given location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createSymbolicLink(atPath:pathContent:)
func (f_ FileManager) CreateSymbolicLinkAtPathPathContent(path string, otherpath string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createSymbolicLinkAtPath:pathContent:"), path, otherpath)
	return rv
}
// Creates a symbolic link that points to the specified destination. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/createSymbolicLink(atPath:withDestinationPath:)
func (f_ FileManager) CreateSymbolicLinkAtPathWithDestinationPathError(path string, destPath string, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("createSymbolicLinkAtPath:withDestinationPath:error:"), path, destPath, error)
	return rv
}
// Returns the path of the item pointed to by a symbolic link. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/destinationOfSymbolicLink(atPath:)
func (f_ FileManager) DestinationOfSymbolicLinkAtPathError(path string, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("destinationOfSymbolicLinkAtPath:error:"), path, error)
	return rv
}
// Returns the directories and files (including symbolic links) contained in a given directory. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/directoryContents(atPath:)
func (f_ FileManager) DirectoryContentsAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("directoryContentsAtPath:"), path)
	return rv
}
// Returns the display name of the file or directory at a specified path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/displayName(atPath:)
func (f_ FileManager) DisplayNameAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("displayNameAtPath:"), path)
	return rv
}
// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/enumerator(atPath:)
func (f_ FileManager) EnumeratorAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("enumeratorAtPath:"), path)
	return rv
}
// Removes the local copy of the specified item that’s stored in iCloud. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/evictUbiquitousItem(at:)
func (f_ FileManager) EvictUbiquitousItemAtURLError(url unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("evictUbiquitousItemAtURL:error:"), url, error)
	return rv
}
// Asynchronously fetches the latest remote version of a given item from the server. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fetchLatestRemoteVersionOfItem(at:completionHandler:)
func (f_ FileManager) FetchLatestRemoteVersionOfItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("fetchLatestRemoteVersionOfItemAtURL:completionHandler:"), url, completionHandler)
}
// Returns a dictionary that describes the POSIX attributes of the file specified at a given. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileAttributes(atPath:traverseLink:)
func (f_ FileManager) FileAttributesAtPathTraverseLink(path string, yorn bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fileAttributesAtPath:traverseLink:"), path, yorn)
	return rv
}
// Returns a Boolean value that indicates whether a file or directory exists at a specified path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileExists(atPath:)
func (f_ FileManager) FileExistsAtPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("fileExistsAtPath:"), path)
	return rv
}
// Returns a Boolean value that indicates whether a file or directory exists at a specified path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileExists(atPath:isDirectory:)
func (f_ FileManager) FileExistsAtPathIsDirectory(path string, isDirectory unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("fileExistsAtPath:isDirectory:"), path, isDirectory)
	return rv
}
// Returns a dictionary that describes the attributes of the mounted file system on which a given path resides. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileSystemAttributes(atPath:)
func (f_ FileManager) FileSystemAttributesAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fileSystemAttributesAtPath:"), path)
	return rv
}
// Returns a C-string representation of a given path that properly encodes Unicode strings for use by the file system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/fileSystemRepresentation(withPath:)
func (f_ FileManager) FileSystemRepresentationWithPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fileSystemRepresentationWithPath:"), path)
	return rv
}
// Returns the services provided by the File Provider extension that manages the item at the given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/getFileProviderServicesForItem(at:completionHandler:)
func (f_ FileManager) GetFileProviderServicesForItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getFileProviderServicesForItemAtURL:completionHandler:"), url, completionHandler)
}
// Determines the type of relationship that exists between a system directory and the specified item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/getRelationship(_:of:in:toItemAt:)
func (f_ FileManager) GetRelationshipOfDirectoryInDomainToItemAtURLError(outRelationship unsafe.Pointer, directory unsafe.Pointer, domainMask unsafe.Pointer, url unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("getRelationship:ofDirectory:inDomain:toItemAtURL:error:"), outRelationship, directory, domainMask, url, error)
	return rv
}
// Determines the type of relationship that exists between a directory and an item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/getRelationship(_:ofDirectoryAt:toItemAt:)
func (f_ FileManager) GetRelationshipOfDirectoryAtURLToItemAtURLError(outRelationship unsafe.Pointer, directoryURL unsafe.Pointer, otherURL unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("getRelationship:ofDirectoryAtURL:toItemAtURL:error:"), outRelationship, directoryURL, otherURL, error)
	return rv
}
// Returns the home directory for the specified user. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/homeDirectory(forUser:)
func (f_ FileManager) HomeDirectoryForUser(userName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("homeDirectoryForUser:"), userName)
	return rv
}
// Returns a Boolean value that indicates whether the invoking object appears able to delete a specified file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isDeletableFile(atPath:)
func (f_ FileManager) IsDeletableFileAtPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDeletableFileAtPath:"), path)
	return rv
}
// Returns a Boolean value that indicates whether the operating system appears able to execute a specified file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isExecutableFile(atPath:)
func (f_ FileManager) IsExecutableFileAtPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isExecutableFileAtPath:"), path)
	return rv
}
// Returns a Boolean value that indicates whether the invoking object appears able to read a specified file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isReadableFile(atPath:)
func (f_ FileManager) IsReadableFileAtPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isReadableFileAtPath:"), path)
	return rv
}
// Returns a Boolean indicating whether the item is targeted for storage in iCloud. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isUbiquitousItem(at:)
func (f_ FileManager) IsUbiquitousItemAtURL(url unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isUbiquitousItemAtURL:"), url)
	return rv
}
// Returns a Boolean value that indicates whether the invoking object appears able to write to a specified file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/isWritableFile(atPath:)
func (f_ FileManager) IsWritableFileAtPath(path string) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isWritableFileAtPath:"), path)
	return rv
}
// Creates a hard link between the items at the specified URLs. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/linkItem(at:to:)
func (f_ FileManager) LinkItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("linkItemAtURL:toURL:error:"), srcURL, dstURL, error)
	return rv
}
// Creates a hard link between the items at the specified paths. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/linkItem(atPath:toPath:)
func (f_ FileManager) LinkItemAtPathToPathError(srcPath string, dstPath string, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("linkItemAtPath:toPath:error:"), srcPath, dstPath, error)
	return rv
}
// Returns an array of URLs that identify the mounted volumes available on the device. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/mountedVolumeURLs(includingResourceValuesForKeys:options:)
func (f_ FileManager) MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("mountedVolumeURLsIncludingResourceValuesForKeys:options:"), propertyKeys, options)
	return rv
}
// Moves the file or directory at the specified URL to a new location synchronously. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/moveItem(at:to:)
func (f_ FileManager) MoveItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("moveItemAtURL:toURL:error:"), srcURL, dstURL, error)
	return rv
}
// Moves the file or directory at the specified path to a new location synchronously. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/moveItem(atPath:toPath:)
func (f_ FileManager) MoveItemAtPathToPathError(srcPath string, dstPath string, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("moveItemAtPath:toPath:error:"), srcPath, dstPath, error)
	return rv
}
// Returns the path of the directory or file that a symbolic link at a given path refers to. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/pathContentOfSymbolicLink(atPath:)
func (f_ FileManager) PathContentOfSymbolicLinkAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("pathContentOfSymbolicLinkAtPath:"), path)
	return rv
}
// Asynchronously pauses sync of an item at the given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/pauseSyncForUbiquitousItem(at:completionHandler:)
func (f_ FileManager) PauseSyncForUbiquitousItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("pauseSyncForUbiquitousItemAtURL:completionHandler:"), url, completionHandler)
}
// Removes the file or directory at the specified URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/removeItem(at:)
func (f_ FileManager) RemoveItemAtURLError(URL unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("removeItemAtURL:error:"), URL, error)
	return rv
}
// Removes the file or directory at the specified path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/removeItem(atPath:)
func (f_ FileManager) RemoveItemAtPathError(path string, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("removeItemAtPath:error:"), path, error)
	return rv
}
// Replaces the contents of the item at the specified URL in a manner that ensures no data loss occurs. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/replaceItem(at:withItemAt:backupItemName:options:resultingItemURL:)
func (f_ FileManager) ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL unsafe.Pointer, newItemURL unsafe.Pointer, backupItemName string, options unsafe.Pointer, resultingURL unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("replaceItemAtURL:withItemAtURL:backupItemName:options:resultingItemURL:error:"), originalItemURL, newItemURL, backupItemName, options, resultingURL, error)
	return rv
}
// Asynchronously resumes the sync on a paused item using the given resume behavior. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/resumeSyncForUbiquitousItem(at:with:completionHandler:)
func (f_ FileManager) ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler(url unsafe.Pointer, behavior unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("resumeSyncForUbiquitousItemAtURL:withBehavior:completionHandler:"), url, behavior, completionHandler)
}
// Sets the attributes of the specified file or directory. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/setAttributes(_:ofItemAtPath:)
func (f_ FileManager) SetAttributesOfItemAtPathError(attributes unsafe.Pointer, path string, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("setAttributes:ofItemAtPath:error:"), attributes, path, error)
	return rv
}
// Indicates whether the item at the specified URL should be stored in iCloud. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/setUbiquitous(_:itemAt:destinationURL:)
func (f_ FileManager) SetUbiquitousItemAtURLDestinationURLError(flag bool, url unsafe.Pointer, destinationURL unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("setUbiquitous:itemAtURL:destinationURL:error:"), flag, url, destinationURL, error)
	return rv
}
// Starts downloading (if necessary) the specified item to the local system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/startDownloadingUbiquitousItem(at:)
func (f_ FileManager) StartDownloadingUbiquitousItemAtURLError(url unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("startDownloadingUbiquitousItemAtURL:error:"), url, error)
	return rv
}
// Returns an object whose contents are derived from the specified C-string path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/string(withFileSystemRepresentation:length:)
func (f_ FileManager) StringWithFileSystemRepresentationLength(str unsafe.Pointer, len uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("stringWithFileSystemRepresentation:length:"), str, len)
	return rv
}
// Returns an array of strings identifying the paths for all items in the specified directory. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/subpaths(atPath:)
func (f_ FileManager) SubpathsAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("subpathsAtPath:"), path)
	return rv
}
// Performs a deep enumeration of the specified directory and returns the paths of all of the contained subdirectories. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/subpathsOfDirectory(atPath:)
func (f_ FileManager) SubpathsOfDirectoryAtPathError(path string, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("subpathsOfDirectoryAtPath:error:"), path, error)
	return rv
}
// Moves an item to the trash. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/trashItem(at:resultingItemURL:)
func (f_ FileManager) TrashItemAtURLResultingItemURLError(url unsafe.Pointer, outResultingURL unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("trashItemAtURL:resultingItemURL:error:"), url, outResultingURL, error)
	return rv
}
// Starts the process of unmounting the specified volume. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/unmountVolume(at:options:completionHandler:)
func (f_ FileManager) UnmountVolumeAtURLOptionsCompletionHandler(url unsafe.Pointer, mask unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("unmountVolumeAtURL:options:completionHandler:"), url, mask, completionHandler)
}
// Asynchronously uploads the local version of the item using the provided conflict resolution policy. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/uploadLocalVersionOfUbiquitousItem(at:withConflictResolutionPolicy:completionHandler:)
func (f_ FileManager) UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler(url unsafe.Pointer, conflictResolutionPolicy unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("uploadLocalVersionOfUbiquitousItemAtURL:withConflictResolutionPolicy:completionHandler:"), url, conflictResolutionPolicy, completionHandler)
}
// Locates and optionally creates the specified common directory in a domain. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/url(for:in:appropriateFor:create:)
func (f_ FileManager) URLForDirectoryInDomainAppropriateForURLCreateError(directory unsafe.Pointer, domain unsafe.Pointer, url unsafe.Pointer, shouldCreate bool, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("URLForDirectory:inDomain:appropriateForURL:create:error:"), directory, domain, url, shouldCreate, error)
	return rv
}
// Returns a URL that can be emailed to users to allow them to download a copy of a flat file item from iCloud. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/url(forPublishingUbiquitousItemAt:expiration:)
func (f_ FileManager) URLForPublishingUbiquitousItemAtURLExpirationDateError(url unsafe.Pointer, outDate unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("URLForPublishingUbiquitousItemAtURL:expirationDate:error:"), url, outDate, error)
	return rv
}
// Returns the URL for the iCloud container associated with the specified identifier and establishes access to that container. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/url(forUbiquityContainerIdentifier:)
func (f_ FileManager) URLForUbiquityContainerIdentifier(containerIdentifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("URLForUbiquityContainerIdentifier:"), containerIdentifier)
	return rv
}
// Returns an array of URLs for the specified common directory in the requested domains. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/urls(for:in:)
func (f_ FileManager) URLsForDirectoryInDomains(directory unsafe.Pointer, domainMask unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("URLsForDirectory:inDomains:"), directory, domainMask)
	return rv
}
// Copies the directory or file specified in a given path to a different location in the file system identified by another path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/copyPath:toPath:handler:
func (f_ FileManager) CopyPathToPathHandler(src string, dest string, handler objc.ID) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("copyPath:toPath:handler:"), src, dest, handler)
	return rv
}
// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:
func (f_ FileManager) EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url unsafe.Pointer, keys unsafe.Pointer, mask unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:"), url, keys, mask, handler)
	return rv
}
// Creates a link from a source to a destination. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/linkPath:toPath:handler:
func (f_ FileManager) LinkPathToPathHandler(src string, dest string, handler objc.ID) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("linkPath:toPath:handler:"), src, dest, handler)
	return rv
}
// Moves the directory or file specified by a given path to a different location in the file system identified by another path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/movePath:toPath:handler:
func (f_ FileManager) MovePathToPathHandler(src string, dest string, handler objc.ID) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("movePath:toPath:handler:"), src, dest, handler)
	return rv
}
// Deletes the file, link, or directory (including, recursively, all subdirectories, files, and links in the directory) identified by a given path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileManager/removeFileAtPath:handler:
func (f_ FileManager) RemoveFileAtPathHandler(path string, handler objc.ID) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("removeFileAtPath:handler:"), path, handler)
	return rv
}


