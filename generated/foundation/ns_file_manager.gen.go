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
	SetCurrentDirectoryPath(value IString)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	HomeDirectoryForCurrentUser() IURL
	SetHomeDirectoryForCurrentUser(value IURL)
	TemporaryDirectory() IURL
	SetTemporaryDirectory(value IURL)
	UbiquityIdentityToken() ObjectProtocol /* not a class type */
	SetUbiquityIdentityToken(value ObjectProtocol /* not a class type */)
	NSFileManagerUnmountDissentingProcessIdentifierErrorKey() IString
	NSFoundationVersionWithFileManagerResourceForkSupport() unsafe.Pointer
	SetNSFoundationVersionWithFileManagerResourceForkSupport(value unsafe.Pointer)
	// methods:
	AttributesOfItemAtPathError(path IString, error_ IError) IDictionary
	GetFileProviderServicesForItemAtURLCompletionHandler(url IURL, completionHandler IDictionary)
	StartDownloadingUbiquitousItemAtURLError(url IURL, error_ IError) bool
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



// Returns the attributes of the item at a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/attributesOfItem(atPath:)
func (f_ FileManager) AttributesOfItemAtPathError(path IString, error_ IError) IDictionary {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("attributesOfItemAtPath:error:"), path, error_)
	return rv
}


// Returns the services provided by the File Provider extension that manages the item at the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/getFileProviderServicesForItem(at:completionHandler:)
func (f_ FileManager) GetFileProviderServicesForItemAtURLCompletionHandler(url IURL, completionHandler IDictionary) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getFileProviderServicesForItemAtURL:completionHandler:"), url, completionHandler)
}


// Starts downloading (if necessary) the specified item to the local system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/startDownloadingUbiquitousItem(at:)
func (f_ FileManager) StartDownloadingUbiquitousItemAtURLError(url IURL, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("startDownloadingUbiquitousItemAtURL:error:"), url, error_)
	return rv
}


// The path to the program’s current directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/currentdirectorypath
func (f_ FileManager) CurrentDirectoryPath() IString {
	rv := objc.Send[String](f_.ID, objc.Sel("currentDirectoryPath"))
	return rv
}


// The path to the program’s current directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/currentdirectorypath
func (f_ FileManager) SetCurrentDirectoryPath(value IString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setCurrentDirectoryPath:"), value)
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


// The home directory for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/homedirectoryforcurrentuser
func (f_ FileManager) HomeDirectoryForCurrentUser() IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("homeDirectoryForCurrentUser"))
	return rv
}


// The home directory for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/homedirectoryforcurrentuser
func (f_ FileManager) SetHomeDirectoryForCurrentUser(value IURL) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setHomeDirectoryForCurrentUser:"), value)
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


// An opaque token that represents the current user’s iCloud Drive Documents identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/ubiquityidentitytoken
func (f_ FileManager) UbiquityIdentityToken() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](f_.ID, objc.Sel("ubiquityIdentityToken"))
	return rv
}


// An opaque token that represents the current user’s iCloud Drive Documents identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filemanager/ubiquityidentitytoken
func (f_ FileManager) SetUbiquityIdentityToken(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUbiquityIdentityToken:"), value)
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



