// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ICDevice */


/* debug [class_header]: Header for ICDevice */
// The class instance for the [ICDevice] class.
var (
	ICDeviceClass     _ICDeviceClass
	ICDeviceClassOnce sync.Once
)

func getICDeviceClass() _ICDeviceClass {
	ICDeviceClassOnce.Do(func() {
		ICDeviceClass = _ICDeviceClass{objc.GetClass("ICDevice")}
	})
	return ICDeviceClass
}

type _ICDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICDevice */
// An interface definition for the [ICDevice] class.
type IICDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ICDevice */
	// properties:
	Capabilities() unsafe.Pointer
	SetCapabilities(value unsafe.Pointer)
	UsbLocationID() unsafe.Pointer
	SetUsbLocationID(value unsafe.Pointer)
	HasOpenSession() unsafe.Pointer
	SetHasOpenSession(value unsafe.Pointer)
	UserData() foundation.MutableDictionary
	SetUserData(value foundation.MutableDictionary)
	UsbProductID() unsafe.Pointer
	SetUsbProductID(value unsafe.Pointer)
	LocationDescription() unsafe.Pointer
	SetLocationDescription(value unsafe.Pointer)
	UsbVendorID() unsafe.Pointer
	SetUsbVendorID(value unsafe.Pointer)
	AutolaunchApplicationPath() unsafe.Pointer
	SetAutolaunchApplicationPath(value unsafe.Pointer)
	Icon() Image get /* not a class type */
	SetIcon(value Image get /* not a class type */)
	TransportType() unsafe.Pointer
	SetTransportType(value unsafe.Pointer)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
	PersistentIDString() unsafe.Pointer
	SetPersistentIDString(value unsafe.Pointer)
	Name() unsafe.Pointer
	SetName(value unsafe.Pointer)
	UuidString() unsafe.Pointer
	SetUuidString(value unsafe.Pointer)
	IsRemote() unsafe.Pointer
	SetIsRemote(value unsafe.Pointer)
	SerialNumberString() unsafe.Pointer
	SetSerialNumberString(value unsafe.Pointer)
	ProductKind() unsafe.Pointer
	SetProductKind(value unsafe.Pointer)
	SystemSymbolName() unsafe.Pointer
	SetSystemSymbolName(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Remote() bool
	ModuleExecutableArchitecture() int
	ModulePath() objc.IObject /* cross-framework: NSString */
	ModuleVersion() objc.IObject /* cross-framework: NSString */
	UUIDString() objc.IObject /* cross-framework: NSString */
	BrowsedDeviceTypeMask() ICDeviceTypeMask
	SetBrowsedDeviceTypeMask(value ICDeviceTypeMask)
	Devices() ICDevice
	SetDevices(value ICDevice)
	IsBrowsing() bool
	SetIsBrowsing(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICDevice */
	// methods:
	RequestOpenSession()
	RequestCloseSession()
	RequestSendMessage()
	RequestCloseSessionWithOptionsCompletion(options objc.IObject, completion func(unsafe.Pointer))
	RequestEject()
	RequestEjectWithCompletion(completion func(unsafe.Pointer))
	RequestOpenSessionWithOptionsCompletion(options objc.IObject, completion func(unsafe.Pointer))
	RequestSendMessageOutDataMaxReturnedDataSizeSendMessageDelegateDidSendMessageSelectorContextInfo(messageCode unsafe.Pointer, data objc.IObject /* cross-framework: NSData */, maxReturnedDataSize unsafe.Pointer, sendMessageDelegate objc.IObject, selector objc.SEL, contextInfo unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICDevice */
// Alloc allocates a new instance without initialization.
func (ic _ICDeviceClass) Alloc() ICDevice {
	rv := objc.Send[ICDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICDeviceClass) New() ICDevice {
	rv := objc.Send[ICDevice](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICDevice) Init() ICDevice {
	rv := objc.Send[ICDevice](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICDevice) Autorelease() ICDevice {
	rv := objc.Send[ICDevice](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICDevice creates a new ICDevice instance.
func NewICDevice() ICDevice {
	return getICDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICDevice */
// An abstract object that represents a device.
//
// The device browser uses the concrete subclasses and to represent the cameras and scanners it finds.


// An abstract object that represents a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice
type ICDevice struct {
	objectivec.Object
}

// ICDeviceFrom constructs a [ICDevice] from an unsafe.Pointer.
//
// An abstract object that represents a device.
func ICDeviceFrom(ptr unsafe.Pointer) ICDevice {
	return ICDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICDevice */

// Requests to open a session on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507649-requestopensession
func (i_ ICDevice) RequestOpenSession() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestOpenSession"))
}/* debug [instance_methods/method]: RequestOpenSession */


// Requests to close an open session on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507882-requestclosesession
func (i_ ICDevice) RequestCloseSession() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestCloseSession"))
}/* debug [instance_methods/method]: RequestCloseSession */


// Asynchronously sends an arbitrary message with optional data to a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1508006-requestsendmessage
func (i_ ICDevice) RequestSendMessage() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSendMessage"))
}/* debug [instance_methods/method]: RequestSendMessage */


// Requests to close an open session on the device, then executes the completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/3142913-requestclosesessionwithoptions
func (i_ ICDevice) RequestCloseSessionWithOptionsCompletion(options objc.IObject, completion func(unsafe.Pointer)) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestCloseSessionWithOptions:completion:"), options, completion)
}/* debug [instance_methods/method]: RequestCloseSessionWithOptionsCompletion */


// Requests to eject the media if permitted by the device, or to disconnect from a remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/3142914-requesteject
func (i_ ICDevice) RequestEject() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestEject"))
}/* debug [instance_methods/method]: RequestEject */


// Requests to eject the media if permitted by the device, or to disconnect from a remote device, then executes the completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/3142915-requestejectwithcompletion
func (i_ ICDevice) RequestEjectWithCompletion(completion func(unsafe.Pointer)) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestEjectWithCompletion:"), completion)
}/* debug [instance_methods/method]: RequestEjectWithCompletion */


// Requests to open a session on the device, then executes the completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/3142916-requestopensessionwithoptions
func (i_ ICDevice) RequestOpenSessionWithOptionsCompletion(options objc.IObject, completion func(unsafe.Pointer)) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestOpenSessionWithOptions:completion:"), options, completion)
}/* debug [instance_methods/method]: RequestOpenSessionWithOptionsCompletion */


// Asynchronously sends an arbitrary message with optional data to a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/requestSendMessage(_:outData:maxReturnedDataSize:sendMessageDelegate:didSendMessageSelector:contextInfo:)
func (i_ ICDevice) RequestSendMessageOutDataMaxReturnedDataSizeSendMessageDelegateDidSendMessageSelectorContextInfo(messageCode unsafe.Pointer, data objc.IObject /* cross-framework: NSData */, maxReturnedDataSize unsafe.Pointer, sendMessageDelegate objc.IObject, selector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSendMessage:outData:maxReturnedDataSize:sendMessageDelegate:didSendMessageSelector:contextInfo:"), messageCode, data, maxReturnedDataSize, sendMessageDelegate, selector, contextInfo)
}/* debug [instance_methods/method]: RequestSendMessageOutDataMaxReturnedDataSizeSendMessageDelegateDidSendMessageSelectorContextInfo */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICDevice */

// The capabilities of the device as reported by the device module.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507594-capabilities
func (i_ ICDevice) Capabilities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("capabilities"))
	return rv
}/* debug [instance_properties/getter]: capabilities */


// The capabilities of the device as reported by the device module.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507594-capabilities
func (i_ ICDevice) SetCapabilities(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCapabilities:"), value)
}/* debug [instance_properties/setter]: capabilities */


// The USB location that the device is occupying.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507614-usblocationid
func (i_ ICDevice) UsbLocationID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("usbLocationID"))
	return rv
}/* debug [instance_properties/getter]: usbLocationID */


// The USB location that the device is occupying.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507614-usblocationid
func (i_ ICDevice) SetUsbLocationID(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUsbLocationID:"), value)
}/* debug [instance_properties/setter]: usbLocationID */


// A Boolean value that indicates whether the device has an open session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507615-hasopensession
func (i_ ICDevice) HasOpenSession() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("hasOpenSession"))
	return rv
}/* debug [instance_properties/getter]: hasOpenSession */


// A Boolean value that indicates whether the device has an open session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507615-hasopensession
func (i_ ICDevice) SetHasOpenSession(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasOpenSession:"), value)
}/* debug [instance_properties/setter]: hasOpenSession */


// A bookkeeping object for client convenience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507676-userdata
func (i_ ICDevice) UserData() foundation.MutableDictionary {
	rv := objc.Send[foundation.MutableDictionary](i_.ID, objc.Sel("userData"))
	return rv
}/* debug [instance_properties/getter]: userData */


// A bookkeeping object for client convenience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507676-userdata
func (i_ ICDevice) SetUserData(value foundation.MutableDictionary) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUserData:"), value)
}/* debug [instance_properties/setter]: userData */


// The USB Product ID (PID) associated with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507716-usbproductid
func (i_ ICDevice) UsbProductID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("usbProductID"))
	return rv
}/* debug [instance_properties/getter]: usbProductID */


