// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mValue */


/* debug [class_header]: Header for mValue */
// The class instance for the [mValue] class.
var (
	MValueClass     _mValueClass
	MValueClassOnce sync.Once
)

func getmValueClass() _mValueClass {
	MValueClassOnce.Do(func() {
		MValueClass = _mValueClass{objc.GetClass("mValue")}
	})
	return MValueClass
}

type _mValueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mValue */
// An interface definition for the [mValue] class.
type ImValue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mValue */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mValue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mValue */
// Alloc allocates a new instance without initialization.
func (mc _mValueClass) Alloc() mValue {
	rv := objc.Send[mValue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mValueClass) New() mValue {
	rv := objc.Send[mValue](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mValue) Init() mValue {
	rv := objc.Send[mValue](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mValue) Autorelease() mValue {
	rv := objc.Send[mValue](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmValue creates a new mValue instance.
func NewmValue() mValue {
	return getmValueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/mValue
type mValue struct {
	objectivec.Object
}

// mValueFrom constructs a [mValue] from an unsafe.Pointer.
func mValueFrom(ptr unsafe.Pointer) mValue {
	return mValue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mValue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mValue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mValue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mValue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mValue */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mValue */



