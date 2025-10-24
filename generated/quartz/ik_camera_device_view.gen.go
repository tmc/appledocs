// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IKCameraDeviceView */


/* debug [class_header]: Header for IKCameraDeviceView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKCameraDeviceView */
// An interface definition for the [IKCameraDeviceView] class.
type IIKCameraDeviceView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for IKCameraDeviceView */
	// properties:
	CameraDevice() objc.IObject
	SetCameraDevice(value objc.IObject)
	CanDeleteSelectedItems() bool
	CanDownloadSelectedItems() bool
	CanRotateSelectedItemsLeft() bool
	CanRotateSelectedItemsRight() bool
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DisplaysDownloadsDirectoryControl() bool
	SetDisplaysDownloadsDirectoryControl(value bool)
	DisplaysPostProcessApplicationControl() bool
	SetDisplaysPostProcessApplicationControl(value bool)
	DownloadAllControlLabel() objc.IObject /* cross-framework: NSString */
	SetDownloadAllControlLabel(value objc.IObject /* cross-framework: NSString */)
	DownloadsDirectory() objc.IObject /* cross-framework: NSURL */
	SetDownloadsDirectory(value objc.IObject /* cross-framework: NSURL */)
	DownloadSelectedControlLabel() objc.IObject /* cross-framework: NSString */
	SetDownloadSelectedControlLabel(value objc.IObject /* cross-framework: NSString */)
	HasDisplayModeIcon() bool
	SetHasDisplayModeIcon(value bool)
	HasDisplayModeTable() bool
	SetHasDisplayModeTable(value bool)
	IconSize() uint
	SetIconSize(value uint)
	Mode() IKCameraDeviceViewDisplayMode
	SetMode(value IKCameraDeviceViewDisplayMode)
	PostProcessApplication() objc.IObject /* cross-framework: NSURL */
	SetPostProcessApplication(value objc.IObject /* cross-framework: NSURL */)
	TransferMode() IKCameraDeviceViewTransferMode
	SetTransferMode(value IKCameraDeviceViewTransferMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKCameraDeviceView */
	// methods:
	DeleteSelectedItems(sender objc.IObject)
	DownloadAllItems(sender objc.IObject)
	DownloadSelectedItems(sender objc.IObject)
	RotateLeft(sender objc.IObject)
	RotateRight(sender objc.IObject)
	SelectIndexesByExtendingSelection(indexes foundation.IndexSet, extend bool)
	SelectedIndexes() foundation.IndexSet
	SetCustomActionControl(control appkit.SegmentedControl)
	SetCustomDeleteControl(control appkit.SegmentedControl)
	SetCustomIconSizeSlider(slider appkit.Slider)
	SetCustomModeControl(control appkit.SegmentedControl)
	SetCustomRotateControl(control appkit.SegmentedControl)
	SetShowStatusInfoAsWindowSubtitle(value bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKCameraDeviceView */
// Alloc allocates a new instance without initialization.
func (ic _IKCameraDeviceViewClass) Alloc() IKCameraDeviceView {
	rv := objc.Send[IKCameraDeviceView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKCameraDeviceView */
// The class displays the contents of the selected camera.


// The class displays the contents of the selected camera.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKCameraDeviceView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKCameraDeviceView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKCameraDeviceView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKCameraDeviceView */

// Deletes the currently selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/deleteSelectedItems(_:)
func (i_ IKCameraDeviceView) DeleteSelectedItems(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("deleteSelectedItems:"), sender)
}/* debug [instance_methods/method]: DeleteSelectedItems */


// Downloads all the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/downloadAllItems(_:)
func (i_ IKCameraDeviceView) DownloadAllItems(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("downloadAllItems:"), sender)
}/* debug [instance_methods/method]: DownloadAllItems */


// Deletes the selected items from the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/downloadSelectedItems(_:)
func (i_ IKCameraDeviceView) DownloadSelectedItems(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("downloadSelectedItems:"), sender)
}/* debug [instance_methods/method]: DownloadSelectedItems */


// Rotates the selected image to the left.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/rotateLeft(_:)
func (i_ IKCameraDeviceView) RotateLeft(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("rotateLeft:"), sender)
}/* debug [instance_methods/method]: RotateLeft */


// Rotates the selected image to the right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/rotateRight(_:)
func (i_ IKCameraDeviceView) RotateRight(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("rotateRight:"), sender)
}/* debug [instance_methods/method]: RotateRight */


