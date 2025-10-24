// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mChannelIsOpen */


/* debug [class_header]: Header for mChannelIsOpen */
// The class instance for the [mChannelIsOpen] class.
var (
	MChannelIsOpenClass     _mChannelIsOpenClass
	MChannelIsOpenClassOnce sync.Once
)

func getmChannelIsOpenClass() _mChannelIsOpenClass {
	MChannelIsOpenClassOnce.Do(func() {
		MChannelIsOpenClass = _mChannelIsOpenClass{objc.GetClass("mChannelIsOpen")}
	})
	return MChannelIsOpenClass
}

type _mChannelIsOpenClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mChannelIsOpen */
// An interface definition for the [mChannelIsOpen] class.
type ImChannelIsOpen interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mChannelIsOpen */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mChannelIsOpen */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mChannelIsOpen */
// Alloc allocates a new instance without initialization.
func (mc _mChannelIsOpenClass) Alloc() mChannelIsOpen {
	rv := objc.Send[mChannelIsOpen](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mChannelIsOpenClass) New() mChannelIsOpen {
	rv := objc.Send[mChannelIsOpen](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mChannelIsOpen) Init() mChannelIsOpen {
	rv := objc.Send[mChannelIsOpen](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mChannelIsOpen) Autorelease() mChannelIsOpen {
	rv := objc.Send[mChannelIsOpen](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmChannelIsOpen creates a new mChannelIsOpen instance.
func NewmChannelIsOpen() mChannelIsOpen {
	return getmChannelIsOpenClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mChannelIsOpen */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mChannelIsOpen
type mChannelIsOpen struct {
	objectivec.Object
}

// mChannelIsOpenFrom constructs a [mChannelIsOpen] from an unsafe.Pointer.
func mChannelIsOpenFrom(ptr unsafe.Pointer) mChannelIsOpen {
	return mChannelIsOpen{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mChannelIsOpen *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mChannelIsOpen */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mChannelIsOpen */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mChannelIsOpen */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mChannelIsOpen */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mChannelIsOpen */



