// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVExternalStorageDevice */


/* debug [class_header]: Header for AVExternalStorageDevice */
// The class instance for the [ExternalStorageDevice] class.
var (
	ExternalStorageDeviceClass     _ExternalStorageDeviceClass
	ExternalStorageDeviceClassOnce sync.Once
)

func getExternalStorageDeviceClass() _ExternalStorageDeviceClass {
	ExternalStorageDeviceClassOnce.Do(func() {
		ExternalStorageDeviceClass = _ExternalStorageDeviceClass{objc.GetClass("AVExternalStorageDevice")}
	})
	return ExternalStorageDeviceClass
}

type _ExternalStorageDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ExternalStorageDevice */
// An interface definition for the [ExternalStorageDevice] class.
type IExternalStorageDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ExternalStorageDevice */
	// properties:
	DisplayName() objc.IObject /* cross-framework: NSString */
	FreeSize() int
	Connected() bool
	NotRecommendedForCaptureUse() bool
	TotalSize() int
	Uuid() foundation.UUID
	IsConnected() bool
	SetIsConnected(value bool)
	IsNotRecommendedForCaptureUse() bool
	SetIsNotRecommendedForCaptureUse(value bool)
	ExternalStorageDevices() IAVExternalStorageDevice
	SetExternalStorageDevices(value IAVExternalStorageDevice)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ExternalStorageDevice */
	// methods:
	NextAvailableURLsWithPathExtensionsError(extensionArray []string, outError objectivec.IObject) []foundation.URL
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ExternalStorageDevice */
// Alloc allocates a new instance without initialization.
func (ec _ExternalStorageDeviceClass) Alloc() ExternalStorageDevice {
	rv := objc.Send[ExternalStorageDevice](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _ExternalStorageDeviceClass) New() ExternalStorageDevice {
	rv := objc.Send[ExternalStorageDevice](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExternalStorageDevice) Init() ExternalStorageDevice {
	rv := objc.Send[ExternalStorageDevice](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExternalStorageDevice) Autorelease() ExternalStorageDevice {
	rv := objc.Send[ExternalStorageDevice](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExternalStorageDevice creates a new ExternalStorageDevice instance.
func NewExternalStorageDevice() ExternalStorageDevice {
	return getExternalStorageDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ExternalStorageDevice */
// Represents a physical external storage device that stores media assets.
//
// Each storage device instance corresponds to a physical external storage device where the system can media assets. You can access all of the currently available external storage devices with the object’s property.


// Represents a physical external storage device that stores media assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice
type ExternalStorageDevice struct {
	objectivec.Object
}

// ExternalStorageDeviceFrom constructs a [ExternalStorageDevice] from an unsafe.Pointer.
//
// Represents a physical external storage device that stores media assets.
func ExternalStorageDeviceFrom(ptr unsafe.Pointer) ExternalStorageDevice {
	return ExternalStorageDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ExternalStorageDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ExternalStorageDevice */

// Requests access to an external storage device on behalf of your app, which can present a dialog to a person on their device’s display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/requestAccess(completionHandler:)
func (ec _ExternalStorageDeviceClass) RequestAccessWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("requestAccessWithCompletionHandler:"), handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestAccessWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ExternalStorageDevice */

// Your app’s authorization status for the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/authorizationStatus
func (ec _ExternalStorageDeviceClass) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(ec.class), objc.Sel("authorizationStatus"))
	return rv
}/* debug [class_properties_class/property]: authorizationStatus */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ExternalStorageDevice */

// Generates an array of security scoped URLs that are compliant for digital camera formats, where each element has a different path extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/nextAvailableURLs(withPathExtensions:)
func (e_ ExternalStorageDevice) NextAvailableURLsWithPathExtensionsError(extensionArray []string, outError objectivec.IObject) []foundation.URL {
	rv := objc.Send[[]foundation.URL](e_.ID, objc.Sel("nextAvailableURLsWithPathExtensions:error:"), extensionArray, outError)
	return rv
}/* debug [instance_methods/method]: NextAvailableURLsWithPathExtensionsError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ExternalStorageDevice */

// Your app’s authorization status for the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/authorizationStatus
func (e_ ExternalStorageDevice) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](e_.ID, objc.Sel("authorizationStatus"))
	return rv
}/* debug [instance_properties/getter]: authorizationStatus */


// The name of an external storage device that’s appropriate for a user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/displayName
func (e_ ExternalStorageDevice) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// The amount of free storage space, in bytes, that’s available on the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/freeSize
func (e_ ExternalStorageDevice) FreeSize() int {
	rv := objc.Send[int](e_.ID, objc.Sel("freeSize"))
	return rv
}/* debug [instance_properties/getter]: freeSize */


// A Boolean value that indicates whether the system has a connection to the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/isConnected
func (e_ ExternalStorageDevice) Connected() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("connected"))
	return rv
}/* debug [instance_properties/getter]: connected */


// A Boolean value that indicates whether the external storage device is suitable for camera capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/isNotRecommendedForCaptureUse
func (e_ ExternalStorageDevice) NotRecommendedForCaptureUse() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("notRecommendedForCaptureUse"))
	return rv
}/* debug [instance_properties/getter]: notRecommendedForCaptureUse */


// The total amount of storage space, in bytes, that’s available on the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/totalSize
func (e_ ExternalStorageDevice) TotalSize() int {
	rv := objc.Send[int](e_.ID, objc.Sel("totalSize"))
	return rv
}/* debug [instance_properties/getter]: totalSize */


// The external storage device’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/uuid
func (e_ ExternalStorageDevice) Uuid() foundation.UUID {
	rv := objc.Send[foundation.UUID](e_.ID, objc.Sel("uuid"))
	return rv
}/* debug [instance_properties/getter]: uuid */


// A Boolean value that indicates whether the system has a connection to the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/isconnected
func (e_ ExternalStorageDevice) IsConnected() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isConnected"))
	return rv
}/* debug [instance_properties/getter]: isConnected */


// A Boolean value that indicates whether the system has a connection to the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/isconnected
func (e_ ExternalStorageDevice) SetIsConnected(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsConnected:"), value)
}/* debug [instance_properties/setter]: isConnected */


// A Boolean value that indicates whether the external storage device is suitable for camera capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/isnotrecommendedforcaptureuse
func (e_ ExternalStorageDevice) IsNotRecommendedForCaptureUse() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isNotRecommendedForCaptureUse"))
	return rv
}/* debug [instance_properties/getter]: isNotRecommendedForCaptureUse */


// A Boolean value that indicates whether the external storage device is suitable for camera capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/isnotrecommendedforcaptureuse
func (e_ ExternalStorageDevice) SetIsNotRecommendedForCaptureUse(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsNotRecommendedForCaptureUse:"), value)
}/* debug [instance_properties/setter]: isNotRecommendedForCaptureUse */


// An array of external storage devices the session updates as individual devices connect or disconnect from the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevicediscoverysession/externalstoragedevices
func (e_ ExternalStorageDevice) ExternalStorageDevices() IAVExternalStorageDevice {
	rv := objc.Send[ExternalStorageDevice](e_.ID, objc.Sel("externalStorageDevices"))
	return rv
}/* debug [instance_properties/getter]: externalStorageDevices */


// An array of external storage devices the session updates as individual devices connect or disconnect from the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevicediscoverysession/externalstoragedevices
func (e_ ExternalStorageDevice) SetExternalStorageDevices(value IAVExternalStorageDevice) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExternalStorageDevices:"), value)
}/* debug [instance_properties/setter]: externalStorageDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVExternalStorageDevice */



