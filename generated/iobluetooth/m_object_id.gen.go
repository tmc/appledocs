// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mObjectID */


/* debug [class_header]: Header for mObjectID */
// The class instance for the [mObjectID] class.
var (
	MObjectIDClass     _mObjectIDClass
	MObjectIDClassOnce sync.Once
)

func getmObjectIDClass() _mObjectIDClass {
	MObjectIDClassOnce.Do(func() {
		MObjectIDClass = _mObjectIDClass{objc.GetClass("mObjectID")}
	})
	return MObjectIDClass
}

type _mObjectIDClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mObjectID */
// An interface definition for the [mObjectID] class.
type ImObjectID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mObjectID */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mObjectID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mObjectID */
// Alloc allocates a new instance without initialization.
func (mc _mObjectIDClass) Alloc() mObjectID {
	rv := objc.Send[mObjectID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mObjectIDClass) New() mObjectID {
	rv := objc.Send[mObjectID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mObjectID) Init() mObjectID {
	rv := objc.Send[mObjectID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mObjectID) Autorelease() mObjectID {
	rv := objc.Send[mObjectID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmObjectID creates a new mObjectID instance.
func NewmObjectID() mObjectID {
	return getmObjectIDClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mObjectID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mObjectID
type mObjectID struct {
	objectivec.Object
}

// mObjectIDFrom constructs a [mObjectID] from an unsafe.Pointer.
func mObjectIDFrom(ptr unsafe.Pointer) mObjectID {
	return mObjectID{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mObjectID *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mObjectID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mObjectID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mObjectID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mObjectID */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mObjectID */



