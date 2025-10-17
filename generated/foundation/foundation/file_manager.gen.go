// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FileManager] class.
var FileManagerClass objc.Class

func init() {
	FileManagerClass = objc.GetClass("NSFileManager")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/init(authorization:)
func (fc FileManager) FileManagerWithAuthorization(authorization unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("fileManagerWithAuthorization:")
	ret := objc.ID(FileManagerClass).Send(sel, authorization)
	return unsafe.Pointer(ret)
}
// Returns a dictionary that describes the attributes of the mounted file system on which a given path resides. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/attributesOfFileSystem(forPath:)
func (f_ FileManager) AttributesOfFileSystemForPathError(path string, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("attributesOfFileSystemForPath:error:")
	ret := f_.ID.Send(sel, path, error)
	return unsafe.Pointer(ret)
}
// Returns the attributes of the item at a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/attributesOfItem(atPath:)
func (f_ FileManager) AttributesOfItemAtPathError(path string, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("attributesOfItemAtPath:error:")
	ret := f_.ID.Send(sel, path, error)
	return unsafe.Pointer(ret)
}
// Changes the path of the current working directory to the specified path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/changeCurrentDirectoryPath(_:)
func (f_ FileManager) ChangeCurrentDirectoryPath(path string) bool {
	sel := objc.RegisterName("changeCurrentDirectoryPath:")
	ret := f_.ID.Send(sel, path)
	return ret != 0
}
// Changes the attributes of a given file or directory. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/changeFileAttributes(_:atPath:)
func (f_ FileManager) ChangeFileAttributesAtPath(attributes unsafe.Pointer, path string) bool {
	sel := objc.RegisterName("changeFileAttributes:atPath:")
	ret := f_.ID.Send(sel, attributes, path)
	return ret != 0
}
// Returns an array of strings representing the user-visible components of a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/componentsToDisplay(forPath:)
func (f_ FileManager) ComponentsToDisplayForPath(path string) unsafe.Pointer {
	sel := objc.RegisterName("componentsToDisplayForPath:")
	ret := f_.ID.Send(sel, path)
	return unsafe.Pointer(ret)
}
// Returns the container directory associated with the specified security application group identifier. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/containerURL(forSecurityApplicationGroupIdentifier:)
func (f_ FileManager) ContainerURLForSecurityApplicationGroupIdentifier(groupIdentifier string) unsafe.Pointer {
	sel := objc.RegisterName("containerURLForSecurityApplicationGroupIdentifier:")
	ret := f_.ID.Send(sel, groupIdentifier)
	return unsafe.Pointer(ret)
}
// Returns the contents of the file at the specified path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/contents(atPath:)
func (f_ FileManager) ContentsAtPath(path string) unsafe.Pointer {
	sel := objc.RegisterName("contentsAtPath:")
	ret := f_.ID.Send(sel, path)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value that indicates whether the files or directories in specified paths have the same contents. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/contentsEqual(atPath:andPath:)
func (f_ FileManager) ContentsEqualAtPathAndPath(path1 string, path2 string) bool {
	sel := objc.RegisterName("contentsEqualAtPath:andPath:")
	ret := f_.ID.Send(sel, path1, path2)
	return ret != 0
}
// Performs a shallow search of the specified directory and returns URLs for the contained items. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/contentsOfDirectory(at:includingPropertiesForKeys:options:)
func (f_ FileManager) ContentsOfDirectoryAtURLIncludingPropertiesForKeysOptionsError(url unsafe.Pointer, keys unsafe.Pointer, mask unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("contentsOfDirectoryAtURL:includingPropertiesForKeys:options:error:")
	ret := f_.ID.Send(sel, url, keys, mask, error)
	return unsafe.Pointer(ret)
}
// Performs a shallow search of the specified directory and returns the paths of any contained items. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/contentsOfDirectory(atPath:)
func (f_ FileManager) ContentsOfDirectoryAtPathError(path string, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("contentsOfDirectoryAtPath:error:")
	ret := f_.ID.Send(sel, path, error)
	return unsafe.Pointer(ret)
}
// Copies the file at the specified URL to a new location synchronously. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/copyItem(at:to:)
func (f_ FileManager) CopyItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("copyItemAtURL:toURL:error:")
	ret := f_.ID.Send(sel, srcURL, dstURL, error)
	return ret != 0
}
// Copies the item at the specified path to a new location synchronously. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/copyItem(atPath:toPath:)
func (f_ FileManager) CopyItemAtPathToPathError(srcPath string, dstPath string, error unsafe.Pointer) bool {
	sel := objc.RegisterName("copyItemAtPath:toPath:error:")
	ret := f_.ID.Send(sel, srcPath, dstPath, error)
	return ret != 0
}
// Creates a directory with the given attributes at the specified URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/createDirectory(at:withIntermediateDirectories:attributes:)
func (f_ FileManager) CreateDirectoryAtURLWithIntermediateDirectoriesAttributesError(url unsafe.Pointer, createIntermediates bool, attributes unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("createDirectoryAtURL:withIntermediateDirectories:attributes:error:")
	ret := f_.ID.Send(sel, url, createIntermediates, attributes, error)
	return ret != 0
}
// Creates a directory (without contents) at a given path with given attributes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/createDirectory(atPath:attributes:)
func (f_ FileManager) CreateDirectoryAtPathAttributes(path string, attributes unsafe.Pointer) bool {
	sel := objc.RegisterName("createDirectoryAtPath:attributes:")
	ret := f_.ID.Send(sel, path, attributes)
	return ret != 0
}
// Creates a directory with given attributes at the specified path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/createDirectory(atPath:withIntermediateDirectories:attributes:)
func (f_ FileManager) CreateDirectoryAtPathWithIntermediateDirectoriesAttributesError(path string, createIntermediates bool, attributes unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("createDirectoryAtPath:withIntermediateDirectories:attributes:error:")
	ret := f_.ID.Send(sel, path, createIntermediates, attributes, error)
	return ret != 0
}
// Creates a file with the specified content and attributes at the given location. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/createFile(atPath:contents:attributes:)
func (f_ FileManager) CreateFileAtPathContentsAttributes(path string, data unsafe.Pointer, attr unsafe.Pointer) bool {
	sel := objc.RegisterName("createFileAtPath:contents:attributes:")
	ret := f_.ID.Send(sel, path, data, attr)
	return ret != 0
}
// Creates a symbolic link at the specified URL that points to an item at the given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/createSymbolicLink(at:withDestinationURL:)
func (f_ FileManager) CreateSymbolicLinkAtURLWithDestinationURLError(url unsafe.Pointer, destURL unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("createSymbolicLinkAtURL:withDestinationURL:error:")
	ret := f_.ID.Send(sel, url, destURL, error)
	return ret != 0
}
// Creates a symbolic link identified by a given path that refers to a given location. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/createSymbolicLink(atPath:pathContent:)
func (f_ FileManager) CreateSymbolicLinkAtPathPathContent(path string, otherpath string) bool {
	sel := objc.RegisterName("createSymbolicLinkAtPath:pathContent:")
	ret := f_.ID.Send(sel, path, otherpath)
	return ret != 0
}
// Creates a symbolic link that points to the specified destination. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/createSymbolicLink(atPath:withDestinationPath:)
func (f_ FileManager) CreateSymbolicLinkAtPathWithDestinationPathError(path string, destPath string, error unsafe.Pointer) bool {
	sel := objc.RegisterName("createSymbolicLinkAtPath:withDestinationPath:error:")
	ret := f_.ID.Send(sel, path, destPath, error)
	return ret != 0
}
// Returns the path of the item pointed to by a symbolic link. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/destinationOfSymbolicLink(atPath:)
func (f_ FileManager) DestinationOfSymbolicLinkAtPathError(path string, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("destinationOfSymbolicLinkAtPath:error:")
	ret := f_.ID.Send(sel, path, error)
	return unsafe.Pointer(ret)
}
// Returns the directories and files (including symbolic links) contained in a given directory. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/directoryContents(atPath:)
func (f_ FileManager) DirectoryContentsAtPath(path string) unsafe.Pointer {
	sel := objc.RegisterName("directoryContentsAtPath:")
	ret := f_.ID.Send(sel, path)
	return unsafe.Pointer(ret)
}
// Returns the display name of the file or directory at a specified path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/displayName(atPath:)
func (f_ FileManager) DisplayNameAtPath(path string) unsafe.Pointer {
	sel := objc.RegisterName("displayNameAtPath:")
	ret := f_.ID.Send(sel, path)
	return unsafe.Pointer(ret)
}
// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/enumerator(atPath:)
func (f_ FileManager) EnumeratorAtPath(path string) unsafe.Pointer {
	sel := objc.RegisterName("enumeratorAtPath:")
	ret := f_.ID.Send(sel, path)
	return unsafe.Pointer(ret)
}
// Removes the local copy of the specified item that’s stored in iCloud. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/evictUbiquitousItem(at:)
func (f_ FileManager) EvictUbiquitousItemAtURLError(url unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("evictUbiquitousItemAtURL:error:")
	ret := f_.ID.Send(sel, url, error)
	return ret != 0
}
// Asynchronously fetches the latest remote version of a given item from the server. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/fetchLatestRemoteVersionOfItem(at:completionHandler:)
func (f_ FileManager) FetchLatestRemoteVersionOfItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("fetchLatestRemoteVersionOfItemAtURL:completionHandler:")
	f_.ID.Send(sel, url, completionHandler)
}
// Returns a dictionary that describes the POSIX attributes of the file specified at a given. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/fileAttributes(atPath:traverseLink:)
func (f_ FileManager) FileAttributesAtPathTraverseLink(path string, yorn bool) unsafe.Pointer {
	sel := objc.RegisterName("fileAttributesAtPath:traverseLink:")
	ret := f_.ID.Send(sel, path, yorn)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value that indicates whether a file or directory exists at a specified path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/fileExists(atPath:)
func (f_ FileManager) FileExistsAtPath(path string) bool {
	sel := objc.RegisterName("fileExistsAtPath:")
	ret := f_.ID.Send(sel, path)
	return ret != 0
}
// Returns a Boolean value that indicates whether a file or directory exists at a specified path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/fileExists(atPath:isDirectory:)
func (f_ FileManager) FileExistsAtPathIsDirectory(path string, isDirectory unsafe.Pointer) bool {
	sel := objc.RegisterName("fileExistsAtPath:isDirectory:")
	ret := f_.ID.Send(sel, path, isDirectory)
	return ret != 0
}
// Returns a dictionary that describes the attributes of the mounted file system on which a given path resides. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/fileSystemAttributes(atPath:)
func (f_ FileManager) FileSystemAttributesAtPath(path string) unsafe.Pointer {
	sel := objc.RegisterName("fileSystemAttributesAtPath:")
	ret := f_.ID.Send(sel, path)
	return unsafe.Pointer(ret)
}
// Returns a C-string representation of a given path that properly encodes Unicode strings for use by the file system. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/fileSystemRepresentation(withPath:)
func (f_ FileManager) FileSystemRepresentationWithPath(path string) unsafe.Pointer {
	sel := objc.RegisterName("fileSystemRepresentationWithPath:")
	ret := f_.ID.Send(sel, path)
	return unsafe.Pointer(ret)
}
// Returns the services provided by the File Provider extension that manages the item at the given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/getFileProviderServicesForItem(at:completionHandler:)
func (f_ FileManager) GetFileProviderServicesForItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("getFileProviderServicesForItemAtURL:completionHandler:")
	f_.ID.Send(sel, url, completionHandler)
}
// Determines the type of relationship that exists between a system directory and the specified item. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/getRelationship(_:of:in:toItemAt:)
func (f_ FileManager) GetRelationshipOfDirectoryInDomainToItemAtURLError(outRelationship unsafe.Pointer, directory unsafe.Pointer, domainMask unsafe.Pointer, url unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("getRelationship:ofDirectory:inDomain:toItemAtURL:error:")
	ret := f_.ID.Send(sel, outRelationship, directory, domainMask, url, error)
	return ret != 0
}
// Determines the type of relationship that exists between a directory and an item. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/getRelationship(_:ofDirectoryAt:toItemAt:)
func (f_ FileManager) GetRelationshipOfDirectoryAtURLToItemAtURLError(outRelationship unsafe.Pointer, directoryURL unsafe.Pointer, otherURL unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("getRelationship:ofDirectoryAtURL:toItemAtURL:error:")
	ret := f_.ID.Send(sel, outRelationship, directoryURL, otherURL, error)
	return ret != 0
}
// Returns the home directory for the specified user. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/homeDirectory(forUser:)
func (f_ FileManager) HomeDirectoryForUser(userName string) unsafe.Pointer {
	sel := objc.RegisterName("homeDirectoryForUser:")
	ret := f_.ID.Send(sel, userName)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value that indicates whether the invoking object appears able to delete a specified file. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/isDeletableFile(atPath:)
func (f_ FileManager) IsDeletableFileAtPath(path string) bool {
	sel := objc.RegisterName("isDeletableFileAtPath:")
	ret := f_.ID.Send(sel, path)
	return ret != 0
}
// Returns a Boolean value that indicates whether the operating system appears able to execute a specified file. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/isExecutableFile(atPath:)
func (f_ FileManager) IsExecutableFileAtPath(path string) bool {
	sel := objc.RegisterName("isExecutableFileAtPath:")
	ret := f_.ID.Send(sel, path)
	return ret != 0
}
// Returns a Boolean value that indicates whether the invoking object appears able to read a specified file. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/isReadableFile(atPath:)
func (f_ FileManager) IsReadableFileAtPath(path string) bool {
	sel := objc.RegisterName("isReadableFileAtPath:")
	ret := f_.ID.Send(sel, path)
	return ret != 0
}
// Returns a Boolean indicating whether the item is targeted for storage in iCloud. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/isUbiquitousItem(at:)
func (f_ FileManager) IsUbiquitousItemAtURL(url unsafe.Pointer) bool {
	sel := objc.RegisterName("isUbiquitousItemAtURL:")
	ret := f_.ID.Send(sel, url)
	return ret != 0
}
// Returns a Boolean value that indicates whether the invoking object appears able to write to a specified file. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/isWritableFile(atPath:)
func (f_ FileManager) IsWritableFileAtPath(path string) bool {
	sel := objc.RegisterName("isWritableFileAtPath:")
	ret := f_.ID.Send(sel, path)
	return ret != 0
}
// Creates a hard link between the items at the specified URLs. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/linkItem(at:to:)
func (f_ FileManager) LinkItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("linkItemAtURL:toURL:error:")
	ret := f_.ID.Send(sel, srcURL, dstURL, error)
	return ret != 0
}
// Creates a hard link between the items at the specified paths. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/linkItem(atPath:toPath:)
func (f_ FileManager) LinkItemAtPathToPathError(srcPath string, dstPath string, error unsafe.Pointer) bool {
	sel := objc.RegisterName("linkItemAtPath:toPath:error:")
	ret := f_.ID.Send(sel, srcPath, dstPath, error)
	return ret != 0
}
// Returns an array of URLs that identify the mounted volumes available on the device. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/mountedVolumeURLs(includingResourceValuesForKeys:options:)
func (f_ FileManager) MountedVolumeURLsIncludingResourceValuesForKeysOptions(propertyKeys unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("mountedVolumeURLsIncludingResourceValuesForKeys:options:")
	ret := f_.ID.Send(sel, propertyKeys, options)
	return unsafe.Pointer(ret)
}
// Moves the file or directory at the specified URL to a new location synchronously. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/moveItem(at:to:)
func (f_ FileManager) MoveItemAtURLToURLError(srcURL unsafe.Pointer, dstURL unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("moveItemAtURL:toURL:error:")
	ret := f_.ID.Send(sel, srcURL, dstURL, error)
	return ret != 0
}
// Moves the file or directory at the specified path to a new location synchronously. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/moveItem(atPath:toPath:)
func (f_ FileManager) MoveItemAtPathToPathError(srcPath string, dstPath string, error unsafe.Pointer) bool {
	sel := objc.RegisterName("moveItemAtPath:toPath:error:")
	ret := f_.ID.Send(sel, srcPath, dstPath, error)
	return ret != 0
}
// Returns the path of the directory or file that a symbolic link at a given path refers to. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/pathContentOfSymbolicLink(atPath:)
func (f_ FileManager) PathContentOfSymbolicLinkAtPath(path string) unsafe.Pointer {
	sel := objc.RegisterName("pathContentOfSymbolicLinkAtPath:")
	ret := f_.ID.Send(sel, path)
	return unsafe.Pointer(ret)
}
// Asynchronously pauses sync of an item at the given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/pauseSyncForUbiquitousItem(at:completionHandler:)
func (f_ FileManager) PauseSyncForUbiquitousItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("pauseSyncForUbiquitousItemAtURL:completionHandler:")
	f_.ID.Send(sel, url, completionHandler)
}
// Removes the file or directory at the specified URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/removeItem(at:)
func (f_ FileManager) RemoveItemAtURLError(URL unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("removeItemAtURL:error:")
	ret := f_.ID.Send(sel, URL, error)
	return ret != 0
}
// Removes the file or directory at the specified path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/removeItem(atPath:)
func (f_ FileManager) RemoveItemAtPathError(path string, error unsafe.Pointer) bool {
	sel := objc.RegisterName("removeItemAtPath:error:")
	ret := f_.ID.Send(sel, path, error)
	return ret != 0
}
// Replaces the contents of the item at the specified URL in a manner that ensures no data loss occurs. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/replaceItem(at:withItemAt:backupItemName:options:resultingItemURL:)
func (f_ FileManager) ReplaceItemAtURLWithItemAtURLBackupItemNameOptionsResultingItemURLError(originalItemURL unsafe.Pointer, newItemURL unsafe.Pointer, backupItemName string, options unsafe.Pointer, resultingURL unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("replaceItemAtURL:withItemAtURL:backupItemName:options:resultingItemURL:error:")
	ret := f_.ID.Send(sel, originalItemURL, newItemURL, backupItemName, options, resultingURL, error)
	return ret != 0
}
// Asynchronously resumes the sync on a paused item using the given resume behavior. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/resumeSyncForUbiquitousItem(at:with:completionHandler:)
func (f_ FileManager) ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler(url unsafe.Pointer, behavior unsafe.Pointer, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("resumeSyncForUbiquitousItemAtURL:withBehavior:completionHandler:")
	f_.ID.Send(sel, url, behavior, completionHandler)
}
// Sets the attributes of the specified file or directory. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/setAttributes(_:ofItemAtPath:)
func (f_ FileManager) SetAttributesOfItemAtPathError(attributes unsafe.Pointer, path string, error unsafe.Pointer) bool {
	sel := objc.RegisterName("setAttributes:ofItemAtPath:error:")
	ret := f_.ID.Send(sel, attributes, path, error)
	return ret != 0
}
// Indicates whether the item at the specified URL should be stored in iCloud. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/setUbiquitous(_:itemAt:destinationURL:)
func (f_ FileManager) SetUbiquitousItemAtURLDestinationURLError(flag bool, url unsafe.Pointer, destinationURL unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("setUbiquitous:itemAtURL:destinationURL:error:")
	ret := f_.ID.Send(sel, flag, url, destinationURL, error)
	return ret != 0
}
// Starts downloading (if necessary) the specified item to the local system. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/startDownloadingUbiquitousItem(at:)
func (f_ FileManager) StartDownloadingUbiquitousItemAtURLError(url unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("startDownloadingUbiquitousItemAtURL:error:")
	ret := f_.ID.Send(sel, url, error)
	return ret != 0
}
// Returns an   object whose contents are derived from the specified C-string path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/string(withFileSystemRepresentation:length:)
func (f_ FileManager) StringWithFileSystemRepresentationLength(str unsafe.Pointer, len uint) unsafe.Pointer {
	sel := objc.RegisterName("stringWithFileSystemRepresentation:length:")
	ret := f_.ID.Send(sel, str, len)
	return unsafe.Pointer(ret)
}
// Returns an array of strings identifying the paths for all items in the specified directory. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/subpaths(atPath:)
func (f_ FileManager) SubpathsAtPath(path string) unsafe.Pointer {
	sel := objc.RegisterName("subpathsAtPath:")
	ret := f_.ID.Send(sel, path)
	return unsafe.Pointer(ret)
}
// Performs a deep enumeration of the specified directory and returns the paths of all of the contained subdirectories. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/subpathsOfDirectory(atPath:)
func (f_ FileManager) SubpathsOfDirectoryAtPathError(path string, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("subpathsOfDirectoryAtPath:error:")
	ret := f_.ID.Send(sel, path, error)
	return unsafe.Pointer(ret)
}
// Moves an item to the trash. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/trashItem(at:resultingItemURL:)
func (f_ FileManager) TrashItemAtURLResultingItemURLError(url unsafe.Pointer, outResultingURL unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("trashItemAtURL:resultingItemURL:error:")
	ret := f_.ID.Send(sel, url, outResultingURL, error)
	return ret != 0
}
// Starts the process of unmounting the specified volume. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/unmountVolume(at:options:completionHandler:)
func (f_ FileManager) UnmountVolumeAtURLOptionsCompletionHandler(url unsafe.Pointer, mask unsafe.Pointer, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("unmountVolumeAtURL:options:completionHandler:")
	f_.ID.Send(sel, url, mask, completionHandler)
}
// Asynchronously uploads the local version of the item using the provided conflict resolution policy. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/uploadLocalVersionOfUbiquitousItem(at:withConflictResolutionPolicy:completionHandler:)
func (f_ FileManager) UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler(url unsafe.Pointer, conflictResolutionPolicy unsafe.Pointer, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("uploadLocalVersionOfUbiquitousItemAtURL:withConflictResolutionPolicy:completionHandler:")
	f_.ID.Send(sel, url, conflictResolutionPolicy, completionHandler)
}
// Locates and optionally creates the specified common directory in a domain. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/url(for:in:appropriateFor:create:)
func (f_ FileManager) URLForDirectoryInDomainAppropriateForURLCreateError(directory unsafe.Pointer, domain unsafe.Pointer, url unsafe.Pointer, shouldCreate bool, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("URLForDirectory:inDomain:appropriateForURL:create:error:")
	ret := f_.ID.Send(sel, directory, domain, url, shouldCreate, error)
	return unsafe.Pointer(ret)
}
// Returns a URL that can be emailed to users to allow them to download a copy of a flat file item from iCloud. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/url(forPublishingUbiquitousItemAt:expiration:)
func (f_ FileManager) URLForPublishingUbiquitousItemAtURLExpirationDateError(url unsafe.Pointer, outDate unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("URLForPublishingUbiquitousItemAtURL:expirationDate:error:")
	ret := f_.ID.Send(sel, url, outDate, error)
	return unsafe.Pointer(ret)
}
// Returns the URL for the iCloud container associated with the specified identifier and establishes access to that container. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/url(forUbiquityContainerIdentifier:)
func (f_ FileManager) URLForUbiquityContainerIdentifier(containerIdentifier string) unsafe.Pointer {
	sel := objc.RegisterName("URLForUbiquityContainerIdentifier:")
	ret := f_.ID.Send(sel, containerIdentifier)
	return unsafe.Pointer(ret)
}
// Returns an array of URLs for the specified common directory in the requested domains. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/FileManager/urls(for:in:)
func (f_ FileManager) URLsForDirectoryInDomains(directory unsafe.Pointer, domainMask unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("URLsForDirectory:inDomains:")
	ret := f_.ID.Send(sel, directory, domainMask)
	return unsafe.Pointer(ret)
}
// Copies the directory or file specified in a given path to a different location in the file system identified by another path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSFileManager/copyPath:toPath:handler:
func (f_ FileManager) CopyPathToPathHandler(src string, dest string, handler objc.ID) bool {
	sel := objc.RegisterName("copyPath:toPath:handler:")
	ret := f_.ID.Send(sel, src, dest, handler)
	return ret != 0
}
// Returns a directory enumerator object that can be used to perform a deep enumeration of the directory at the specified URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSFileManager/enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:
func (f_ FileManager) EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler(url unsafe.Pointer, keys unsafe.Pointer, mask unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("enumeratorAtURL:includingPropertiesForKeys:options:errorHandler:")
	ret := f_.ID.Send(sel, url, keys, mask, handler)
	return unsafe.Pointer(ret)
}
// Creates a link from a source to a destination. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSFileManager/linkPath:toPath:handler:
func (f_ FileManager) LinkPathToPathHandler(src string, dest string, handler objc.ID) bool {
	sel := objc.RegisterName("linkPath:toPath:handler:")
	ret := f_.ID.Send(sel, src, dest, handler)
	return ret != 0
}
// Moves the directory or file specified by a given path to a different location in the file system identified by another path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSFileManager/movePath:toPath:handler:
func (f_ FileManager) MovePathToPathHandler(src string, dest string, handler objc.ID) bool {
	sel := objc.RegisterName("movePath:toPath:handler:")
	ret := f_.ID.Send(sel, src, dest, handler)
	return ret != 0
}
// Deletes the file, link, or directory (including, recursively, all subdirectories, files, and links in the directory) identified by a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSFileManager/removeFileAtPath:handler:
func (f_ FileManager) RemoveFileAtPathHandler(path string, handler objc.ID) bool {
	sel := objc.RegisterName("removeFileAtPath:handler:")
	ret := f_.ID.Send(sel, path, handler)
	return ret != 0
}


