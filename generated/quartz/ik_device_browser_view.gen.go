// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/mlcompute"
)

// The class instance for the [IKDeviceBrowserView] class.
var (
	IKDeviceBrowserViewClass     _IKDeviceBrowserViewClass
	IKDeviceBrowserViewClassOnce sync.Once
)

func getIKDeviceBrowserViewClass() _IKDeviceBrowserViewClass {
	IKDeviceBrowserViewClassOnce.Do(func() {
		IKDeviceBrowserViewClass = _IKDeviceBrowserViewClass{objc.GetClass("IKDeviceBrowserView")}
	})
	return IKDeviceBrowserViewClass
}

type _IKDeviceBrowserViewClass struct {
	class objc.Class
}

// An interface definition for the [IKDeviceBrowserView] class.
type IIKDeviceBrowserView interface {
	appkit.IView
}

// The allows you to select a camera or scanner from a list of the available devices.
//
// The delegate must conform to the protocol. The delegate provides methods to inform you of selection changes in the browser as well as errors encountered when creating the browser list.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView
type IKDeviceBrowserView struct {
	appkit.View
}

// IKDeviceBrowserViewFrom constructs a [IKDeviceBrowserView] from an unsafe.Pointer.
//
// The allows you to select a camera or scanner from a list of the available devices.
func IKDeviceBrowserViewFrom(ptr unsafe.Pointer) IKDeviceBrowserView {
	return IKDeviceBrowserView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKDeviceBrowserViewClass) Alloc() IKDeviceBrowserView {
	rv := objc.Send[IKDeviceBrowserView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKDeviceBrowserViewClass) New() IKDeviceBrowserView {
	rv := objc.Send[IKDeviceBrowserView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKDeviceBrowserView) Init() IKDeviceBrowserView {
	rv := objc.Send[IKDeviceBrowserView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKDeviceBrowserView) Autorelease() IKDeviceBrowserView {
	rv := objc.Send[IKDeviceBrowserView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKDeviceBrowserView creates a new IKDeviceBrowserView instance.
func NewIKDeviceBrowserView() IKDeviceBrowserView {
	return getIKDeviceBrowserViewClass().New()
}


// Specifies the delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/delegate
func (i_ IKDeviceBrowserView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// Specifies the delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/delegate
func (i_ IKDeviceBrowserView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}

// Specifies whether local cameras are displayed by the browser.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/displayslocalcameras
func (i_ IKDeviceBrowserView) DisplaysLocalCameras() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysLocalCameras"))
	return rv
}


// SetDisplaysLocalCameras sets the value of the displaysLocalCameras property.
// Specifies whether local cameras are displayed by the browser.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/displayslocalcameras
func (i_ IKDeviceBrowserView) SetDisplaysLocalCameras(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysLocalCameras:"), value)
}

// Specifies whether local scanners are displayed by the browser.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/displayslocalscanners
func (i_ IKDeviceBrowserView) DisplaysLocalScanners() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysLocalScanners"))
	return rv
}


// SetDisplaysLocalScanners sets the value of the displaysLocalScanners property.
// Specifies whether local scanners are displayed by the browser.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/displayslocalscanners
func (i_ IKDeviceBrowserView) SetDisplaysLocalScanners(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysLocalScanners:"), value)
}

// Specifies whether network cameras are displayed by the browser.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/displaysnetworkcameras
func (i_ IKDeviceBrowserView) DisplaysNetworkCameras() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysNetworkCameras"))
	return rv
}


// SetDisplaysNetworkCameras sets the value of the displaysNetworkCameras property.
// Specifies whether network cameras are displayed by the browser.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/displaysnetworkcameras
func (i_ IKDeviceBrowserView) SetDisplaysNetworkCameras(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysNetworkCameras:"), value)
}

// Specifies whether network scanners are displayed by the browser.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/displaysnetworkscanners
func (i_ IKDeviceBrowserView) DisplaysNetworkScanners() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysNetworkScanners"))
	return rv
}


// SetDisplaysNetworkScanners sets the value of the displaysNetworkScanners property.
// Specifies whether network scanners are displayed by the browser.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/displaysnetworkscanners
func (i_ IKDeviceBrowserView) SetDisplaysNetworkScanners(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysNetworkScanners:"), value)
}

// Specifies the browser display mode.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/mode
func (i_ IKDeviceBrowserView) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
// Specifies the browser display mode.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/mode
func (i_ IKDeviceBrowserView) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMode:"), value)
}

// Returns the selected device.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/selecteddevice
func (i_ IKDeviceBrowserView) SelectedDevice() mlcompute.ICDevice {
	rv := objc.Send[mlcompute.ICDevice](i_.ID, objc.Sel("selectedDevice"))
	return rv
}


// SetSelectedDevice sets the value of the selectedDevice property.
// Returns the selected device.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/selecteddevice
func (i_ IKDeviceBrowserView) SetSelectedDevice(value mlcompute.ICDevice) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSelectedDevice:"), value)
}