// The USB Product ID (PID) associated with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507716-usbproductid
func (i_ ICDevice) SetUsbProductID(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUsbProductID:"), value)
}/* debug [instance_properties/setter]: usbProductID */


// A nonlocalized location description for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507717-locationdescription
func (i_ ICDevice) LocationDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("locationDescription"))
	return rv
}/* debug [instance_properties/getter]: locationDescription */


// A nonlocalized location description for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507717-locationdescription
func (i_ ICDevice) SetLocationDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLocationDescription:"), value)
}/* debug [instance_properties/setter]: locationDescription */


// The USB Vendor ID (VID) associated with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507787-usbvendorid
func (i_ ICDevice) UsbVendorID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("usbVendorID"))
	return rv
}/* debug [instance_properties/getter]: usbVendorID */


// The USB Vendor ID (VID) associated with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507787-usbvendorid
func (i_ ICDevice) SetUsbVendorID(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUsbVendorID:"), value)
}/* debug [instance_properties/setter]: usbVendorID */


// The file system path of an application to launch automatically when this device is added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507829-autolaunchapplicationpath
func (i_ ICDevice) AutolaunchApplicationPath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("autolaunchApplicationPath"))
	return rv
}/* debug [instance_properties/getter]: autolaunchApplicationPath */


// The file system path of an application to launch automatically when this device is added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507829-autolaunchapplicationpath
func (i_ ICDevice) SetAutolaunchApplicationPath(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAutolaunchApplicationPath:"), value)
}/* debug [instance_properties/setter]: autolaunchApplicationPath */


// The device’s icon image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507860-icon
func (i_ ICDevice) Icon() Image get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("icon"))
	return rv
}/* debug [instance_properties/getter]: icon */


// The device’s icon image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507860-icon
func (i_ ICDevice) SetIcon(value Image get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIcon:"), value)
}/* debug [instance_properties/setter]: icon */


// The hardware connection type the device is using.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507904-transporttype
func (i_ ICDevice) TransportType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("transportType"))
	return rv
}/* debug [instance_properties/getter]: transportType */


// The hardware connection type the device is using.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507904-transporttype
func (i_ ICDevice) SetTransportType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransportType:"), value)
}/* debug [instance_properties/setter]: transportType */


// A combination of the device’s type and its location type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507909-type
func (i_ ICDevice) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// A combination of the device’s type and its location type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507909-type
func (i_ ICDevice) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// A string representation of the device’s persistent ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507918-persistentidstring
func (i_ ICDevice) PersistentIDString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("persistentIDString"))
	return rv
}/* debug [instance_properties/getter]: persistentIDString */


// A string representation of the device’s persistent ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507918-persistentidstring
func (i_ ICDevice) SetPersistentIDString(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPersistentIDString:"), value)
}/* debug [instance_properties/setter]: persistentIDString */


// The device’s name as reported by the device module, or if no device module is in control of this device, by the device transport.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507989-name
func (i_ ICDevice) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The device’s name as reported by the device module, or if no device module is in control of this device, by the device transport.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507989-name
func (i_ ICDevice) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// A string representation of the device’s universally unique identifier (UUID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507999-uuidstring
func (i_ ICDevice) UuidString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("uuidString"))
	return rv
}/* debug [instance_properties/getter]: uuidString */


// A string representation of the device’s universally unique identifier (UUID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1507999-uuidstring
func (i_ ICDevice) SetUuidString(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUuidString:"), value)
}/* debug [instance_properties/setter]: uuidString */


// A Boolean value indicating whether the device is published by the Image Capture device-sharing facility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1508075-isremote
func (i_ ICDevice) IsRemote() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("isRemote"))
	return rv
}/* debug [instance_properties/getter]: isRemote */


// A Boolean value indicating whether the device is published by the Image Capture device-sharing facility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1508075-isremote
func (i_ ICDevice) SetIsRemote(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsRemote:"), value)
}/* debug [instance_properties/setter]: isRemote */


// The device’s serial number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1508137-serialnumberstring
func (i_ ICDevice) SerialNumberString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("serialNumberString"))
	return rv
}/* debug [instance_properties/getter]: serialNumberString */


// The device’s serial number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/1508137-serialnumberstring
func (i_ ICDevice) SetSerialNumberString(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSerialNumberString:"), value)
}/* debug [instance_properties/setter]: serialNumberString */


// The device’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/3142912-productkind
func (i_ ICDevice) ProductKind() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("productKind"))
	return rv
}/* debug [instance_properties/getter]: productKind */


// The device’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/3142912-productkind
func (i_ ICDevice) SetProductKind(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setProductKind:"), value)
}/* debug [instance_properties/setter]: productKind */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/3571394-systemsymbolname
func (i_ ICDevice) SystemSymbolName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("systemSymbolName"))
	return rv
}/* debug [instance_properties/getter]: systemSymbolName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevice/3571394-systemsymbolname
func (i_ ICDevice) SetSystemSymbolName(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSystemSymbolName:"), value)
}/* debug [instance_properties/setter]: systemSymbolName */


// The delegate to receive messages once a session is opened on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/delegate
func (i_ ICDevice) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate to receive messages once a session is opened on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/delegate
func (i_ ICDevice) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value indicating whether the device is published by the Image Capture device-sharing facility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/isRemote
func (i_ ICDevice) Remote() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("remote"))
	return rv
}/* debug [instance_properties/getter]: remote */


// The executable architecture of the device module servicing the requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/moduleExecutableArchitecture
func (i_ ICDevice) ModuleExecutableArchitecture() int {
	rv := objc.Send[int](i_.ID, objc.Sel("moduleExecutableArchitecture"))
	return rv
}/* debug [instance_properties/getter]: moduleExecutableArchitecture */


// The file system path of the device module associated with this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/modulePath
func (i_ ICDevice) ModulePath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("modulePath"))
	return rv
}/* debug [instance_properties/getter]: modulePath */


// The bundle version of the device module associated with this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/moduleVersion
func (i_ ICDevice) ModuleVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("moduleVersion"))
	return rv
}/* debug [instance_properties/getter]: moduleVersion */


// A string representation of the device’s universally unique identifier (UUID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDevice/uuidString
func (i_ ICDevice) UUIDString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("UUIDString"))
	return rv
}/* debug [instance_properties/getter]: UUIDString */


// A mask whose set bits indicate the type of devices being browsed after the delegate receives the start message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/browseddevicetypemask
func (i_ ICDevice) BrowsedDeviceTypeMask() ICDeviceTypeMask {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("browsedDeviceTypeMask"))
	return rv
}/* debug [instance_properties/getter]: browsedDeviceTypeMask */


// A mask whose set bits indicate the type of devices being browsed after the delegate receives the start message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/browseddevicetypemask
func (i_ ICDevice) SetBrowsedDeviceTypeMask(value ICDeviceTypeMask) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBrowsedDeviceTypeMask:"), value)
}/* debug [instance_properties/setter]: browsedDeviceTypeMask */


// All devices found by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/devices
func (i_ ICDevice) Devices() ICDevice {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("devices"))
	return rv
}/* debug [instance_properties/getter]: devices */


// All devices found by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/devices
func (i_ ICDevice) SetDevices(value ICDevice) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDevices:"), value)
}/* debug [instance_properties/setter]: devices */


// A Boolean value indicating whether the device browser is browsing for devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/isbrowsing
func (i_ ICDevice) IsBrowsing() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isBrowsing"))
	return rv
}/* debug [instance_properties/getter]: isBrowsing */


// A Boolean value indicating whether the device browser is browsing for devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/isbrowsing
func (i_ ICDevice) SetIsBrowsing(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsBrowsing:"), value)
}/* debug [instance_properties/setter]: isBrowsing */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICDevice */



