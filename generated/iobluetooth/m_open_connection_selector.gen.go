// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mOpenConnectionSelector */


/* debug [class_header]: Header for mOpenConnectionSelector */
// The class instance for the [mOpenConnectionSelector] class.
var (
	MOpenConnectionSelectorClass     _mOpenConnectionSelectorClass
	MOpenConnectionSelectorClassOnce sync.Once
)

func getmOpenConnectionSelectorClass() _mOpenConnectionSelectorClass {
	MOpenConnectionSelectorClassOnce.Do(func() {
		MOpenConnectionSelectorClass = _mOpenConnectionSelectorClass{objc.GetClass("mOpenConnectionSelector")}
	})
	return MOpenConnectionSelectorClass
}

type _mOpenConnectionSelectorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mOpenConnectionSelector */
// An interface definition for the [mOpenConnectionSelector] class.
type ImOpenConnectionSelector interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mOpenConnectionSelector */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mOpenConnectionSelector */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mOpenConnectionSelector */
// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionSelectorClass) Alloc() mOpenConnectionSelector {
	rv := objc.Send[mOpenConnectionSelector](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mOpenConnectionSelectorClass) New() mOpenConnectionSelector {
	rv := objc.Send[mOpenConnectionSelector](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionSelector) Init() mOpenConnectionSelector {
	rv := objc.Send[mOpenConnectionSelector](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionSelector) Autorelease() mOpenConnectionSelector {
	rv := objc.Send[mOpenConnectionSelector](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionSelector creates a new mOpenConnectionSelector instance.
func NewmOpenConnectionSelector() mOpenConnectionSelector {
	return getmOpenConnectionSelectorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mOpenConnectionSelector */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionSelector
type mOpenConnectionSelector struct {
	objectivec.Object
}

// mOpenConnectionSelectorFrom constructs a [mOpenConnectionSelector] from an unsafe.Pointer.
func mOpenConnectionSelectorFrom(ptr unsafe.Pointer) mOpenConnectionSelector {
	return mOpenConnectionSelector{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mOpenConnectionSelector *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mOpenConnectionSelector */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mOpenConnectionSelector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mOpenConnectionSelector */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mOpenConnectionSelector */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mOpenConnectionSelector */



