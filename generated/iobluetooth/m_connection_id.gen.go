// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mConnectionID */


/* debug [class_header]: Header for mConnectionID */
// The class instance for the [mConnectionID] class.
var (
	MConnectionIDClass     _mConnectionIDClass
	MConnectionIDClassOnce sync.Once
)

func getmConnectionIDClass() _mConnectionIDClass {
	MConnectionIDClassOnce.Do(func() {
		MConnectionIDClass = _mConnectionIDClass{objc.GetClass("mConnectionID")}
	})
	return MConnectionIDClass
}

type _mConnectionIDClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mConnectionID */
// An interface definition for the [mConnectionID] class.
type ImConnectionID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mConnectionID */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mConnectionID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mConnectionID */
// Alloc allocates a new instance without initialization.
func (mc _mConnectionIDClass) Alloc() mConnectionID {
	rv := objc.Send[mConnectionID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mConnectionIDClass) New() mConnectionID {
	rv := objc.Send[mConnectionID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mConnectionID) Init() mConnectionID {
	rv := objc.Send[mConnectionID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mConnectionID) Autorelease() mConnectionID {
	rv := objc.Send[mConnectionID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmConnectionID creates a new mConnectionID instance.
func NewmConnectionID() mConnectionID {
	return getmConnectionIDClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mConnectionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mConnectionID
type mConnectionID struct {
	objectivec.Object
}

// mConnectionIDFrom constructs a [mConnectionID] from an unsafe.Pointer.
func mConnectionIDFrom(ptr unsafe.Pointer) mConnectionID {
	return mConnectionID{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mConnectionID *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mConnectionID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mConnectionID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mConnectionID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mConnectionID */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mConnectionID */



