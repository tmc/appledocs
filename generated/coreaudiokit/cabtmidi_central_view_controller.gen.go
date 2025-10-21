// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BTMIDICentralViewController] class.
var (
	BTMIDICentralViewControllerClass     _BTMIDICentralViewControllerClass
	BTMIDICentralViewControllerClassOnce sync.Once
)

func getBTMIDICentralViewControllerClass() _BTMIDICentralViewControllerClass {
	BTMIDICentralViewControllerClassOnce.Do(func() {
		BTMIDICentralViewControllerClass = _BTMIDICentralViewControllerClass{objc.GetClass("CABTMIDICentralViewController")}
	})
	return BTMIDICentralViewControllerClass
}

type _BTMIDICentralViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [BTMIDICentralViewController] class.
type IBTMIDICentralViewController interface {
	objectivec.IObject
}

// A view controller that displays nearby Bluetooth-based MIDI peripherals.
//
// To let the user search for nearby MIDI peripherals, create a new object and then either present it modally or push it onto a view controller. No other configuration of the object is necessary. Once the user interface is visible, the iOS device finds nearby peripherals and displays them to the user. If the user selects a peripheral, it’s automatically paired with this iOS device. The object manages its own user interface and is dismissed automatically. Once connected, the peripheral appears as a MIDI device, just like any other connected MIDI device. MIDI commands sent to the peripheral are automatically played. For more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CABTMIDICentralViewController
type BTMIDICentralViewController struct {
	objectivec.Object
}

// BTMIDICentralViewControllerFrom constructs a [BTMIDICentralViewController] from an unsafe.Pointer.
//
// A view controller that displays nearby Bluetooth-based MIDI peripherals.
func BTMIDICentralViewControllerFrom(ptr unsafe.Pointer) BTMIDICentralViewController {
	return BTMIDICentralViewController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BTMIDICentralViewControllerClass) Alloc() BTMIDICentralViewController {
	rv := objc.Send[BTMIDICentralViewController](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BTMIDICentralViewControllerClass) New() BTMIDICentralViewController {
	rv := objc.Send[BTMIDICentralViewController](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BTMIDICentralViewController) Init() BTMIDICentralViewController {
	rv := objc.Send[BTMIDICentralViewController](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BTMIDICentralViewController) Autorelease() BTMIDICentralViewController {
	rv := objc.Send[BTMIDICentralViewController](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBTMIDICentralViewController creates a new BTMIDICentralViewController instance.
func NewBTMIDICentralViewController() BTMIDICentralViewController {
	return getBTMIDICentralViewControllerClass().New()
}




