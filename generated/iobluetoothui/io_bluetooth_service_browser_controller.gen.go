// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [BluetoothServiceBrowserController] class.
type IBluetoothServiceBrowserController interface {
	appkit.IWindowController
	BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow appkit.IWindow, modalDelegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) unsafe.Pointer
	DiscoverWithDeviceAttributesServiceListServiceRecord(deviceAttributes unsafe.Pointer, serviceArray objectivec.IObject, outRecord unsafe.Pointer) unsafe.Pointer
	GetResults() foundation.Array
}

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

// Alloc allocates a new instance without initialization.
func (bc _BluetoothServiceBrowserControllerClass) Alloc() BluetoothServiceBrowserController {
	rv := objc.Send[BluetoothServiceBrowserController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/browseDevices:options:
func (bc _BluetoothServiceBrowserControllerClass) BrowseDevicesOptions(outRecord unsafe.Pointer, inOptions BluetoothServiceBrowserControllerOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("browseDevices:options:"), outRecord, inOptions)
	return rv
}


// Runs the service browser panel as a sheet on the target window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/beginSheetModal(for:modalDelegate:didEnd:contextInfo:)
func (b_ BluetoothServiceBrowserController) BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow appkit.IWindow, modalDelegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("beginSheetModalForWindow:modalDelegate:didEndSelector:contextInfo:"), sheetWindow, modalDelegate, didEndSelector, contextInfo)
	return rv
}


// Invoke an already created window controller to display, and run the modal dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/discoverWithDeviceAttributes:serviceList:serviceRecord:
func (b_ BluetoothServiceBrowserController) DiscoverWithDeviceAttributesServiceListServiceRecord(deviceAttributes unsafe.Pointer, serviceArray objectivec.IObject, outRecord unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("discoverWithDeviceAttributes:serviceList:serviceRecord:"), deviceAttributes, serviceArray, outRecord)
	return rv
}


// Returns the result of the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothServiceBrowserController/getResults()
func (b_ BluetoothServiceBrowserController) GetResults() foundation.Array {
	rv := objc.Send[foundation.Array](b_.ID, objc.Sel("getResults"))
	return rv
}




