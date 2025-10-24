// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mIsClosed */


/* debug [class_header]: Header for mIsClosed */
// The class instance for the [mIsClosed] class.
var (
	MIsClosedClass     _mIsClosedClass
	MIsClosedClassOnce sync.Once
)

func getmIsClosedClass() _mIsClosedClass {
	MIsClosedClassOnce.Do(func() {
		MIsClosedClass = _mIsClosedClass{objc.GetClass("mIsClosed")}
	})
	return MIsClosedClass
}

type _mIsClosedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mIsClosed */
// An interface definition for the [mIsClosed] class.
type ImIsClosed interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mIsClosed */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mIsClosed */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mIsClosed */
// Alloc allocates a new instance without initialization.
func (mc _mIsClosedClass) Alloc() mIsClosed {
	rv := objc.Send[mIsClosed](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mIsClosedClass) New() mIsClosed {
	rv := objc.Send[mIsClosed](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIsClosed) Init() mIsClosed {
	rv := objc.Send[mIsClosed](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIsClosed) Autorelease() mIsClosed {
	rv := objc.Send[mIsClosed](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIsClosed creates a new mIsClosed instance.
func NewmIsClosed() mIsClosed {
	return getmIsClosedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mIsClosed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mIsClosed
type mIsClosed struct {
	objectivec.Object
}

// mIsClosedFrom constructs a [mIsClosed] from an unsafe.Pointer.
func mIsClosedFrom(ptr unsafe.Pointer) mIsClosed {
	return mIsClosed{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mIsClosed *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mIsClosed */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mIsClosed */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mIsClosed */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mIsClosed */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mIsClosed */



