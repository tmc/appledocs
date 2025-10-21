// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [BluetoothPasskeyDisplay] class.
var (
	BluetoothPasskeyDisplayClass     _BluetoothPasskeyDisplayClass
	BluetoothPasskeyDisplayClassOnce sync.Once
)

func getBluetoothPasskeyDisplayClass() _BluetoothPasskeyDisplayClass {
	BluetoothPasskeyDisplayClassOnce.Do(func() {
		BluetoothPasskeyDisplayClass = _BluetoothPasskeyDisplayClass{objc.GetClass("IOBluetoothPasskeyDisplay")}
	})
	return BluetoothPasskeyDisplayClass
}

type _BluetoothPasskeyDisplayClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothPasskeyDisplay] class.
type IBluetoothPasskeyDisplay interface {
	appkit.IView
	RetreatPasskeyIndicator()
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay
type BluetoothPasskeyDisplay struct {
	appkit.View
}

// BluetoothPasskeyDisplayFrom constructs a [BluetoothPasskeyDisplay] from an unsafe.Pointer.
func BluetoothPasskeyDisplayFrom(ptr unsafe.Pointer) BluetoothPasskeyDisplay {
	return BluetoothPasskeyDisplay{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothPasskeyDisplayClass) Alloc() BluetoothPasskeyDisplay {
	rv := objc.Send[BluetoothPasskeyDisplay](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothPasskeyDisplayClass) New() BluetoothPasskeyDisplay {
	rv := objc.Send[BluetoothPasskeyDisplay](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothPasskeyDisplay) Init() BluetoothPasskeyDisplay {
	rv := objc.Send[BluetoothPasskeyDisplay](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothPasskeyDisplay) Autorelease() BluetoothPasskeyDisplay {
	rv := objc.Send[BluetoothPasskeyDisplay](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothPasskeyDisplay creates a new BluetoothPasskeyDisplay instance.
func NewBluetoothPasskeyDisplay() BluetoothPasskeyDisplay {
	return getBluetoothPasskeyDisplayClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/retreatPasskeyIndicator()
func (b_ BluetoothPasskeyDisplay) RetreatPasskeyIndicator() {
	objc.Send[objc.ID](b_.ID, objc.Sel("retreatPasskeyIndicator"))
}



