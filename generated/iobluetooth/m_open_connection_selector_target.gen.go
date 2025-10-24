// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mOpenConnectionSelectorTarget */


/* debug [class_header]: Header for mOpenConnectionSelectorTarget */
// The class instance for the [mOpenConnectionSelectorTarget] class.
var (
	MOpenConnectionSelectorTargetClass     _mOpenConnectionSelectorTargetClass
	MOpenConnectionSelectorTargetClassOnce sync.Once
)

func getmOpenConnectionSelectorTargetClass() _mOpenConnectionSelectorTargetClass {
	MOpenConnectionSelectorTargetClassOnce.Do(func() {
		MOpenConnectionSelectorTargetClass = _mOpenConnectionSelectorTargetClass{objc.GetClass("mOpenConnectionSelectorTarget")}
	})
	return MOpenConnectionSelectorTargetClass
}

type _mOpenConnectionSelectorTargetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mOpenConnectionSelectorTarget */
// An interface definition for the [mOpenConnectionSelectorTarget] class.
type ImOpenConnectionSelectorTarget interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mOpenConnectionSelectorTarget */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mOpenConnectionSelectorTarget */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mOpenConnectionSelectorTarget */
// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionSelectorTargetClass) Alloc() mOpenConnectionSelectorTarget {
	rv := objc.Send[mOpenConnectionSelectorTarget](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mOpenConnectionSelectorTargetClass) New() mOpenConnectionSelectorTarget {
	rv := objc.Send[mOpenConnectionSelectorTarget](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionSelectorTarget) Init() mOpenConnectionSelectorTarget {
	rv := objc.Send[mOpenConnectionSelectorTarget](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionSelectorTarget) Autorelease() mOpenConnectionSelectorTarget {
	rv := objc.Send[mOpenConnectionSelectorTarget](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionSelectorTarget creates a new mOpenConnectionSelectorTarget instance.
func NewmOpenConnectionSelectorTarget() mOpenConnectionSelectorTarget {
	return getmOpenConnectionSelectorTargetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mOpenConnectionSelectorTarget */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionSelectorTarget
type mOpenConnectionSelectorTarget struct {
	objectivec.Object
}

// mOpenConnectionSelectorTargetFrom constructs a [mOpenConnectionSelectorTarget] from an unsafe.Pointer.
func mOpenConnectionSelectorTargetFrom(ptr unsafe.Pointer) mOpenConnectionSelectorTarget {
	return mOpenConnectionSelectorTarget{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mOpenConnectionSelectorTarget *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mOpenConnectionSelectorTarget */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mOpenConnectionSelectorTarget */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mOpenConnectionSelectorTarget */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mOpenConnectionSelectorTarget */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mOpenConnectionSelectorTarget */



