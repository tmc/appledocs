// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mLastNameUpdate */


/* debug [class_header]: Header for mLastNameUpdate */
// The class instance for the [mLastNameUpdate] class.
var (
	MLastNameUpdateClass     _mLastNameUpdateClass
	MLastNameUpdateClassOnce sync.Once
)

func getmLastNameUpdateClass() _mLastNameUpdateClass {
	MLastNameUpdateClassOnce.Do(func() {
		MLastNameUpdateClass = _mLastNameUpdateClass{objc.GetClass("mLastNameUpdate")}
	})
	return MLastNameUpdateClass
}

type _mLastNameUpdateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mLastNameUpdate */
// An interface definition for the [mLastNameUpdate] class.
type ImLastNameUpdate interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mLastNameUpdate */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mLastNameUpdate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mLastNameUpdate */
// Alloc allocates a new instance without initialization.
func (mc _mLastNameUpdateClass) Alloc() mLastNameUpdate {
	rv := objc.Send[mLastNameUpdate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mLastNameUpdateClass) New() mLastNameUpdate {
	rv := objc.Send[mLastNameUpdate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mLastNameUpdate) Init() mLastNameUpdate {
	rv := objc.Send[mLastNameUpdate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mLastNameUpdate) Autorelease() mLastNameUpdate {
	rv := objc.Send[mLastNameUpdate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmLastNameUpdate creates a new mLastNameUpdate instance.
func NewmLastNameUpdate() mLastNameUpdate {
	return getmLastNameUpdateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mLastNameUpdate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mLastNameUpdate
type mLastNameUpdate struct {
	objectivec.Object
}

// mLastNameUpdateFrom constructs a [mLastNameUpdate] from an unsafe.Pointer.
func mLastNameUpdateFrom(ptr unsafe.Pointer) mLastNameUpdate {
	return mLastNameUpdate{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mLastNameUpdate *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mLastNameUpdate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mLastNameUpdate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mLastNameUpdate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mLastNameUpdate */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mLastNameUpdate */



