// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mIsIncoming */


/* debug [class_header]: Header for mIsIncoming */
// The class instance for the [mIsIncoming] class.
var (
	MIsIncomingClass     _mIsIncomingClass
	MIsIncomingClassOnce sync.Once
)

func getmIsIncomingClass() _mIsIncomingClass {
	MIsIncomingClassOnce.Do(func() {
		MIsIncomingClass = _mIsIncomingClass{objc.GetClass("mIsIncoming")}
	})
	return MIsIncomingClass
}

type _mIsIncomingClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mIsIncoming */
// An interface definition for the [mIsIncoming] class.
type ImIsIncoming interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mIsIncoming */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mIsIncoming */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mIsIncoming */
// Alloc allocates a new instance without initialization.
func (mc _mIsIncomingClass) Alloc() mIsIncoming {
	rv := objc.Send[mIsIncoming](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mIsIncomingClass) New() mIsIncoming {
	rv := objc.Send[mIsIncoming](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIsIncoming) Init() mIsIncoming {
	rv := objc.Send[mIsIncoming](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIsIncoming) Autorelease() mIsIncoming {
	rv := objc.Send[mIsIncoming](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIsIncoming creates a new mIsIncoming instance.
func NewmIsIncoming() mIsIncoming {
	return getmIsIncomingClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mIsIncoming */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mIsIncoming
type mIsIncoming struct {
	objectivec.Object
}

// mIsIncomingFrom constructs a [mIsIncoming] from an unsafe.Pointer.
func mIsIncomingFrom(ptr unsafe.Pointer) mIsIncoming {
	return mIsIncoming{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mIsIncoming *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mIsIncoming */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mIsIncoming */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mIsIncoming */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mIsIncoming */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mIsIncoming */



