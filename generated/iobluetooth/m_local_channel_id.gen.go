// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mLocalChannelID */


/* debug [class_header]: Header for mLocalChannelID */
// The class instance for the [mLocalChannelID] class.
var (
	MLocalChannelIDClass     _mLocalChannelIDClass
	MLocalChannelIDClassOnce sync.Once
)

func getmLocalChannelIDClass() _mLocalChannelIDClass {
	MLocalChannelIDClassOnce.Do(func() {
		MLocalChannelIDClass = _mLocalChannelIDClass{objc.GetClass("mLocalChannelID")}
	})
	return MLocalChannelIDClass
}

type _mLocalChannelIDClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mLocalChannelID */
// An interface definition for the [mLocalChannelID] class.
type ImLocalChannelID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mLocalChannelID */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mLocalChannelID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mLocalChannelID */
// Alloc allocates a new instance without initialization.
func (mc _mLocalChannelIDClass) Alloc() mLocalChannelID {
	rv := objc.Send[mLocalChannelID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mLocalChannelIDClass) New() mLocalChannelID {
	rv := objc.Send[mLocalChannelID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mLocalChannelID) Init() mLocalChannelID {
	rv := objc.Send[mLocalChannelID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mLocalChannelID) Autorelease() mLocalChannelID {
	rv := objc.Send[mLocalChannelID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmLocalChannelID creates a new mLocalChannelID instance.
func NewmLocalChannelID() mLocalChannelID {
	return getmLocalChannelIDClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mLocalChannelID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mLocalChannelID
type mLocalChannelID struct {
	objectivec.Object
}

// mLocalChannelIDFrom constructs a [mLocalChannelID] from an unsafe.Pointer.
func mLocalChannelIDFrom(ptr unsafe.Pointer) mLocalChannelID {
	return mLocalChannelID{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mLocalChannelID *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mLocalChannelID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mLocalChannelID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mLocalChannelID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mLocalChannelID */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mLocalChannelID */



