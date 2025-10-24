// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mRFCOMMChannelID */


/* debug [class_header]: Header for mRFCOMMChannelID */
// The class instance for the [mRFCOMMChannelID] class.
var (
	MRFCOMMChannelIDClass     _mRFCOMMChannelIDClass
	MRFCOMMChannelIDClassOnce sync.Once
)

func getmRFCOMMChannelIDClass() _mRFCOMMChannelIDClass {
	MRFCOMMChannelIDClassOnce.Do(func() {
		MRFCOMMChannelIDClass = _mRFCOMMChannelIDClass{objc.GetClass("mRFCOMMChannelID")}
	})
	return MRFCOMMChannelIDClass
}

type _mRFCOMMChannelIDClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mRFCOMMChannelID */
// An interface definition for the [mRFCOMMChannelID] class.
type ImRFCOMMChannelID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mRFCOMMChannelID */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mRFCOMMChannelID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mRFCOMMChannelID */
// Alloc allocates a new instance without initialization.
func (mc _mRFCOMMChannelIDClass) Alloc() mRFCOMMChannelID {
	rv := objc.Send[mRFCOMMChannelID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mRFCOMMChannelIDClass) New() mRFCOMMChannelID {
	rv := objc.Send[mRFCOMMChannelID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mRFCOMMChannelID) Init() mRFCOMMChannelID {
	rv := objc.Send[mRFCOMMChannelID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mRFCOMMChannelID) Autorelease() mRFCOMMChannelID {
	rv := objc.Send[mRFCOMMChannelID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmRFCOMMChannelID creates a new mRFCOMMChannelID instance.
func NewmRFCOMMChannelID() mRFCOMMChannelID {
	return getmRFCOMMChannelIDClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mRFCOMMChannelID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mRFCOMMChannelID
type mRFCOMMChannelID struct {
	objectivec.Object
}

// mRFCOMMChannelIDFrom constructs a [mRFCOMMChannelID] from an unsafe.Pointer.
func mRFCOMMChannelIDFrom(ptr unsafe.Pointer) mRFCOMMChannelID {
	return mRFCOMMChannelID{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mRFCOMMChannelID *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mRFCOMMChannelID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mRFCOMMChannelID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mRFCOMMChannelID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mRFCOMMChannelID */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mRFCOMMChannelID */



