// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mRemoteChannelID */


/* debug [class_header]: Header for mRemoteChannelID */
// The class instance for the [mRemoteChannelID] class.
var (
	MRemoteChannelIDClass     _mRemoteChannelIDClass
	MRemoteChannelIDClassOnce sync.Once
)

func getmRemoteChannelIDClass() _mRemoteChannelIDClass {
	MRemoteChannelIDClassOnce.Do(func() {
		MRemoteChannelIDClass = _mRemoteChannelIDClass{objc.GetClass("mRemoteChannelID")}
	})
	return MRemoteChannelIDClass
}

type _mRemoteChannelIDClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mRemoteChannelID */
// An interface definition for the [mRemoteChannelID] class.
type ImRemoteChannelID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mRemoteChannelID */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mRemoteChannelID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mRemoteChannelID */
// Alloc allocates a new instance without initialization.
func (mc _mRemoteChannelIDClass) Alloc() mRemoteChannelID {
	rv := objc.Send[mRemoteChannelID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mRemoteChannelIDClass) New() mRemoteChannelID {
	rv := objc.Send[mRemoteChannelID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mRemoteChannelID) Init() mRemoteChannelID {
	rv := objc.Send[mRemoteChannelID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mRemoteChannelID) Autorelease() mRemoteChannelID {
	rv := objc.Send[mRemoteChannelID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmRemoteChannelID creates a new mRemoteChannelID instance.
func NewmRemoteChannelID() mRemoteChannelID {
	return getmRemoteChannelIDClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mRemoteChannelID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mRemoteChannelID
type mRemoteChannelID struct {
	objectivec.Object
}

// mRemoteChannelIDFrom constructs a [mRemoteChannelID] from an unsafe.Pointer.
func mRemoteChannelIDFrom(ptr unsafe.Pointer) mRemoteChannelID {
	return mRemoteChannelID{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mRemoteChannelID *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mRemoteChannelID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mRemoteChannelID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mRemoteChannelID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mRemoteChannelID */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mRemoteChannelID */



