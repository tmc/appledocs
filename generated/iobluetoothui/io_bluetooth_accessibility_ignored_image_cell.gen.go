// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [BluetoothAccessibilityIgnoredImageCell] class.
var (
	BluetoothAccessibilityIgnoredImageCellClass     _BluetoothAccessibilityIgnoredImageCellClass
	BluetoothAccessibilityIgnoredImageCellClassOnce sync.Once
)

func getBluetoothAccessibilityIgnoredImageCellClass() _BluetoothAccessibilityIgnoredImageCellClass {
	BluetoothAccessibilityIgnoredImageCellClassOnce.Do(func() {
		BluetoothAccessibilityIgnoredImageCellClass = _BluetoothAccessibilityIgnoredImageCellClass{objc.GetClass("IOBluetoothAccessibilityIgnoredImageCell")}
	})
	return BluetoothAccessibilityIgnoredImageCellClass
}

type _BluetoothAccessibilityIgnoredImageCellClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothAccessibilityIgnoredImageCell] class.
type IBluetoothAccessibilityIgnoredImageCell interface {
	appkit.IImageCell
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothAccessibilityIgnoredImageCell
type BluetoothAccessibilityIgnoredImageCell struct {
	appkit.ImageCell
}

// BluetoothAccessibilityIgnoredImageCellFrom constructs a [BluetoothAccessibilityIgnoredImageCell] from an unsafe.Pointer.
func BluetoothAccessibilityIgnoredImageCellFrom(ptr unsafe.Pointer) BluetoothAccessibilityIgnoredImageCell {
	return BluetoothAccessibilityIgnoredImageCell{
		ImageCell: appkit.ImageCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothAccessibilityIgnoredImageCellClass) Alloc() BluetoothAccessibilityIgnoredImageCell {
	rv := objc.Send[BluetoothAccessibilityIgnoredImageCell](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothAccessibilityIgnoredImageCellClass) New() BluetoothAccessibilityIgnoredImageCell {
	rv := objc.Send[BluetoothAccessibilityIgnoredImageCell](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothAccessibilityIgnoredImageCell) Init() BluetoothAccessibilityIgnoredImageCell {
	rv := objc.Send[BluetoothAccessibilityIgnoredImageCell](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothAccessibilityIgnoredImageCell) Autorelease() BluetoothAccessibilityIgnoredImageCell {
	rv := objc.Send[BluetoothAccessibilityIgnoredImageCell](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothAccessibilityIgnoredImageCell creates a new BluetoothAccessibilityIgnoredImageCell instance.
func NewBluetoothAccessibilityIgnoredImageCell() BluetoothAccessibilityIgnoredImageCell {
	return getBluetoothAccessibilityIgnoredImageCellClass().New()
}




