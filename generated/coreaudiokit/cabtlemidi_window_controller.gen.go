// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)





// The class instance for the [BTLEMIDIWindowController] class.
var (
	BTLEMIDIWindowControllerClass     _BTLEMIDIWindowControllerClass
	BTLEMIDIWindowControllerClassOnce sync.Once
)

func getBTLEMIDIWindowControllerClass() _BTLEMIDIWindowControllerClass {
	BTLEMIDIWindowControllerClassOnce.Do(func() {
		BTLEMIDIWindowControllerClass = _BTLEMIDIWindowControllerClass{objc.GetClass("CABTLEMIDIWindowController")}
	})
	return BTLEMIDIWindowControllerClass
}

type _BTLEMIDIWindowControllerClass struct {
	class objc.Class
}





// An interface definition for the [BTLEMIDIWindowController] class.
type IBTLEMIDIWindowController interface {
	appkit.IWindowController
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (bc _BTLEMIDIWindowControllerClass) Alloc() BTLEMIDIWindowController {
	rv := objc.Send[BTLEMIDIWindowController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BTLEMIDIWindowControllerClass) New() BTLEMIDIWindowController {
	rv := objc.Send[BTLEMIDIWindowController](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BTLEMIDIWindowController) Init() BTLEMIDIWindowController {
	rv := objc.Send[BTLEMIDIWindowController](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BTLEMIDIWindowController) Autorelease() BTLEMIDIWindowController {
	rv := objc.Send[BTLEMIDIWindowController](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBTLEMIDIWindowController creates a new BTLEMIDIWindowController instance.
func NewBTLEMIDIWindowController() BTLEMIDIWindowController {
	return getBTLEMIDIWindowControllerClass().New()
}





// A window controller that displays nearby Bluetooth-based MIDI peripherals.


// A window controller that displays nearby Bluetooth-based MIDI peripherals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CABTLEMIDIWindowController
type BTLEMIDIWindowController struct {
	appkit.WindowController
}

// BTLEMIDIWindowControllerFrom constructs a [BTLEMIDIWindowController] from an unsafe.Pointer.
//
// A window controller that displays nearby Bluetooth-based MIDI peripherals.
func BTLEMIDIWindowControllerFrom(ptr unsafe.Pointer) BTLEMIDIWindowController {
	return BTLEMIDIWindowController{
		WindowController: appkit.WindowControllerFrom(ptr),
	}
}































