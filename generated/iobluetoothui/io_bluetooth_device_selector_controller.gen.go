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

/* debug [class.gen.go]: Generating class IOBluetoothDeviceSelectorController */


/* debug [class_header]: Header for IOBluetoothDeviceSelectorController */
// The class instance for the [BluetoothDeviceSelectorController] class.
var (
	BluetoothDeviceSelectorControllerClass     _BluetoothDeviceSelectorControllerClass
	BluetoothDeviceSelectorControllerClassOnce sync.Once
)

func getBluetoothDeviceSelectorControllerClass() _BluetoothDeviceSelectorControllerClass {
	BluetoothDeviceSelectorControllerClassOnce.Do(func() {
		BluetoothDeviceSelectorControllerClass = _BluetoothDeviceSelectorControllerClass{objc.GetClass("IOBluetoothDeviceSelectorController")}
	})
	return BluetoothDeviceSelectorControllerClass
}

type _BluetoothDeviceSelectorControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothDeviceSelectorController */
// An interface definition for the [BluetoothDeviceSelectorController] class.
type IBluetoothDeviceSelectorController interface {
	appkit.IWindowController
	
/* debug [class_interface_properties]: Properties for BluetoothDeviceSelectorController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothDeviceSelectorController */
	// methods:
	AddAllowedUUID(allowedUUID iobluetooth.BluetoothSDPUUID)
	AddAllowedUUIDArray(allowedUUIDArray objc.IObject /* cross-framework: NSArray */)
	BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow appkit.Window, modalDelegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) int
	ClearAllowedUUIDs()
	GetCancel() foundation.String
	GetDescriptionText() foundation.String
	GetHeader() foundation.String
	GetOptions() BluetoothServiceBrowserControllerOptions /* typedef */
	GetPrompt() foundation.String
	GetResults() foundation.Array
	GetSearchAttributes() objc.IObject /* cross-framework: BluetoothDeviceSearchAttributes */
	GetTitle() foundation.String
	RunModal() int
	SetCancel(prompt objc.IObject /* cross-framework: NSString */)
	SetDescriptionText(descriptionText objc.IObject /* cross-framework: NSString */)
	SetHeader(headerText objc.IObject /* cross-framework: NSString */)
	SetOptions(options BluetoothServiceBrowserControllerOptions /* typedef */)
	SetPrompt(prompt objc.IObject /* cross-framework: NSString */)
	SetSearchAttributes(searchAttributes iobluetooth.IOBluetoothDeviceSearchAttributes)
	SetTitle(windowTitle objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothDeviceSelectorController */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothDeviceSelectorControllerClass) Alloc() BluetoothDeviceSelectorController {
	rv := objc.Send[BluetoothDeviceSelectorController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BluetoothDeviceSelectorControllerClass) New() BluetoothDeviceSelectorController {
	rv := objc.Send[BluetoothDeviceSelectorController](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothDeviceSelectorController) Init() BluetoothDeviceSelectorController {
	rv := objc.Send[BluetoothDeviceSelectorController](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothDeviceSelectorController) Autorelease() BluetoothDeviceSelectorController {
	rv := objc.Send[BluetoothDeviceSelectorController](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothDeviceSelectorController creates a new BluetoothDeviceSelectorController instance.
func NewBluetoothDeviceSelectorController() BluetoothDeviceSelectorController {
	return getBluetoothDeviceSelectorControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothDeviceSelectorController */
// A NSWindowController subclass to display a window to initiate pairing to other bluetooth devices.
//
// Implementation of a window controller to return a NSArray of selected bluetooth devices. This class will handle connecting to the Bluetooth Daemon for the purposes of searches, and displaying the results. This controller will return a NSArray of IOBluetoothDevice objects to the user.


// A NSWindowController subclass to display a window to initiate pairing to other bluetooth devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController
type BluetoothDeviceSelectorController struct {
	appkit.WindowController
}

// BluetoothDeviceSelectorControllerFrom constructs a [BluetoothDeviceSelectorController] from an unsafe.Pointer.
//
// A NSWindowController subclass to display a window to initiate pairing to other bluetooth devices.
func BluetoothDeviceSelectorControllerFrom(ptr unsafe.Pointer) BluetoothDeviceSelectorController {
	return BluetoothDeviceSelectorController{
		WindowController: appkit.WindowControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothDeviceSelectorController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothDeviceSelectorController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/deviceSelector()
func (bc _BluetoothDeviceSelectorControllerClass) DeviceSelector() IBluetoothDeviceSelectorController {
	rv := objc.Send[BluetoothDeviceSelectorController](objc.ID(bc.class), objc.Sel("deviceSelector"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceSelector) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothDeviceSelectorController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothDeviceSelectorController */

// Adds a UUID to the list of UUIDs that are used to validate the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/addAllowedUUID(_:)
func (b_ BluetoothDeviceSelectorController) AddAllowedUUID(allowedUUID iobluetooth.BluetoothSDPUUID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("addAllowedUUID:"), allowedUUID)
}/* debug [instance_methods/method]: AddAllowedUUID */


// Adds an array of UUIDs to the list of UUIDs that are used to validate the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/addAllowedUUIDArray(_:)
func (b_ BluetoothDeviceSelectorController) AddAllowedUUIDArray(allowedUUIDArray objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("addAllowedUUIDArray:"), allowedUUIDArray)
}/* debug [instance_methods/method]: AddAllowedUUIDArray */


// Runs the device selector panel as a sheet on the target window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/beginSheetModal(for:modalDelegate:didEnd:contextInfo:)
func (b_ BluetoothDeviceSelectorController) BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow appkit.Window, modalDelegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) int {
	rv := objc.Send[int](b_.ID, objc.Sel("beginSheetModalForWindow:modalDelegate:didEndSelector:contextInfo:"), sheetWindow, modalDelegate, didEndSelector, contextInfo)
	return rv
}/* debug [instance_methods/method]: BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo */


// Resets the controller back to the default state where it will accept any device the user selects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/clearAllowedUUIDs()
func (b_ BluetoothDeviceSelectorController) ClearAllowedUUIDs() {
	objc.Send[objc.ID](b_.ID, objc.Sel("clearAllowedUUIDs"))
}/* debug [instance_methods/method]: ClearAllowedUUIDs */


// Returns the title of the default/cancel button in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/getCancel()
func (b_ BluetoothDeviceSelectorController) GetCancel() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getCancel"))
	return rv
}/* debug [instance_methods/method]: GetCancel */


// Returns the description text that appears in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/getDescriptionText()
func (b_ BluetoothDeviceSelectorController) GetDescriptionText() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getDescriptionText"))
	return rv
}/* debug [instance_methods/method]: GetDescriptionText */


// Returns the header text that appears in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/getHeader()
func (b_ BluetoothDeviceSelectorController) GetHeader() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getHeader"))
	return rv
}/* debug [instance_methods/method]: GetHeader */


// Returns the option bits that control the panel’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/getOptions()
func (b_ BluetoothDeviceSelectorController) GetOptions() BluetoothServiceBrowserControllerOptions /* typedef */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("getOptions"))
	return rv
}/* debug [instance_methods/method]: GetOptions */


// Returns the title of the default/select button in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/getPrompt()
func (b_ BluetoothDeviceSelectorController) GetPrompt() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getPrompt"))
	return rv
}/* debug [instance_methods/method]: GetPrompt */


// Returns the result of the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/getResults()
func (b_ BluetoothDeviceSelectorController) GetResults() foundation.Array {
	rv := objc.Send[foundation.Array](b_.ID, objc.Sel("getResults"))
	return rv
}/* debug [instance_methods/method]: GetResults */


// Returns the search attributes that control the panel’s search/inquiry behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/getSearchAttributes()
func (b_ BluetoothDeviceSelectorController) GetSearchAttributes() objc.IObject /* cross-framework: BluetoothDeviceSearchAttributes */ {
	rv := objc.Send[iobluetooth.BluetoothDeviceSearchAttributes](b_.ID, objc.Sel("getSearchAttributes"))
	return rv
}/* debug [instance_methods/method]: GetSearchAttributes */


// Returns the title of the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/getTitle()
func (b_ BluetoothDeviceSelectorController) GetTitle() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getTitle"))
	return rv
}/* debug [instance_methods/method]: GetTitle */


// Runs the device selector panel in a modal session to allow the user to select a Bluetooth device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/runModal()
func (b_ BluetoothDeviceSelectorController) RunModal() int {
	rv := objc.Send[int](b_.ID, objc.Sel("runModal"))
	return rv
}/* debug [instance_methods/method]: RunModal */


// Sets the title of the default/cancel button in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/setCancel(_:)
func (b_ BluetoothDeviceSelectorController) SetCancel(prompt objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCancel:"), prompt)
}/* debug [instance_methods/method]: SetCancel */


// Sets the description text that appears in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/setDescriptionText(_:)
func (b_ BluetoothDeviceSelectorController) SetDescriptionText(descriptionText objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDescriptionText:"), descriptionText)
}/* debug [instance_methods/method]: SetDescriptionText */


// Sets the header text that appears in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/setHeader(_:)
func (b_ BluetoothDeviceSelectorController) SetHeader(headerText objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setHeader:"), headerText)
}/* debug [instance_methods/method]: SetHeader */


// Sets the option bits that control the panel’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/setOptions(_:)
func (b_ BluetoothDeviceSelectorController) SetOptions(options BluetoothServiceBrowserControllerOptions /* typedef */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOptions:"), options)
}/* debug [instance_methods/method]: SetOptions */


// Sets the title of the default/select button in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/setPrompt(_:)
func (b_ BluetoothDeviceSelectorController) SetPrompt(prompt objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrompt:"), prompt)
}/* debug [instance_methods/method]: SetPrompt */


// Sets the search attributes that control the panel’s search/inquiry behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/setSearchAttributes(_:)
func (b_ BluetoothDeviceSelectorController) SetSearchAttributes(searchAttributes iobluetooth.IOBluetoothDeviceSearchAttributes) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSearchAttributes:"), searchAttributes)
}/* debug [instance_methods/method]: SetSearchAttributes */


// Sets the title of the panel when not run as a sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/setTitle(_:)
func (b_ BluetoothDeviceSelectorController) SetTitle(windowTitle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), windowTitle)
}/* debug [instance_methods/method]: SetTitle */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothDeviceSelectorController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothDeviceSelectorController */



