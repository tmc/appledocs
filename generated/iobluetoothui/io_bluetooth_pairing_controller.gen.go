// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [BluetoothPairingController] class.
type IBluetoothPairingController interface {
	appkit.IWindowController
	GetPrompt() string
	GetTitle() string
	SetPrompt(prompt string)
}

// A NSWindowController subclass to display a window to initiate pairing to other bluetooth devices.
//
// Implementation of a window controller to handle pairing with a bluetooth device. This class will handle connecting to the Bluetooth Daemon for the purposes of searches, and displaying the results. When necessary this class will display a sheet asking the user for a PIN code. This window will not return anything to the caller if it is canceled or if pairing occurs.
//
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

// Alloc allocates a new instance without initialization.
func (bc _BluetoothPairingControllerClass) Alloc() BluetoothPairingController {
	rv := objc.Send[BluetoothPairingController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the title of the default/select button in the device selector panel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/getPrompt()
func (b_ BluetoothPairingController) GetPrompt() string {
	rv := objc.Send[string](b_.ID, objc.Sel("getPrompt"))
	return rv
}

// Returns the title of the device selector panel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/getTitle()
func (b_ BluetoothPairingController) GetTitle() string {
	rv := objc.Send[string](b_.ID, objc.Sel("getTitle"))
	return rv
}

// Sets the title of the default/select button in the device selector panel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPairingController/setPrompt(_:)
func (b_ BluetoothPairingController) SetPrompt(prompt string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrompt:"), objc.String(prompt))
}



