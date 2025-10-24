// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothObject */


/* debug [class_header]: Header for IOBluetoothObject */
// The class instance for the [BluetoothObject] class.
var (
	BluetoothObjectClass     _BluetoothObjectClass
	BluetoothObjectClassOnce sync.Once
)

func getBluetoothObjectClass() _BluetoothObjectClass {
	BluetoothObjectClassOnce.Do(func() {
		BluetoothObjectClass = _BluetoothObjectClass{objc.GetClass("IOBluetoothObject")}
	})
	return BluetoothObjectClass
}

type _BluetoothObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothObject */
// An interface definition for the [BluetoothObject] class.
type IBluetoothObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BluetoothObject */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothObject */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothObjectClass) Alloc() BluetoothObject {
	rv := objc.Send[BluetoothObject](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BluetoothObjectClass) New() BluetoothObject {
	rv := objc.Send[BluetoothObject](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothObject) Init() BluetoothObject {
	rv := objc.Send[BluetoothObject](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothObject) Autorelease() BluetoothObject {
	rv := objc.Send[BluetoothObject](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothObject creates a new BluetoothObject instance.
func NewBluetoothObject() BluetoothObject {
	return getBluetoothObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothObject
type BluetoothObject struct {
	objectivec.Object
}

// BluetoothObjectFrom constructs a [BluetoothObject] from an unsafe.Pointer.
func BluetoothObjectFrom(ptr unsafe.Pointer) BluetoothObject {
	return BluetoothObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothObject */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothObject */



