// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ICDeviceBrowser */


/* debug [class_header]: Header for ICDeviceBrowser */
// The class instance for the [ICDeviceBrowser] class.
var (
	ICDeviceBrowserClass     _ICDeviceBrowserClass
	ICDeviceBrowserClassOnce sync.Once
)

func getICDeviceBrowserClass() _ICDeviceBrowserClass {
	ICDeviceBrowserClassOnce.Do(func() {
		ICDeviceBrowserClass = _ICDeviceBrowserClass{objc.GetClass("ICDeviceBrowser")}
	})
	return ICDeviceBrowserClass
}

type _ICDeviceBrowserClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICDeviceBrowser */
// An interface definition for the [ICDeviceBrowser] class.
type IICDeviceBrowser interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ICDeviceBrowser */
	// properties:
	BrowsedDeviceTypeMask() unsafe.Pointer
	SetBrowsedDeviceTypeMask(value unsafe.Pointer)
	IsBrowsing() unsafe.Pointer
	SetIsBrowsing(value unsafe.Pointer)
	Devices() ICDevice
	SetDevices(value ICDevice)
	PreferredDevice() ICDevice
	SetPreferredDevice(value ICDevice)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Browsing() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICDeviceBrowser */
	// methods:
	Stop()
	Start()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICDeviceBrowser */
// Alloc allocates a new instance without initialization.
func (ic _ICDeviceBrowserClass) Alloc() ICDeviceBrowser {
	rv := objc.Send[ICDeviceBrowser](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICDeviceBrowserClass) New() ICDeviceBrowser {
	rv := objc.Send[ICDeviceBrowser](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICDeviceBrowser) Init() ICDeviceBrowser {
	rv := objc.Send[ICDeviceBrowser](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICDeviceBrowser) Autorelease() ICDeviceBrowser {
	rv := objc.Send[ICDeviceBrowser](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICDeviceBrowser creates a new ICDeviceBrowser instance.
func NewICDeviceBrowser() ICDeviceBrowser {
	return getICDeviceBrowserClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICDeviceBrowser */
// An object for finding digital cameras and scanners.


// An object for finding digital cameras and scanners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser
type ICDeviceBrowser struct {
	objectivec.Object
}

// ICDeviceBrowserFrom constructs a [ICDeviceBrowser] from an unsafe.Pointer.
//
// An object for finding digital cameras and scanners.
func ICDeviceBrowserFrom(ptr unsafe.Pointer) ICDeviceBrowser {
	return ICDeviceBrowser{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICDeviceBrowser */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICDeviceBrowser */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICDeviceBrowser */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICDeviceBrowser */

// Tells the delegate to stop looking for devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/1508085-stop
func (i_ ICDeviceBrowser) Stop() {
	objc.Send[objc.ID](i_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */


// Tells the delegate to start looking for devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/1508087-start
func (i_ ICDeviceBrowser) Start() {
	objc.Send[objc.ID](i_.ID, objc.Sel("start"))
}/* debug [instance_methods/method]: Start */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICDeviceBrowser */

// A mask whose set bits indicate the type of devices being browsed after the delegate receives the start message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/1507613-browseddevicetypemask
func (i_ ICDeviceBrowser) BrowsedDeviceTypeMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("browsedDeviceTypeMask"))
	return rv
}/* debug [instance_properties/getter]: browsedDeviceTypeMask */


// A mask whose set bits indicate the type of devices being browsed after the delegate receives the start message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/1507613-browseddevicetypemask
func (i_ ICDeviceBrowser) SetBrowsedDeviceTypeMask(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBrowsedDeviceTypeMask:"), value)
}/* debug [instance_properties/setter]: browsedDeviceTypeMask */


// A Boolean value indicating whether the device browser is browsing for devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/1507627-isbrowsing
func (i_ ICDeviceBrowser) IsBrowsing() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("isBrowsing"))
	return rv
}/* debug [instance_properties/getter]: isBrowsing */


// A Boolean value indicating whether the device browser is browsing for devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/1507627-isbrowsing
func (i_ ICDeviceBrowser) SetIsBrowsing(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsBrowsing:"), value)
}/* debug [instance_properties/setter]: isBrowsing */


// All devices found by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/1507765-devices
func (i_ ICDeviceBrowser) Devices() ICDevice {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("devices"))
	return rv
}/* debug [instance_properties/getter]: devices */


// All devices found by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/1507765-devices
func (i_ ICDeviceBrowser) SetDevices(value ICDevice) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDevices:"), value)
}/* debug [instance_properties/setter]: devices */


// Returns a device object that the client application should select when it launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/1507781-preferreddevice
func (i_ ICDeviceBrowser) PreferredDevice() ICDevice {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("preferredDevice"))
	return rv
}/* debug [instance_properties/getter]: preferredDevice */


// Returns a device object that the client application should select when it launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/1507781-preferreddevice
func (i_ ICDeviceBrowser) SetPreferredDevice(value ICDevice) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredDevice:"), value)
}/* debug [instance_properties/setter]: preferredDevice */


// The object that acts as the delegate of the device browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/delegate
func (i_ ICDeviceBrowser) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The object that acts as the delegate of the device browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/delegate
func (i_ ICDeviceBrowser) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value indicating whether the device browser is browsing for devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/isBrowsing
func (i_ ICDeviceBrowser) Browsing() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("browsing"))
	return rv
}/* debug [instance_properties/getter]: browsing */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICDeviceBrowser */


