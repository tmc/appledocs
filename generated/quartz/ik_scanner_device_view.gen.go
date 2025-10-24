// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class IKScannerDeviceView */


/* debug [class_header]: Header for IKScannerDeviceView */
// The class instance for the [IKScannerDeviceView] class.
var (
	IKScannerDeviceViewClass     _IKScannerDeviceViewClass
	IKScannerDeviceViewClassOnce sync.Once
)

func getIKScannerDeviceViewClass() _IKScannerDeviceViewClass {
	IKScannerDeviceViewClassOnce.Do(func() {
		IKScannerDeviceViewClass = _IKScannerDeviceViewClass{objc.GetClass("IKScannerDeviceView")}
	})
	return IKScannerDeviceViewClass
}

type _IKScannerDeviceViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKScannerDeviceView */
// An interface definition for the [IKScannerDeviceView] class.
type IIKScannerDeviceView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for IKScannerDeviceView */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DisplaysDownloadsDirectoryControl() bool
	SetDisplaysDownloadsDirectoryControl(value bool)
	DisplaysPostProcessApplicationControl() bool
	SetDisplaysPostProcessApplicationControl(value bool)
	DocumentName() objc.IObject /* cross-framework: NSString */
	SetDocumentName(value objc.IObject /* cross-framework: NSString */)
	DownloadsDirectory() objc.IObject /* cross-framework: NSURL */
	SetDownloadsDirectory(value objc.IObject /* cross-framework: NSURL */)
	HasDisplayModeAdvanced() bool
	SetHasDisplayModeAdvanced(value bool)
	HasDisplayModeSimple() bool
	SetHasDisplayModeSimple(value bool)
	Mode() IKScannerDeviceViewDisplayMode
	SetMode(value IKScannerDeviceViewDisplayMode)
	OverviewControlLabel() objc.IObject /* cross-framework: NSString */
	SetOverviewControlLabel(value objc.IObject /* cross-framework: NSString */)
	PostProcessApplication() objc.IObject /* cross-framework: NSURL */
	SetPostProcessApplication(value objc.IObject /* cross-framework: NSURL */)
	ScanControlLabel() objc.IObject /* cross-framework: NSString */
	SetScanControlLabel(value objc.IObject /* cross-framework: NSString */)
	ScannerDevice() objc.IObject
	SetScannerDevice(value objc.IObject)
	TransferMode() IKScannerDeviceViewTransferMode
	SetTransferMode(value IKScannerDeviceViewTransferMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKScannerDeviceView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKScannerDeviceView */
// Alloc allocates a new instance without initialization.
func (ic _IKScannerDeviceViewClass) Alloc() IKScannerDeviceView {
	rv := objc.Send[IKScannerDeviceView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IKScannerDeviceViewClass) New() IKScannerDeviceView {
	rv := objc.Send[IKScannerDeviceView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKScannerDeviceView) Init() IKScannerDeviceView {
	rv := objc.Send[IKScannerDeviceView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKScannerDeviceView) Autorelease() IKScannerDeviceView {
	rv := objc.Send[IKScannerDeviceView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKScannerDeviceView creates a new IKScannerDeviceView instance.
func NewIKScannerDeviceView() IKScannerDeviceView {
	return getIKScannerDeviceViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKScannerDeviceView */
// The class displays a view that allows scanning. It can be customized by specifying the display mode. The delegate receives the scanned data and must implement the protocol.


// The class displays a view that allows scanning. It can be customized by specifying the display mode. The delegate receives the scanned data and must implement the protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView
type IKScannerDeviceView struct {
	appkit.View
}

// IKScannerDeviceViewFrom constructs a [IKScannerDeviceView] from an unsafe.Pointer.
//
// The class displays a view that allows scanning. It can be customized by specifying the display mode. The delegate receives the scanned data and must implement the protocol.
func IKScannerDeviceViewFrom(ptr unsafe.Pointer) IKScannerDeviceView {
	return IKScannerDeviceView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKScannerDeviceView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKScannerDeviceView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKScannerDeviceView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKScannerDeviceView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKScannerDeviceView */

// The scanner device delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/delegate
func (i_ IKScannerDeviceView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The scanner device delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/delegate
func (i_ IKScannerDeviceView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Determines whether the downloads directory control is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/displaysDownloadsDirectoryControl
func (i_ IKScannerDeviceView) DisplaysDownloadsDirectoryControl() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysDownloadsDirectoryControl"))
	return rv
}/* debug [instance_properties/getter]: displaysDownloadsDirectoryControl */


// Determines whether the downloads directory control is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/displaysDownloadsDirectoryControl
func (i_ IKScannerDeviceView) SetDisplaysDownloadsDirectoryControl(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysDownloadsDirectoryControl:"), value)
}/* debug [instance_properties/setter]: displaysDownloadsDirectoryControl */


// Specifies whether the post processing application control is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/displaysPostProcessApplicationControl
func (i_ IKScannerDeviceView) DisplaysPostProcessApplicationControl() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysPostProcessApplicationControl"))
	return rv
}/* debug [instance_properties/getter]: displaysPostProcessApplicationControl */


// Specifies whether the post processing application control is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/displaysPostProcessApplicationControl
func (i_ IKScannerDeviceView) SetDisplaysPostProcessApplicationControl(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysPostProcessApplicationControl:"), value)
}/* debug [instance_properties/setter]: displaysPostProcessApplicationControl */


// Returns the document name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/documentName
func (i_ IKScannerDeviceView) DocumentName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("documentName"))
	return rv
}/* debug [instance_properties/getter]: documentName */


// Returns the document name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/documentName
func (i_ IKScannerDeviceView) SetDocumentName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentName:"), value)
}/* debug [instance_properties/setter]: documentName */


// The directory where scans are saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/downloadsDirectory
func (i_ IKScannerDeviceView) DownloadsDirectory() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](i_.ID, objc.Sel("downloadsDirectory"))
	return rv
}/* debug [instance_properties/getter]: downloadsDirectory */


// The directory where scans are saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/downloadsDirectory
func (i_ IKScannerDeviceView) SetDownloadsDirectory(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDownloadsDirectory:"), value)
}/* debug [instance_properties/setter]: downloadsDirectory */


// The property that determines whether the scanner view uses the advanced display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/hasDisplayModeAdvanced
func (i_ IKScannerDeviceView) HasDisplayModeAdvanced() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasDisplayModeAdvanced"))
	return rv
}/* debug [instance_properties/getter]: hasDisplayModeAdvanced */


// The property that determines whether the scanner view uses the advanced display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/hasDisplayModeAdvanced
func (i_ IKScannerDeviceView) SetHasDisplayModeAdvanced(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasDisplayModeAdvanced:"), value)
}/* debug [instance_properties/setter]: hasDisplayModeAdvanced */


// The property that determines whether the scanner view uses the simple display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/hasDisplayModeSimple
func (i_ IKScannerDeviceView) HasDisplayModeSimple() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasDisplayModeSimple"))
	return rv
}/* debug [instance_properties/getter]: hasDisplayModeSimple */


// The property that determines whether the scanner view uses the simple display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/hasDisplayModeSimple
func (i_ IKScannerDeviceView) SetHasDisplayModeSimple(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasDisplayModeSimple:"), value)
}/* debug [instance_properties/setter]: hasDisplayModeSimple */


// The display mode used by the device view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/mode
func (i_ IKScannerDeviceView) Mode() IKScannerDeviceViewDisplayMode {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// The display mode used by the device view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/mode
func (i_ IKScannerDeviceView) SetMode(value IKScannerDeviceViewDisplayMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */


// Allows customization of the “Overview” label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/overviewControlLabel
func (i_ IKScannerDeviceView) OverviewControlLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("overviewControlLabel"))
	return rv
}/* debug [instance_properties/getter]: overviewControlLabel */


// Allows customization of the “Overview” label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/overviewControlLabel
func (i_ IKScannerDeviceView) SetOverviewControlLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOverviewControlLabel:"), value)
}/* debug [instance_properties/setter]: overviewControlLabel */


// The URL of the application to use for post processing of the scan.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/postProcessApplication
func (i_ IKScannerDeviceView) PostProcessApplication() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](i_.ID, objc.Sel("postProcessApplication"))
	return rv
}/* debug [instance_properties/getter]: postProcessApplication */


// The URL of the application to use for post processing of the scan.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/postProcessApplication
func (i_ IKScannerDeviceView) SetPostProcessApplication(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPostProcessApplication:"), value)
}/* debug [instance_properties/setter]: postProcessApplication */


// Allows customization of the “Scan” label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/scanControlLabel
func (i_ IKScannerDeviceView) ScanControlLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("scanControlLabel"))
	return rv
}/* debug [instance_properties/getter]: scanControlLabel */


// Allows customization of the “Scan” label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/scanControlLabel
func (i_ IKScannerDeviceView) SetScanControlLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScanControlLabel:"), value)
}/* debug [instance_properties/setter]: scanControlLabel */


// The device used for scanning
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/scannerDevice
func (i_ IKScannerDeviceView) ScannerDevice() objc.IObject {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("scannerDevice"))
	return rv
}/* debug [instance_properties/getter]: scannerDevice */


// The device used for scanning
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/scannerDevice
func (i_ IKScannerDeviceView) SetScannerDevice(value objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScannerDevice:"), value)
}/* debug [instance_properties/setter]: scannerDevice */


// Determines how the scanned content is provided to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/transferMode
func (i_ IKScannerDeviceView) TransferMode() IKScannerDeviceViewTransferMode {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("transferMode"))
	return rv
}/* debug [instance_properties/getter]: transferMode */


// Determines how the scanned content is provided to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceView/transferMode
func (i_ IKScannerDeviceView) SetTransferMode(value IKScannerDeviceViewTransferMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransferMode:"), value)
}/* debug [instance_properties/setter]: transferMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKScannerDeviceView */



