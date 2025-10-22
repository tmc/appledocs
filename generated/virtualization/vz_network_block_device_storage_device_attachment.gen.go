// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZNetworkBlockDeviceStorageDeviceAttachment] class.
var (
	VZNetworkBlockDeviceStorageDeviceAttachmentClass     _VZNetworkBlockDeviceStorageDeviceAttachmentClass
	VZNetworkBlockDeviceStorageDeviceAttachmentClassOnce sync.Once
)

func getVZNetworkBlockDeviceStorageDeviceAttachmentClass() _VZNetworkBlockDeviceStorageDeviceAttachmentClass {
	VZNetworkBlockDeviceStorageDeviceAttachmentClassOnce.Do(func() {
		VZNetworkBlockDeviceStorageDeviceAttachmentClass = _VZNetworkBlockDeviceStorageDeviceAttachmentClass{objc.GetClass("VZNetworkBlockDeviceStorageDeviceAttachment")}
	})
	return VZNetworkBlockDeviceStorageDeviceAttachmentClass
}

type _VZNetworkBlockDeviceStorageDeviceAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [VZNetworkBlockDeviceStorageDeviceAttachment] class.
type IVZNetworkBlockDeviceStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	ForcedReadOnly() bool
	IsForcedReadOnly() bool
	SetIsForcedReadOnly(value bool)
	SynchronizationMode() VZDiskSynchronizationMode
	SetSynchronizationMode(value VZDiskSynchronizationMode)
	Timeout() unsafe.Pointer
	SetTimeout(value unsafe.Pointer)
	Url() foundation.URL
	SetUrl(value foundation.IURL)
}

// A storage device attachment backed by a Network Block Device (NBD) client.
//
// This storage device attachment provides a Network Block Device (NBD) client implementation. The NBD client connects to an NBD server referred to by an NBD Uniform Resource Indicator (URI), represented as an URL in this API. The NBD server runs outside of and isn’t controlled by the Virtualization framework. The NBD client forwards the guest’s I/O operations to the NBD server, which handles the I/O operations. The NBD client attempts to connect to the NBD server referred to by the URL used when you started the VM with . However, it’s important to note that a connection attempt isn’t made when the framework initializes the attachment object. Reconnection attempts take place throughout the life cycle of the VM when the NBD client encounters a recoverable error such as connection timeout and unexpected connection errors. The NBD client disconnects from the server when the VM shuts down. Using this attachment requires the app to have the entitlement because this attachment opens an outgoing network connection. To create a device that uses an NBD service, first initialize a with the URI of an NBD server, then use the attachment to configure a as shown in the example below (the attachment works with any subclass of , not just ): For more information about Network Block Devices, see the on GitHub. For more information about the NBD URL format, see the on GitHub.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment
type VZNetworkBlockDeviceStorageDeviceAttachment struct {
	VZStorageDeviceAttachment
}

// VZNetworkBlockDeviceStorageDeviceAttachmentFrom constructs a [VZNetworkBlockDeviceStorageDeviceAttachment] from an unsafe.Pointer.
//
// A storage device attachment backed by a Network Block Device (NBD) client.
func VZNetworkBlockDeviceStorageDeviceAttachmentFrom(ptr unsafe.Pointer) VZNetworkBlockDeviceStorageDeviceAttachment {
	return VZNetworkBlockDeviceStorageDeviceAttachment{
		VZStorageDeviceAttachment: VZStorageDeviceAttachmentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZNetworkBlockDeviceStorageDeviceAttachmentClass) Alloc() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZNetworkBlockDeviceStorageDeviceAttachmentClass) New() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) Init() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) Autorelease() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkBlockDeviceStorageDeviceAttachment creates a new VZNetworkBlockDeviceStorageDeviceAttachment instance.
func NewVZNetworkBlockDeviceStorageDeviceAttachment() VZNetworkBlockDeviceStorageDeviceAttachment {
	return getVZNetworkBlockDeviceStorageDeviceAttachmentClass().New()
}




// Creates a new network block device (NBD) storage attachment from an NDB Uniform Resource Indicator (URI) represented as a URL that you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/init(url:)
func NewVZNetworkBlockDeviceStorageDeviceAttachmentWithURLError(URL foundation.IURL, error_ unsafe.Pointer) VZNetworkBlockDeviceStorageDeviceAttachment {
	instance := getVZNetworkBlockDeviceStorageDeviceAttachmentClass().Alloc()
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](instance.ID, objc.Sel("initWithURL:error:"), URL, error_)
	rv.Autorelease()
	return rv
}



// Creates a new network block device storage attachment from an NBD Uniform Resource Indicator (URI) represented as a URL, timeout value, and read-only and synchronization modes that you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/init(url:timeout:isForcedReadOnly:synchronizationMode:)
func NewVZNetworkBlockDeviceStorageDeviceAttachmentWithURLTimeoutForcedReadOnlySynchronizationModeError(URL foundation.IURL, timeout foundation.ITimeInterval, forcedReadOnly bool, synchronizationMode VZDiskSynchronizationMode, error_ unsafe.Pointer) VZNetworkBlockDeviceStorageDeviceAttachment {
	instance := getVZNetworkBlockDeviceStorageDeviceAttachmentClass().Alloc()
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](instance.ID, objc.Sel("initWithURL:timeout:forcedReadOnly:synchronizationMode:error:"), URL, timeout, forcedReadOnly, synchronizationMode, error_)
	rv.Autorelease()
	return rv
}


// Checks if the URL is a valid network block device URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/validate(_:)
func (vc _VZNetworkBlockDeviceStorageDeviceAttachmentClass) ValidateURLError(URL foundation.IURL, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("validateURL:error:"), URL, error_)
	return rv
}

// The object that receives messages about changes to the network block device attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/delegate
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) Delegate() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The object that receives messages about changes to the network block device attachment.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/delegate
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDelegate:"), value)
}

// Returns a Boolean value that indicates whether the underlying disk attachment network is in a read-only state.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/isForcedReadOnly
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) ForcedReadOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("forcedReadOnly"))
	return rv
}

// Returns a Boolean value that indicates whether the underlying disk attachment network is in a read-only state.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkblockdevicestoragedeviceattachment/isforcedreadonly
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) IsForcedReadOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isForcedReadOnly"))
	return rv
}


// SetIsForcedReadOnly sets the value of the isForcedReadOnly property.
// Returns a Boolean value that indicates whether the underlying disk attachment network is in a read-only state.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkblockdevicestoragedeviceattachment/isforcedreadonly
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) SetIsForcedReadOnly(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsForcedReadOnly:"), value)
}

// The mode in which the NBD client synchronizes data with the NBD server.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkblockdevicestoragedeviceattachment/synchronizationmode
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) SynchronizationMode() VZDiskSynchronizationMode {
	rv := objc.Send[VZDiskSynchronizationMode](v_.ID, objc.Sel("synchronizationMode"))
	return rv
}


// SetSynchronizationMode sets the value of the synchronizationMode property.
// The mode in which the NBD client synchronizes data with the NBD server.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkblockdevicestoragedeviceattachment/synchronizationmode
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) SetSynchronizationMode(value VZDiskSynchronizationMode) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSynchronizationMode:"), value)
}

// The timeout value in seconds for the connection between the client and server.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkblockdevicestoragedeviceattachment/timeout
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) Timeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("timeout"))
	return rv
}


// SetTimeout sets the value of the timeout property.
// The timeout value in seconds for the connection between the client and server.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkblockdevicestoragedeviceattachment/timeout
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) SetTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTimeout:"), value)
}

// The URL that refers to the NBD server to which the NBD client will connect.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkblockdevicestoragedeviceattachment/url
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) Url() foundation.URL {
	rv := objc.Send[foundation.URL](v_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL that refers to the NBD server to which the NBD client will connect.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkblockdevicestoragedeviceattachment/url
func (v_ VZNetworkBlockDeviceStorageDeviceAttachment) SetUrl(value foundation.IURL) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUrl:"), value)
}


