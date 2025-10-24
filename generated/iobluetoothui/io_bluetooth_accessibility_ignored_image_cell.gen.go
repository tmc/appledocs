// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class IOBluetoothAccessibilityIgnoredImageCell */


/* debug [class_header]: Header for IOBluetoothAccessibilityIgnoredImageCell */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothAccessibilityIgnoredImageCell */
// An interface definition for the [BluetoothAccessibilityIgnoredImageCell] class.
type IBluetoothAccessibilityIgnoredImageCell interface {
	appkit.ImageCell
	
/* debug [class_interface_properties]: Properties for BluetoothAccessibilityIgnoredImageCell */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothAccessibilityIgnoredImageCell */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothAccessibilityIgnoredImageCell */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothAccessibilityIgnoredImageCellClass) Alloc() BluetoothAccessibilityIgnoredImageCell {
	rv := objc.Send[BluetoothAccessibilityIgnoredImageCell](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothAccessibilityIgnoredImageCell */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothAccessibilityIgnoredImageCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothAccessibilityIgnoredImageCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothAccessibilityIgnoredImageCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothAccessibilityIgnoredImageCell */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothAccessibilityIgnoredImageCell */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothAccessibilityIgnoredImageCell */



