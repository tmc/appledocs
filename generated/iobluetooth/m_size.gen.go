// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mSize */


/* debug [class_header]: Header for mSize */
// The class instance for the [mSize] class.
var (
	MSizeClass     _mSizeClass
	MSizeClassOnce sync.Once
)

func getmSizeClass() _mSizeClass {
	MSizeClassOnce.Do(func() {
		MSizeClass = _mSizeClass{objc.GetClass("mSize")}
	})
	return MSizeClass
}

type _mSizeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mSize */
// An interface definition for the [mSize] class.
type ImSize interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mSize */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mSize */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mSize */
// Alloc allocates a new instance without initialization.
func (mc _mSizeClass) Alloc() mSize {
	rv := objc.Send[mSize](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mSizeClass) New() mSize {
	rv := objc.Send[mSize](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mSize) Init() mSize {
	rv := objc.Send[mSize](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mSize) Autorelease() mSize {
	rv := objc.Send[mSize](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmSize creates a new mSize instance.
func NewmSize() mSize {
	return getmSizeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/mSize
type mSize struct {
	objectivec.Object
}

// mSizeFrom constructs a [mSize] from an unsafe.Pointer.
func mSizeFrom(ptr unsafe.Pointer) mSize {
	return mSize{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mSize *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mSize */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mSize */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mSize */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mSize */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mSize */



