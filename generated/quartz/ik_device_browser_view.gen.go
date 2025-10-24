// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class IKDeviceBrowserView */


/* debug [class_header]: Header for IKDeviceBrowserView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKDeviceBrowserView */
// An interface definition for the [IKDeviceBrowserView] class.
type IIKDeviceBrowserView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for IKDeviceBrowserView */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DisplaysLocalCameras() bool
	SetDisplaysLocalCameras(value bool)
	DisplaysLocalScanners() bool
	SetDisplaysLocalScanners(value bool)
	DisplaysNetworkCameras() bool
	SetDisplaysNetworkCameras(value bool)
	DisplaysNetworkScanners() bool
	SetDisplaysNetworkScanners(value bool)
	Mode() IKDeviceBrowserViewDisplayMode
	SetMode(value IKDeviceBrowserViewDisplayMode)
	SelectedDevice() objc.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKDeviceBrowserView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKDeviceBrowserView */
// Alloc allocates a new instance without initialization.
func (ic _IKDeviceBrowserViewClass) Alloc() IKDeviceBrowserView {
	rv := objc.Send[IKDeviceBrowserView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKDeviceBrowserView */
// The allows you to select a camera or scanner from a list of the available devices.
//
// The delegate must conform to the protocol. The delegate provides methods to inform you of selection changes in the browser as well as errors encountered when creating the browser list.


// The allows you to select a camera or scanner from a list of the available devices.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKDeviceBrowserView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKDeviceBrowserView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKDeviceBrowserView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKDeviceBrowserView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKDeviceBrowserView */

// Specifies the delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/delegate
func (i_ IKDeviceBrowserView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Specifies the delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/delegate
func (i_ IKDeviceBrowserView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Specifies whether local cameras are displayed by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/displaysLocalCameras
func (i_ IKDeviceBrowserView) DisplaysLocalCameras() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysLocalCameras"))
	return rv
}/* debug [instance_properties/getter]: displaysLocalCameras */


// Specifies whether local cameras are displayed by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/displaysLocalCameras
func (i_ IKDeviceBrowserView) SetDisplaysLocalCameras(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysLocalCameras:"), value)
}/* debug [instance_properties/setter]: displaysLocalCameras */


// Specifies whether local scanners are displayed by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/displaysLocalScanners
func (i_ IKDeviceBrowserView) DisplaysLocalScanners() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysLocalScanners"))
	return rv
}/* debug [instance_properties/getter]: displaysLocalScanners */


// Specifies whether local scanners are displayed by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/displaysLocalScanners
func (i_ IKDeviceBrowserView) SetDisplaysLocalScanners(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysLocalScanners:"), value)
}/* debug [instance_properties/setter]: displaysLocalScanners */


// Specifies whether network cameras are displayed by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/displaysNetworkCameras
func (i_ IKDeviceBrowserView) DisplaysNetworkCameras() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysNetworkCameras"))
	return rv
}/* debug [instance_properties/getter]: displaysNetworkCameras */


// Specifies whether network cameras are displayed by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/displaysNetworkCameras
func (i_ IKDeviceBrowserView) SetDisplaysNetworkCameras(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysNetworkCameras:"), value)
}/* debug [instance_properties/setter]: displaysNetworkCameras */


// Specifies whether network scanners are displayed by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/displaysNetworkScanners
func (i_ IKDeviceBrowserView) DisplaysNetworkScanners() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysNetworkScanners"))
	return rv
}/* debug [instance_properties/getter]: displaysNetworkScanners */


// Specifies whether network scanners are displayed by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/displaysNetworkScanners
func (i_ IKDeviceBrowserView) SetDisplaysNetworkScanners(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysNetworkScanners:"), value)
}/* debug [instance_properties/setter]: displaysNetworkScanners */


// Specifies the browser display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/mode
func (i_ IKDeviceBrowserView) Mode() IKDeviceBrowserViewDisplayMode {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// Specifies the browser display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/mode
func (i_ IKDeviceBrowserView) SetMode(value IKDeviceBrowserViewDisplayMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */


// Returns the selected device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserView/selectedDevice
func (i_ IKDeviceBrowserView) SelectedDevice() objc.IObject {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("selectedDevice"))
	return rv
}/* debug [instance_properties/getter]: selectedDevice */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKDeviceBrowserView */



