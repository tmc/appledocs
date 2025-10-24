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

/* debug [class.gen.go]: Generating class IOBluetoothServiceBrowserController */


/* debug [class_header]: Header for IOBluetoothServiceBrowserController */
// The class instance for the [BluetoothServiceBrowserController] class.
var (
	BluetoothServiceBrowserControllerClass     _BluetoothServiceBrowserControllerClass
	BluetoothServiceBrowserControllerClassOnce sync.Once
)

func getBluetoothServiceBrowserControllerClass() _BluetoothServiceBrowserControllerClass {
	BluetoothServiceBrowserControllerClassOnce.Do(func() {
		BluetoothServiceBrowserControllerClass = _BluetoothServiceBrowserControllerClass{objc.GetClass("IOBluetoothServiceBrowserController")}
	})
	return BluetoothServiceBrowserControllerClass
}

type _BluetoothServiceBrowserControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothServiceBrowserController */
// An interface definition for the [BluetoothServiceBrowserController] class.
type IBluetoothServiceBrowserController interface {
	appkit.IWindowController
	
/* debug [class_interface_properties]: Properties for BluetoothServiceBrowserController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothServiceBrowserController */
	// methods:
	AddAllowedUUID(allowedUUID iobluetooth.BluetoothSDPUUID)
	AddAllowedUUIDArray(allowedUUIDArray objc.IObject /* cross-framework: NSArray */)
	BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow appkit.Window, modalDelegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) int
	ClearAllowedUUIDs()
	GetDescriptionText() foundation.String
	GetOptions() BluetoothServiceBrowserControllerOptions /* typedef */
	GetPrompt() foundation.String
	GetServiceBrowserControllerRef() BluetoothServiceBrowserControllerRef /* typedef */
	GetResults() foundation.Array
	GetSearchAttributes() objc.IObject /* cross-framework: BluetoothDeviceSearchAttributes */
	GetTitle() foundation.String
	RunModal() int
	SetDescriptionText(descriptionText objc.IObject /* cross-framework: NSString */)
	SetOptions(inOptions BluetoothServiceBrowserControllerOptions /* typedef */)
	SetPrompt(prompt objc.IObject /* cross-framework: NSString */)
	SetSearchAttributes(searchAttributes iobluetooth.IOBluetoothDeviceSearchAttributes)
	SetTitle(windowTitle objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothServiceBrowserController */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothServiceBrowserControllerClass) Alloc() BluetoothServiceBrowserController {
	rv := objc.Send[BluetoothServiceBrowserController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BluetoothServiceBrowserControllerClass) New() BluetoothServiceBrowserController {
	rv := objc.Send[BluetoothServiceBrowserController](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothServiceBrowserController) Init() BluetoothServiceBrowserController {
	rv := objc.Send[BluetoothServiceBrowserController](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothServiceBrowserController) Autorelease() BluetoothServiceBrowserController {
	rv := objc.Send[BluetoothServiceBrowserController](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothServiceBrowserController creates a new BluetoothServiceBrowserController instance.
func NewBluetoothServiceBrowserController() BluetoothServiceBrowserController {
	return getBluetoothServiceBrowserControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothServiceBrowserController */
// A NSWindowController subclass to display a window to search for and perform SDP queries on bluetooth devices within range.
//
// This NSWindowController subclass will bring up a generic Bluetooth search and SDP browsing window allowing the user to find devices within range, perform SDP queries on a particular device, and select a SDP service to connect to. The client application can provide NSArrays of valid service UUIDs to allow, and an NSArray of valid device types to allow. The device type filter is not yet implemented.


// A NSWindowController subclass to display a window to search for and perform SDP queries on bluetooth devices within range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController
type BluetoothServiceBrowserController struct {
	appkit.WindowController
}

// BluetoothServiceBrowserControllerFrom constructs a [BluetoothServiceBrowserController] from an unsafe.Pointer.
//
// A NSWindowController subclass to display a window to search for and perform SDP queries on bluetooth devices within range.
func BluetoothServiceBrowserControllerFrom(ptr unsafe.Pointer) BluetoothServiceBrowserController {
	return BluetoothServiceBrowserController{
		WindowController: appkit.WindowControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothServiceBrowserController */

// Allocator work Bluetooth Service Browser window controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/init(_:)
func NewBluetoothServiceBrowserController(inOptions BluetoothServiceBrowserControllerOptions /* typedef */) BluetoothServiceBrowserController {
	rv := objc.Send[BluetoothServiceBrowserController](objc.ID(getBluetoothServiceBrowserControllerClass().class), objc.Sel("serviceBrowserController:"), inOptions)
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothServiceBrowserController */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothServiceBrowserController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/browseDevices:options:
func (bc _BluetoothServiceBrowserControllerClass) BrowseDevicesOptions(outRecord unsafe.Pointer, inOptions BluetoothServiceBrowserControllerOptions /* typedef */) int {
	rv := objc.Send[int](objc.ID(bc.class), objc.Sel("browseDevices:options:"), outRecord, inOptions)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BrowseDevicesOptions) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/browseDevicesAsSheetForWindow:options:window:
func (bc _BluetoothServiceBrowserControllerClass) BrowseDevicesAsSheetForWindowOptionsWindow(outRecord unsafe.Pointer, inOptions BluetoothServiceBrowserControllerOptions /* typedef */, inWindow appkit.Window) int {
	rv := objc.Send[int](objc.ID(bc.class), objc.Sel("browseDevicesAsSheetForWindow:options:window:"), outRecord, inOptions, inWindow)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BrowseDevicesAsSheetForWindowOptionsWindow) */


// Allocator work Bluetooth Service Browser window controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/init(_:)
func (bc _BluetoothServiceBrowserControllerClass) ServiceBrowserController(inOptions BluetoothServiceBrowserControllerOptions /* typedef */) IBluetoothServiceBrowserController {
	rv := objc.Send[BluetoothServiceBrowserController](objc.ID(bc.class), objc.Sel("serviceBrowserController:"), inOptions)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ServiceBrowserController) */


// Method call to convert an IOBluetoothServiceBrowserControllerRef into an IOBluetoothServiceBrowserController *.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/withServiceBrowserControllerRef(_:)
func (bc _BluetoothServiceBrowserControllerClass) WithServiceBrowserControllerRef(serviceBrowserControllerRef BluetoothServiceBrowserControllerRef /* typedef */) IBluetoothServiceBrowserController {
	rv := objc.Send[BluetoothServiceBrowserController](objc.ID(bc.class), objc.Sel("withServiceBrowserControllerRef:"), serviceBrowserControllerRef)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WithServiceBrowserControllerRef) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothServiceBrowserController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothServiceBrowserController */

// Adds a UUID to the list of UUIDs that are used to validate the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/addAllowedUUID(_:)
func (b_ BluetoothServiceBrowserController) AddAllowedUUID(allowedUUID iobluetooth.BluetoothSDPUUID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("addAllowedUUID:"), allowedUUID)
}/* debug [instance_methods/method]: AddAllowedUUID */


// Adds an array of UUIDs to the list of UUIDs that are used to validate the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/addAllowedUUIDArray(_:)
func (b_ BluetoothServiceBrowserController) AddAllowedUUIDArray(allowedUUIDArray objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("addAllowedUUIDArray:"), allowedUUIDArray)
}/* debug [instance_methods/method]: AddAllowedUUIDArray */


// Runs the service browser panel as a sheet on the target window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/beginSheetModal(for:modalDelegate:didEnd:contextInfo:)
func (b_ BluetoothServiceBrowserController) BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow appkit.Window, modalDelegate objc.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) int {
	rv := objc.Send[int](b_.ID, objc.Sel("beginSheetModalForWindow:modalDelegate:didEndSelector:contextInfo:"), sheetWindow, modalDelegate, didEndSelector, contextInfo)
	return rv
}/* debug [instance_methods/method]: BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo */


// Resets the controller back to the default state where it will accept any device the user selects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/clearAllowedUUIDs()
func (b_ BluetoothServiceBrowserController) ClearAllowedUUIDs() {
	objc.Send[objc.ID](b_.ID, objc.Sel("clearAllowedUUIDs"))
}/* debug [instance_methods/method]: ClearAllowedUUIDs */


// Returns the description text that appears in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/getDescriptionText()
func (b_ BluetoothServiceBrowserController) GetDescriptionText() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getDescriptionText"))
	return rv
}/* debug [instance_methods/method]: GetDescriptionText */


// Returns the option bits that control the panel’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/getOptions()
func (b_ BluetoothServiceBrowserController) GetOptions() BluetoothServiceBrowserControllerOptions /* typedef */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("getOptions"))
	return rv
}/* debug [instance_methods/method]: GetOptions */


// Returns the title of the default/select button in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/getPrompt()
func (b_ BluetoothServiceBrowserController) GetPrompt() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getPrompt"))
	return rv
}/* debug [instance_methods/method]: GetPrompt */


// Returns an IOBluetoothServiceBrowserControllerRef representation of the target IOBluetoothServiceBrowserController object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/getRef()
func (b_ BluetoothServiceBrowserController) GetServiceBrowserControllerRef() BluetoothServiceBrowserControllerRef /* typedef */ {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getServiceBrowserControllerRef"))
	return rv
}/* debug [instance_methods/method]: GetServiceBrowserControllerRef */


// Returns the result of the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/getResults()
func (b_ BluetoothServiceBrowserController) GetResults() foundation.Array {
	rv := objc.Send[foundation.Array](b_.ID, objc.Sel("getResults"))
	return rv
}/* debug [instance_methods/method]: GetResults */


// Returns the search attributes that control the panel’s search/inquiry behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/getSearchAttributes()
func (b_ BluetoothServiceBrowserController) GetSearchAttributes() objc.IObject /* cross-framework: BluetoothDeviceSearchAttributes */ {
	rv := objc.Send[iobluetooth.BluetoothDeviceSearchAttributes](b_.ID, objc.Sel("getSearchAttributes"))
	return rv
}/* debug [instance_methods/method]: GetSearchAttributes */


// Returns the title of the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/getTitle()
func (b_ BluetoothServiceBrowserController) GetTitle() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getTitle"))
	return rv
}/* debug [instance_methods/method]: GetTitle */


// Runs the service browser panel in a modal session to allow the user to select a service on a Bluetooth device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/runModal()
func (b_ BluetoothServiceBrowserController) RunModal() int {
	rv := objc.Send[int](b_.ID, objc.Sel("runModal"))
	return rv
}/* debug [instance_methods/method]: RunModal */


// Sets the description text that appears in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/setDescriptionText(_:)
func (b_ BluetoothServiceBrowserController) SetDescriptionText(descriptionText objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDescriptionText:"), descriptionText)
}/* debug [instance_methods/method]: SetDescriptionText */


// Modify the options for the window controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/setOptions(_:)
func (b_ BluetoothServiceBrowserController) SetOptions(inOptions BluetoothServiceBrowserControllerOptions /* typedef */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOptions:"), inOptions)
}/* debug [instance_methods/method]: SetOptions */


// Sets the title of the default/select button in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/setPrompt(_:)
func (b_ BluetoothServiceBrowserController) SetPrompt(prompt objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrompt:"), prompt)
}/* debug [instance_methods/method]: SetPrompt */


// Sets the search attributes that control the panel’s search/inquiry behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/setSearchAttributes(_:)
func (b_ BluetoothServiceBrowserController) SetSearchAttributes(searchAttributes iobluetooth.IOBluetoothDeviceSearchAttributes) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSearchAttributes:"), searchAttributes)
}/* debug [instance_methods/method]: SetSearchAttributes */


// Sets the title of the panel when not run as a sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/setTitle(_:)
func (b_ BluetoothServiceBrowserController) SetTitle(windowTitle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), windowTitle)
}/* debug [instance_methods/method]: SetTitle */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothServiceBrowserController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothServiceBrowserController */


