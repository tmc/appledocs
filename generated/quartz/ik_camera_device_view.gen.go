// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [IKCameraDeviceView] class.
var (
	IKCameraDeviceViewClass     _IKCameraDeviceViewClass
	IKCameraDeviceViewClassOnce sync.Once
)

func getIKCameraDeviceViewClass() _IKCameraDeviceViewClass {
	IKCameraDeviceViewClassOnce.Do(func() {
		IKCameraDeviceViewClass = _IKCameraDeviceViewClass{objc.GetClass("IKCameraDeviceView")}
	})
	return IKCameraDeviceViewClass
}

type _IKCameraDeviceViewClass struct {
	class objc.Class
}

// An interface definition for the [IKCameraDeviceView] class.
type IIKCameraDeviceView interface {
	appkit.IView
}

// The class displays the contents of the selected camera.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView
type IKCameraDeviceView struct {
	appkit.View
}

// IKCameraDeviceViewFrom constructs a [IKCameraDeviceView] from an unsafe.Pointer.
//
// The class displays the contents of the selected camera.
func IKCameraDeviceViewFrom(ptr unsafe.Pointer) IKCameraDeviceView {
	return IKCameraDeviceView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKCameraDeviceViewClass) Alloc() IKCameraDeviceView {
	rv := objc.Send[IKCameraDeviceView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKCameraDeviceViewClass) New() IKCameraDeviceView {
	rv := objc.Send[IKCameraDeviceView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKCameraDeviceView) Init() IKCameraDeviceView {
	rv := objc.Send[IKCameraDeviceView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKCameraDeviceView) Autorelease() IKCameraDeviceView {
	rv := objc.Send[IKCameraDeviceView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKCameraDeviceView creates a new IKCameraDeviceView instance.
func NewIKCameraDeviceView() IKCameraDeviceView {
	return getIKCameraDeviceViewClass().New()
}


// The current camera device.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/cameradevice
func (i_ IKCameraDeviceView) CameraDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("cameraDevice"))
	return rv
}


// SetCameraDevice sets the value of the cameraDevice property.
// The current camera device.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/cameradevice
func (i_ IKCameraDeviceView) SetCameraDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCameraDevice:"), value)
}

// Returns whether the selected items can be deleted.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/candeleteselecteditems
func (i_ IKCameraDeviceView) CanDeleteSelectedItems() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canDeleteSelectedItems"))
	return rv
}


// SetCanDeleteSelectedItems sets the value of the canDeleteSelectedItems property.
// Returns whether the selected items can be deleted.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/candeleteselecteditems
func (i_ IKCameraDeviceView) SetCanDeleteSelectedItems(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCanDeleteSelectedItems:"), value)
}

// Returns whether the selected items can be downloaded
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/candownloadselecteditems
func (i_ IKCameraDeviceView) CanDownloadSelectedItems() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canDownloadSelectedItems"))
	return rv
}


// SetCanDownloadSelectedItems sets the value of the canDownloadSelectedItems property.
// Returns whether the selected items can be downloaded

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/candownloadselecteditems
func (i_ IKCameraDeviceView) SetCanDownloadSelectedItems(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCanDownloadSelectedItems:"), value)
}

// Returns whether the selected items can be rotated left.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/canrotateselecteditemsleft
func (i_ IKCameraDeviceView) CanRotateSelectedItemsLeft() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canRotateSelectedItemsLeft"))
	return rv
}


// SetCanRotateSelectedItemsLeft sets the value of the canRotateSelectedItemsLeft property.
// Returns whether the selected items can be rotated left.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/canrotateselecteditemsleft
func (i_ IKCameraDeviceView) SetCanRotateSelectedItemsLeft(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCanRotateSelectedItemsLeft:"), value)
}

// Returns whether the selected items can be rotated right.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/canrotateselecteditemsright
func (i_ IKCameraDeviceView) CanRotateSelectedItemsRight() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canRotateSelectedItemsRight"))
	return rv
}


// SetCanRotateSelectedItemsRight sets the value of the canRotateSelectedItemsRight property.
// Returns whether the selected items can be rotated right.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/canrotateselecteditemsright
func (i_ IKCameraDeviceView) SetCanRotateSelectedItemsRight(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCanRotateSelectedItemsRight:"), value)
}

// The camera device view delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/delegate
func (i_ IKCameraDeviceView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The camera device view delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/delegate
func (i_ IKCameraDeviceView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}

// Specifies whether the downloads directory control should be displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/displaysdownloadsdirectorycontrol
func (i_ IKCameraDeviceView) DisplaysDownloadsDirectoryControl() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysDownloadsDirectoryControl"))
	return rv
}


// SetDisplaysDownloadsDirectoryControl sets the value of the displaysDownloadsDirectoryControl property.
// Specifies whether the downloads directory control should be displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/displaysdownloadsdirectorycontrol
func (i_ IKCameraDeviceView) SetDisplaysDownloadsDirectoryControl(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysDownloadsDirectoryControl:"), value)
}

// Displays whether the post process application control should be displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/displayspostprocessapplicationcontrol
func (i_ IKCameraDeviceView) DisplaysPostProcessApplicationControl() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysPostProcessApplicationControl"))
	return rv
}


// SetDisplaysPostProcessApplicationControl sets the value of the displaysPostProcessApplicationControl property.
// Displays whether the post process application control should be displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/displayspostprocessapplicationcontrol
func (i_ IKCameraDeviceView) SetDisplaysPostProcessApplicationControl(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysPostProcessApplicationControl:"), value)
}

// Allows the “Download All” control to be renamed.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/downloadallcontrollabel
func (i_ IKCameraDeviceView) DownloadAllControlLabel() string {
	rv := objc.Send[string](i_.ID, objc.Sel("downloadAllControlLabel"))
	return rv
}


// SetDownloadAllControlLabel sets the value of the downloadAllControlLabel property.
// Allows the “Download All” control to be renamed.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/downloadallcontrollabel
func (i_ IKCameraDeviceView) SetDownloadAllControlLabel(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDownloadAllControlLabel:"), objc.String(value))
}

// Allows the “Download Selected” control to be renamed.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/downloadselectedcontrollabel
func (i_ IKCameraDeviceView) DownloadSelectedControlLabel() string {
	rv := objc.Send[string](i_.ID, objc.Sel("downloadSelectedControlLabel"))
	return rv
}


// SetDownloadSelectedControlLabel sets the value of the downloadSelectedControlLabel property.
// Allows the “Download Selected” control to be renamed.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/downloadselectedcontrollabel
func (i_ IKCameraDeviceView) SetDownloadSelectedControlLabel(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDownloadSelectedControlLabel:"), objc.String(value))
}

// Specifies the directory where files are downloaded
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/downloadsdirectory
func (i_ IKCameraDeviceView) DownloadsDirectory() foundation.URL {
	rv := objc.Send[foundation.URL](i_.ID, objc.Sel("downloadsDirectory"))
	return rv
}


// SetDownloadsDirectory sets the value of the downloadsDirectory property.
// Specifies the directory where files are downloaded

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/downloadsdirectory
func (i_ IKCameraDeviceView) SetDownloadsDirectory(value foundation.URL) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDownloadsDirectory:"), value)
}

// Returns whether the device view is being displayed in icon mode.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/hasdisplaymodeicon
func (i_ IKCameraDeviceView) HasDisplayModeIcon() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasDisplayModeIcon"))
	return rv
}


// SetHasDisplayModeIcon sets the value of the hasDisplayModeIcon property.
// Returns whether the device view is being displayed in icon mode.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/hasdisplaymodeicon
func (i_ IKCameraDeviceView) SetHasDisplayModeIcon(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasDisplayModeIcon:"), value)
}

// Returns whether the device view is being displayed in table mode.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/hasdisplaymodetable
func (i_ IKCameraDeviceView) HasDisplayModeTable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasDisplayModeTable"))
	return rv
}


// SetHasDisplayModeTable sets the value of the hasDisplayModeTable property.
// Returns whether the device view is being displayed in table mode.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/hasdisplaymodetable
func (i_ IKCameraDeviceView) SetHasDisplayModeTable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasDisplayModeTable:"), value)
}

// Specifies the icon size.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/iconsize
func (i_ IKCameraDeviceView) IconSize() int {
	rv := objc.Send[int](i_.ID, objc.Sel("iconSize"))
	return rv
}


// SetIconSize sets the value of the iconSize property.
// Specifies the icon size.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/iconsize
func (i_ IKCameraDeviceView) SetIconSize(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIconSize:"), value)
}

// Specifies the display mode of the camera device view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/mode
func (i_ IKCameraDeviceView) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
// Specifies the display mode of the camera device view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/mode
func (i_ IKCameraDeviceView) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMode:"), value)
}

// The URL of the application used to post process the image.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/postprocessapplication
func (i_ IKCameraDeviceView) PostProcessApplication() foundation.URL {
	rv := objc.Send[foundation.URL](i_.ID, objc.Sel("postProcessApplication"))
	return rv
}


// SetPostProcessApplication sets the value of the postProcessApplication property.
// The URL of the application used to post process the image.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/postprocessapplication
func (i_ IKCameraDeviceView) SetPostProcessApplication(value foundation.URL) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPostProcessApplication:"), value)
}

// Determines how the contents are saved by the delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/transfermode
func (i_ IKCameraDeviceView) TransferMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("transferMode"))
	return rv
}


// SetTransferMode sets the value of the transferMode property.
// Determines how the contents are saved by the delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceview/transfermode
func (i_ IKCameraDeviceView) SetTransferMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransferMode:"), value)
}



