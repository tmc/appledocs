// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	// methods:
	BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow objc.IObject /* cross-framework Window */, modalDelegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) Return /* not a class type */
}

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



// Runs the transfer UI as a sheet on the target window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothObjectPushUIController/beginSheetModal(for:modalDelegate:didEnd:contextInfo:)
func (b_ BluetoothObjectPushUIController) BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheetWindow objc.IObject /* cross-framework Window */, modalDelegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) Return /* not a class type */ {
	rv := objc.Send[Return](b_.ID, objc.Sel("beginSheetModalForWindow:modalDelegate:didEndSelector:contextInfo:"), sheetWindow, modalDelegate, didEndSelector, contextInfo)
	return rv
}



