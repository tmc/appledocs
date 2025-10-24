// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class IOBluetoothAccessibilityIgnoredTextFieldCell */


/* debug [class_header]: Header for IOBluetoothAccessibilityIgnoredTextFieldCell */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothAccessibilityIgnoredTextFieldCell */
// An interface definition for the [BluetoothAccessibilityIgnoredTextFieldCell] class.
type IBluetoothAccessibilityIgnoredTextFieldCell interface {
	appkit.ITextFieldCell
	
/* debug [class_interface_properties]: Properties for BluetoothAccessibilityIgnoredTextFieldCell */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothAccessibilityIgnoredTextFieldCell */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothAccessibilityIgnoredTextFieldCell */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothAccessibilityIgnoredTextFieldCellClass) Alloc() BluetoothAccessibilityIgnoredTextFieldCell {
	rv := objc.Send[BluetoothAccessibilityIgnoredTextFieldCell](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothAccessibilityIgnoredTextFieldCell */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothAccessibilityIgnoredTextFieldCell
type BluetoothAccessibilityIgnoredTextFieldCell struct {
	appkit.TextFieldCell
}

// BluetoothAccessibilityIgnoredTextFieldCellFrom constructs a [BluetoothAccessibilityIgnoredTextFieldCell] from an unsafe.Pointer.
func BluetoothAccessibilityIgnoredTextFieldCellFrom(ptr unsafe.Pointer) BluetoothAccessibilityIgnoredTextFieldCell {
	return BluetoothAccessibilityIgnoredTextFieldCell{
		TextFieldCell: appkit.TextFieldCellFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothAccessibilityIgnoredTextFieldCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothAccessibilityIgnoredTextFieldCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothAccessibilityIgnoredTextFieldCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothAccessibilityIgnoredTextFieldCell */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothAccessibilityIgnoredTextFieldCell */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothAccessibilityIgnoredTextFieldCell */



