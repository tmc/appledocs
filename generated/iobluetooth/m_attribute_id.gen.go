// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mAttributeID */


/* debug [class_header]: Header for mAttributeID */
// The class instance for the [mAttributeID] class.
var (
	MAttributeIDClass     _mAttributeIDClass
	MAttributeIDClassOnce sync.Once
)

func getmAttributeIDClass() _mAttributeIDClass {
	MAttributeIDClassOnce.Do(func() {
		MAttributeIDClass = _mAttributeIDClass{objc.GetClass("mAttributeID")}
	})
	return MAttributeIDClass
}

type _mAttributeIDClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mAttributeID */
// An interface definition for the [mAttributeID] class.
type ImAttributeID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mAttributeID */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mAttributeID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mAttributeID */
// Alloc allocates a new instance without initialization.
func (mc _mAttributeIDClass) Alloc() mAttributeID {
	rv := objc.Send[mAttributeID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mAttributeIDClass) New() mAttributeID {
	rv := objc.Send[mAttributeID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mAttributeID) Init() mAttributeID {
	rv := objc.Send[mAttributeID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mAttributeID) Autorelease() mAttributeID {
	rv := objc.Send[mAttributeID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmAttributeID creates a new mAttributeID instance.
func NewmAttributeID() mAttributeID {
	return getmAttributeIDClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mAttributeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/mAttributeID
type mAttributeID struct {
	objectivec.Object
}

// mAttributeIDFrom constructs a [mAttributeID] from an unsafe.Pointer.
func mAttributeIDFrom(ptr unsafe.Pointer) mAttributeID {
	return mAttributeID{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mAttributeID *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mAttributeID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mAttributeID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mAttributeID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mAttributeID */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mAttributeID */



