// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/mlcompute"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ICDeviceBrowser] class.
type IICDeviceBrowser interface {
	objectivec.IObject
	// properties:
	BrowsedDeviceTypeMask() unsafe.Pointer
	SetBrowsedDeviceTypeMask(value unsafe.Pointer)
	ContentsAuthorizationStatus() ICAuthorizationStatus /* already interface */
	SetContentsAuthorizationStatus(value ICAuthorizationStatus /* already interface */)
	ControlAuthorizationStatus() ICAuthorizationStatus /* already interface */
	SetControlAuthorizationStatus(value ICAuthorizationStatus /* already interface */)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Devices() mlcompute.ICDevice /* already interface */
	SetDevices(value mlcompute.ICDevice /* already interface */)
	IsBrowsing() bool /* primitive/slice/pointer. */
	SetIsBrowsing(value bool /* primitive/slice/pointer. */)
	IsSuspended() bool /* primitive/slice/pointer. */
	SetIsSuspended(value bool /* primitive/slice/pointer. */)
	PreferredDevice() mlcompute.ICDevice /* already interface */
	SetPreferredDevice(value mlcompute.ICDevice /* already interface */)
	// methods:
	RequestControlAuthorizationWithCompletion(completion unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (ic _ICDeviceBrowserClass) Alloc() ICDeviceBrowser {
	rv := objc.Send[ICDeviceBrowser](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/requestControlAuthorization(completion:)
func (i_ ICDeviceBrowser) RequestControlAuthorizationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestControlAuthorizationWithCompletion:"), completion)
}


// A mask whose set bits indicate the type of devices being browsed after the delegate receives the start message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/browsedDeviceTypeMask
func (i_ ICDeviceBrowser) BrowsedDeviceTypeMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("browsedDeviceTypeMask"))
	return rv
}


// A mask whose set bits indicate the type of devices being browsed after the delegate receives the start message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/browsedDeviceTypeMask
func (i_ ICDeviceBrowser) SetBrowsedDeviceTypeMask(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBrowsedDeviceTypeMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/contentsauthorizationstatus
func (i_ ICDeviceBrowser) ContentsAuthorizationStatus() ICAuthorizationStatus /* already interface */ {
	rv := objc.Send[ICAuthorizationStatus](i_.ID, objc.Sel("contentsAuthorizationStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/contentsauthorizationstatus
func (i_ ICDeviceBrowser) SetContentsAuthorizationStatus(value ICAuthorizationStatus /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContentsAuthorizationStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/controlauthorizationstatus
func (i_ ICDeviceBrowser) ControlAuthorizationStatus() ICAuthorizationStatus /* already interface */ {
	rv := objc.Send[ICAuthorizationStatus](i_.ID, objc.Sel("controlAuthorizationStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/controlauthorizationstatus
func (i_ ICDeviceBrowser) SetControlAuthorizationStatus(value ICAuthorizationStatus /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setControlAuthorizationStatus:"), value)
}


// The object that acts as the delegate of the device browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/delegate
func (i_ ICDeviceBrowser) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}


// The object that acts as the delegate of the device browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/delegate
func (i_ ICDeviceBrowser) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}


// All devices found by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/devices
func (i_ ICDeviceBrowser) Devices() mlcompute.ICDevice /* already interface */ {
	rv := objc.Send[mlcompute.ICDevice](i_.ID, objc.Sel("devices"))
	return rv
}


// All devices found by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/devices
func (i_ ICDeviceBrowser) SetDevices(value mlcompute.ICDevice /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDevices:"), value)
}


// A Boolean value indicating whether the device browser is browsing for devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/isbrowsing
func (i_ ICDeviceBrowser) IsBrowsing() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isBrowsing"))
	return rv
}


// A Boolean value indicating whether the device browser is browsing for devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/isbrowsing
func (i_ ICDeviceBrowser) SetIsBrowsing(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsBrowsing:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/issuspended
func (i_ ICDeviceBrowser) IsSuspended() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isSuspended"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/issuspended
func (i_ ICDeviceBrowser) SetIsSuspended(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsSuspended:"), value)
}


// Returns a device object that the client application should select when it launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/preferreddevice
func (i_ ICDeviceBrowser) PreferredDevice() mlcompute.ICDevice /* already interface */ {
	rv := objc.Send[mlcompute.ICDevice](i_.ID, objc.Sel("preferredDevice"))
	return rv
}


// Returns a device object that the client application should select when it launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicebrowser/preferreddevice
func (i_ ICDeviceBrowser) SetPreferredDevice(value mlcompute.ICDevice /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredDevice:"), value)
}



