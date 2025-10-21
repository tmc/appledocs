// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	Start()
}

// An object for finding digital cameras and scanners.
//
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


// Tells the delegate to start looking for devices.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/start()
func (i_ ICDeviceBrowser) Start() {
	objc.Send[objc.ID](i_.ID, objc.Sel("start"))
}

// A mask whose set bits indicate the type of devices being browsed after the delegate receives the start message.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/browsedDeviceTypeMask
func (i_ ICDeviceBrowser) BrowsedDeviceTypeMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("browsedDeviceTypeMask"))
	return rv
}


// SetBrowsedDeviceTypeMask sets the value of the browsedDeviceTypeMask property.
// A mask whose set bits indicate the type of devices being browsed after the delegate receives the start message.

//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/browsedDeviceTypeMask
func (i_ ICDeviceBrowser) SetBrowsedDeviceTypeMask(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBrowsedDeviceTypeMask:"), value)
}

// The object that acts as the delegate of the device browser.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/delegate
func (i_ ICDeviceBrowser) Delegate() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The object that acts as the delegate of the device browser.

//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/delegate
func (i_ ICDeviceBrowser) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}



