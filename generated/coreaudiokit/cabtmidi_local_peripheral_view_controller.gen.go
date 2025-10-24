// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CABTMIDILocalPeripheralViewController */


/* debug [class_header]: Header for CABTMIDILocalPeripheralViewController */
// The class instance for the [BTMIDILocalPeripheralViewController] class.
var (
	BTMIDILocalPeripheralViewControllerClass     _BTMIDILocalPeripheralViewControllerClass
	BTMIDILocalPeripheralViewControllerClassOnce sync.Once
)

func getBTMIDILocalPeripheralViewControllerClass() _BTMIDILocalPeripheralViewControllerClass {
	BTMIDILocalPeripheralViewControllerClassOnce.Do(func() {
		BTMIDILocalPeripheralViewControllerClass = _BTMIDILocalPeripheralViewControllerClass{objc.GetClass("CABTMIDILocalPeripheralViewController")}
	})
	return BTMIDILocalPeripheralViewControllerClass
}

type _BTMIDILocalPeripheralViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BTMIDILocalPeripheralViewController */
// An interface definition for the [BTMIDILocalPeripheralViewController] class.
type IBTMIDILocalPeripheralViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for BTMIDILocalPeripheralViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BTMIDILocalPeripheralViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BTMIDILocalPeripheralViewController */
// Alloc allocates a new instance without initialization.
func (bc _BTMIDILocalPeripheralViewControllerClass) Alloc() BTMIDILocalPeripheralViewController {
	rv := objc.Send[BTMIDILocalPeripheralViewController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BTMIDILocalPeripheralViewControllerClass) New() BTMIDILocalPeripheralViewController {
	rv := objc.Send[BTMIDILocalPeripheralViewController](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BTMIDILocalPeripheralViewController) Init() BTMIDILocalPeripheralViewController {
	rv := objc.Send[BTMIDILocalPeripheralViewController](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BTMIDILocalPeripheralViewController) Autorelease() BTMIDILocalPeripheralViewController {
	rv := objc.Send[BTMIDILocalPeripheralViewController](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBTMIDILocalPeripheralViewController creates a new BTMIDILocalPeripheralViewController instance.
func NewBTMIDILocalPeripheralViewController() BTMIDILocalPeripheralViewController {
	return getBTMIDILocalPeripheralViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BTMIDILocalPeripheralViewController */
// A view controller that advertises an iOS device as a Bluetooth-based MIDI peripheral.
//
// To advertise the iOS device as a Bluetooth MIDI peripheral, create a new object and then either present it modally or push it onto a view controller. No other configuration of the object is necessary. Once the user interface is displayed, the iOS device is discoverable by another device looking for Bluetooth MIDI peripherals, such as an iOS device displaying a object. The object manages its own user interface and is dismissed automatically. Once connected, the peripheral appears as a MIDI device, just like any other connected MIDI device. MIDI commands sent to the peripheral are automatically played. For more information, see .


// A view controller that advertises an iOS device as a Bluetooth-based MIDI peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CABTMIDILocalPeripheralViewController
type BTMIDILocalPeripheralViewController struct {
	ViewController
}

// BTMIDILocalPeripheralViewControllerFrom constructs a [BTMIDILocalPeripheralViewController] from an unsafe.Pointer.
//
// A view controller that advertises an iOS device as a Bluetooth-based MIDI peripheral.
func BTMIDILocalPeripheralViewControllerFrom(ptr unsafe.Pointer) BTMIDILocalPeripheralViewController {
	return BTMIDILocalPeripheralViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BTMIDILocalPeripheralViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BTMIDILocalPeripheralViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BTMIDILocalPeripheralViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BTMIDILocalPeripheralViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BTMIDILocalPeripheralViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CABTMIDILocalPeripheralViewController */



