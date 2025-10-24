// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/iobluetooth"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothObjectPushUIController */


/* debug [class_header]: Header for IOBluetoothObjectPushUIController */
// The class instance for the [BluetoothObjectPushUIController] class.
var (
	BluetoothObjectPushUIControllerClass     _BluetoothObjectPushUIControllerClass
	BluetoothObjectPushUIControllerClassOnce sync.Once
)

func getBluetoothObjectPushUIControllerClass() _BluetoothObjectPushUIControllerClass {
	BluetoothObjectPushUIControllerClassOnce.Do(func() {
		BluetoothObjectPushUIControllerClass = _BluetoothObjectPushUIControllerClass{objc.GetClass("IOBluetoothObjectPushUIController")}
	})
	return BluetoothObjectPushUIControllerClass
}

type _BluetoothObjectPushUIControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothObjectPushUIController */
// An interface definition for the [BluetoothObjectPushUIController] class.
type IBluetoothObjectPushUIController interface {
	appkit.IWindowController
	
/* debug [class_interface_properties]: Properties for BluetoothObjectPushUIController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothObjectPushUIController */
	// methods:
	BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow appkit.Window, modalDelegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) int
	GetDevice() iobluetooth.BluetoothDevice
	GetTitle() foundation.String
	IsTransferInProgress() bool
	RunModal()
	RunPanel()
	SetIconImage(image appkit.Image)
	SetTitle(windowTitle objc.IObject /* cross-framework: NSString */)
	Stop()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothObjectPushUIController */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothObjectPushUIControllerClass) Alloc() BluetoothObjectPushUIController {
	rv := objc.Send[BluetoothObjectPushUIController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BluetoothObjectPushUIControllerClass) New() BluetoothObjectPushUIController {
	rv := objc.Send[BluetoothObjectPushUIController](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothObjectPushUIController) Init() BluetoothObjectPushUIController {
	rv := objc.Send[BluetoothObjectPushUIController](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothObjectPushUIController) Autorelease() BluetoothObjectPushUIController {
	rv := objc.Send[BluetoothObjectPushUIController](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothObjectPushUIController creates a new BluetoothObjectPushUIController instance.
func NewBluetoothObjectPushUIController() BluetoothObjectPushUIController {
	return getBluetoothObjectPushUIControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothObjectPushUIController */
// An NSWindowController subclass that supports the creation of an IOBluetoothObjectPushUIController object.


// An NSWindowController subclass that supports the creation of an IOBluetoothObjectPushUIController object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController
type BluetoothObjectPushUIController struct {
	appkit.WindowController
}

// BluetoothObjectPushUIControllerFrom constructs a [BluetoothObjectPushUIController] from an unsafe.Pointer.
//
// An NSWindowController subclass that supports the creation of an IOBluetoothObjectPushUIController object.
func BluetoothObjectPushUIControllerFrom(ptr unsafe.Pointer) BluetoothObjectPushUIController {
	return BluetoothObjectPushUIController{
		WindowController: appkit.WindowControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothObjectPushUIController */

// Creates and returns a new IOBluetoothObjectPush object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/init(objectPushWith:withFiles:delegate:)
func NewBluetoothObjectPushUIControllerObjectPushWithBluetoothDeviceWithFilesDelegate(inDevice iobluetooth.BluetoothDevice, inFiles objc.IObject /* cross-framework: NSArray */, inDelegate objc.IObject) BluetoothObjectPushUIController {
	instance := getBluetoothObjectPushUIControllerClass().Alloc()
	rv := objc.Send[BluetoothObjectPushUIController](instance.ID, objc.Sel("initObjectPushWithBluetoothDevice:withFiles:delegate:"), inDevice, inFiles, inDelegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothObjectPushUIControllerObjectPushWithBluetoothDeviceWithFilesDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothObjectPushUIController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothObjectPushUIController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothObjectPushUIController */

// Runs the transfer UI as a sheet on the target window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/beginSheetModal(for:modalDelegate:didEnd:contextInfo:)
func (b_ BluetoothObjectPushUIController) BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow appkit.Window, modalDelegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) int {
	rv := objc.Send[int](b_.ID, objc.Sel("beginSheetModalForWindow:modalDelegate:didEndSelector:contextInfo:"), sheetWindow, modalDelegate, didEndSelector, contextInfo)
	return rv
}/* debug [instance_methods/method]: BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo */


// Gets the object representing the remote target device in the transfer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/getDevice()
func (b_ BluetoothObjectPushUIController) GetDevice() iobluetooth.BluetoothDevice {
	rv := objc.Send[iobluetooth.BluetoothDevice](b_.ID, objc.Sel("getDevice"))
	return rv
}/* debug [instance_methods/method]: GetDevice */


// Returns the title of the transfer panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/getTitle()
func (b_ BluetoothObjectPushUIController) GetTitle() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getTitle"))
	return rv
}/* debug [instance_methods/method]: GetTitle */


// Gets state of the transfer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/isTransferInProgress()
func (b_ BluetoothObjectPushUIController) IsTransferInProgress() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransferInProgress"))
	return rv
}/* debug [instance_methods/method]: IsTransferInProgress */


// Runs the transfer UI panel in a modal session
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/runModal()
func (b_ BluetoothObjectPushUIController) RunModal() {
	objc.Send[objc.ID](b_.ID, objc.Sel("runModal"))
}/* debug [instance_methods/method]: RunModal */


// Runs the transfer UI as a panel with no modal session
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/runPanel()
func (b_ BluetoothObjectPushUIController) RunPanel() {
	objc.Send[objc.ID](b_.ID, objc.Sel("runPanel"))
}/* debug [instance_methods/method]: RunPanel */


// Manually sets the icon used in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/setIconImage(_:)
func (b_ BluetoothObjectPushUIController) SetIconImage(image appkit.Image) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIconImage:"), image)
}/* debug [instance_methods/method]: SetIconImage */


// Sets the title of the panel when not run as a sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/setTitle(_:)
func (b_ BluetoothObjectPushUIController) SetTitle(windowTitle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), windowTitle)
}/* debug [instance_methods/method]: SetTitle */


// Stops the transfer UI
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/stop()
func (b_ BluetoothObjectPushUIController) Stop() {
	objc.Send[objc.ID](b_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothObjectPushUIController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothObjectPushUIController */


