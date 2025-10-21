// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [BluetoothDeviceSelectorController] class.
type IBluetoothDeviceSelectorController interface {
	appkit.IWindowController
	AddAllowedUUID(allowedUUID unsafe.Pointer)
	GetOptions() unsafe.Pointer
	GetPrompt() unsafe.Pointer
	GetSearchAttributes() unsafe.Pointer
	SetSearchAttributes(searchAttributes unsafe.Pointer)
	SetTitle(windowTitle string)
}

// A NSWindowController subclass to display a window to initiate pairing to other bluetooth devices.
//
// Implementation of a window controller to return a NSArray of selected bluetooth devices. This class will handle connecting to the Bluetooth Daemon for the purposes of searches, and displaying the results. This controller will return a NSArray of IOBluetoothDevice objects to the user.
//
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

// Alloc allocates a new instance without initialization.
func (bc _BluetoothDeviceSelectorControllerClass) Alloc() BluetoothDeviceSelectorController {
	rv := objc.Send[BluetoothDeviceSelectorController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Adds a UUID to the list of UUIDs that are used to validate the user’s selection.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/addAllowedUUID(_:)
func (b_ BluetoothDeviceSelectorController) AddAllowedUUID(allowedUUID unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("addAllowedUUID:"), allowedUUID)
}

// Returns the option bits that control the panel’s behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/getOptions()
func (b_ BluetoothDeviceSelectorController) GetOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getOptions"))
	return rv
}

// Returns the title of the default/select button in the device selector panel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/getPrompt()
func (b_ BluetoothDeviceSelectorController) GetPrompt() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getPrompt"))
	return rv
}

// Returns the search attributes that control the panel’s search/inquiry behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/getSearchAttributes()
func (b_ BluetoothDeviceSelectorController) GetSearchAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getSearchAttributes"))
	return rv
}

// Sets the search attributes that control the panel’s search/inquiry behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/setSearchAttributes(_:)
func (b_ BluetoothDeviceSelectorController) SetSearchAttributes(searchAttributes unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSearchAttributes:"), searchAttributes)
}

// Sets the title of the panel when not run as a sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothDeviceSelectorController/setTitle(_:)
func (b_ BluetoothDeviceSelectorController) SetTitle(windowTitle string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), objc.String(windowTitle))
}



