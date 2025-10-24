// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/iobluetooth"
)

/* debug [class.gen.go]: Generating class IOBluetoothPairingController */


/* debug [class_header]: Header for IOBluetoothPairingController */
// The class instance for the [BluetoothPairingController] class.
var (
	BluetoothPairingControllerClass     _BluetoothPairingControllerClass
	BluetoothPairingControllerClassOnce sync.Once
)

func getBluetoothPairingControllerClass() _BluetoothPairingControllerClass {
	BluetoothPairingControllerClassOnce.Do(func() {
		BluetoothPairingControllerClass = _BluetoothPairingControllerClass{objc.GetClass("IOBluetoothPairingController")}
	})
	return BluetoothPairingControllerClass
}

type _BluetoothPairingControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothPairingController */
// An interface definition for the [BluetoothPairingController] class.
type IBluetoothPairingController interface {
	appkit.IWindowController
	
/* debug [class_interface_properties]: Properties for BluetoothPairingController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothPairingController */
	// methods:
	AddAllowedUUID(allowedUUID iobluetooth.BluetoothSDPUUID)
	AddAllowedUUIDArray(allowedUUIDArray objc.IObject /* cross-framework: NSArray */)
	ClearAllowedUUIDs()
	GetDescriptionText() foundation.String
	GetOptions() BluetoothServiceBrowserControllerOptions /* typedef */
	GetPrompt() foundation.String
	GetResults() foundation.Array
	GetSearchAttributes() objc.IObject /* cross-framework: BluetoothDeviceSearchAttributes */
	GetTitle() foundation.String
	RunModal() int
	SetDescriptionText(descriptionText objc.IObject /* cross-framework: NSString */)
	SetOptions(options BluetoothServiceBrowserControllerOptions /* typedef */)
	SetPrompt(prompt objc.IObject /* cross-framework: NSString */)
	SetSearchAttributes(searchAttributes iobluetooth.IOBluetoothDeviceSearchAttributes)
	SetTitle(windowTitle objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothPairingController */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothPairingControllerClass) Alloc() BluetoothPairingController {
	rv := objc.Send[BluetoothPairingController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BluetoothPairingControllerClass) New() BluetoothPairingController {
	rv := objc.Send[BluetoothPairingController](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothPairingController) Init() BluetoothPairingController {
	rv := objc.Send[BluetoothPairingController](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothPairingController) Autorelease() BluetoothPairingController {
	rv := objc.Send[BluetoothPairingController](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothPairingController creates a new BluetoothPairingController instance.
func NewBluetoothPairingController() BluetoothPairingController {
	return getBluetoothPairingControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothPairingController */
// A NSWindowController subclass to display a window to initiate pairing to other bluetooth devices.
//
// Implementation of a window controller to handle pairing with a bluetooth device. This class will handle connecting to the Bluetooth Daemon for the purposes of searches, and displaying the results. When necessary this class will display a sheet asking the user for a PIN code. This window will not return anything to the caller if it is canceled or if pairing occurs.


// A NSWindowController subclass to display a window to initiate pairing to other bluetooth devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController
type BluetoothPairingController struct {
	appkit.WindowController
}

// BluetoothPairingControllerFrom constructs a [BluetoothPairingController] from an unsafe.Pointer.
//
// A NSWindowController subclass to display a window to initiate pairing to other bluetooth devices.
func BluetoothPairingControllerFrom(ptr unsafe.Pointer) BluetoothPairingController {
	return BluetoothPairingController{
		WindowController: appkit.WindowControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothPairingController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothPairingController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/pairingController
func (bc _BluetoothPairingControllerClass) PairingController() IBluetoothPairingController {
	rv := objc.Send[BluetoothPairingController](objc.ID(bc.class), objc.Sel("pairingController"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PairingController) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothPairingController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothPairingController */

// Adds a UUID to the list of UUIDs that are used to validate the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/addAllowedUUID(_:)
func (b_ BluetoothPairingController) AddAllowedUUID(allowedUUID iobluetooth.BluetoothSDPUUID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("addAllowedUUID:"), allowedUUID)
}/* debug [instance_methods/method]: AddAllowedUUID */


// Adds an array of UUIDs to the list of UUIDs that are used to validate the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/addAllowedUUIDArray(_:)
func (b_ BluetoothPairingController) AddAllowedUUIDArray(allowedUUIDArray objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("addAllowedUUIDArray:"), allowedUUIDArray)
}/* debug [instance_methods/method]: AddAllowedUUIDArray */


// Resets the controller back to the default state where it will accept any device the user selects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/clearAllowedUUIDs()
func (b_ BluetoothPairingController) ClearAllowedUUIDs() {
	objc.Send[objc.ID](b_.ID, objc.Sel("clearAllowedUUIDs"))
}/* debug [instance_methods/method]: ClearAllowedUUIDs */


// Returns the description text that appears in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/getDescriptionText()
func (b_ BluetoothPairingController) GetDescriptionText() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getDescriptionText"))
	return rv
}/* debug [instance_methods/method]: GetDescriptionText */


// Returns the option bits that control the panel’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/getOptions()
func (b_ BluetoothPairingController) GetOptions() BluetoothServiceBrowserControllerOptions /* typedef */ {
	rv := objc.Send[uint32](b_.ID, objc.Sel("getOptions"))
	return rv
}/* debug [instance_methods/method]: GetOptions */


// Returns the title of the default/select button in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/getPrompt()
func (b_ BluetoothPairingController) GetPrompt() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getPrompt"))
	return rv
}/* debug [instance_methods/method]: GetPrompt */


// Returns an NSArray of the devices that were paired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/getResults()
func (b_ BluetoothPairingController) GetResults() foundation.Array {
	rv := objc.Send[foundation.Array](b_.ID, objc.Sel("getResults"))
	return rv
}/* debug [instance_methods/method]: GetResults */


// Returns the search attributes that control the panel’s search/inquiry behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/getSearchAttributes()
func (b_ BluetoothPairingController) GetSearchAttributes() objc.IObject /* cross-framework: BluetoothDeviceSearchAttributes */ {
	rv := objc.Send[iobluetooth.BluetoothDeviceSearchAttributes](b_.ID, objc.Sel("getSearchAttributes"))
	return rv
}/* debug [instance_methods/method]: GetSearchAttributes */


// Returns the title of the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/getTitle()
func (b_ BluetoothPairingController) GetTitle() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getTitle"))
	return rv
}/* debug [instance_methods/method]: GetTitle */


// Runs the pairing panel in a modal session to allow the user to select a Bluetooth device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/runModal()
func (b_ BluetoothPairingController) RunModal() int {
	rv := objc.Send[int](b_.ID, objc.Sel("runModal"))
	return rv
}/* debug [instance_methods/method]: RunModal */


// Sets the description text that appears in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/setDescriptionText(_:)
func (b_ BluetoothPairingController) SetDescriptionText(descriptionText objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDescriptionText:"), descriptionText)
}/* debug [instance_methods/method]: SetDescriptionText */


// Sets the option bits that control the panel’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/setOptions(_:)
func (b_ BluetoothPairingController) SetOptions(options BluetoothServiceBrowserControllerOptions /* typedef */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOptions:"), options)
}/* debug [instance_methods/method]: SetOptions */


// Sets the title of the default/select button in the device selector panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/setPrompt(_:)
func (b_ BluetoothPairingController) SetPrompt(prompt objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrompt:"), prompt)
}/* debug [instance_methods/method]: SetPrompt */


// Sets the search attributes that control the panel’s search/inquiry behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/setSearchAttributes(_:)
func (b_ BluetoothPairingController) SetSearchAttributes(searchAttributes iobluetooth.IOBluetoothDeviceSearchAttributes) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSearchAttributes:"), searchAttributes)
}/* debug [instance_methods/method]: SetSearchAttributes */


// Sets the title of the panel when not run as a sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/setTitle(_:)
func (b_ BluetoothPairingController) SetTitle(windowTitle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), windowTitle)
}/* debug [instance_methods/method]: SetTitle */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothPairingController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothPairingController */



