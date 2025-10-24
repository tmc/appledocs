// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class CABTLEMIDIWindowController */


/* debug [class_header]: Header for CABTLEMIDIWindowController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BTLEMIDIWindowController */
// An interface definition for the [BTLEMIDIWindowController] class.
type IBTLEMIDIWindowController interface {
	appkit.IWindowController
	
/* debug [class_interface_properties]: Properties for BTLEMIDIWindowController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BTLEMIDIWindowController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BTLEMIDIWindowController */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BTLEMIDIWindowController */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BTLEMIDIWindowController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BTLEMIDIWindowController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BTLEMIDIWindowController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BTLEMIDIWindowController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BTLEMIDIWindowController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CABTLEMIDIWindowController */



