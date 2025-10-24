// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mChannelID */


/* debug [class_header]: Header for mChannelID */
// The class instance for the [mChannelID] class.
var (
	MChannelIDClass     _mChannelIDClass
	MChannelIDClassOnce sync.Once
)

func getmChannelIDClass() _mChannelIDClass {
	MChannelIDClassOnce.Do(func() {
		MChannelIDClass = _mChannelIDClass{objc.GetClass("mChannelID")}
	})
	return MChannelIDClass
}

type _mChannelIDClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mChannelID */
// An interface definition for the [mChannelID] class.
type ImChannelID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mChannelID */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mChannelID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mChannelID */
// Alloc allocates a new instance without initialization.
func (mc _mChannelIDClass) Alloc() mChannelID {
	rv := objc.Send[mChannelID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mChannelIDClass) New() mChannelID {
	rv := objc.Send[mChannelID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mChannelID) Init() mChannelID {
	rv := objc.Send[mChannelID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mChannelID) Autorelease() mChannelID {
	rv := objc.Send[mChannelID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmChannelID creates a new mChannelID instance.
func NewmChannelID() mChannelID {
	return getmChannelIDClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mChannelID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mChannelID
type mChannelID struct {
	objectivec.Object
}

// mChannelIDFrom constructs a [mChannelID] from an unsafe.Pointer.
func mChannelIDFrom(ptr unsafe.Pointer) mChannelID {
	return mChannelID{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mChannelID *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mChannelID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mChannelID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mChannelID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mChannelID */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mChannelID */



