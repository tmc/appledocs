// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/imagecapturecore"
)

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

// An interface definition for the [IKScannerDeviceView] class.
type IIKScannerDeviceView interface {
	appkit.IView
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DisplaysDownloadsDirectoryControl() bool
	SetDisplaysDownloadsDirectoryControl(value bool)
	DisplaysPostProcessApplicationControl() bool
	SetDisplaysPostProcessApplicationControl(value bool)
	DocumentName() objc.IObject /* cross-framework: NSString */
	SetDocumentName(value objc.IObject /* cross-framework: NSString */)
	DownloadsDirectory() objc.IObject /* cross-framework: URL */
	SetDownloadsDirectory(value objc.IObject /* cross-framework: URL */)
	HasDisplayModeAdvanced() bool
	SetHasDisplayModeAdvanced(value bool)
	HasDisplayModeSimple() bool
	SetHasDisplayModeSimple(value bool)
	Mode() unsafe.Pointer
	SetMode(value unsafe.Pointer)
	OverviewControlLabel() objc.IObject /* cross-framework: NSString */
	SetOverviewControlLabel(value objc.IObject /* cross-framework: NSString */)
	PostProcessApplication() objc.IObject /* cross-framework: URL */
	SetPostProcessApplication(value objc.IObject /* cross-framework: URL */)
	ScanControlLabel() objc.IObject /* cross-framework: NSString */
	SetScanControlLabel(value objc.IObject /* cross-framework: NSString */)
	ScannerDevice() imagecapturecore.ICScannerDevice
	SetScannerDevice(value imagecapturecore.ICScannerDevice)
	TransferMode() unsafe.Pointer
	SetTransferMode(value unsafe.Pointer)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ic _IKScannerDeviceViewClass) Alloc() IKScannerDeviceView {
	rv := objc.Send[IKScannerDeviceView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The scanner device delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/delegate
func (i_ IKScannerDeviceView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}


// The scanner device delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/delegate
func (i_ IKScannerDeviceView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}


// Determines whether the downloads directory control is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/displaysdownloadsdirectorycontrol
func (i_ IKScannerDeviceView) DisplaysDownloadsDirectoryControl() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysDownloadsDirectoryControl"))
	return rv
}


// Determines whether the downloads directory control is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/displaysdownloadsdirectorycontrol
func (i_ IKScannerDeviceView) SetDisplaysDownloadsDirectoryControl(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysDownloadsDirectoryControl:"), value)
}


// Specifies whether the post processing application control is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/displayspostprocessapplicationcontrol
func (i_ IKScannerDeviceView) DisplaysPostProcessApplicationControl() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysPostProcessApplicationControl"))
	return rv
}


// Specifies whether the post processing application control is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/displayspostprocessapplicationcontrol
func (i_ IKScannerDeviceView) SetDisplaysPostProcessApplicationControl(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysPostProcessApplicationControl:"), value)
}


// Returns the document name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/documentname
func (i_ IKScannerDeviceView) DocumentName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("documentName"))
	return rv
}


// Returns the document name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/documentname
func (i_ IKScannerDeviceView) SetDocumentName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentName:"), value)
}


// The directory where scans are saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/downloadsdirectory
func (i_ IKScannerDeviceView) DownloadsDirectory() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](i_.ID, objc.Sel("downloadsDirectory"))
	return rv
}


// The directory where scans are saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/downloadsdirectory
func (i_ IKScannerDeviceView) SetDownloadsDirectory(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDownloadsDirectory:"), value)
}


// The property that determines whether the scanner view uses the advanced display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/hasdisplaymodeadvanced
func (i_ IKScannerDeviceView) HasDisplayModeAdvanced() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasDisplayModeAdvanced"))
	return rv
}


// The property that determines whether the scanner view uses the advanced display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/hasdisplaymodeadvanced
func (i_ IKScannerDeviceView) SetHasDisplayModeAdvanced(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasDisplayModeAdvanced:"), value)
}


// The property that determines whether the scanner view uses the simple display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/hasdisplaymodesimple
func (i_ IKScannerDeviceView) HasDisplayModeSimple() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasDisplayModeSimple"))
	return rv
}


// The property that determines whether the scanner view uses the simple display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/hasdisplaymodesimple
func (i_ IKScannerDeviceView) SetHasDisplayModeSimple(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasDisplayModeSimple:"), value)
}


// The display mode used by the device view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/mode
func (i_ IKScannerDeviceView) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mode"))
	return rv
}


// The display mode used by the device view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/mode
func (i_ IKScannerDeviceView) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMode:"), value)
}


// Allows customization of the “Overview” label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/overviewcontrollabel
func (i_ IKScannerDeviceView) OverviewControlLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("overviewControlLabel"))
	return rv
}


// Allows customization of the “Overview” label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/overviewcontrollabel
func (i_ IKScannerDeviceView) SetOverviewControlLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOverviewControlLabel:"), value)
}


// The URL of the application to use for post processing of the scan.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/postprocessapplication
func (i_ IKScannerDeviceView) PostProcessApplication() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](i_.ID, objc.Sel("postProcessApplication"))
	return rv
}


// The URL of the application to use for post processing of the scan.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/postprocessapplication
func (i_ IKScannerDeviceView) SetPostProcessApplication(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPostProcessApplication:"), value)
}


// Allows customization of the “Scan” label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/scancontrollabel
func (i_ IKScannerDeviceView) ScanControlLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("scanControlLabel"))
	return rv
}


// Allows customization of the “Scan” label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/scancontrollabel
func (i_ IKScannerDeviceView) SetScanControlLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScanControlLabel:"), value)
}


// The device used for scanning
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/scannerdevice
func (i_ IKScannerDeviceView) ScannerDevice() imagecapturecore.ICScannerDevice {
	rv := objc.Send[imagecapturecore.ICScannerDevice](i_.ID, objc.Sel("scannerDevice"))
	return rv
}


// The device used for scanning
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/scannerdevice
func (i_ IKScannerDeviceView) SetScannerDevice(value imagecapturecore.ICScannerDevice) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScannerDevice:"), value)
}


// Determines how the scanned content is provided to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/transfermode
func (i_ IKScannerDeviceView) TransferMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("transferMode"))
	return rv
}


// Determines how the scanned content is provided to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/transfermode
func (i_ IKScannerDeviceView) SetTransferMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransferMode:"), value)
}



