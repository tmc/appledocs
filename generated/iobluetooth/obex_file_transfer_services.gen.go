// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [OBEXFileTransferServices] class.
var (
	OBEXFileTransferServicesClass     _OBEXFileTransferServicesClass
	OBEXFileTransferServicesClassOnce sync.Once
)

func getOBEXFileTransferServicesClass() _OBEXFileTransferServicesClass {
	OBEXFileTransferServicesClassOnce.Do(func() {
		OBEXFileTransferServicesClass = _OBEXFileTransferServicesClass{objc.GetClass("OBEXFileTransferServices")}
	})
	return OBEXFileTransferServicesClass
}

type _OBEXFileTransferServicesClass struct {
	class objc.Class
}

// An interface definition for the [OBEXFileTransferServices] class.
type IOBEXFileTransferServices interface {
	objectivec.IObject
	Abort() unsafe.Pointer
	ChangeCurrentFolderBackward() unsafe.Pointer
	ChangeCurrentFolderForwardToPath(inDirName string) unsafe.Pointer
	ChangeCurrentFolderToRoot() unsafe.Pointer
	ConnectToFTPService() unsafe.Pointer
	ConnectToObjectPushService() unsafe.Pointer
	CopyRemoteFileToLocalPath(inRemoteFileName string, inLocalPathAndName string) unsafe.Pointer
	CreateFolder(inDirName string) unsafe.Pointer
	CurrentPath() string
	Disconnect() unsafe.Pointer
	GetDefaultVCard(inLocalPathAndName string) unsafe.Pointer
	IsBusy() bool
	IsConnected() bool
	RemoveItem(inItemName string) unsafe.Pointer
	RetrieveFolderListing() unsafe.Pointer
	SendDataTypeName(inData unsafe.Pointer, inType string, inName string) unsafe.Pointer
	SendFile(inLocalPathAndName string) unsafe.Pointer
}

// Implements advanced OBEX operations in addition to simple PUT and GET.
//
// All operations are asynchronous and will callback over a respective delegate method if the initial return value is successful. The initial return value usually concerns the state of this object where as the delegate return value reflects the response of the remote device.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices
type OBEXFileTransferServices struct {
	objectivec.Object
}

// OBEXFileTransferServicesFrom constructs a [OBEXFileTransferServices] from an unsafe.Pointer.
//
// Implements advanced OBEX operations in addition to simple PUT and GET.
func OBEXFileTransferServicesFrom(ptr unsafe.Pointer) OBEXFileTransferServices {
	return OBEXFileTransferServices{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OBEXFileTransferServicesClass) Alloc() OBEXFileTransferServices {
	rv := objc.Send[OBEXFileTransferServices](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OBEXFileTransferServicesClass) New() OBEXFileTransferServices {
	rv := objc.Send[OBEXFileTransferServices](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OBEXFileTransferServices) Init() OBEXFileTransferServices {
	rv := objc.Send[OBEXFileTransferServices](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OBEXFileTransferServices) Autorelease() OBEXFileTransferServices {
	rv := objc.Send[OBEXFileTransferServices](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOBEXFileTransferServices creates a new OBEXFileTransferServices instance.
func NewOBEXFileTransferServices() OBEXFileTransferServices {
	return getOBEXFileTransferServicesClass().New()
}




// Create a new OBEXFileTransferServices object
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/init(obexSession:)
func NewOBEXFileTransferServicesWithOBEXSession(inOBEXSession unsafe.Pointer) OBEXFileTransferServices {
	instance := getOBEXFileTransferServicesClass().Alloc()
	rv := objc.Send[OBEXFileTransferServices](instance.ID, objc.Sel("initWithOBEXSession:"), inOBEXSession)
	rv.Autorelease()
	return rv
}


// Create a new OBEXFileTransferServices object
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/withOBEXSession(_:)
func (oc _OBEXFileTransferServicesClass) WithOBEXSession(inOBEXSession unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("withOBEXSession:"), inOBEXSession)
	return rv
}

// Abort the current operation
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/abort()
func (o_ OBEXFileTransferServices) Abort() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("abort"))
	return rv
}

// Change to the directory above the current level if not at the root
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/changeCurrentFolderBackward()
func (o_ OBEXFileTransferServices) ChangeCurrentFolderBackward() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("changeCurrentFolderBackward"))
	return rv
}

// Change the remote path
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/changeCurrentFolderForward(toPath:)
func (o_ OBEXFileTransferServices) ChangeCurrentFolderForwardToPath(inDirName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("changeCurrentFolderForwardToPath:"), objc.String(inDirName))
	return rv
}

// Asynchronously change to the remote root directory
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/changeCurrentFolderToRoot()
func (o_ OBEXFileTransferServices) ChangeCurrentFolderToRoot() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("changeCurrentFolderToRoot"))
	return rv
}

// Connect to a remote device for FTP operations
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/connectToFTPService()
func (o_ OBEXFileTransferServices) ConnectToFTPService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("connectToFTPService"))
	return rv
}

// Connect to a remote device for ObjectPush operations. Most of the FTP functionality of this object will be disabled.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/connectToObjectPushService()
func (o_ OBEXFileTransferServices) ConnectToObjectPushService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("connectToObjectPushService"))
	return rv
}

// Copy a remote file to a local path
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/copyRemoteFile(_:toLocalPath:)
func (o_ OBEXFileTransferServices) CopyRemoteFileToLocalPath(inRemoteFileName string, inLocalPathAndName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("copyRemoteFile:toLocalPath:"), objc.String(inRemoteFileName), objc.String(inLocalPathAndName))
	return rv
}

// Create a folder on the remote target
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/createFolder(_:)
func (o_ OBEXFileTransferServices) CreateFolder(inDirName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("createFolder:"), objc.String(inDirName))
	return rv
}

// Get the remote current directory path during an FTP session
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/currentPath()
func (o_ OBEXFileTransferServices) CurrentPath() string {
	rv := objc.Send[string](o_.ID, objc.Sel("currentPath"))
	return rv
}

// Disconnect from the remote device
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/disconnect()
func (o_ OBEXFileTransferServices) Disconnect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("disconnect"))
	return rv
}

// Get the remote default VCard, if it is supported
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/getDefaultVCard(_:)
func (o_ OBEXFileTransferServices) GetDefaultVCard(inLocalPathAndName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("getDefaultVCard:"), objc.String(inLocalPathAndName))
	return rv
}

// Get the action state of the module
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/isBusy()
func (o_ OBEXFileTransferServices) IsBusy() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isBusy"))
	return rv
}

// Get the connected state of this module.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/isConnected()
func (o_ OBEXFileTransferServices) IsConnected() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isConnected"))
	return rv
}

// Remove a remote item.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/removeItem(_:)
func (o_ OBEXFileTransferServices) RemoveItem(inItemName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("removeItem:"), objc.String(inItemName))
	return rv
}

// Get a remote directory listing
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/retrieveFolderListing()
func (o_ OBEXFileTransferServices) RetrieveFolderListing() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("retrieveFolderListing"))
	return rv
}

// Send data to a remote target
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/send(_:type:name:)
func (o_ OBEXFileTransferServices) SendDataTypeName(inData unsafe.Pointer, inType string, inName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("sendData:type:name:"), inData, objc.String(inType), objc.String(inName))
	return rv
}

// Put a local file to the remote target
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/sendFile(_:)
func (o_ OBEXFileTransferServices) SendFile(inLocalPathAndName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("sendFile:"), objc.String(inLocalPathAndName))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/delegate
func (o_ OBEXFileTransferServices) Delegate() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/delegate
func (o_ OBEXFileTransferServices) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDelegate:"), value)
}