// Invoked to select the specified files, extending the selection if specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/select(_:byExtendingSelection:)
func (i_ IKCameraDeviceView) SelectIndexesByExtendingSelection(indexes foundation.IndexSet, extend bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("selectIndexes:byExtendingSelection:"), indexes, extend)
}/* debug [instance_methods/method]: SelectIndexesByExtendingSelection */


// The selected indexes of the camera files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/selectedIndexes()
func (i_ IKCameraDeviceView) SelectedIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](i_.ID, objc.Sel("selectedIndexes"))
	return rv
}/* debug [instance_methods/method]: SelectedIndexes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/setCustomActionControl(_:)
func (i_ IKCameraDeviceView) SetCustomActionControl(control appkit.SegmentedControl) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCustomActionControl:"), control)
}/* debug [instance_methods/method]: SetCustomActionControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/setCustomDelete(_:)
func (i_ IKCameraDeviceView) SetCustomDeleteControl(control appkit.SegmentedControl) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCustomDeleteControl:"), control)
}/* debug [instance_methods/method]: SetCustomDeleteControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/setCustomIconSizeSlider(_:)
func (i_ IKCameraDeviceView) SetCustomIconSizeSlider(slider appkit.Slider) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCustomIconSizeSlider:"), slider)
}/* debug [instance_methods/method]: SetCustomIconSizeSlider */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/setCustomModeControl(_:)
func (i_ IKCameraDeviceView) SetCustomModeControl(control appkit.SegmentedControl) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCustomModeControl:"), control)
}/* debug [instance_methods/method]: SetCustomModeControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/setCustomRotateControl(_:)
func (i_ IKCameraDeviceView) SetCustomRotateControl(control appkit.SegmentedControl) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCustomRotateControl:"), control)
}/* debug [instance_methods/method]: SetCustomRotateControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/setShowStatusInfoAsWindowSubtitle(_:)
func (i_ IKCameraDeviceView) SetShowStatusInfoAsWindowSubtitle(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setShowStatusInfoAsWindowSubtitle:"), value)
}/* debug [instance_methods/method]: SetShowStatusInfoAsWindowSubtitle */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKCameraDeviceView */

// The current camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/cameraDevice
func (i_ IKCameraDeviceView) CameraDevice() objc.IObject {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("cameraDevice"))
	return rv
}/* debug [instance_properties/getter]: cameraDevice */


// The current camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/cameraDevice
func (i_ IKCameraDeviceView) SetCameraDevice(value objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCameraDevice:"), value)
}/* debug [instance_properties/setter]: cameraDevice */


// Returns whether the selected items can be deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/canDeleteSelectedItems
func (i_ IKCameraDeviceView) CanDeleteSelectedItems() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canDeleteSelectedItems"))
	return rv
}/* debug [instance_properties/getter]: canDeleteSelectedItems */


// Returns whether the selected items can be downloaded
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/canDownloadSelectedItems
func (i_ IKCameraDeviceView) CanDownloadSelectedItems() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canDownloadSelectedItems"))
	return rv
}/* debug [instance_properties/getter]: canDownloadSelectedItems */


// Returns whether the selected items can be rotated left.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/canRotateSelectedItemsLeft
func (i_ IKCameraDeviceView) CanRotateSelectedItemsLeft() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canRotateSelectedItemsLeft"))
	return rv
}/* debug [instance_properties/getter]: canRotateSelectedItemsLeft */


// Returns whether the selected items can be rotated right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/canRotateSelectedItemsRight
func (i_ IKCameraDeviceView) CanRotateSelectedItemsRight() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canRotateSelectedItemsRight"))
	return rv
}/* debug [instance_properties/getter]: canRotateSelectedItemsRight */


// The camera device view delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/delegate
func (i_ IKCameraDeviceView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The camera device view delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/delegate
func (i_ IKCameraDeviceView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Specifies whether the downloads directory control should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/displaysDownloadsDirectoryControl
func (i_ IKCameraDeviceView) DisplaysDownloadsDirectoryControl() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysDownloadsDirectoryControl"))
	return rv
}/* debug [instance_properties/getter]: displaysDownloadsDirectoryControl */


// Specifies whether the downloads directory control should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/displaysDownloadsDirectoryControl
func (i_ IKCameraDeviceView) SetDisplaysDownloadsDirectoryControl(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysDownloadsDirectoryControl:"), value)
}/* debug [instance_properties/setter]: displaysDownloadsDirectoryControl */


// Displays whether the post process application control should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/displaysPostProcessApplicationControl
func (i_ IKCameraDeviceView) DisplaysPostProcessApplicationControl() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("displaysPostProcessApplicationControl"))
	return rv
}/* debug [instance_properties/getter]: displaysPostProcessApplicationControl */


// Displays whether the post process application control should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/displaysPostProcessApplicationControl
func (i_ IKCameraDeviceView) SetDisplaysPostProcessApplicationControl(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplaysPostProcessApplicationControl:"), value)
}/* debug [instance_properties/setter]: displaysPostProcessApplicationControl */


// Allows the “Download All” control to be renamed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/downloadAllControlLabel
func (i_ IKCameraDeviceView) DownloadAllControlLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("downloadAllControlLabel"))
	return rv
}/* debug [instance_properties/getter]: downloadAllControlLabel */


// Allows the “Download All” control to be renamed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/downloadAllControlLabel
func (i_ IKCameraDeviceView) SetDownloadAllControlLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDownloadAllControlLabel:"), value)
}/* debug [instance_properties/setter]: downloadAllControlLabel */


// Specifies the directory where files are downloaded
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/downloadsDirectory
func (i_ IKCameraDeviceView) DownloadsDirectory() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](i_.ID, objc.Sel("downloadsDirectory"))
	return rv
}/* debug [instance_properties/getter]: downloadsDirectory */


// Specifies the directory where files are downloaded
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/downloadsDirectory
func (i_ IKCameraDeviceView) SetDownloadsDirectory(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDownloadsDirectory:"), value)
}/* debug [instance_properties/setter]: downloadsDirectory */


// Allows the “Download Selected” control to be renamed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/downloadSelectedControlLabel
func (i_ IKCameraDeviceView) DownloadSelectedControlLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("downloadSelectedControlLabel"))
	return rv
}/* debug [instance_properties/getter]: downloadSelectedControlLabel */


// Allows the “Download Selected” control to be renamed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/downloadSelectedControlLabel
func (i_ IKCameraDeviceView) SetDownloadSelectedControlLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDownloadSelectedControlLabel:"), value)
}/* debug [instance_properties/setter]: downloadSelectedControlLabel */


// Returns whether the device view is being displayed in icon mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/hasDisplayModeIcon
func (i_ IKCameraDeviceView) HasDisplayModeIcon() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasDisplayModeIcon"))
	return rv
}/* debug [instance_properties/getter]: hasDisplayModeIcon */


// Returns whether the device view is being displayed in icon mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/hasDisplayModeIcon
func (i_ IKCameraDeviceView) SetHasDisplayModeIcon(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasDisplayModeIcon:"), value)
}/* debug [instance_properties/setter]: hasDisplayModeIcon */


// Returns whether the device view is being displayed in table mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/hasDisplayModeTable
func (i_ IKCameraDeviceView) HasDisplayModeTable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasDisplayModeTable"))
	return rv
}/* debug [instance_properties/getter]: hasDisplayModeTable */


// Returns whether the device view is being displayed in table mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/hasDisplayModeTable
func (i_ IKCameraDeviceView) SetHasDisplayModeTable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasDisplayModeTable:"), value)
}/* debug [instance_properties/setter]: hasDisplayModeTable */


// Specifies the icon size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/iconSize
func (i_ IKCameraDeviceView) IconSize() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("iconSize"))
	return rv
}/* debug [instance_properties/getter]: iconSize */


// Specifies the icon size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/iconSize
func (i_ IKCameraDeviceView) SetIconSize(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIconSize:"), value)
}/* debug [instance_properties/setter]: iconSize */


// Specifies the display mode of the camera device view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/mode
func (i_ IKCameraDeviceView) Mode() IKCameraDeviceViewDisplayMode {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// Specifies the display mode of the camera device view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/mode
func (i_ IKCameraDeviceView) SetMode(value IKCameraDeviceViewDisplayMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */


// The URL of the application used to post process the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/postProcessApplication
func (i_ IKCameraDeviceView) PostProcessApplication() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](i_.ID, objc.Sel("postProcessApplication"))
	return rv
}/* debug [instance_properties/getter]: postProcessApplication */


// The URL of the application used to post process the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/postProcessApplication
func (i_ IKCameraDeviceView) SetPostProcessApplication(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPostProcessApplication:"), value)
}/* debug [instance_properties/setter]: postProcessApplication */


// Determines how the contents are saved by the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/transferMode
func (i_ IKCameraDeviceView) TransferMode() IKCameraDeviceViewTransferMode {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("transferMode"))
	return rv
}/* debug [instance_properties/getter]: transferMode */


// Determines how the contents are saved by the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceView/transferMode
func (i_ IKCameraDeviceView) SetTransferMode(value IKCameraDeviceViewTransferMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransferMode:"), value)
}/* debug [instance_properties/setter]: transferMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKCameraDeviceView */



