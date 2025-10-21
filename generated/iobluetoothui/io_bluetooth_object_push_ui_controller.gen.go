// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [BluetoothObjectPushUIController] class.
type IBluetoothObjectPushUIController interface {
	appkit.IWindowController
	BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow unsafe.Pointer, modalDelegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer) unsafe.Pointer
	GetDevice() unsafe.Pointer
	GetTitle() unsafe.Pointer
	IsTransferInProgress() bool
	SetIconImage(image unsafe.Pointer)
	SetTitle(windowTitle string)
	Stop()
}

// An NSWindowController subclass that supports the creation of an IOBluetoothObjectPushUIController object.
//
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

// Alloc allocates a new instance without initialization.
func (bc _BluetoothObjectPushUIControllerClass) Alloc() BluetoothObjectPushUIController {
	rv := objc.Send[BluetoothObjectPushUIController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates and returns a new IOBluetoothObjectPush object
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/init(objectPushWith:withFiles:delegate:)
func NewBluetoothObjectPushUIControllerObjectPushWithBluetoothDeviceWithFilesDelegate(inDevice unsafe.Pointer, inFiles objc.ID, inDelegate objc.ID) BluetoothObjectPushUIController {
	instance := getBluetoothObjectPushUIControllerClass().Alloc()
	rv := objc.Send[BluetoothObjectPushUIController](instance.ID, objc.Sel("initObjectPushWithBluetoothDevice:withFiles:delegate:"), inDevice, inFiles, inDelegate)
	rv.Autorelease()
	return rv
}


// Runs the transfer UI as a sheet on the target window.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/beginSheetModal(for:modalDelegate:didEnd:contextInfo:)
func (b_ BluetoothObjectPushUIController) BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow unsafe.Pointer, modalDelegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("beginSheetModalForWindow:modalDelegate:didEndSelector:contextInfo:"), sheetWindow, modalDelegate, didEndSelector, contextInfo)
	return rv
}

// Gets the object representing the remote target device in the transfer.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/getDevice()
func (b_ BluetoothObjectPushUIController) GetDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getDevice"))
	return rv
}

// Returns the title of the transfer panel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/getTitle()
func (b_ BluetoothObjectPushUIController) GetTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getTitle"))
	return rv
}

// Gets state of the transfer
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/isTransferInProgress()
func (b_ BluetoothObjectPushUIController) IsTransferInProgress() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransferInProgress"))
	return rv
}

// Manually sets the icon used in the panel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/setIconImage(_:)
func (b_ BluetoothObjectPushUIController) SetIconImage(image unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIconImage:"), image)
}

// Sets the title of the panel when not run as a sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/setTitle(_:)
func (b_ BluetoothObjectPushUIController) SetTitle(windowTitle string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), objc.String(windowTitle))
}

// Stops the transfer UI
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/stop()
func (b_ BluetoothObjectPushUIController) Stop() {
	objc.Send[objc.ID](b_.ID, objc.Sel("stop"))
}


