// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BluetoothAccessibilityIgnoredTextFieldCell] class.
var (
	BluetoothAccessibilityIgnoredTextFieldCellClass     _BluetoothAccessibilityIgnoredTextFieldCellClass
	BluetoothAccessibilityIgnoredTextFieldCellClassOnce sync.Once
)

func getBluetoothAccessibilityIgnoredTextFieldCellClass() _BluetoothAccessibilityIgnoredTextFieldCellClass {
	BluetoothAccessibilityIgnoredTextFieldCellClassOnce.Do(func() {
		BluetoothAccessibilityIgnoredTextFieldCellClass = _BluetoothAccessibilityIgnoredTextFieldCellClass{objc.GetClass("IOBluetoothAccessibilityIgnoredTextFieldCell")}
	})
	return BluetoothAccessibilityIgnoredTextFieldCellClass
}

type _BluetoothAccessibilityIgnoredTextFieldCellClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothAccessibilityIgnoredTextFieldCell] class.
type IBluetoothAccessibilityIgnoredTextFieldCell interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothAccessibilityIgnoredTextFieldCell
type BluetoothAccessibilityIgnoredTextFieldCell struct {
	objectivec.Object
}

// BluetoothAccessibilityIgnoredTextFieldCellFrom constructs a [BluetoothAccessibilityIgnoredTextFieldCell] from an unsafe.Pointer.
func BluetoothAccessibilityIgnoredTextFieldCellFrom(ptr unsafe.Pointer) BluetoothAccessibilityIgnoredTextFieldCell {
	return BluetoothAccessibilityIgnoredTextFieldCell{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothAccessibilityIgnoredTextFieldCellClass) Alloc() BluetoothAccessibilityIgnoredTextFieldCell {
	rv := objc.Send[BluetoothAccessibilityIgnoredTextFieldCell](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothAccessibilityIgnoredTextFieldCellClass) New() BluetoothAccessibilityIgnoredTextFieldCell {
	rv := objc.Send[BluetoothAccessibilityIgnoredTextFieldCell](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothAccessibilityIgnoredTextFieldCell) Init() BluetoothAccessibilityIgnoredTextFieldCell {
	rv := objc.Send[BluetoothAccessibilityIgnoredTextFieldCell](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothAccessibilityIgnoredTextFieldCell) Autorelease() BluetoothAccessibilityIgnoredTextFieldCell {
	rv := objc.Send[BluetoothAccessibilityIgnoredTextFieldCell](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothAccessibilityIgnoredTextFieldCell creates a new BluetoothAccessibilityIgnoredTextFieldCell instance.
func NewBluetoothAccessibilityIgnoredTextFieldCell() BluetoothAccessibilityIgnoredTextFieldCell {
	return getBluetoothAccessibilityIgnoredTextFieldCellClass().New()
}




